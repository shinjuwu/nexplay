package notification

import (
	"backend/api/game/model"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

var (
	//default
	ApiNotification_baseAddr = "http://10.1.0.12:9642/"
)

const (
	ApiNotification_deposit                = "deposit"                // 上下分通知
	ApiNotification_depositupCreateAccount = "depositupCreateAccount" // 上分創帳號通知
	ApiNotification_querygold              = "querygold"              // 取指定user目前剩餘分數
	ApiNotification_querypluralgold        = "querypluralgold"        // 取指定多位user目前剩餘分數
	ApiNotification_setgamelist            = "setgamelist"            // 設定遊戲開放列表
	ApiNotification_setgameinfo            = "setgameinfo"            // 設定代理開放遊戲列表
	ApiNotification_setlobbyinfo           = "setlobbyinfo"           // 設定大廳遊戲列表
	ApiNotification_marquee                = "marquee"                // 更新跑馬燈設定
	ApiNotification_setGameServerState     = "setgameserverstate"     // 設定遊戲全局維護
	ApiNotification_getdefaultkilldiveinfo = "getdefaultkilldiveinfo" // 拿預設的殺放參數
	ApiNotification_setkilldiveinfo        = "setkilldiveinfo"        // 設定殺放
	ApiNotification_setplayerkilldive      = "setplayerkilldive"      // 設定遊戲用戶殺放
	ApiNotification_getdefaultgamesetting  = "getdefaultgamesetting"  // 拿預設的遊戲基礎設定參數
	ApiNotification_getrealtimegameratio   = "getrealtimegameratio"   // 取得即時遊戲殺率
	ApiNotification_getgameusernewbie      = "getgameusernewbie"      // 取指定user目前遊戲局數判斷是否為新手期
	ApiNotification_getplinkoballmaxodds   = "getplinkoballmaxodds"   // 取得Plinko球倍率上限
	ApiNotification_setplinkoballmaxodds   = "setplinkoballmaxodds"   // 設定Plinko球倍率上限
	/*
		本地端發生錯誤通用碼

		非遊戲端回傳之錯誤碼都一律使用-1
	*/
	API_OTHER_ERROR                = -1
	API_ERROR_MARQUEE_DO_NOT_EXIST = 200108 // game server 跑馬燈資料不存在錯誤代碼

	API_CHAT_MESSAGE_SUBJECT_MESSAGE      = "message"
	API_CHAT_MESSAGE_SUBJECT_ANNOUNCEMENT = "announcement"
)

const (
	ChatNotification_broadcast = "broadcast"
)

func SetNotificationAddress(addr string) {
	// 去除空格
	trimAddr := strings.ReplaceAll(addr, " ", "")
	// 如果沒有斜線要補上斜線
	l := len(trimAddr)
	if addr[l-1] != '/' {
		addr += addr + "/"
	}
	ApiNotification_baseAddr = addr
}

func TransformApiCodeToModelErrorCode(code int) int {
	switch code {
	case 0:
		return model.Response_Success
	case 200100:
		fallthrough
	case 200103:
		return model.Response_ParseParam_Error
	case 200101:
		return model.Response_GameServeDepositFailed_Error
	case 200102:
		return model.Response_AccountExist_Error
	default:
		return model.Response_GameServer_Error
	}
}

// 上下分通知
func SendDeposit(method string, gold float64, userid int64) (float64, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_deposit

	form := url.Values{}
	form.Add("method", method)
	form.Add("gold", strconv.FormatFloat(gold, 'f', -1, 64))
	form.Add("userid", strconv.FormatInt(userid, 10))

	resultStr, err := utils.PostAPI(addr, "application/x-www-form-urlencoded", "", form.Encode())
	if err != nil {
		return .0, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	finalGold := float64(0)

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			// var goldFlag, lockGoldFlag bool
			if gold, ok := result["gold"].(float64); ok {
				// return gold, utils.ToInt(code), nil
				finalGold += gold
				// goldFlag = ok
			}
			if lockgold, ok := result["lockgold"].(float64); ok {
				// return gold, utils.ToInt(code), nil
				finalGold += lockgold
				// lockGoldFlag = ok
			}
			return finalGold, utils.ToInt(code), nil
		} else {
			return .0, utils.ToInt(code), fmt.Errorf("SendDeposit has error, code is: %v", code)
		}
	}

	return .0, API_OTHER_ERROR, fmt.Errorf("SendDeposit has error, response body is: %v", resultStr)
}

