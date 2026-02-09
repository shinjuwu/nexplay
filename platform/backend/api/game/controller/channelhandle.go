package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/server/global"
	"definition"
	"fmt"
	"strconv"
	"sync"
	"time"

	"backend/pkg/redis"
	"backend/pkg/utils"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
)

/*
	3.2.1 登录游戏
		此接口用以验证游戏账号，如果账号不存在则创建游戏账号并为账号上分。
*/
func GameLogin(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	/*
		https://api.ky34.com/channelHandle?agent=10001&timestamp=1488781836949&param=ng
		tgiYCl26%2FgBmGvf9Euj2c1MOpzIzy4VWru%2Fsv3jao88cUlrENQTXz6pAeS3I2FqR7%2FPJFUIoT
		h%0D%0Ae%2FFnAkdbw2TxTkbhPCi5yjGJVVdY2C4%3D&key=f3afd416a0bb1b183eed8ef6cac30d7
		5

		s：操作子类型：0
		account：会员帐号(64 位字符)
		money：金额(上分的金额,如果不
		携带分数传 0)
		orderid：流水号（格式：代理编
		号+yyyyMMddHHmmssSSS+ account,
		长度不能超过 100 字符串）
		kind：遊戲 ID
	*/
	lock.Lock()
	defer lock.Unlock()

	type GameLoginChannelHandleResponseData struct {
		Code int    `json:"code"`
		Url  string `json:"url"`
	}

	returnData := GameLoginChannelHandleResponseData{
		Code: 0,
		Url:  "",
	}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_GameLogin, &returnData)

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")
	money, err := strconv.ParseFloat(data.ParamMap["money"][0], 64)
	if err != nil {
		return true, returnTmp, err
	} else {
		if money < 0 {
			returnData.Code = model.Response_CoinInOutValueFailed_Error
			return true, returnTmp, nil
		}
	}
	orderId := utils.ToString(data.ParamMap["orderid"][0], "")
	gameId := utils.ToInt64(data.ParamMap["kind"][0], 0)

	if thirdAccount == "" || orderId == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}
	/**
		1. 檢查遊戲連結是否有設定
	  		i. game -> "h5_link", "code"
		2. 檢查代理是否存在
			i. agent -> "id", "code", "level_code"
		3. 創建帳號
			i. 帳號不存在 -> 通知 game server 創建帳號並加款
			ii. 帳號存在 -> 通知 game server 加款
	*/
	// check game kind id is exist and get h5 link
	// http://${domain}/client/vue/apps/blackjack/
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("h5_link", "code", "server_info_code", "state", "type").
		From("game").
		Where(sq.Eq{"id": gameId}).
		ToSql()

	h5Link := ""
	gameCode := ""
	serverInfoCode := ""
	gameState := 0
	gameType := 0
	err = db.QueryRow(query, args...).Scan(&h5Link, &gameCode, &serverInfoCode, &gameState, &gameType)
	if err != nil || h5Link == "" || gameState != definition.GAME_STATE_ONLINE {
		returnData.Code = model.Response_GameNotOpen_Error
		return true, returnTmp, err
	}

	if gameType != definition.GAME_TYPE_LOBBY {
		h5Link = fmt.Sprintf("%s/%s", h5Link, gameCode)
	}

	// check agent code is exist
	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "code", "level_code", "cooperation", "name",
			"is_not_kill_dive_cal", "wallet_type", "lobby_switch_info").
		From("agent").
		Where(sq.And{
			sq.Eq{"id": data.Agent},
			sq.Eq{"is_enabled": 1},
		}).
		ToSql()

	agentId := 0
	agentCode := ""
	agentLevelCode := ""
	agentCooperation := 0
	agentname := ""
	agentIsNotKillDiveCal := false
	agentWalletType := definition.AGENT_WALLET_UNKNOW
	agentLobbySwitchInfo := 0
	err = db.QueryRow(query, args...).Scan(&agentId, &agentCode, &agentLevelCode, &agentCooperation, &agentname,
		&agentIsNotKillDiveCal, &agentWalletType, &agentLobbySwitchInfo)
	if err != nil || agentId <= 0 || len(agentLevelCode) <= 4 || agentWalletType < definition.AGENT_WALLET_TRANSFER {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	if gameType != definition.GAME_TYPE_LOBBY {
		// 非大廳類直接進入要在檢查代理是否有開啟遊戲
		// check agent game state is online
		query, args, _ = sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("state").
			From("agent_game").
			Where(sq.And{
				sq.Eq{"agent_id": data.Agent},
				sq.Eq{"game_id": gameId},
			}).
			ToSql()

		agentGameState := 0
		err = db.QueryRow(query, args...).Scan(&agentGameState)
		if err != nil {
			returnData.Code = model.Response_QueryRow_Error
			return true, returnTmp, err
		}

		if agentGameState != definition.GAME_STATE_ONLINE {
			returnData.Code = model.Response_GameNotOpen_Error
			return true, returnTmp, err
		}

	} else {
		lobbySwitch, find := definition.GameIdToLobbySwitch[int(gameId)]
		if !find {
			return true, returnTmp, fmt.Errorf("lobby not found, gameId=%d", gameId)
		}

		// 大廳類檢查代理是否有開啟開關
		lobbyOpen := agentLobbySwitchInfo&lobbySwitch == lobbySwitch
		if !lobbyOpen {
			returnData.Code = model.Response_GameNotOpen_Error
			return true, returnTmp, err
		}
	}

	defaultUserMetadata := model.NewUserMetadataEmpty().ToJson()
	transUsername := global.CreateNewThirdUsername(data.AgentCode, thirdAccount)
	userId, username, userMetadata, riskControlStatus, createNew, isEnabled, killDiveState, err := global.UdfCheckGameUserData(
		db, thirdAccount, transUsername, defaultUserMetadata, agentLevelCode,
		agentId, .0)
	if err != nil {
		logger.Printf("CheckCoinOutLimit exec QueryRow has error: %v", err)
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	errCode := 0
	// 該帳號是封停狀態
	if !isEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, err
	} else if string(riskControlStatus[definition.RISK_CONTROL_LOGIN_IDX]) == definition.RISK_CONTROL_STATUS_ENABLED {
		returnData.Code = model.Response_RiskControlLogin_Error
		return true, returnTmp, err
	}

	// 自動風控檢測
	autoRiskControlSetting := global.AutoRiskControlSettingCache.Get()
	if autoRiskControlSetting == nil {
		logger.Info("GameLogin AutoRiskControlSettingCache Get nil, userId=%d", userId)
	}

	checkAutoRiskControl := autoRiskControlSetting != nil && autoRiskControlSetting.IsEnabled
	if checkAutoRiskControl {
		if isEnabled = checkGameUserApiRequest(logger, db, rdb, agentId, userId,
			thirdAccount, agentLevelCode, autoRiskControlSetting.GameUserApiRequestPerSecondLimit); !isEnabled {
			returnData.Code = model.Response_AccountBlock_Error
			return true, returnTmp, err
		}
	}

	kind := definition.WALLET_LEDGER_KIND_API_UP
	info := ""
	creator := ""
	request := utils.ToJSON(data)

	loginTime := time.Now().UTC().String()
	lastLoginTime := "" // set default, if not new account, get lastlogintime from db

	// 通知 game server 上下分或是新增帳號
	if createNew {
		_, errCode, _ = createNewAccount(logger, db, userId, agentId, agentCooperation,
			money, data.ToCoin, data.Currency, orderId, thirdAccount,
			username, agentLevelCode, info, creator, request,
			agentname, createNew)
	} else if money > 0 {
		_, errCode, _ = coinIn(logger, db, agentId, agentCooperation, userId,
			kind, money, data.ToCoin, data.Currency, orderId,
			thirdAccount, agentLevelCode, info, creator, request, agentname)
	}

	if errCode == 0 {
		global.AutoRiskControlStatCache.IncrGameUserTotalCoinIn(userId, money)
	}

	// 非新增帳號, 取上次登入時間
	// 如果發生錯誤, 將上次登入時間改為本次登入時間
	if !createNew {
		lastLoginTime, _, err = getLastLoginTinme(logger, db, userId)
		if err != nil {
			lastLoginTime = ""
		}
	}

	// 查詢代理遊戲icon list
	args = make([]interface{}, 0)
	part := len(agentLevelCode) / 4
	partQuery := ""
	if len(agentLevelCode) == 4 {
		args = append(args, agentLevelCode)
	} else {
		for i := 1; i < part; i++ {
			args = append(args, agentLevelCode[0:i*4])
			partQuery += fmt.Sprintf(`$%d,`, i)
		}
		args = append(args, agentLevelCode)
		partQuery += fmt.Sprintf(`$%d`, part)
	}

	query = fmt.Sprintf(`select level_code, is_default, icon_list from agent_game_icon_list where level_code IN (%s) order by level_code desc`, partQuery)

	rows, err := db.Query(query, args...)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	var iconList string
	for rows.Next() {
		var dbLevelCode string
		var dbIsDefault bool
		var dbIconList string
		if err := rows.Scan(&dbLevelCode, &dbIsDefault, &dbIconList); err != nil {
			returnData.Code = model.Response_QueryRow_Error
			return true, returnTmp, err
		}

		// 吃預設就往上層找
		// 最上層 is_default 參數必須為false
		if !dbIsDefault {
			iconList = dbIconList
			break
		}

		// 防呆 避免最上層 is_default 參數不為false
		if len(dbLevelCode) == 4 {
			iconList = dbIconList
		}
	}

	rows.Close()

	// after create new account or send deposit add coin correct, record login info in redis and response to
	uuidToken, _ := uuid.NewV4()
	/**
	* 創建帳號的資訊
	* 此結構要與遊戲內部呼叫的 LoginGame api 服務相同
	 */
	response := model.LoginGameResponse{
		Id:               userId,
		AgentId:          agentId,
		AgentCode:        agentCode,
		LevelCode:        agentLevelCode,
		Username:         thirdAccount,
		TransUsername:    username,
		OrderId:          orderId,
		Coin:             money,
		KindId:           gameId,
		GameLink:         h5Link,
		Token:            uuidToken.String(),
		IsNew:            createNew,
		LoginTime:        loginTime,
		LastLoginTime:    lastLoginTime,
		UserMetadata:     userMetadata,
		KillDiveState:    killDiveState,
		IsRelogin:        false,
		GameIconList:     iconList,
		ServerInfoCode:   serverInfoCode,
		IsNotKillDiveCal: agentIsNotKillDiveCal,
		WalletType:       agentWalletType,
		LobbySwitchInfo:  createLobbySwitchInfoString(int(gameId), agentLobbySwitchInfo),
	}

	// check if exist old token in redis
	oldToken, _ := rdb.LoadHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, fmt.Sprintf("%d_%s", response.Id, response.Username))
	// record user info into redis
	if oldToken == "" {
		// store new one
		rdb.StoreHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, fmt.Sprintf("%d_%s", response.Id, response.Username), response.Token)
		rdb.StoreValue(global.REDIS_IDX_VERIFY_INFO, response.Token, utils.ToJSON(response), global.REDIS_LOGIN_TOKEN_LIFETIME)
	} else {
		// store new one
		rdb.StoreHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, fmt.Sprintf("%d_%s", response.Id, response.Username), response.Token)
		// remove old token & store new one
		rdb.DeleteValue(global.REDIS_IDX_VERIFY_INFO, oldToken)
		rdb.StoreValue(global.REDIS_IDX_VERIFY_INFO, response.Token, utils.ToJSON(response), global.REDIS_LOGIN_TOKEN_LIFETIME)
	}

	returnData.Url = h5Link + "?token=" + response.Token + "&lang=" + data.Lang
	returnData.Code = errCode
	returnTmp.D = returnData

	return true, returnTmp, nil
}

