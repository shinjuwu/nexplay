package notification

import (
	"backend/pkg/utils"
	"encoding/json"
	"log"
	"net/url"
)

const (
	MaintainNotification_action_write = "write" // 修改維護資料頁面顯示資訊
	MaintainNotification_action_read  = "read"  // 讀取維護資料頁面顯示資訊
)

// 取得維護頁顯示資訊
func SendNotifyToGetMaintainPageSetting(connInfo map[string]interface{}, action string) (bool, string) {
	/*
	   http://172.30.0.152/etetools/maintain_time?
	   action=write&
	   timezone=Asia&
	   starttime=2022-06-15 08:00&
	   endtime=2022-06-15 12:00
	*/
	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")
	// apiKey := utils.ToString(connInfo["api_key"], "")

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("action", action)

	u := url.URL{Scheme: scheme, Host: domain, Path: path, RawQuery: rawApiQuery.Encode()}

	body, err := utils.GetBasicAuthAPI(u.String(), "", "")
	if err != nil {
		log.Println(err)
		return false, err.Error()
	}

	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return false, err.Error()
	}

	return true, string(body)
}

// 設定維護頁顯示資訊
func SendNotifyToModifyMaintainPageSetting(connInfo map[string]interface{}, action, timeZone, stratTime, endTime string) (bool, string) {
	/*
	   http://172.30.0.152/etetools/maintain_time?
	   action=write&
	   timezone=Asia&
	   starttime=2022-06-15 08:00&
	   endtime=2022-06-15 12:00
	*/
	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")
	apiKey := utils.ToString(connInfo["api_key"], "")

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("action", action)
	rawApiQuery.Add("timezone", timeZone)
	rawApiQuery.Add("starttime", stratTime)
	rawApiQuery.Add("endtime", endTime)

	u := url.URL{Scheme: scheme, Host: domain, Path: path, RawQuery: rawApiQuery.Encode()}

	body, err := utils.GetBasicAuthAPI(u.String(), apiKey, "")
	if err != nil {
		log.Println(err)
		return false, err.Error()
	}

	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return false, err.Error()
	}

	return true, string(body)
}

// ping game server
func SendNotifyToPing(connInfo map[string]interface{}) (bool, string) {
	/*
	   http://172.30.0.154/ping
	*/
	pingApiPath := utils.ToString(connInfo["notification"], "")

	body, err := utils.GetBasicAuthAPI(pingApiPath+"ping", "", "")
	if err != nil {
		log.Println(err)
		return false, err.Error()
	}

	bodyMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return false, err.Error()
	}

	return true, string(body)
}