// 創帳號通知
func SendCreateAccount(method string, gold float64, agentId int, userId int64, username, trans_username string) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_depositupCreateAccount

	form := url.Values{}
	strCoin := fmt.Sprint(gold)
	form.Add("userid", strconv.FormatInt(userId, 10))
	form.Add("agentid", strconv.Itoa(agentId))
	form.Add("username", username)
	form.Add("trans_username", trans_username)
	form.Add("coin", strCoin)

	resultStr, err := utils.PostAPI(addr, "application/x-www-form-urlencoded", "", form.Encode())
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return utils.ToInt(code), nil
		} else {
			return utils.ToInt(code), fmt.Errorf("SendCreateAccount has error, code is: %v, form is: %v", code, form)
		}
	}

	return API_OTHER_ERROR, fmt.Errorf("SendCreateAccount has error, response body is: %v", resultStr)
}

// 跑馬燈修改通知
func SendMarquee(param map[string]interface{}) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_marquee

	resultStr, err := utils.PostAPI(addr, "application/json", "", utils.ToJSON(param))
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return utils.ToInt(code), nil
		} else {
			return utils.ToInt(code), fmt.Errorf("SendMarquee has error, code is: %v, param is: %v", code, param)
		}
	}

	return API_OTHER_ERROR, fmt.Errorf("SendMarquee has error, response body is: %v", resultStr)
}

// 取指定user目前剩餘分數
func GetGameServerGold(userid int64) (float64, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_querygold

	resultStr, err := utils.GetAPI(addr, fmt.Sprintf("userid=%v", userid), "")
	if err != nil {
		fmt.Printf("GetGameServerGold() response: %s, err: %v", resultStr, err)
		return .0, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	finalGold := float64(0)

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			// var goldFlag, lockGoldFlag bool
			if gold, ok := result["gold"].(float64); ok {
				// return gold, utils.ToInt(code), nil
				finalGold += gold
				// goldFlag = ok
			}
			if lockgold, ok := result["lockgold"].(float64); ok {
				// return gold, utils.ToInt(code), nil
				finalGold += lockgold
				// lockGoldFlag = ok
			}
			return finalGold, utils.ToInt(code), nil
		} else {
			return .0, utils.ToInt(code), fmt.Errorf("GetGameServerGold has error, code is: %v", code)
		}
	}

	return .0, API_OTHER_ERROR, fmt.Errorf("GetGameServerGold has error, response body is: %v", resultStr)
}

// 取指定user目前剩餘分數細項
func GetGameServerGoldDetail(userid int64) (float64, float64, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_querygold

	resultStr, err := utils.GetAPI(addr, fmt.Sprintf("userid=%v", userid), "")
	if err != nil {
		fmt.Printf("GetGameServerGold() response: %s, err: %v", resultStr, err)
		return .0, .0, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	totlalGold := float64(0)
	freeGold := float64(0)
	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			// var goldFlag, lockGoldFlag bool
			if freeGoldTmp, ok := result["gold"].(float64); ok {
				freeGold = freeGoldTmp
				totlalGold += freeGold
			}
			if lockgold, ok := result["lockgold"].(float64); ok {
				totlalGold += lockgold
			}
			return totlalGold, freeGold, utils.ToInt(code), nil
		} else {
			return totlalGold, freeGold, utils.ToInt(code), fmt.Errorf("GetGameServerGold has error, code is: %v", code)
		}
	}

	return totlalGold, freeGold, API_OTHER_ERROR, fmt.Errorf("GetGameServerGold has error, response body is: %v", resultStr)
}

