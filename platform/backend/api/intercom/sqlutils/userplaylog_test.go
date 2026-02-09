package sqlutils_test

import (
	"backend/api/intercom/sqlutils"
	"backend/pkg/utils"
	"testing"
)

func TestParseUserPlayLogRobot(t *testing.T) {
	playlog := `{"result":[12,8],"cardtype":"none","cardpoint":19,"playerlog":[{"tax":0,"cards":[[2,44,20],[]],"seatId":1,"user_id":0,"bet_area":[0,300,0,0,0],"cardtype":["none","none"],"de_score":0,"isDouble":[false,false],"is_robot":1,"jp_token":null,"username":"AveryClaire","ya_score":300,"cardpoint":[17,0],"end_score":0,"ownername":"AveryClaire","start_score":0,"valid_score":300,"jp_commission":0.001,"jp_commission _score":0},{"tax":0,"cards":[[13,3,41,48],[]],"seatId":2,"user_id":1530,"bet_area":[0,0,3,0,0],"cardtype":["none","none"],"de_score":0,"isDouble":[false,false],"is_robot":0,"jp_token":null,"username":"ddkktest","ya_score":3,"cardpoint":[18,0],"end_score":7,"ownername":"ddkktest","start_score":10,"valid_score":3,"jp_commission":0.001,"jp_commission _score":0},{"tax":0,"cards":[[35,44,9],[]],"seatId":3,"user_id":0,"bet_area":[0,0,0,75,0],"cardtype":["bust","none"],"de_score":0,"isDouble":[false,false],"is_robot":1,"jp_token":null,"username":"Blithe","ya_score":75,"cardpoint":[26,0],"end_score":0,"ownername":"Blithe","start_score":0,"valid_score":75,"jp_commission":0.001,"jp_commission _score":0}],"taxpercent":0.05}`

	tempMap := utils.ToMap([]byte(playlog))

	playerDataTemp, ok := tempMap["playerlog"].([]interface{})
	if !ok {
		t.Logf("playerData is not array")
	}

	// vals := []interface{}{}
	// for i := 0; i < 100000; i++ {
	for _, val := range playerDataTemp {

		valMap, ok := val.(map[string]interface{})
		if ok {
			isRobotF, ok := valMap["is_robot"].(float64)
			if !ok {
				isRobotF = .0
				// t.Logf("loop process not found is_robot, isRobot=%v", isRobotF)
			}
			isRobot := int(isRobotF)

			// t.Logf("loop process, isRobot=%v", isRobotF)
			// 不新增機器人的個人遊戲紀錄
			if isRobot == 1 {
				continue
			}

			t.Logf("loop process, isRobot=%v, user_id=%v", isRobotF, valMap["user_id"])
		}
	}
	// }
}

func TestParseUserPlayLog(t *testing.T) {
	playlog := `{"basebet": 8, "playerlog": [{"tax": 0, "seatId": 0, "user_id": 0, "de_score": 0, "is_robot": 1, "jp_token": null, "username": "Lawrence", "ya_score": 48, "end_score": 107512.3, "start_score": 107560.3, "valid_score": 48, "jp_commission": 0.001, "jp_commission _score": 0}, {"tax": 2.4, "seatId": 1, "user_id": 0, "de_score": 96, "is_robot": 1, "jp_token": null, "username": "Michael22", "ya_score": 48, "end_score": 125498.6, "start_score": 125453, "valid_score": 48, "jp_commission": 0.001, "jp_commission _score": 0}, {"tax": 3.2, "seatId": 2, "user_id": 0, "de_score": 128, "is_robot": 1, "jp_token": null, "username": "Merlin", "ya_score": 64, "end_score": 150840.8, "start_score": 150780, "valid_score": 64, "jp_commission": 0.001, "jp_commission _score": 0}, {"tax": 0, "seatId": 3, "user_id": 1029, "de_score": 0, "is_robot": 0, "jp_token": null, "username": "action", "ya_score": 64, "end_score": 1599694.0535, "kill_type": 1, "start_score": 1599758.0535, "valid_score": 64, "jp_commission": 0.001, "jp_commission _score": 0.064}], "taxpercent": 0.05, "playerrecord": {"shoot": {"1": 1, "2": 1}, "getshoot": {"0": 1, "3": 1}, "playerRecord": [{"point": -6, "seatId": 0, "playInfo": {"arrayType": [0, 3, 4], "cardsArray": [[21, 25, 26], [50, 24, 11, 40, 15], [16, 43, 5, 32, 7]], "handCardType": 0}, "resultScore": -48}, {"point": 6, "seatId": 1, "playInfo": {"arrayType": [0, 3, 5], "cardsArray": [[10, 12, 39], [46, 33, 20, 1, 28], [14, 18, 19, 22, 13]], "handCardType": 0}, "resultScore": 48}, {"point": 8, "seatId": 2, "playInfo": {"arrayType": [1, 4, 5], "cardsArray": [[51, 38, 47], [41, 29, 30, 31, 45], [2, 3, 4, 8, 9]], "handCardType": 0}, "resultScore": 64}, {"point": -8, "seatId": 3, "playInfo": {"arrayType": [0, 1, 5], "cardsArray": [[6, 48, 0], [49, 23, 42, 17, 44], [27, 34, 35, 36, 37]], "handCardType": 0}, "resultScore": -64}]}}`

	gameId := 2009
	roomType := 3
	deskId := 0
	playlogList, err := sqlutils.ParseUserPlayLog(playlog, gameId, roomType, deskId)
	if err != nil {
		t.Logf("have error: %v", err)
	} else {
		for _, v := range playlogList {
			t.Logf("playLogData is %v", v)
		}
		// t.Logf("playLogData is %v", playlogList)
	}
}
