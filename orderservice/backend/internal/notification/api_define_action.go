package notification

import (
	"backend/pkg/utils"
	"fmt"
)

var (
	ApiNotification_setbackendinfo = "setbackendinfo" // 後台送給Server使用共通的AP(封停玩家、設置玩家身分、設置總代理殺率)

	ApiAction_blockplayer     = "blockplayer"     // 封停玩家
	ApiAction_setplayerstatus = "setplayerstatus" // 設置玩家身分
	ApiAction_setkilladmin    = "setkilladmin"    // 設置總代理殺率
)

/*
設定遊戲用戶封停(smallest unit: user id)
*/
func SendSetBlockPlayer(userIds []int, isBlock bool) (string, int, error) {
	/*
		2.封停玩家
		(1)以陣列的方式送，可一次設多筆，Action為blockplayer
		(2)以目前的規格，IsBlock設true只會把玩家踢離遊戲，設false不會有動作
		(3)單筆結構為
		type PlayerBlock struct {
			UserID  int  `json:"userid"`       玩家ID
			IsBlock bool `json:"isblock"`      現在是否封鎖該玩家
		}
		(4)範例
		{
		    "action":"blockplayer",
		    "data":[{"userid":86569765,"isblock": true},{"userid":81333840,"isblock": false}]
		}
		Server回給後台
		{"code":0,"message":"success"}
	*/

	tmpArrayMap := make([]map[string]interface{}, 0)

	for _, userId := range userIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["userid"] = userId
		tmpMap["isblock"] = isBlock
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_blockplayer, tmpArrayMap)
}

/*
設定遊戲用戶身分(smallest unit: user id)
*/
func SendSetPlayerStatus(userIds []int, statuses []int) (string, int, error) {
	/*
		3.設置玩家身分
		(1)以陣列的方式送，可一次設多筆，Action為setplayerstatus
		(2)單筆結構為
		type PlayerStatus struct {
			UserID int `json:"userid"`   玩家ID
			Status int `json:"status"`   玩家身分
		}
		(3)Status中可填入這四種玩家身分
		一般玩家   : 0
		定點玩家   : 1
		黑名單玩家 : 2

		(4)範例
		{
		    "action":"setplayerstatus",
		    "data":[{"userid":86569765,"status": 2},{"userid":81333840,"status": 1}]
		}
		Server回給後台
		{"code":0,"message":"success"}

	*/
	if len(statuses) != len(userIds) {
		return "", -1, fmt.Errorf("SendSetPlayerStatus() item length not equal")
	}
	tmpArrayMap := make([]map[string]interface{}, 0)

	for idx, userId := range userIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["userid"] = userId
		tmpMap["status"] = statuses[idx]
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_setplayerstatus, tmpArrayMap)
}

/*
設置總代理殺率
*/
func SendSetKillAdminInfo(levelCode string, killRate float64, isEnable int) (string, int, error) {
	/*
		4.設置總代理殺率
		(1)以陣列的方式送，可一次設多筆，Action為setkilladmin
		(2)單筆結構為
		type Killadmininfo struct {
			Level    string  `json:"level"      總代理Level code
			KillRate float64 `json:"killrate"   設置殺率，設0視為無效
			Enable   int     `json:"enable"     是否開啟
		}
		(3)Server回覆時，會以string的方式，同步目前server所有的設定
		(4)範例
		{
		    "action":"setkilladmin",
		    "data":[{"level":"0002","killrate": 0.05,"enable":1}]
		}
		Server回給後台
		{"action":"KillAdmin","code":0,"message":"success","data":"[{\"level\":\"0001\",\"killrate\":0.02,\"enable\":0},{\"level\":\"0002\",\"killrate\":0.05,\"enable\":1}]"}
	*/

	tmpMap := make(map[string]interface{}, 0)
	tmpMap["level"] = levelCode
	tmpMap["killrate"] = killRate
	tmpMap["enable"] = isEnable

	tmpArrayMap := make([]map[string]interface{}, 0)
	tmpArrayMap = append(tmpArrayMap, tmpMap)

	return SendSetBackendInfo(ApiAction_setkilladmin, tmpArrayMap)
}

/*
後台送給Server使用共通的AP
	1. 封停玩家
	2. 設置玩家身分
	3. 設置總代理殺率
*/
func SendSetBackendInfo(action string, data []map[string]interface{}) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setbackendinfo

	/*
		type BackEndInfo struct {
			Action string      `json:"action"`     操作指令
			Data   interface{} `json:"data"`       操作資料
		}
	*/

	tmpMap := make(map[string]interface{}, 0)
	tmpMap["action"] = action
	tmpMap["data"] = data

	paramJsonStr := utils.ToJSON(tmpMap)

	resultStr, err := utils.PostAPI(addr, "application/json", "", paramJsonStr)
	if err != nil {
		return resultStr, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return utils.ToJSON(result), utils.ToInt(code), nil
		} else {
			return utils.ToJSON(result), utils.ToInt(code), fmt.Errorf("SendSetBackendInfo has error, code is: %v, form is: %v", code, tmpMap)
		}
	}

	return utils.ToJSON(result), API_OTHER_ERROR, fmt.Errorf("SendSetBackendInfo has error, response body is: %v", resultStr)
}