// 取指定多位user目前剩餘分數
func GetGameServerGolds(users []int) (map[int]map[string]interface{}, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_querypluralgold

	userid := strings.Trim(strings.Replace(fmt.Sprint(users), " ", ",", -1), "[]")

	resultStr, err := utils.GetAPI(addr, fmt.Sprintf("userid=%s", userid), "")
	if err != nil {
		return nil, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	code, ok := result["code"].(float64)
	if ok {
		if code == 0 {
			resp := make(map[int]map[string]interface{})
			for _, obj := range result["data"].([]interface{}) {
				mObj := obj.(map[string]interface{})
				userid := utils.ToInt(mObj["userid"])
				resp[userid] = mObj
			}
			if len(resp) > 0 {
				return resp, utils.ToInt(code), nil
			}
		}
		return nil, utils.ToInt(code), fmt.Errorf("GetGameServerGolds trans data has error, api code is %v, response body is: %v", code, resultStr)
	}
	return nil, API_OTHER_ERROR, fmt.Errorf("GetGameServerGolds has error, response body is: %v", resultStr)
}

// 設定遊戲列表(root)
// {"gamelist":[{"Id":1,"GameId":1001,"GameCode":"baccarat","Status":0},{"Id":3,"GameId":1002,"GameCode":"foo1234","Status":0}]}
func SendSetGameList(gamelists []map[string]interface{}) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setgamelist

	tmpMap := make(map[string]interface{}, 0)

	tmpMap["gamelist"] = gamelists

	resultStr, err := utils.PostAPI(addr, "application/json", "", utils.ToJSON(tmpMap))
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		return utils.ToInt(code), nil
	}

	return API_OTHER_ERROR, fmt.Errorf("SendSetGameList has error, response body is: %v", resultStr)
}

// 設定代理開放遊戲列表
// {"gameinfo":[{"Id":3,"AgentId":11,"GameId":1001,"GameCode":"baccarat","Status":1},{"Id":2,"AgentId":0,"GameId":1002,"GameCode":"foo1234","Status":1}]}
func SendSetGameInfo(gameInfos []map[string]interface{}) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setgameinfo

	tmpMap := make(map[string]interface{}, 0)

	tmpMap["gameinfo"] = gameInfos

	resultStr, err := utils.PostAPI(addr, "application/json", "", utils.ToJSON(tmpMap))
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		return utils.ToInt(code), nil
	}

	return API_OTHER_ERROR, fmt.Errorf("SendSetGameInfo has error, response body is: %v", resultStr)
}

// 設定大廳遊戲列表
// {"lobbyinfo":[{"Id":17,"AgentId":0,"GameId":1002,"TableId":100201,"GameCode":"qweert","GameType":0,"RoomType":0,"Status":1},{"Id":18,"AgentId":0,"GameId":1002,"TableId":100201,"GameCode":"sdaf","GameType":0,"RoomType":0,"Status":0}]}
func SendSetLobbyInfo(lobbyInfos []map[string]interface{}) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setlobbyinfo

	tmpMap := make(map[string]interface{}, 0)

	tmpMap["lobbyinfo"] = lobbyInfos

	resultStr, err := utils.PostAPI(addr, "application/json", "", utils.ToJSON(tmpMap))
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		return utils.ToInt(code), nil
	}

	return API_OTHER_ERROR, fmt.Errorf("SendSetLobbyInfo has error, response body is: %v", resultStr)
}

// 設定即時通知後台前端新訊息發佈(即時通訊公告)
func SendNotifyToFrontend(platform string, msg string, method string, connInfo map[string]interface{}, subject string) bool {

	scheme := utils.ToString(connInfo["scheme"], "")
	domain := utils.ToString(connInfo["domain"], "")
	path := utils.ToString(connInfo["path"], "")
	apiKey := utils.ToString(connInfo["api_key"], "")

	rawApiQuery := make(url.Values)
	rawApiQuery.Add("platform", platform) //channel
	rawApiQuery.Add("data", msg)
	rawApiQuery.Add("subject", subject)

	u := url.URL{Scheme: scheme, Host: domain, Path: path + method, RawQuery: rawApiQuery.Encode()}

	body, err := utils.PostBasicAuthAPI(u.String(), apiKey, "", "")
	if err != nil {
		log.Println(err)
		return false
	}

	bodyMap := make(map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(body), &bodyMap)

	return utils.ToInt(bodyMap["code"], -1) == 0
}

