package module

import (
	"chatservice/internal/api"
	"chatservice/pkg/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

func Testfunc(db *sql.DB, it api.ITequila, userInfo map[string]interface{}, payload string) (string, error) {

	msg := api.NewEnvelope(CTWC_MARQUEE, payload, api.ProtocolSubjectMessage, 0)
	it.BroadcastServer(msg)
	return "this is test func", nil
}

func GetUnreadCount(db *sql.DB, it api.ITequila, userInfo map[string]interface{}, payload string) (string, error) {

	var returnJson string
	userId := utils.ToString(userInfo[api.TEQUILA_CTX_RUNTIME_USER_ID], "")
	platform := utils.ToString(userInfo[api.TEQUILA_CTX_RUNTIME_PLATFORM], "")
	if userId == "" || platform == "" {
		return returnJson, errors.New("no user")
	}

	unReadCount := 0

	query := `SELECT unread FROM users WHERE id = $1;`
	err := db.QueryRow(query, userId).Scan(&unReadCount)
	if err != nil {
		returnJson = "Error update users unread count."
	}

	query = `SELECT create_time,content 
	FROM notification WHERE platform=$1 AND create_time > now() - INTERVAL '1 DAY' * $2 
	ORDER BY create_time DESC 
	LIMIT $3;`
	rows, err := db.Query(query, platform, ReadContentDayLimit, unReadCount)
	if err != nil {
		return returnJson, err
	}

	contextMap := make([]map[string]interface{}, 0)
	for rows.Next() {
		var createTime time.Time
		var jsonContent string
		tmpMap := make(map[string]interface{})
		if err = rows.Scan(&createTime, &jsonContent); err != nil {
			rows.Close()
			return returnJson, err
		}
		tmpMap["create_time"] = createTime
		tmpMap["content"] = jsonContent
		contextMap = append(contextMap, tmpMap)
	}

	_ = rows.Close()

	tmp := make(map[string]interface{}, 0)
	tmp["unread_count"] = unReadCount
	tmp["content_map"] = contextMap

	byteJson, err := json.Marshal(tmp)
	if err != nil {
		returnJson = "json marshal failed."
	}

	returnJson = string(byteJson)

	return returnJson, err
}

func ReadNotification(db *sql.DB, it api.ITequila, userInfo map[string]interface{}, payload string) (string, error) {
	var returnJson string

	userId := utils.ToString(userInfo[api.TEQUILA_CTX_RUNTIME_USER_ID], "")
	if userId == "" {
		return returnJson, errors.New("no user")
	}

	query := `UPDATE users SET unread=0 WHERE id = $1;`
	_, err := db.Exec(query, userId)
	if err != nil {
		returnJson = "Error update users unread count."
	}
	return returnJson, err
}

// 取得指定數量訊息
func GetContent(db *sql.DB, it api.ITequila, userInfo map[string]interface{}, payload string) (string, error) {

	var returnJson string
	userId := utils.ToString(userInfo[api.TEQUILA_CTX_RUNTIME_USER_ID], "")
	platform := utils.ToString(userInfo[api.TEQUILA_CTX_RUNTIME_PLATFORM], "")
	if userId == "" || platform == "" {
		return returnJson, errors.New("no user")
	}

	payloadMap := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(payload), &payloadMap)
	if err != nil {
		return returnJson, err
	}

	offset := utils.ToInt(payloadMap["offset"], -1)
	if offset < 0 {
		return returnJson, errors.New("params is illegal")
	}

	query := `SELECT create_time,content 
	FROM notification 
	WHERE platform=$1 AND create_time > now() - INTERVAL '1 DAY' * $2 
	ORDER BY create_time DESC 
	LIMIT 10 OFFSET $3;`
	rows, err := db.Query(query, platform, ReadContentDayLimit, offset)
	if err != nil {
		return returnJson, err
	}

	contextMap := make([]map[string]interface{}, 0)
	for rows.Next() {
		var createTime time.Time
		var jsonContent string
		tmpMap := make(map[string]interface{})
		if err = rows.Scan(&createTime, &jsonContent); err != nil {
			rows.Close()
			return returnJson, err
		}
		tmpMap["create_time"] = createTime
		tmpMap["content"] = jsonContent
		contextMap = append(contextMap, tmpMap)
	}

	_ = rows.Close()

	allContent := 0
	// 只查詢單一平台內30天以內的總訊息數量
	query = `SELECT count(*) 
	FROM notification 
	WHERE platform=$1 AND create_time > now() - INTERVAL '1 DAY' * $2;`
	err = db.QueryRow(query, platform, ReadContentDayLimit).Scan(&allContent)
	if err != nil {
		returnJson = "Error select all context count."
	}

	tmp := make(map[string]interface{}, 0)
	tmp["content_count"] = allContent
	tmp["content_map"] = contextMap

	byteJson, err := json.Marshal(tmp)
	if err != nil {
		returnJson = "json marshal failed."
	}

	returnJson = string(byteJson)

	return returnJson, err
}
