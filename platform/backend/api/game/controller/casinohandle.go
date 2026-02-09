package controller

import (
	"backend/api/game/model"
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"fmt"
	"strconv"
	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid"
)

func CasinoGameLogin(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.CasinoHandleRequest, lock *sync.Mutex) (bool, *model.CasinoHandleResponse, error) {
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
	returnTmp := model.CreateCasinoHandleResponse(model.ChannelHandle_GameLogin, &returnData)

	thirdAccount := data.Account

	if thirdAccount == "" {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	// check game user exist.
	var userId, killDiveState int
	var userMetadata, riskControlStatus, transUsername string
	var isEnabled bool
	query := `SELECT id, username, user_metadata, risk_control_status, is_enabled, kill_dive_state 
		FROM game_users 
		WHERE original_username=$1 AND agent_id=$2;`
	err := db.QueryRow(query, data.Account, data.Agent).Scan(&userId, &transUsername, &userMetadata, &riskControlStatus, &isEnabled,
		&killDiveState)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}
	if err == sql.ErrNoRows {
		returnData.Code = model.Response_GameUserNotExist_Error
		return true, returnTmp, err
	}

	// 該帳號是封停狀態
	if !isEnabled {
		returnData.Code = model.Response_AccountBlock_Error
		return true, returnTmp, err
	} else if string(riskControlStatus[definition.RISK_CONTROL_LOGIN_IDX]) == definition.RISK_CONTROL_STATUS_ENABLED {
		returnData.Code = model.Response_RiskControlLogin_Error
		return true, returnTmp, err
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
		Select("h5_link", "code", "server_info_code").
		From("game").
		Where(sq.And{
			sq.Eq{"id": 0}, // Lobby: 0
			sq.Eq{"state": 1},
		}).
		ToSql()

	h5Link := ""
	gameCode := ""
	serverInfoCode := ""
	err = db.QueryRow(query, args...).Scan(&h5Link, &gameCode, &serverInfoCode)
	if err != nil || h5Link == "" {
		return true, returnTmp, err
	}

	// if gameId != definition.GAME_ID_LOBBY {
	// 	h5Link = fmt.Sprintf("%s/%s", h5Link, gameCode)
	// }

	// check agent code is exist
	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "code", "level_code", "cooperation", "name").
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
	err = db.QueryRow(query, args...).Scan(&agentId, &agentCode, &agentLevelCode, &agentCooperation, &agentname)
	if err != nil || agentId <= 0 || len(agentLevelCode) <= 4 {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	errCode := 0

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

	loginTime := time.Now().UTC().String()
	lastLoginTime := "" // set default, if not new account, get lastlogintime from db
	orderId := ""
	money := float64(0)

	// 如果發生錯誤, 將上次登入時間改為本次登入時間
	lastLoginTime, _, err = getLastLoginTinme(logger, db, userId)
	if err != nil {
		lastLoginTime = ""
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
		Id:            userId,
		AgentId:       agentId,
		AgentCode:     agentCode,
		LevelCode:     agentLevelCode,
		Username:      thirdAccount,
		TransUsername: transUsername,
		OrderId:       orderId,
		Coin:          money,
		// KindId:         gameId,
		GameLink:       h5Link,
		Token:          uuidToken.String(),
		IsNew:          false,
		LoginTime:      loginTime,
		LastLoginTime:  lastLoginTime,
		UserMetadata:   userMetadata,
		KillDiveState:  killDiveState,
		IsRelogin:      false,
		GameIconList:   iconList,
		ServerInfoCode: serverInfoCode,
	}

	// check if exist old token in redis
	oldToken, _ := rdb.LoadHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, response.Username)
	// record user info into redis
	if oldToken == "" {
		// store new one
		rdb.StoreHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, response.Username, response.Token)
		rdb.StoreValue(global.REDIS_IDX_VERIFY_INFO, response.Token, utils.ToJSON(response), global.REDIS_LOGIN_TOKEN_LIFETIME)
	} else {
		// store new one
		rdb.StoreHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, response.Username, response.Token)
		// remove old token & store new one
		rdb.DeleteValue(global.REDIS_IDX_VERIFY_INFO, oldToken)
		rdb.StoreValue(global.REDIS_IDX_VERIFY_INFO, response.Token, utils.ToJSON(response), global.REDIS_LOGIN_TOKEN_LIFETIME)
	}

	returnData.Url = h5Link + "?token=" + response.Token + "&lang=" + data.Lang
	returnData.Code = errCode
	returnTmp.D = returnData

	return true, returnTmp, nil
}