// 設定遊戲全局開關
// state: 狀態 (1: 開啟, 2: 關閉; 預設: 關閉)
func SetGameServerState(state int) (int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setGameServerState

	tmpMap := make(map[string]interface{}, 0)

	tmpMap["state"] = state

	resultStr, err := utils.PostAPI(addr, "application/json", "", utils.ToJSON(tmpMap))
	if err != nil {
		return API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		return utils.ToInt(code), nil
	}

	return API_OTHER_ERROR, fmt.Errorf("SetGameServerState has error, response body is: %v", resultStr)
}

/*
取得遊戲預設殺放
不過濾無效資料，完整保留原始資料
*/
func Getdefaultkilldiveinfo() ([]map[string]interface{}, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_getdefaultkilldiveinfo

	// game server struct
	// type Killdiveinfo struct {
	// 	AgentId     int     `json:"AgentId,omitempty" xorm:"pk notnull default(0)"`
	// 	GameId      int     `json:"GameId" xorm:"notnull default(0)"`
	// 	RoomId      int     `json:"RoomId" xorm:"pk notnull default(0)"`
	// 	KillRate    float64 `json:"Killrate" xorm:"notnull default(0)"`
	// 	NewKillRate float64 `json:"Newkillrate" xorm:"notnull default(0)"`
	// 	ActiveNum   int     `json:"Activenum" xorm:"notnull default(0)"`
	// }

	resultStr, err := utils.GetAPI(addr, "application/x-www-form-urlencoded", "")
	if err != nil {
		return nil, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	code, ok := result["code"].(float64)
	if ok {
		if code == 0 {
			val, ok := result["data"].(string)
			if ok {
				resp := utils.ToArrayMap([]byte(val))
				if resp != nil {
					return resp, utils.ToInt(code), nil
				}
			}
		}
		return nil, utils.ToInt(code), fmt.Errorf("Getdefaultkilldiveinfo trans data has error, api code is %v, response body is: %v", code, resultStr)
	}
	return nil, API_OTHER_ERROR, fmt.Errorf("Getdefaultkilldiveinfo has error, response body is: %v", resultStr)
}

/*
設定遊戲殺放(smallest unit: roomType)
*/
func SendSetKillDive(agnetId, gameId, roomType int64, killRate, newKillRate float64, activeNum int) (string, int, error) {
	// game server struct
	// // KillDiveInfo represents the information of dive and kill rate
	// type Killdiveinfo struct {
	// 	AgentId     int     `json:"AgentId"`
	// 	GameId      int     `json:"GameId"`
	// 	RoomId      int     `json:"RoomId"`
	// 	KillRate    float64 `json:"Killrate"`
	// 	NewKillRate float64 `json:"Newkillrate"`
	// 	ActiveNum   int     `json:"Activenum"`
	// }

	roomId, err := strconv.Atoi(fmt.Sprintf("%d%d", gameId, roomType))
	if err != nil {
		return "", API_OTHER_ERROR, err
	}

	tmpMap := make(map[string]interface{}, 0)
	tmpMap["AgentId"] = agnetId
	tmpMap["GameId"] = gameId
	tmpMap["RoomId"] = roomId
	tmpMap["KillRate"] = killRate //0 ~ 0.1
	tmpMap["Newkillrate"] = newKillRate
	tmpMap["Activenum"] = activeNum

	tmpArrayMap := make([]map[string]interface{}, 0)
	tmpArrayMap = append(tmpArrayMap, tmpMap)

	return SendSetKillDives(tmpArrayMap)
}

/*
設定遊戲殺放
*/
func SendSetKillDives(killDiveInfo []map[string]interface{}) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setkilldiveinfo

	// game server struct
	// KillDiveInfo represents the information of dive and kill rate
	// type Killdiveinfo struct {
	// 	AgentId     int     `json:"AgentId"`
	// 	GameId      int     `json:"GameId"`
	// 	RoomId      int     `json:"RoomId"`
	// 	KillRate    float64 `json:"Killrate"`
	// 	NewKillRate float64 `json:"Newkillrate"`
	// 	ActiveNum   int     `json:"Activenum"`
	// }

	resMap := make(map[string]interface{}, 0)
	resMap["killdiveinfo"] = killDiveInfo

	paramJsonStr := utils.ToJSON(resMap)

	resultStr, err := utils.PostAPI(addr, "application/json", "", paramJsonStr)
	if err != nil {
		return resultStr, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return utils.ToJSON(result), utils.ToInt(code), nil
		} else {
			return utils.ToJSON(result), utils.ToInt(code), fmt.Errorf("SendSetKillDive has error, code is: %v, form is: %v", code, resMap)
		}
	}

	return utils.ToJSON(result), API_OTHER_ERROR, fmt.Errorf("SendSetKillDive has error, response body is: %v", resultStr)
}