/*
	3.2.2 查询可下分
		此接口用来查询玩家的可下分余额
*/
func CheckCoinOutLimit(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	// type CheckCoinOutLimitResponse struct {
	// 	Money int64 `json"money"`
	// 	Code  int   `json"code"`
	// }
	/*
		参数加密字符串 param=
		(s=1&account=111111)
		s:操作子类型:1 account:会员帐号)
		Encrypt.AESEncrypt(param,DESKey);
		DESKey:平台提供

		{"s":101,"m":"/channelHandle","d":{"money":100,"code":0}}
	*/
	lock.Lock()
	defer lock.Unlock()

	type CheckCoinOutLimitChannelHandleResponseData struct {
		Balance float64 `json:"balance"` // 最後餘額
		Code    int     `json:"code"`
	}

	returnData := CheckCoinOutLimitChannelHandleResponseData{}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CheckCoinOutLimit, &returnData)

	errCode := 0
	gameGold := float64(0)

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")

	if thirdAccount == "" {
		errCode = model.Response_ParseParam_Error
	}

	userId := int64(0)
	transUsername := global.CreateNewThirdUsername(data.AgentCode, thirdAccount)

	agentId := utils.ToInt(data.Agent, 0)

	defaultUserMetadata := model.NewUserMetadataEmpty().ToJson()
	uid, username, _, _, createNew, _, _, err := global.UdfCheckGameUserData(db, thirdAccount, transUsername, defaultUserMetadata, data.LevelCode, agentId, .0)
	if err != nil {
		errCode = model.Response_QueryRow_Error
		logger.Printf("CheckCoinOutLimit exec QueryRow has error: %v", err)
	}

	// // 該帳號是封停狀態
	// if !isEnabled {
	// 	errCode = definition.ERROR_CODE_ERROR_GAME_USERS_BE_BLOCK
	// }

	userId = int64(uid)
	if errCode == 0 {
		if createNew {
			if userId > 0 {
				if apiCode, err := notification.SendCreateAccount("", 0, agentId, userId, thirdAccount, username); err != nil {
					logger.Printf("SendCreateAccount has error, err: %v", err)
					errCode = apiCode
				}
			}
		} else {
			// call game server api
			gold, apiCode, err := notification.GetGameServerGold(userId)
			if err != nil {
				gameGold = 0
				errCode = apiCode
				logger.Printf("CheckCoinOutLimit exec notification.GetGameServerGold() has error: %v", err)
			} else {
				gameGold = gold
			}
		}
	}

	returnData.Balance = gameGold
	returnData.Code = errCode

	return true, returnTmp, nil
}

