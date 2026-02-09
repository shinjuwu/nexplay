package notification

import (
	"backend/pkg/encrypt/base64url"
	"backend/pkg/utils"
	"encoding/json"
	"log"
	"net/url"
)

const (
	ApiNotification_im = "im"
)

// 設定維護頁顯示資訊
func SendNotifyToIM(connInfo map[string]interface{}, im, jsonData string) (bool, string) {
	/*
		127.0.0.1:8896/chatservice.api/v1/im?
		im=1&
		data=eyJ0b2tlbiI6IjU5OTAyMTM0ODQ6QUFGZ3l5cGQwR0tOc2lnNDY1TE1KTV8tckdqX0l0YXBxdXciLCJjaGF0X2lkIjotNzUzNzM1NDc0LCJtc2dfZGF0YSI6IuePvuWcqOaZgumWk-eCuiAyMDIzLTA2LTE5IDExOjM2OjQyLjI3ODYwMDkgKzA4MDAgQ1NUIG09KzAuMDA3NDUyMjAxIn0

		type tg struct {
			Token   string `json:"token"`
			ChatId  int64  `json:"chat_id"`
			MsgData string `json:"msg_data"`
		}
	*/
	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")
	apiKey := utils.ToString(connInfo["api_key"], "")

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("im", im)
	rawApiQuery.Add("data", base64url.Encode([]byte(jsonData)))

	u := url.URL{Scheme: scheme, Host: domain, Path: path + "im", RawQuery: rawApiQuery.Encode()}

	body, err := utils.GetBasicAuthAPI(u.String(), apiKey, "")
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
