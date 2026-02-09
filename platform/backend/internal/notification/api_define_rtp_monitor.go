package notification

import (
	"backend/pkg/utils"
	"fmt"
	"net/url"
	"time"

	table_model "backend/server/table/model"
)

var (
	Apirtpmonitor_collectorgamelist           = "collectorgamelist"           // 同步遊戲列表
	Apirtpmonitor_collectorcoininout          = "collectorcoininout"          // 玩家上下分資料收集
	Apirtpmonitor_collectorabnormalwinandlose = "collectorabnormalwinandlose" // 異常輸贏資料收集
	Apirtpmonitor_collectorrtpstat            = "collectorrtpstat"            // 統計資料收集
)

// RTP監控 同步遊戲列表
func SendGameListToMonitorService(connInfo map[string]interface{}, patform string, gameA []*table_model.Game) (bool, string) {

	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")

	u := url.URL{Scheme: scheme, Host: domain, Path: path + Apirtpmonitor_collectorgamelist}

	// create data of game list
	type CollectorGameListRequest struct {
		ID   string        `json:"id"`
		Data []interface{} `json:"data"`
	}

	var res CollectorGameListRequest

	res.Data = make([]interface{}, 0)
	res.ID = patform
	for _, v := range gameA {
		if v.Id <= 0 {
			continue
		} else {
			tmp := make(map[string]interface{}, 0)
			tmp["id"] = v.Id
			tmp["code"] = v.Code
			tmp["name"] = v.Name
			tmp["type"] = v.Type

			res.Data = append(res.Data, tmp)
		}
	}

	paramJsonStr := utils.ToJSON(res)

	resultStr, err := utils.PostAPI(u.String(), "application/json", "", paramJsonStr)
	if err != nil {
		return false, resultStr
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return true, ""
		} else {
			return false, fmt.Errorf("SendGameListToMonitorService has error, code is: %v, form is: %v", code, result).Error()
		}
	}

	return false, fmt.Errorf("SendGameListToMonitorService has error, response body is: %v", resultStr).Error()
}

// create data
type TempWalletLedger struct {
	ID         string    `json:"id"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	Kind       int       `json:"kind"`
	Status     int       `json:"status"`
	ChangeSet  string    `json:"changeset"` // You can use map[string]interface{} for JSON
	CreateTime time.Time `json:"create_time"`
}

type TempCollectorCoinInOutRequest struct {
	ID   string              `json:"id"`
	Data []*TempWalletLedger `json:"data"`
}

func NewTempCollectorCoinInOutRequest() *TempCollectorCoinInOutRequest {
	return &TempCollectorCoinInOutRequest{
		ID:   "",
		Data: make([]*TempWalletLedger, 0),
	}
}

// RTP監控 玩家上下分資料收集
func SendCollectorCoinInOutToMonitorService(connInfo map[string]interface{}, data *TempCollectorCoinInOutRequest) (bool, string) {

	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")

	u := url.URL{Scheme: scheme, Host: domain, Path: path + Apirtpmonitor_collectorcoininout}

	paramJsonStr := utils.ToJSON(data)
	resultStr, err := utils.PostAPI(u.String(), "application/json", "", paramJsonStr)
	if err != nil {
		return false, resultStr
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return true, ""
		} else {
			return false, fmt.Errorf("SendCollectorCoinInOutToMonitorService has error, code is: %v, form is: %v", code, result).Error()
		}
	}

	return false, fmt.Errorf("SendCollectorCoinInOutToMonitorService has error, response body is: %v", resultStr).Error()
}

type TmpUserPlayLog struct {
	LogNumber  string    `json:"lognumber"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	GameID     int       `json:"game_id"`
	GameName   string    `json:"game_name"`
	RoomType   int       `json:"room_type"`
	RoomName   string    `json:"room_name"`
	DeScore    float64   `json:"de_score"`
	Bonus      float64   `json:"bonus"`
	BetID      string    `json:"bet_id"`
	BetTime    time.Time `json:"bet_time"`
	CreateTime time.Time `json:"create_time"`
}

type TmpCollectorAbnormalWinAndLoseRequest struct {
	ID   string            `json:"id"`
	Data []*TmpUserPlayLog `json:"data"`
}

func NewTmpCollectorAbnormalWinAndLoseRequest() *TmpCollectorAbnormalWinAndLoseRequest {
	return &TmpCollectorAbnormalWinAndLoseRequest{
		ID:   "",
		Data: make([]*TmpUserPlayLog, 0),
	}
}

// RTP監控 異常輸贏資料收集
func SendCollectorAbnormalWinAndLoseToMonitorService(connInfo map[string]interface{}, data *TmpCollectorAbnormalWinAndLoseRequest) (bool, string) {

	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")

	u := url.URL{Scheme: scheme, Host: domain, Path: path + Apirtpmonitor_collectorabnormalwinandlose}

	paramJsonStr := utils.ToJSON(data)
	resultStr, err := utils.PostAPI(u.String(), "application/json", "", paramJsonStr)
	if err != nil {
		return false, resultStr
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return true, ""
		} else {
			return false, fmt.Errorf("SendCollectorAbnormalWinAndLoseToMonitorService has error, code is: %v, form is: %v", code, result).Error()
		}
	}

	return false, fmt.Errorf("SendCollectorAbnormalWinAndLoseToMonitorService has error, response body is: %v", resultStr).Error()
}

type TmpCollectorRTPStat struct {
	GameId     int       `json:"game_id"`
	GameType   int       `json:"game_type"`
	RoomType   int       `json:"room_type"`
	Ya         float64   `json:"ya"`
	VaildYa    float64   `json:"vaild_ya"`
	De         float64   `json:"de"`
	Tax        float64   `json:"tax"`
	Bonus      float64   `json:"bonus"`
	PlayCount  int       `json:"play_count"`
	UpdateTime time.Time `json:"update_time"`
}

type TmpCollectorRTPStatRequest struct {
	ID   string                 `json:"id"`
	Data []*TmpCollectorRTPStat `json:"data"`
}

func NewTmpCollectorRTPStatRequest() *TmpCollectorRTPStatRequest {
	return &TmpCollectorRTPStatRequest{
		ID:   "",
		Data: make([]*TmpCollectorRTPStat, 0),
	}
}

// RTP監控 RTP 統計資料收集
func SendCollectorRTPStatToMonitorService(connInfo map[string]interface{}, data *TmpCollectorRTPStatRequest) (bool, string) {

	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")

	u := url.URL{Scheme: scheme, Host: domain, Path: path + Apirtpmonitor_collectorrtpstat}

	paramJsonStr := utils.ToJSON(data)
	resultStr, err := utils.PostAPI(u.String(), "application/json", "", paramJsonStr)
	if err != nil {
		return false, resultStr
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return true, ""
		} else {
			return false, fmt.Errorf("SendCollectorRTPStatToMonitorService has error, code is: %v, form is: %v", code, result).Error()
		}
	}

	return false, fmt.Errorf("SendCollectorRTPStatToMonitorService has error, response body is: %v", resultStr).Error()
}