/*
	3.2.3 上分
		此接口用来为账号上分
*/
func CoinIn(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	/*
		https://api.ky34.com/channelHandle?agent=10001&timestamp=1488791553051&param=nS
		42zzqT3fHQLBEfbB4ok2c1MOpzIzy4VWru%2Fsv3jao88cUlrENQTXz6pAeS3I2F8SI5db8tTG20%0D
		%0AWQDY9LQPMW9Xfy%2F1boz0REbE957bAvk%3D&key=8aeef9ff9b32f5f746ca663e8676a412

		参数加密字符串 param=
		(s=2&account=111111&money=10
		0&orderid=1000120170306143036
		949111111)

		s:操作子类型:2
		account:会员帐号
		money:金额(上分的金额)
		orderid:流水号(格式:代理编号+yyyyMMddHHmmssSSS+ account,长度不能超过 100 字符串)
		Encrypt.AESEncrypt(param,DESKey);
		DESKey:平台提供

		test data:
		agent: 1
		timestamp: 1234567890123
		param: XAU+Xui26JaAAXLRUDbNbX0ryg8SLVx+Kzw8QTrH5oESwo9Rq4delwbaFm0P+VTk8dX4+q2ebk3vzBlWjPcLLdkKBoOsx10xqOP8INGsSkw=
		key: 3e00d6fae2b56f3dc20f73108a9dbd1d
	*/
	lock.Lock()
	defer lock.Unlock()

	type CoinInChannelHandleResponseData struct {
		Code    int     `json:"code"`
		Balance float64 `json:"balance"`
	}

	returnData := CoinInChannelHandleResponseData{}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CoinIn, &returnData)

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")
	money, err := strconv.ParseFloat(data.ParamMap["money"][0], 64)
	if err != nil {
		returnData.Code = model.Response_CoinInOutValueFailed_Error
		return true, returnTmp, err
	} else {
		if money < 1 {
			returnData.Code = model.Response_CoinInOutValueFailed_Error
			return true, returnTmp, fmt.Errorf("exchange value can't less than 1")
		}
	}

	orderId := utils.ToString(data.ParamMap["orderid"][0], "")

	if thirdAccount == "" || orderId == "" || money <= 0 {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	// notice: 後台server的特殊參數請勿提供外部
	info := ""
	if _, ok := data.ParamMap["beinfo"]; ok {
		info = utils.ToString(data.ParamMap["beinfo"][0], "")
	}

	creator := ""
	if _, ok := data.ParamMap["becreator"]; ok {
		creator = utils.ToString(data.ParamMap["becreator"][0], "")
	}

	kind := 0
	if _, ok := data.ParamMap["bekind"]; ok {
		kind = utils.ToInt(data.ParamMap["bekind"][0], 0)
	}

	if kind != 0 && (kind != definition.WALLET_LEDGER_KIND_BACKEND_UP || info == "" || creator == "") {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	} else if kind == 0 {
		kind = definition.WALLET_LEDGER_KIND_API_UP
	}

	agentId := utils.ToInt(data.Agent)
	agentLevelCode := ""
	agentCooperation := 0
	agentname := ""

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("level_code", "cooperation", "name").
		From("agent").
		Where(sq.Eq{"id": agentId}).
		ToSql()
	err = db.QueryRow(query, args...).Scan(&agentLevelCode, &agentCooperation, &agentname)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	defaultUserMetadata := model.NewUserMetadataEmpty().ToJson()
	transUsername := global.CreateNewThirdUsername(data.AgentCode, thirdAccount)
	userId, username, _, riskControlStatus, createNew, isEnabled, _, err := global.UdfCheckGameUserData(
		db, thirdAccount, transUsername, defaultUserMetadata, agentLevelCode,
		agentId, .0)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		logger.Printf("CheckCoinOutLimit exec QueryRow has error: %v", err)
		return true, returnTmp, err
	}

	// 該帳號是封停狀態
	if !isEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, err
	} else if string(riskControlStatus[definition.RISK_CONTROL_COIN_IN_IDX]) == definition.RISK_CONTROL_STATUS_ENABLED {
		returnData.Code = model.Response_RiskControlCoinIn_Error
		return true, returnTmp, err
	}

	autoRiskControlSetting := global.AutoRiskControlSettingCache.Get()
	if autoRiskControlSetting == nil {
		logger.Info("CoinIn AutoRiskControlSettingCache Get nil, userId=%d", userId)
	}

	// API上分需要自動風控檢測
	checkAutoRiskControl := kind == definition.WALLET_LEDGER_KIND_API_UP && autoRiskControlSetting != nil && autoRiskControlSetting.IsEnabled
	if checkAutoRiskControl {
		if isEnabled = checkGameUserApiRequest(logger, db, rdb, agentId, userId, thirdAccount, agentLevelCode, autoRiskControlSetting.GameUserApiRequestPerSecondLimit); !isEnabled {
			returnData.Code = model.Response_AccountBlock_Error
			return true, returnTmp, err
		}

		if isTooManyRequests := checkGameUserCoinInAndOutRequest(logger, db, agentId, userId, thirdAccount, autoRiskControlSetting.GameUserCoinInAndOutRequestPerMinuteLimit); isTooManyRequests {
			returnData.Code = model.Response_TooManyRequests_Error
			return true, returnTmp, err
		}
	}

	request := utils.ToJSON(data)

	errCode := 0
	gameGold := float64(0)
	// 通知 game server 上下分或是新增帳號
	if createNew {
		gameGold, errCode, _ = createNewAccount(logger, db, userId, agentId, agentCooperation,
			money, data.ToCoin, data.Currency, orderId, thirdAccount,
			username, agentLevelCode, info, creator, request,
			agentname, createNew)
	} else {
		gold, _, _ := notification.GetGameServerGold(int64(userId))

		if utils.DecimalAdd(money, gold) >= 10000000 {
			returnData.Code = model.Response_CoinInOutValueFailed_Error
			return true, returnTmp, nil
		}
		gameGold, errCode, _ = coinIn(logger, db, agentId, agentCooperation, userId, kind,
			money, data.ToCoin, data.Currency, orderId, thirdAccount,
			agentLevelCode, info, creator, request, agentname)
	}

	if errCode == 0 {
		global.AutoRiskControlStatCache.IncrGameUserTotalCoinIn(userId, money)
	}

	returnData.Balance = gameGold
	returnData.Code = errCode
	returnTmp.D = returnData

	return true, returnTmp, nil
}

