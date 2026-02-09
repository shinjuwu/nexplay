package notification_test

import (
	"backend/internal/notification"
	"backend/pkg/utils"
	"log"
	"testing"
)

func TestGetGameServerGold(t *testing.T) {
	// agentId := 1
	userId := int64(998)
	// thirdAccount := "kinco123"
	// username := "gggg_kinco123"
	money := int64(1000)

	log.Printf("SendCreateAccount int64 money: %v", money)
	log.Printf("SendCreateAccount int money: %v", int(money))
	log.Printf("SendCreateAccount float64 money: %v", float64(int(money)))

	// code, err := notification.SendCreateAccount("", float64(int(money)), agentId, userId, thirdAccount, username)
	// if err != nil {
	// 	log.Printf("SendCreateAccount error: %v", err)
	// }

	// log.Printf("SendCreateAccount code: %v", code)

	userId = 1005

	gold, _, err := notification.GetGameServerGold(userId)
	log.Printf("result: %v", gold)
	log.Printf("GetGameServerGold error: %v", err)

	// money := int64(100)
	// user_id := int64(2)

	// if g, code, err := notification.SendDeposit("up", money, int64(user_id)); err != nil {
	// 	log.Printf("SendDeposit has error, userid: %v, money: %v, code: %v, err: %v", user_id, money, code, err)
	// } else {
	// 	log.Printf("SendDeposit success, userid: %v, money: %v, code: %v, returnGold: %v", user_id, money, code, g)
	// }
}

func TestSendNotifyToFrontend(t *testing.T) {

	address := []byte(`{"host": "127.0.0.1", "port": "8986", "scheme": "http", "api_key": "defaultkey", "ws_conn_path": "/chatservice.ws", "path": "/chatservice.api/v1/"}`)

	connInfo := utils.ToMap(address)

	message := "12345我我我"

	notification.SendNotifyToFrontend("chat", message, notification.ChatNotification_broadcast, connInfo, notification.API_CHAT_MESSAGE_SUBJECT_MESSAGE)
}

func TestSendSetKillDive(t *testing.T) {

	notification.SendSetKillDive(1, 1001, 0, 0.02)
}

func TestGetdefaultkilldiveinfo(t *testing.T) {

	res, code, err := notification.Getdefaultkilldiveinfo()
	if err != nil {
		t.Logf("TestGetdefaultkilldiveinfo() code: %d, err: %v", code, err)
	} else {
		t.Logf("TestGetdefaultkilldiveinfo() res: %v", res)
	}

	// resMap := utils.ToMap([]byte(res))

	// t.Logf("TestGetdefaultkilldiveinfo() resMap: %v", resMap)
}

func TestSendSetPlayerKilldive(t *testing.T) {

	_, _, err := notification.SendSetPlayerKilldive(1006, 0)
	if err != nil {
		t.Logf("TestSendSetPlayerKilldive err: %v", err)
	}
}

func TestSendSetBackendInfo(t *testing.T) {

	_, _, err := notification.SendSetKillAdminInfo("00010003", 0.05, 0)
	if err != nil {
		t.Logf("SendSetKillAdminInfo err: %v", err)
	}
}