/*
設定遊戲用戶殺放(smallest unit: user id)
*/
func SendSetPlayerKilldive(userid int64, status int) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setplayerkilldive

	/*
		(2)API資料格式，以陣列的方式送，以達到可一次設多筆：
		{
		    "playerkilldiveinfo":
		  [
		    {
		        userid int 玩家ID
		        status int 玩家身分，詳情請看(2)
		    },...
		  ]
		}

		(2)Status中可填入這四種玩家身分
		一般玩家   : 0
		定點玩家   : 1
		黑名單玩家 : 2
	*/

	tmpMap := make(map[string]interface{}, 0)
	tmpMap["userid"] = userid
	tmpMap["status"] = status

	tmpArrayMap := make([]map[string]interface{}, 0)
	tmpArrayMap = append(tmpArrayMap, tmpMap)

	resMap := make(map[string]interface{}, 0)
	resMap["playerkilldiveinfo"] = tmpArrayMap

	paramJsonStr := utils.ToJSON(resMap)

	resultStr, err := utils.PostAPI(addr, "application/json", "", paramJsonStr)
	if err != nil {
		return resultStr, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			return utils.ToJSON(result), utils.ToInt(code), nil
		} else {
			return utils.ToJSON(result), utils.ToInt(code), fmt.Errorf("SendSetplayerkilldive has error, code is: %v, form is: %v", code, resMap)
		}
	}

	return utils.ToJSON(result), API_OTHER_ERROR, fmt.Errorf("SendSetplayerkilldive has error, response body is: %v", resultStr)
}

/*
取得遊戲預設殺放
不過濾無效資料，完整保留原始資料
*/
func Getdefaultgamesetting() (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_getdefaultgamesetting

	// game server struct
	// type Gamesetting struct {
	// 	GameId                  int     `json:"gameId" xorm:"pk notnull default(0)"`
	// 	MatchGameRtp            float64 `json:"MatchGameRTP" xorm:"notnull default(0)"`
	// 	MatchGameKillRate       float64 `json:"MatchGameKillRate" xorm:"notnull default(0)"`
	// 	MatchGames              int     `json:"MatchGames" xorm:"notnull default(0)"`
	// 	NormalMatchGameRtp      float64 `json:"NormalMatchGameRTP" xorm:"notnull default(0)"`
	// 	NormalMatchGameKillRate float64 `json:"NormalMatchGameKillRate" xorm:"notnull default(0)"`
	// 	LowBoundRtp             float64 `json:"LowBoundRTP" xorm:"notnull default(0)"`
	// }

	resultStr, err := utils.GetAPI(addr, "application/x-www-form-urlencoded", "")
	if err != nil {
		return "", API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	code, ok := result["code"].(float64)
	if ok {
		if code == 0 {
			val, ok := result["data"].(string)
			if ok {
				return val, utils.ToInt(code), nil
			}
		}
		return "", utils.ToInt(code), fmt.Errorf("Getdefaultgamesetting trans data has error, api code is %v, response body is: %v", code, resultStr)
	}
	return "", API_OTHER_ERROR, fmt.Errorf("Getdefaultgamesetting has error, response body is: %v", resultStr)
}

// 取得即時遊戲殺率
func GetRealtimeGameRatio(gameid int) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_getrealtimegameratio

	resultStr, err := utils.GetAPI(addr, fmt.Sprintf("gameid=%v", gameid), "")
	if err != nil {
		fmt.Printf("GetRealtimeGameRatio() response: %s, err: %v", resultStr, err)
		return "", API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			val, ok := result["data"].(string)
			if ok {
				return val, utils.ToInt(code), nil
			}
		}
		return "", utils.ToInt(code), fmt.Errorf("GetRealtimeGameRatio trans data has error, api code is %v, response body is: %v", code, resultStr)
	}

	return "", API_OTHER_ERROR, fmt.Errorf("GetRealtimeGameRatio has error, response body is: %v", resultStr)
}