/*
	3.2.4 下分
		此接口用来为账号下分
*/
func CoinOut(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	/*
	   参数加密字符串 param=
	   (s=3&account=111111&money=100&orderid=1000120170306143036949111111)
	   s:操作子类型:3
	   account:会员帐号
	   money:金额(下分的金额,不要超过可下分数)
	   orderid:流水号(格式:代理编号+yyyyMMddHHmmssSSS+ account,长度不能超过 100 字符串)
	   Encrypt.AESEncrypt(param,DESKey);
	   DESKey:平台提供

	   {"s":103,"m":"/channelHandle","d":{"account":"111111","code":0}}

	   test data:
	   agent: 1
	   timestamp: 1234567890123
	   param: 8+pn0W7BgSNTAEydcMCumUSqXRbzCsouf8fTDCA11uTHay09T/3DQZzk1FEcCfLmpcUYoe6H5z5lF3tKrfWna9ncRBKVpE2nFG1ehrgsRTA=
	   key: 3e00d6fae2b56f3dc20f73108a9dbd1d
	*/
	lock.Lock()
	defer lock.Unlock()

	type CoinOutChannelHandleResponseData struct {
		Account string  `json:"account"` // 會員帳號
		Code    int     `json:"code"`    // 錯誤碼
		Balance float64 `json:"balance"` // 最後餘額
	}

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")

	returnData := CoinOutChannelHandleResponseData{Account: thirdAccount}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CoinOut, &returnData)

	money, err := strconv.ParseFloat(data.ParamMap["money"][0], 64)
	if err != nil {
		returnData.Code = model.Response_CoinInOutValueFailed_Error
		return true, returnTmp, err
	} else {
		if money < 1 {
			returnData.Code = model.Response_CoinInOutValueFailed_Error
			return true, returnTmp, fmt.Errorf("exchange value can't less than 1")
		}
	}

	orderId := utils.ToString(data.ParamMap["orderid"][0], "")

	if thirdAccount == "" || orderId == "" || money <= 0 {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	// notice: 後台server的特殊參數請勿提供外部
	info := ""
	if _, ok := data.ParamMap["beinfo"]; ok {
		info = utils.ToString(data.ParamMap["beinfo"][0], "")
	}

	creator := ""
	if _, ok := data.ParamMap["becreator"]; ok {
		creator = utils.ToString(data.ParamMap["becreator"][0], "")
	}

	kind := 0
	if _, ok := data.ParamMap["bekind"]; ok {
		kind = utils.ToInt(data.ParamMap["bekind"][0], 0)
	}

	if kind != 0 && (kind != definition.WALLET_LEDGER_KIND_BACKEND_DOWN || info == "" || creator == "") {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	} else if kind == 0 {
		kind = definition.WALLET_LEDGER_KIND_API_DOWN
	}

	agentId := utils.ToInt(data.Agent)
	agentLevelCode := ""
	agentCooperation := 0
	userId := 0
	isEnabled := false
	riskControlStatus := ""
	agentname := ""

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("gu.id", "a.level_code", "a.cooperation", "gu.is_enabled", "gu.risk_control_status", "a.name").
		From("game_users AS gu").
		InnerJoin("agent AS a ON gu.agent_id = a.id").
		Where(sq.And{
			sq.Eq{"gu.agent_id": agentId},
			sq.Eq{"gu.original_username": thirdAccount},
		}).
		ToSql()
	err = db.QueryRow(query, args...).Scan(&userId, &agentLevelCode, &agentCooperation, &isEnabled, &riskControlStatus,
		&agentname)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	if !isEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, err
	} else if string(riskControlStatus[definition.RISK_CONTROL_COIN_OUT_IDX]) == definition.RISK_CONTROL_STATUS_ENABLED {
		returnData.Code = model.Response_RiskControlCoinOut_Error
		return true, returnTmp, err
	}

	autoRiskControlSetting := global.AutoRiskControlSettingCache.Get()
	if autoRiskControlSetting == nil {
		logger.Info("CoinOut AutoRiskControlSettingCache Get nil, userId=%d", userId)
	}

	// API下分需要自動風控檢測
	checkAutoRiskControl := kind == definition.WALLET_LEDGER_KIND_API_DOWN && autoRiskControlSetting != nil && autoRiskControlSetting.IsEnabled
	if checkAutoRiskControl {
		if isEnabled = checkGameUserApiRequest(logger, db, rdb, agentId, userId, thirdAccount, agentLevelCode, autoRiskControlSetting.GameUserApiRequestPerSecondLimit); !isEnabled {
			returnData.Code = model.Response_AccountBlock_Error
			return true, returnTmp, err
		}

		if isTooManyRequests := checkGameUserCoinInAndOutRequest(logger, db, agentId, userId, thirdAccount, autoRiskControlSetting.GameUserCoinInAndOutRequestPerMinuteLimit); isTooManyRequests {
			returnData.Code = model.Response_TooManyRequests_Error
			return true, returnTmp, err
		}
	}

	gold, apiCode, err := notification.GetGameServerGold(int64(userId))
	if err != nil {
		logger.Info("CoinOut notification GetGameServerGold failed, userId=%d, username=%s, err=%v", userId, thirdAccount, err)
		returnData.Code = notification.TransformApiCodeToModelErrorCode(apiCode)
		return true, returnTmp, err
	}

	if gold < money {
		returnData.Code = model.Response_GameServeDepositFailed_Error
		return true, returnTmp, err
	}

	if checkAutoRiskControl {
		if isUserCoinOutEnabled := checkGameUserCoinOutRequest(logger, db, agentId, userId, thirdAccount, agentLevelCode, money, autoRiskControlSetting.GameUserCoinInAndOutDiffLimit); !isUserCoinOutEnabled {
			returnData.Code = model.Response_RiskControlCoinOut_Error
			return true, returnTmp, err
		}
	}

	// TODO: check orderId format
	// if errCode == 0 {
	// 	if _, check := global.CheckOrderIdFormat(orderId, utils.ToString(user_id), thirdAccount); check {
	// 		errCode = 19 + 100
	// 	}
	// }

	request := utils.ToJSON(data)

	gameGold, errCode, _ := coinOut(logger, db, agentId, agentCooperation, userId,
		kind, money, data.ToCoin, data.Currency, orderId,
		thirdAccount, agentLevelCode, info, creator, request,
		agentname)

	if errCode == 0 {
		global.AutoRiskControlStatCache.IncrGameUserTotalCoinOut(userId, money)
	}

	returnData.Balance = gameGold
	returnData.Code = errCode
	returnTmp.D = returnData

	return true, returnTmp, nil
}

