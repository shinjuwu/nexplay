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
	ApiAction_softblockplayer = "softblockplayer" // 軟封鎖玩家
	ApiAction_japotinfo       = "japotinfo"       // 設置代理jackpot資訊
	ApiAction_japotaddcoin    = "japotaddcoin"    // 建立jackpot token
	ApiAction_gamesetting     = "gamesetting"     // 設置遊戲基礎
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
設置玩家軟封鎖
*/
func SendSoftBlockPlayer(userIds []int, isSoftBlocks []int) (string, int, error) {
	/*
		5.設置玩家軟封鎖
		(1)以陣列的方式送，可一次設多筆，Action為softblockplayer
		(2)單筆結構為
		type PlayerSoftBlock struct {
			UserID      int `json:"userid"`
			IsSoftBlock int `json:"issoftblock"`
		}
		(3)Server回覆時，會以string的方式，同步目前server所有的設定
		(4)範例
		{
		    "action":"softblockplayer",
		    "data":[{"userid":1001,"softblock":1}]
		}
		Server回給後台
		{"action":"softblock","code":0,"message":"success","data":"[{\"userid\":1001,\"softblock\":1},{\"userid\":1002,\"softblock\":1}]"}
	*/
	if len(userIds) != len(isSoftBlocks) {
		return "", -1, fmt.Errorf("SendSoftBlockPlayer() item length not equal")
	}
	tmpArrayMap := make([]map[string]interface{}, 0)

	for idx, userId := range userIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["userid"] = userId
		tmpMap["softblock"] = isSoftBlocks[idx]
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_softblockplayer, tmpArrayMap)
}

/*
設置代理jackpot資訊
*/
func SendJapotInfo(agentIds []int, starts, ends []int64) (string, int, error) {
	/*
		6.設置代理jackpot資訊
		(1)以陣列的方式送，可一次設多筆，Action為japotinfo
		(2)單筆結構為
		type Japotinfo struct {
			AgentId int   `json:"agentId" xorm:"pk notnull default(0)"`
			Start   int64 `json:"start" xorm:"notnull default(0)"`
			End     int64 `json:"end" xorm:"notnull default(0)"`
		}
		(3)Server回覆時，會以string的方式，同步目前server所有的設定
		(4)範例
		{
		    "action":"japotinfo",
		    "data":[{"agentId":3,"start":0,"end":0}]
		}
		Server回給後台
		{"action":"japotinfo","code":0,"message":"success","data":"[{\"agentId\":3,\"start\":0,\"end\":0}]"}
	*/
	if len(agentIds) <= 0 || len(agentIds) != len(starts) || len(agentIds) != len(ends) {
		return "", -1, fmt.Errorf("SendJapotInfo() item length is error")
	}

	tmpArrayMap := make([]map[string]interface{}, 0)

	for idx, userId := range agentIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["agentId"] = userId
		tmpMap["start"] = starts[idx]
		tmpMap["end"] = ends[idx]
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_japotinfo, tmpArrayMap)
}

/*
建立jackpot token
*/
func SendJapotAddCoin(userIds []string, coinTokens []string, coinBets []float64, createTimes []int64) (string, int, error) {
	/*
		7.建立jackpot token
		(1)以陣列的方式送，可一次設多筆，Action為japotaddcoin
		(2)單筆結構為
		type JapotAddCoin struct {
			UserId     string  `json:"userId"`
			CoinId     string  `json:"cointoken"`
			CoinBet    float64 `json:"coinbet"`
			CreateTime int64   `json:"createtime"`
		}
		(3)Server回覆時，會以string的方式，同步目前server所有的設定
		(4)範例
		{
		    "action":"japotaddcoin",
		    "data":[{"userId":"1001","cointoken":"520230920023349229693462221","coinbet":100.5,"createtime":1695176820968}]
		}
		Server回給後台
		{"action":"japotaddcoin","data":[{"coinbet":100.5,"cointoken":"520230920022700820968459068","createtime":1695176820968,"userId":"1001"}]}"
	*/
	if len(userIds) <= 0 || len(userIds) != len(coinTokens) || len(userIds) != len(coinBets) || len(userIds) != len(createTimes) {
		return "", -1, fmt.Errorf("SendJapotAddCoin() item length is error")
	}

	tmpArrayMap := make([]map[string]interface{}, 0)

	for idx, userId := range userIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["userId"] = userId
		tmpMap["cointoken"] = coinTokens[idx]
		tmpMap["coinbet"] = coinBets[idx]
		tmpMap["createtime"] = createTimes[idx]
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_japotaddcoin, tmpArrayMap)
}

