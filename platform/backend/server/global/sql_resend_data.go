package global

import (
	"backend/api/intercom/model"
	"backend/internal/ginweb"
	"backend/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

var (
	// 重送 API 類型定義列舉

	// 單一錢包 request 類型
	RESEND_TYPE_SINGLEWALLET = 1
)

// 訊息重送紀錄表
type ReSendData struct {
	Id               string                    `json:"id"`           // 訂單號(orderid)
	AgentId          int                       `json:"agent_id"`     // 代理編號
	LevelCode        string                    `json:"level_code"`   // 代理層級碼
	RequestType      int                       `json:"request_type"` // type of api request
	IsReSend         bool                      `json:"is_resend"`    // 重試狀態
	Retries          int                       `json:"retries"`      // 重試次數
	RequestFrom      string                    `json:"request_from"` // API 原始請求參數
	RequestTo        string                    `json:"request_to"`   // API 回調請求參數
	RequestFromParam model.SingleWalletRequest `json:"-"`
	RequestToURL     url.URL                   `json:"-"`
	ApiKey           string                    `json:"-"`
	CreateTime       time.Time                 `json:"create_time"` // 創建時間
	UpdateTime       time.Time                 `json:"update_time"` // 更新時間
}

// table resend_data 指定查詢條件,
// dataType: 指定重送 API 類型(類型參照 RESEND_TYPE),
// retries: 指定重試次數
func SelectReSendData(logger ginweb.ILogger, db *sql.DB, dataType int, retries int) []*ReSendData {
	tmps := make([]*ReSendData, 0)

	query := `select "id", "agent_id", "level_code", "request_type", "is_resend", "retries", "request_from", "request_to", "create_time", "update_time" 
	from resend_data 
	where request_type=$1 AND is_resend=false AND retries<$2
	order by create_time DESC
	limit 100;`

	rows, err := db.Query(query, dataType, retries)
	if err != nil {
		logger.Printf("SelectReSendData() has error to select resend_data, err  is: %v", err)
		return tmps
	}

	for rows.Next() {
		tmp := &ReSendData{}

		if err := rows.Scan(&tmp.Id, &tmp.AgentId, &tmp.LevelCode, &tmp.RequestType, &tmp.IsReSend,
			&tmp.Retries, &tmp.RequestFrom, &tmp.RequestTo, &tmp.CreateTime, &tmp.UpdateTime); err != nil {
			logger.Printf("SelectReSendData() has error to Scan resend_data, err  is: %v", err)
			continue
		}

		if err = utils.ToStruct([]byte(tmp.RequestTo), &tmp.RequestToURL); err != nil {
			logger.Printf("SelectReSendData() has error to struct RequestTo, err  is: %v", err)
			continue
		}

		if err = json.Unmarshal([]byte(tmp.RequestFrom), &tmp.RequestFromParam); err != nil {
			// if err = utils.ToStruct([]byte(tmp.RequestFrom), &tmp.RequestFromParam); err != nil {
			logger.Printf("SelectReSendData() has error to struct RequestFrom, err  is: %v", err)
			continue
		}

		aObj := AgentCache.Get(tmp.AgentId)
		if aObj == nil {
			continue
		}
		// 更換最新的 api walletConnectInfo
		tmp.RequestToURL.Scheme = aObj.WalletConnInfo.Scheme
		tmp.RequestToURL.Host = aObj.WalletConnInfo.Domain
		tmp.RequestToURL.Path = aObj.WalletConnInfo.Path
		tmp.ApiKey = aObj.WalletConnInfo.ApiKey

		tmps = append(tmps, tmp)
	}

	return tmps
}

// table resend_data 指定更新條件,
// finished: 是否已成功重送,
// dataType: 指定重送 API 類型(類型參照 RESEND_TYPE),
// orderIds: 要更新的注單id,
// tip: 每次更新都會增加 retries 的次數
func UpdateReSendData(logger ginweb.ILogger, db *sql.DB, finished bool, dataType int, orderIds []string) {
	if len(orderIds) > 0 {
		query := `UPDATE "public"."resend_data" 
			SET "is_resend" = $1, "retries" = "retries"+1, "update_time"=now()
			WHERE request_type = $2 AND "id" IN (%s);`

		part := ""
		for i := 0; i < len(orderIds); i++ {
			part += "'" + orderIds[i] + "'"
			if i+1 < len(orderIds) {
				part += ", "
			}
		}

		finishQuery := fmt.Sprintf(query, part)

		result, err := db.Exec(finishQuery, finished, dataType)
		if err != nil {
			logger.Printf("updateReSendData() has error to insert resend_data, err  is: %v", err)
		} else {
			if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
				logger.Printf("updateReSendData() insert resend_data rowsAffectedCount is 0, query is %s", query)
			}
		}
	} else {
		logger.Println("updateReSendData() params orderId length is 0")
	}
}

// table resend_data 指定插入條件,
// dataType: 指定重送 API 類型(類型參照 RESEND_TYPE),
// orderId: 要更新的注單id,
// request: API request 原始參數字串
func InsertReSendData(logger ginweb.ILogger, db *sql.DB, agentId, dataType int, levelCode, orderId, requestFrom, requestTo string) {
	query := `INSERT INTO "public"."resend_data"
	("id", "agent_id", "level_code", "request_type", "request_from", "request_to")
	VALUES ($1, $2, $3, $4, $5, $6);`

	result, err := db.Exec(query, orderId, agentId, levelCode, dataType, requestFrom, requestTo)
	if err != nil {
		logger.Printf("InsertReSendData() has error to insert resend_data, err  is: %v", err)
	} else {
		if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
			logger.Printf("InsertReSendData() insert resend_data rowsAffectedCount is 0, query is %s", query)
		}
	}
}