/*
	3.2.5 查询订单(單筆)
		此接口用来查询玩家上下分的订单信息，通过 status 状态来判断上下分是否成功。
*/
func CheckUserOrder(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CheckUserOrderChannelHandleResponseData struct {
		Code   int     `json:"code"`   // 錯誤碼
		Status int     `json:"status"` // 狀態碼（-1：不存在、0：成功、2:失敗）
		Kind   int     `json:"kind"`   // 上下分類型(1:上分；2:下分)
		Money  float64 `json:"money"`  // 交易金額
	}

	returnData := CheckUserOrderChannelHandleResponseData{Status: -1}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CheckUserOrder, &returnData)

	orderid := utils.ToString(data.ParamMap["orderid"][0], "")
	if orderid == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	agentId := utils.ToInt(data.Agent)

	jsonChangeset := ""
	orderStatus := 0
	kind := 0
	errorCode := 0
	// check gameuser exist
	query := `SELECT changeset, status, kind, error_code FROM wallet_ledger WHERE id = $1 and agent_id = $2;`
	err := db.QueryRow(query, orderid, agentId).Scan(&jsonChangeset, &orderStatus, &kind, &errorCode)
	if err != nil {
		if err == sql.ErrNoRows {
			returnData.Code = model.Response_NoRowResultSet_Error
		} else {
			returnData.Code = model.Response_QueryRow_Error
		}
		return true, returnTmp, nil
	}

	if kind > definition.WALLET_LEDGER_KIND_API_DOWN {
		returnData.Code = model.Response_NoRowResultSet_Error
		return true, returnTmp, nil
	}

	returnData.Kind = kind

	if errorCode != 0 {
		returnData.Status = 2
	} else {
		returnData.Status = 0
	}

	changesetMap := utils.ToMap([]byte(jsonChangeset))
	if len(changesetMap) <= 0 {
		returnData.Code = model.Response_ParseJsonFailed_Error
		return true, returnTmp, nil
	}

	//{"add_coin": 100, "after_coin": 200, "before_coin": 100}
	transMoney, ok := changesetMap["add_coin"].(float64)
	if !ok {
		returnData.Code = model.Response_TypeTransFailed_Error
	}
	returnData.Money = transMoney

	return true, returnTmp, nil
}

