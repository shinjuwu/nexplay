package notification_test

import (
	"backend/api/v1/model"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"log"
	"testing"
	"time"

	table_model "backend/server/table/model"

	sq "github.com/Masterminds/squirrel"
)

func TestGetGameServerGold(t *testing.T) {
	// agentId := 1
	userId := int64(1010)
	// thirdAccount := "kinco123"
	// username := "gggg_kinco123"
	// money := int64(1000)

	// log.Printf("SendCreateAccount int64 money: %v", money)
	// log.Printf("SendCreateAccount int money: %v", int(money))
	// log.Printf("SendCreateAccount float64 money: %v", float64(int(money)))

	// code, err := notification.SendCreateAccount("", float64(int(money)), agentId, userId, thirdAccount, username)
	// if err != nil {
	// 	log.Printf("SendCreateAccount error: %v", err)
	// }

	// log.Printf("SendCreateAccount code: %v", code)

	userId = 1010

	// gold, _, err := notification.GetGameServerGold(userId)
	// log.Printf("result: %v", gold)
	// log.Printf("GetGameServerGold error: %v", err)

	totalGold, freeGold, _, err := notification.GetGameServerGoldDetail(userId)
	log.Printf("totalGold: %v", totalGold)
	log.Printf("freeGold: %v", freeGold)
	log.Printf("GetGameServerGoldDetail error: %v", err)

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

	notification.SendSetKillDive(1, 1001, 0, 0.02, 0.03, 1)
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

func TestSendNotifyToMaintain(t *testing.T) {
	/*
	   http://172.30.0.152/etetools/maintain_time?
	   action=write&
	   timezone=Asia&
	   starttime=2022-06-15 08:00&
	   endtime=2022-06-15 12:00
	*/
	address := []byte(`{"path": "/etetools/maintain_time", "domain": "172.30.0.152", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}`)

	connInfo := utils.ToMap(address)

	timeZone := "Asia"
	startTime := "2022-06-15 08:00"
	endTime := "2022-06-15 12:00"

	t.Log(notification.SendNotifyToModifyMaintainPageSetting(
		connInfo,
		notification.MaintainNotification_action_write,
		timeZone,
		startTime,
		endTime))

	t.Log(notification.SendNotifyToGetMaintainPageSetting(connInfo, notification.MaintainNotification_action_read))
}

func TestSendNotifyToIMTelegram(t *testing.T) {
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
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}
	ggg := global.NewStorageDatabaseCache(db)
	ggg.InitCacheFromDB()

	s, ok := ggg.SelectOne(definition.STORAGE_KEY_IM_ALERT_TG)
	if ok {
		t.Log(s.Value)
	}

	imNotifys := make([]string, 0)
	imNotifys = append(imNotifys,
		fmt.Sprintf("代理id: %v, 用戶: %v, 上下分太頻繁，目前次數: %v, 限制: %v",
			1,
			2,
			3,
			4))

	address := []byte(`{"path": "/chatservice.api/v1/", "domain": "172.30.0.155:8896", "scheme": "http", "api_key": "defaultkey", "channel": "dev", "ws_conn_path": "/chatservice.ws"}`)

	connInfo := utils.ToMap(address)

	type tg struct {
		Token   string `json:"token"`
		ChatId  int64  `json:"chat_id" mapstructure:"chat_id"`
		MsgData string `json:"msg_data"`
	}

	storage, ok := ggg.SelectOne(definition.STORAGE_KEY_IM_ALERT_TG)
	if !ok {
		return
	}

	var tgData tg

	err = utils.ToStruct([]byte(storage.Value), &tgData)
	if err != nil {
		return
	}

	tgData.MsgData = "本次被自動風控的用戶資訊如下：" + utils.ToJSON(imNotifys)

	notification.SendNotifyToIM(connInfo, definition.IM_TYPE_TELEGRAM, utils.ToJSON(tgData))

	// type tg struct {
	// 	Token   string `json:"token"`
	// 	ChatId  int64  `json:"chat_id"`
	// 	MsgData string `json:"msg_data"`
	// }

	// var tgData tg

	// tgData.Token = "5990213484:AAFgyypd0GKNsig465LMJM_-rGj_Itapquw"
	// tgData.ChatId = -753735474
	// tgData.MsgData = fmt.Sprintf("現在時間為 %v", time.Now())

	// sendJson := utils.ToJSON(tgData)

	// address := []byte(`{"path": "/chatservice.api/v1/", "domain": "127.0.0.1:8896", "scheme": "http", "api_key": "defaultkey", "channel": "dev", "ws_conn_path": "/chatservice.ws"}`)

	// connInfo := utils.ToMap(address)

	// t.Log(notification.SendNotifyToIM(
	// 	connInfo,
	// 	"1",
	// 	sendJson))
}

func TestTssss(t *testing.T) {
	imNotifys := make([]string, 0)
	imNotifys = append(imNotifys, "123")
	imNotifys = append(imNotifys, "456")
	imNotifys = append(imNotifys, "789")
	imNotifys = append(imNotifys, "123156")
	utils.ToJSON(imNotifys)

	t.Log(utils.ToJSON(imNotifys))

}