/*
建立jackpot token
*/
func SendGameSetting(gameIds []int, matchGameRTPs []float64, matchGameKillRates []float64, matchGames []int, normalMatchGameRtps []float64,
	normalMatchGameKillRates []float64, lowBoundRtps []float64, limitOdds []float64, limitMoneys []float64) (string, int, error) {
	/*
		7.建立jackpot token
		(1)以陣列的方式送，可一次設多筆，Action為gamesetting
		(2)單筆結構為
		type Gamesetting struct {
			GameId                  int     `json:"gameId"`
			MatchGameRtp            float64 `json:"MatchGameRTP"`
			MatchGameKillRate       float64 `json:"MatchGameKillRate"`
			MatchGames              int     `json:"MatchGames"`
			NormalMatchGameRtp      float64 `json:"NormalMatchGameRTP"`
			NormalMatchGameKillRate float64 `json:"NormalMatchGameKillRate"`
			LowBoundRtp             float64 `json:"LowBoundRTP"`
			LimitOdds               float64 `json:"LimitOdds"`
			LimitMoney              float64 `json:"LimitMoney"`
		}
		(3)Server回覆時，會以string的方式，同步目前server所有的設定
		(4)範例
		{
		    "action":"gamesetting",
		    "data":[{"gameId":1008,"MatchGameRTP":0,"MatchGameKillRate":0,"MatchGames":0,"NormalMatchGameRTP":0.95,"NormalMatchGameKillRate":0,"LowBoundRTP":0.8,"LimitOdds":0,"LimitMoney":0}]
		}
		Server回給後台
		{"action":"gamesetting","data":[{"gameId":1008,"MatchGameRTP":0,"MatchGameKillRate":0,"MatchGames":0,"NormalMatchGameRTP":0.95,"NormalMatchGameKillRate":0,"LowBoundRTP":0.8,"LimitOdds":0,"LimitMoney":0}]}"
	*/
	if len(gameIds) <= 0 ||
		len(gameIds) != len(matchGameRTPs) ||
		len(gameIds) != len(matchGameKillRates) ||
		len(gameIds) != len(matchGames) ||
		len(gameIds) != len(normalMatchGameRtps) ||
		len(gameIds) != len(normalMatchGameKillRates) ||
		len(gameIds) != len(lowBoundRtps) ||
		len(gameIds) != len(limitOdds) ||
		len(gameIds) != len(limitMoneys) {
		return "", -1, fmt.Errorf("SendGameSetting() item length is error")
	}

	tmpArrayMap := make([]map[string]interface{}, 0)

	for idx, gameId := range gameIds {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["gameId"] = gameId
		tmpMap["MatchGameRTP"] = matchGameRTPs[idx]
		tmpMap["MatchGameKillRate"] = matchGameKillRates[idx]
		tmpMap["MatchGames"] = matchGames[idx]
		tmpMap["NormalMatchGameRTP"] = normalMatchGameRtps[idx]
		tmpMap["NormalMatchGameKillRate"] = normalMatchGameKillRates[idx]
		tmpMap["LowBoundRTP"] = lowBoundRtps[idx]
		tmpMap["LimitOdds"] = limitOdds[idx]
		tmpMap["LimitMoney"] = limitMoneys[idx]
		tmpArrayMap = append(tmpArrayMap, tmpMap)
	}

	return SendSetBackendInfo(ApiAction_gamesetting, tmpArrayMap)
}

/*
後台送給Server使用共通的AP
	1. 封停玩家
	2. 設置玩家身分
	3. 設置總代理殺率
	4. 軟封鎖玩家
	5. 代理jackpot設定
	6. 建立jackpot token
	7. 設置遊戲基礎
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