func CasinoRegister(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.CasinoHandleRequest, lock *sync.Mutex) (bool, *model.CasinoHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CasinoRegisterResponseData struct {
		Code    int  `json:"code"`
		Success bool `json:"success"`
	}

	returnData := CasinoRegisterResponseData{
		Code:    0,
		Success: false,
	}
	returnTmp := model.CreateCasinoHandleResponse(model.ChannelHandle_GameLogin, &returnData)

	thirdAccount := data.Account

	if thirdAccount == "" {
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
		Select("h5_link", "code", "server_info_code").
		From("game").
		Where(sq.And{
			sq.Eq{"id": 0}, // Lobby: 0
			sq.Eq{"state": 1},
		}).
		ToSql()

	h5Link := ""
	gameCode := ""
	serverInfoCode := ""
	err := db.QueryRow(query, args...).Scan(&h5Link, &gameCode, &serverInfoCode)
	if err != nil || h5Link == "" {
		return true, returnTmp, err
	}

	// check agent code is exist
	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "code", "level_code", "cooperation", "name").
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
	err = db.QueryRow(query, args...).Scan(&agentId, &agentCode, &agentLevelCode, &agentCooperation, &agentname)
	if err != nil || agentId <= 0 || len(agentLevelCode) <= 4 {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}

	defaultUserMetadata := model.NewUserMetadataEmpty().ToJson()
	transUsername := global.CreateNewThirdUsername(data.AgentCode, thirdAccount)
	_, _, _, _, createNew, _, _, err := global.UdfCheckGameUserData(
		db, thirdAccount, transUsername, defaultUserMetadata, agentLevelCode,
		agentId, .0)
	if err != nil {
		logger.Printf("UdfCheckGameUserData exec QueryRow has error: %v", err)
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}
	if !createNew {
		returnData.Code = model.Response_AccountExist_Error
	}
	returnData.Success = createNew
	returnTmp.D = returnData

	return true, returnTmp, nil
}

func CasinoCredit(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.CasinoHandleRequest, lock *sync.Mutex) (bool, *model.CasinoHandleResponse, error) {
	lock.Lock()
	defer lock.Unlock()

	type CasinoCreditResponseData struct {
		Code    int     `json:"code"`
		Balance float64 `json:"balance"`
	}

	returnData := CasinoCreditResponseData{}
	returnTmp := model.CreateCasinoHandleResponse(model.ChannelHandle_CoinIn, &returnData)

	thirdAccount := data.Account
	money, err := strconv.ParseFloat(data.Coin, 64)
	if err != nil {
		returnData.Code = model.Response_CoinInOutValueFailed_Error
		return true, returnTmp, err
	} else {
		if money < 1 {
			returnData.Code = model.Response_CoinInOutValueFailed_Error
			return true, returnTmp, fmt.Errorf("exchange value can't less than 1")
		}
	}

	// create order id
	// 代理號+帳號+今日日期
	orderId := data.Agent + thirdAccount + utils.Get18UnsignedTimeNowUTC()

	if thirdAccount == "" || orderId == "" || money <= 0 {
		returnData.Code = model.Response_ParseParam_Error
		return true, returnTmp, nil
	}

	// notice: 後台server的特殊參數請勿提供外部
	info := "casino coin in"
	creator := data.Account
	kind := definition.WALLET_LEDGER_KIND_BACKEND_UP

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

	var userId, killDiveState int
	var userMetadata, riskControlStatus, transUsername string
	var isEnabled bool
	query = `SELECT id, username, user_metadata, risk_control_status, is_enabled, kill_dive_state 
		FROM game_users 
		WHERE original_username=$1 AND agent_id=$2;`
	err = db.QueryRow(query, data.Account, data.Agent).Scan(&userId, &transUsername, &userMetadata, &riskControlStatus, &isEnabled,
		&killDiveState)
	if err != nil {
		returnData.Code = model.Response_QueryRow_Error
		return true, returnTmp, err
	}
	if err == sql.ErrNoRows {
		returnData.Code = model.Response_GameUserNotExist_Error
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

	gold, _, _ := notification.GetGameServerGold(int64(userId))

	if utils.DecimalAdd(money, gold) >= 10000000 {
		returnData.Code = model.Response_CoinInOutValueFailed_Error
		return true, returnTmp, nil
	}
	gameGold, errCode, _ = coinIn(logger, db, agentId, agentCooperation, userId, kind,
		money, data.ToCoin, data.Currency, orderId, thirdAccount,
		agentLevelCode, info, creator, request, agentname)

	if errCode == 0 {
		global.AutoRiskControlStatCache.IncrGameUserTotalCoinIn(userId, money)
	}

	returnData.Balance = gameGold
	returnData.Code = errCode
	returnTmp.D = returnData

	return true, returnTmp, nil
}