// 同步平台遊戲列表
func TestRTPMonitorSyncGameList(t *testing.T) {
	// database/sql
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}

	GameCache := global.NewGlobalGameCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "server_info_code", "name", "code", "state",
			"image", "h5_link", "create_time", "update_time", "type",
			"room_number", "table_number", "cal_state").
		From("game").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.Game

		if err := rows.Scan(&temp.Id, &temp.ServerInfoCode, &temp.Name, &temp.Code, &temp.State,
			&temp.Image, &temp.H5Link, &temp.CreateTime, &temp.UpdateTime, &temp.Type,
			&temp.RoomNumber, &temp.TableNumber, &temp.CalState); err != nil {
			return
		}

		GameCache.Add(&temp)
	}

	gameA := GameCache.GetAll()

	connStr := `{"path": "/monitor/collector.api/v1/", "domain": "127.0.0.1:17782", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}`
	utils.ToMap([]byte(connStr))
	connInfo := utils.ToMap([]byte(connStr))

	success, result := notification.SendGameListToMonitorService(connInfo, "dev", gameA)
	if !success {
		t.Errorf("result: %v", result)
	} else {
		t.Logf("result: %v", result)
	}
}

// 玩家上下分資料收集
func TestRTPMonitorSyncCoinInOut(t *testing.T) {

	connStr := `{"path": "/monitor/collector.api/v1/", "domain": "127.0.0.1:17782", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}`
	utils.ToMap([]byte(connStr))
	connInfo := utils.ToMap([]byte(connStr))

	d := notification.NewTempCollectorCoinInOutRequest()
	d.ID = "dev"

	tmp := new(notification.TempWalletLedger)
	tmp.ID = "7533967"
	tmp.AgentID = 3
	tmp.AgentName = "test2"
	tmp.ChangeSet = "{}"
	tmp.CreateTime = time.Now()
	tmp.Kind = 1
	tmp.Status = 1
	tmp.UserID = 1001
	tmp.Username = "kinco"

	d.Data = append(d.Data, tmp)

	success, result := notification.SendCollectorCoinInOutToMonitorService(connInfo, d)
	if !success {
		t.Errorf("result: %v", result)
	} else {
		t.Logf("result: %v", result)
	}
}

// 異常輸贏資料收集
func TestRTPMonitorSyncAbnormalWinAndLose(t *testing.T) {

	connStr := `{"path": "/monitor/collector.api/v1/", "domain": "127.0.0.1:17782", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}`
	utils.ToMap([]byte(connStr))
	connInfo := utils.ToMap([]byte(connStr))

	res := notification.NewTmpCollectorAbnormalWinAndLoseRequest()

	success, result := notification.SendCollectorAbnormalWinAndLoseToMonitorService(connInfo, res)
	if !success {
		t.Errorf("result: %v", result)
	} else {
		t.Logf("result: %v", result)
	}
}

// RTP 統計資料收集
func TestRTPMonitorSyncRTPStat(t *testing.T) {

	connStr := `{"path": "/monitor/collector.api/v1/", "domain": "127.0.0.1:17782", "scheme": "http", "api_key": "", "channel": "", "ws_conn_path": ""}`
	utils.ToMap([]byte(connStr))
	connInfo := utils.ToMap([]byte(connStr))

	res := notification.NewTmpCollectorRTPStatRequest()

	success, result := notification.SendCollectorRTPStatToMonitorService(connInfo, res)
	if !success {
		t.Errorf("result: %v", result)
	} else {
		t.Logf("result: %v", result)
	}
}

// ping 遊戲server，判斷是否存活
func TestPingGameServer(t *testing.T) {

	connStr := `{"notification": "http://172.30.0.154:9642/"}`
	utils.ToMap([]byte(connStr))
	connInfo := utils.ToMap([]byte(connStr))

	success, _ := notification.SendNotifyToPing(connInfo)
	if !success {
		t.Errorf("TestPingGameServer failed")
	} else {
		t.Logf("TestPingGameServer success")
	}

}

func TestGetRealtimeGameRatio(t *testing.T) {

	apiBody, code, err := notification.GetRealtimeGameRatio(1001)
	if err != nil {
		t.Errorf("TestGetRealtimeGameRatio failed, body: %v, code: %v, err: %v", apiBody, code, err)
	} else {
		t.Logf("TestGetRealtimeGameRatio success")
	}
}

func TestGetGameServerUserIsNewbie(t *testing.T) {

	apiBody, limit, code, err := notification.GetGameServerUserIsNewbie(1001)
	if err != nil {
		t.Errorf("TestGetGameServerUserIsNewbie failed, body: %v, code: %v, err: %v", apiBody, code, err)
	} else {
		t.Logf("TestGetGameServerUserIsNewbie success")
	}

	resp := new(model.GetGameUserPlayCountDataResponse)
	// resp.UserId = 0
	resp.TotalNewbieLimit = limit
	if succ := resp.DataConvert(apiBody); !succ {
		// c.Logger().Info("GetGameUserPlayCountData DataConvert failed, userId=%d, playCountData=%v", req.UserId, playCountData)
		t.Error(apiBody, limit, code)
	}

	t.Log(apiBody, limit, code)
}