// 取指定user目前遊戲局數判斷是否為新手期
func GetGameServerUserIsNewbie(userid int64) (string, int, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_getgameusernewbie

	resultStr, err := utils.GetAPI(addr, fmt.Sprintf("userid=%v", userid), "")
	if err != nil {
		fmt.Printf("GetGameServerUserIsNewbie() response: %s, err: %v", resultStr, err)
		return "", 0, API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			val, ok := result["data"].(string)
			if ok {
				resultInside := utils.ToMap([]byte(val))
				if limit, ok := resultInside["total_newbie_limit"].(float64); ok {
					return utils.ToJSON(resultInside["data"]), utils.ToInt(limit), utils.ToInt(code), nil
				}
			}
		}
		return "", 0, utils.ToInt(code), fmt.Errorf("GetGameServerUserIsNewbie trans data has error, api code is %v, response body is: %v", code, resultStr)
	}

	return "", 0, API_OTHER_ERROR, fmt.Errorf("GetGameServerUserIsNewbie has error, response body is: %v", resultStr)
}

// 取得Plinko球倍率上限
func GetPlinktoBallMaxOdds(agentName string) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_getplinkoballmaxodds

	params := fmt.Sprintf("agentName=%s", agentName)
	resultStr, err := utils.GetAPI(addr, params, "")
	if err != nil {
		fmt.Printf("GetPlinktoBallMaxOdds() response: %s, err: %v", resultStr, err)
		return "", API_OTHER_ERROR, err
	}

	result := utils.ToMap([]byte(resultStr))

	if code, ok := result["code"].(float64); ok {
		if code == 0 {
			val, ok := result["data"].(string)
			if ok {
				return val, utils.ToInt(code), nil
			}
		}
		return "", utils.ToInt(code), fmt.Errorf("GetPlinktoBallMaxOdds trans data has error, api code is: %v, response body is: %v", code, resultStr)
	}

	return "", API_OTHER_ERROR, fmt.Errorf("GetPlinktoBallMaxOdds has error, response body is: %v", resultStr)
}

// 設定Plinko球倍率上限
func SetPlinktoBallMaxOdds(agentName string, maxOdds float64) (string, int, error) {
	addr := ApiNotification_baseAddr + ApiNotification_setplinkoballmaxodds

	tmpMap := make(map[string]interface{}, 0)
	tmpMap["agent_name"] = agentName
	tmpMap["max_odds"] = maxOdds

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
			return utils.ToJSON(result), utils.ToInt(code), fmt.Errorf("SetPlinktoBallMaxOdds has error, code is: %v, agentName is: %v, maxOdds is: %v", code, agentName, maxOdds)
		}
	}

	return utils.ToJSON(result), API_OTHER_ERROR, fmt.Errorf("SetPlinktoBallMaxOdds has error, response body is: %v", resultStr)
}