/*
	3.2.6 查询玩家在线状态
		此接口用来查询玩家是否在线
*/
func CheckUserOnlineStatus(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CheckUserOnlineStatusChannelHandleResponseData struct {
		Code   int `json:"code"`   // 錯誤碼
		Status int `json:"status"` // 狀態碼（-1:不存在，0:不在線,1:在線）

	}

	returnData := CheckUserOnlineStatusChannelHandleResponseData{Status: -1}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CheckUserOnlineStatus, &returnData)

	dbUserId := 0
	dbIsEnabled := false

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")
	if thirdAccount == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	query := `SELECT id, is_enabled
	            FROM game_users
				WHERE original_username = $1 AND level_code = $2;`
	err := db.QueryRow(query, thirdAccount, data.LevelCode).Scan(&dbUserId, &dbIsEnabled)
	if err != nil {
		if err == sql.ErrNoRows {
			returnData.Status = -1
			return true, returnTmp, nil
		} else {
			returnData.Code = model.Response_QueryRow_Error
			return true, returnTmp, nil
		}
	}

	// check user online
	jsonStr, _ := rdb.LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, utils.IntToString(dbUserId))
	if jsonStr == "" {
		returnData.Status = 0
	} else {
		returnData.Status = 1
	}

	return true, returnTmp, nil
}

/*
	3.2.8 查询玩家总分
		此接口用来查询玩家的游戏内总分、玩家可下分余额、玩家在线状态
*/
func CheckUserCoin(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CheckUserCoinChannelHandleResponseData struct {
		Code       int     `json:"code"`        // 錯誤碼
		Status     int     `json:"status"`      // 狀態碼（-1:不存在，0:不在線,1:在線）
		TotalMoney float64 `json:"total_money"` // 總餘額
		FreeMoney  float64 `json:"free_money"`  // 可下分餘額
	}

	returnData := CheckUserCoinChannelHandleResponseData{Status: -1}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CheckUserCoin, &returnData)

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")
	if thirdAccount == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	dbUserId := 0
	dbIsEnabled := false

	// check gameuser exist
	query := `SELECT id, is_enabled
	            FROM game_users
				WHERE original_username = $1 AND level_code = $2;`
	err := db.QueryRow(query, thirdAccount, data.LevelCode).Scan(&dbUserId, &dbIsEnabled)
	if err != nil {
		if err == sql.ErrNoRows {
			returnData.Code = model.Response_GameUserNotExist_Error
		} else {
			returnData.Code = model.Response_QueryRow_Error
		}
		return true, returnTmp, nil
	}

	if !dbIsEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, nil
	}

	// check user online
	jsonStr, _ := rdb.LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, utils.IntToString(dbUserId))
	if jsonStr == "" {
		returnData.Status = 0
	} else {
		returnData.Status = 1
	}

	totalGold, freeGold, code, err := notification.GetGameServerGoldDetail(int64(dbUserId))
	if err != nil || code != 0 {
		returnData.Code = model.Response_GameServerGold_Error
		return true, returnTmp, nil
	}

	returnData.TotalMoney = totalGold
	returnData.FreeMoney = freeGold

	return true, returnTmp, nil
}

/*
	3.2.9 踢玩家下线
		此接口用以将在线的玩家强制离线
*/
func KickUser(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type KickUserChannelHandleResponseData struct {
		Code int `json:"code"` // 錯誤碼
	}

	returnData := KickUserChannelHandleResponseData{}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_KickUser, &returnData)

	thirdAccount := utils.ToString(data.ParamMap["account"][0], "")
	if thirdAccount == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	dbUserId := 0
	dbIsEnabled := false

	query := `SELECT id, is_enabled
	            FROM game_users
				WHERE original_username = $1 AND level_code = $2;`
	err := db.QueryRow(query, thirdAccount, data.LevelCode).Scan(&dbUserId, &dbIsEnabled)
	if err != nil {
		if err == sql.ErrNoRows {
			returnData.Code = model.Response_GameUserNotExist_Error
			return true, returnTmp, nil
		} else {
			returnData.Code = model.Response_QueryRow_Error
			return true, returnTmp, nil
		}
	}

	if !dbIsEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, nil
	}

	userId := utils.IntToString(dbUserId)

	// check user online
	serverInfoCode, err := rdb.LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, userId)
	if err != nil {
		returnData.Code = model.Response_PlayerNotOnline_Error
		return true, returnTmp, nil
	}

	_ = rdb.DeleteHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, userId)
	_ = rdb.DeleteHValue(
		global.REDIS_IDX_LOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, serverInfoCode),
		userId)

	// sent notification to game server and check response
	_, code, err := notification.SendSetBlockPlayer([]int{dbUserId}, true)
	if err != nil || code != 0 {
		returnData.Code = model.Response_GameServerBlockPlayer_Error
		logger.Printf("SendSetBlockPlayer has error code: %v, error: %v", code, err)
		return true, returnTmp, nil
	}

	return true, returnTmp, nil
}

/*
	3.2.10 查询代理余额(非必要)
		此接口用以查询代理余额
*/
func CheckAgentCoin(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CheckAgentCoinChannelHandleResponseData struct {
		Money float64 `json:"money"` // 代理餘額
		Code  int     `json:"code"`  // 錯誤碼
	}

	returnData := CheckAgentCoinChannelHandleResponseData{}
	returnTmp := model.CreateChannelHandleResponse(model.ChannelHandle_CheckAgentCoin, &returnData)

	dbAmount := float64(0)

	query := `select amount from agent_wallet where agent_id = $1;`
	err := db.QueryRow(query, data.Agent).Scan(&dbAmount)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, nil
	}

	returnData.Money = dbAmount

	return true, returnTmp, nil
}
