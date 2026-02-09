package intercom

import (
	game_model "backend/api/game/model"
	"backend/api/intercom/model"
	"backend/api/intercom/module"
	"backend/api/intercom/sqlutils"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"
	"backend/server/global"
	"definition"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

/*
內部溝通 api
與遊戲溝通使用
*/

/*
if you want create new one, copy that and paste on file.
*/

// /*
// 	lock for intercom api module
// */
// LockOfintercomTokenSyncLock *sync.Mutex
// // [token, sync.lock]
// IntercomTokenSyncLock sync.Map

type IntercomApi struct {
	BasePath    string
	userLockMap *sync.Map
}

/** create new api group instance
 * basePath: path of base group router
 */
func NewIntercomApi(basePath string) api_cluster.IApiEach {
	return &IntercomApi{
		BasePath:    basePath,
		userLockMap: new(sync.Map),
	}
}

/** get base group path */
func (p *IntercomApi) GetGroupPath() string {
	return p.BasePath
}

// if get nil, create new one
func (p *IntercomApi) GetUserLock(key string) *sync.Mutex {
	tmp, _ := p.userLockMap.LoadOrStore(key, new(sync.Mutex))
	return tmp.(*sync.Mutex)
}

/** 註冊 api */
func (p *IntercomApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.GET("/logingame", ginHandler.Handle(p.LoginGame))                     //10101
	g.GET("/logoutgame", ginHandler.Handle(p.LogoutGame))                   //10102
	g.POST("/creategamerecord", ginHandler.Handle(p.CreateGameRecord))      //10103
	g.GET("/getmarqueesetting", ginHandler.Handle(p.GetMarqueeSettingList)) //10104
	g.GET("/getgameserverstate", ginHandler.Handle(p.GetGameServerState))
	g.GET("/getagentlist", ginHandler.Handle(p.GetAgentList))
	g.POST("/servicerestart", ginHandler.Handle(p.ServiceRestart))
	g.POST("/singlewallet", ginHandler.Handle(p.SingleWallet))               //10106
	g.POST("/createjackpotrecord", ginHandler.Handle(p.CreateJackpotRecord)) //10105
	g.POST("/recalculategamerecord", ginHandler.Handle(p.RecalculateGameRecord))
	g.POST("/createfriendroomrecord", ginHandler.Handle(p.CreateFriendRoomRecord))

	go module.T_ResendSingleWalletRequest(ginHandler.Logger, ginHandler.DB.GetDefaultDB())
}

// @Tags IntercomApi
// @Summary login game
// @Description 此接口用以驗證遊戲帳號，如果帳號不存在則創建遊戲帳號並為帳號上分。(server use)
// @Produce  application/json
// @param token query string true "token"
// @Success 200 {object} response.Response{data=model.LoginGameResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/logingame [get]
func (p *IntercomApi) LoginGame(c *ginweb.Context) {

	token := c.Request.URL.Query().Get("token")

	var loginInfo model.LoginGameResponse

	isRelogin := false
	// get login info from redis
	jsonStr, err := c.Redis().LoadValue(global.REDIS_IDX_VERIFY_INFO, token)
	if err != nil {
		//search from relogin info
		jsonStr, err = c.Redis().LoadValue(global.REDIS_IDX_RELOGIN_INFO, token)
		if err != nil {
			c.Fail(definition.INTERCOME_ERROR_CODE_GET_VERIFY_INFO_FAILED)
			return
		}
		//delete relogin token
		c.Redis().DeleteValue(global.REDIS_IDX_RELOGIN_INFO, token)
		isRelogin = true
	}

	err = utils.ToStruct([]byte(jsonStr), &loginInfo)
	if err != nil {
		c.Fail(definition.INTERCOME_ERROR_CODE_PARAMS_PARSE_FAILED)
		return
	}

	// check use have relogin token else?
	// if token exist, delete it.
	reloginToken, _ := c.Redis().LoadHValue(
		global.REDIS_IDX_RELOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_RELOGIN_TOKEN, loginInfo.ServerInfoCode),
		utils.IntToString(loginInfo.Id))
	if reloginToken != "" {
		c.Redis().DeleteHValue(
			global.REDIS_IDX_RELOGIN_INFO,
			global.GenRedisHashName(global.REDIS_HASH_RELOGIN_TOKEN, loginInfo.ServerInfoCode),
			utils.IntToString(loginInfo.Id))
		c.Redis().DeleteValue(global.REDIS_IDX_RELOGIN_INFO, reloginToken)
	}

	loginInfo.IsRelogin = isRelogin // 是否為重連

	loginTime := time.Now().UTC()
	loginInfo.LoginTime = loginTime.String()

	finalJsonStr := utils.ToJSON(loginInfo)

	query := `INSERT INTO user_login_log (level_code, agent_id, game_user_id, token, coin, user_info, is_new, login_time)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	result, err := c.DB().ExecContext(c.Request.Context(), query, loginInfo.LevelCode, loginInfo.AgentId, loginInfo.Id, token, loginInfo.Coin, finalJsonStr, loginInfo.IsNew, loginTime)
	if err != nil {
		c.Logger().Printf("LoginGame() has error to insert user_login_log, err  is: %v", err)
	} else {
		if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
			c.Logger().Printf("LoginGame() insert user_login_log rowsAffectedCount is 0, query is %s", query)
		}
	}

	query = `UPDATE game_users SET last_login_time = $1 WHERE "id" = $2 `
	result, err = c.DB().ExecContext(c.Request.Context(), query, loginTime, loginInfo.Id)
	if err != nil {
		c.Logger().Printf("LoginGame() has error to update game_users, err  is: %v", err)
	} else {
		if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
			c.Logger().Printf("LoginGame() update game_users rowsAffectedCount is 0, query is %s", query)
		}
	}

	// if !c.GetIsDebugMode() {
	c.Redis().DeleteHValue(global.REDIS_IDX_VERIFY_INFO, global.REDIS_HASH_LOGIN_TOKEN, fmt.Sprintf("%d_%s", loginInfo.Id, loginInfo.Username))
	c.Redis().DeleteValue(global.REDIS_IDX_VERIFY_INFO, token)
	// }

	// 紀錄用戶登入資訊
	c.Redis().StoreHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, utils.IntToString(loginInfo.Id), loginInfo.ServerInfoCode)
	c.Redis().StoreHValue(global.REDIS_IDX_LOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, loginInfo.ServerInfoCode),
		utils.IntToString(loginInfo.Id), utils.ToJSON(loginInfo))

	c.Ok(loginInfo)
}

// @Tags IntercomApi
// @Summary logout game
// @Description 此接口用以登出遊戲帳號(server use)
// @Produce  application/json
// @param user_id query string true "用戶id"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/logoutgame [get]
func (p *IntercomApi) LogoutGame(c *ginweb.Context) {
	reqUserId := c.Request.URL.Query().Get("user_id")
	reqToken := c.Request.URL.Query().Get("token")

	serverInfoCode, err := c.Redis().LoadHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
	if err != nil {
		c.Logger().Printf("LogoutGame() has error to update user_login_log, err  is: %v", err)
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_INGAMEUSER_FAILED)
		return
	}

	jsonStr, err := c.Redis().LoadHValue(
		global.REDIS_IDX_LOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, serverInfoCode),
		reqUserId)
	if err != nil {
		c.Logger().Printf("LogoutGame() has error to update user_login_log, err  is: %v", err)
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_INGAMEUSER_FAILED)
		return
	}

	tmp := utils.ToMap([]byte(jsonStr))
	token := utils.ToString(tmp["token"], "")
	id := utils.StringToInt(reqUserId, -1)

	if token == "" || id == -1 {
		c.Logger().Printf("LogoutGame() has error to get user info, userId is %s", reqUserId)
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARAMS_PARSE_FAILED)
		return
	}

	logoutTime := time.Now().UTC()

	query := `UPDATE user_login_log SET logout_time = $1 WHERE game_user_id = $2 AND token = $3;`
	_, err = c.DB().ExecContext(c.Request.Context(), query, logoutTime, id, token)
	if err != nil {
		c.Logger().Printf("LogoutGame() has error to update user_login_log, err  is: %v", err)
	}

	query = `UPDATE game_users SET last_logout_time = $1 WHERE "id" = $2 `
	_, err = c.DB().ExecContext(c.Request.Context(), query, logoutTime, id)
	if err != nil {
		c.Logger().Printf("LogoutGame() has error to update game_users, err  is: %v", err)
	}

	// 只要遊戲呼叫就更新最後登入時間
	// token 不一樣，代表用戶已經用其他token 登入
	// 且傳入的token 跟記錄在redis 內的token 不為空值
	if token != "" && reqToken != "" && token != reqToken {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_INGAMEUSER_FAILED)
		return
	}

	// 取得自己代理當下的icon list setting
	agentLevelCode := utils.ToString(tmp["level_code"], "")
	// 查詢代理遊戲icon list
	args := make([]interface{}, 0)
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

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.Logger().Printf("LogoutGame() has error to select agent_game_icon_list, err  is: %v", err)
	}

	var iconList string
	for rows.Next() {
		var dbLevelCode string
		var dbIsDefault bool
		var dbIconList string
		if err := rows.Scan(&dbLevelCode, &dbIsDefault, &dbIconList); err != nil {
			c.Logger().Printf("LogoutGame() has error to Scan agent_game_icon_list, err  is: %v", err)
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

	if iconList != "" {
		tmp["game_icon_list"] = iconList
	}

	rows.Close()

	_ = c.Redis().DeleteHValue(global.REDIS_IDX_LOGIN_INFO, global.REDIS_HASH_INGAME_USER, reqUserId)
	_ = c.Redis().DeleteHValue(
		global.REDIS_IDX_LOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, serverInfoCode),
		reqUserId)

	if len(tmp) > 0 {
		// save relogin game user
		err = c.Redis().StoreValue(global.REDIS_IDX_RELOGIN_INFO, token, utils.ToJSON(tmp), global.REDIS_RELOGIN_TOKEN_LIFETIME)
		if err != nil {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_RELOGIN_TOKEN_SAVE_FAILED)
			return
		}

		_ = c.Redis().StoreHValue(global.REDIS_IDX_RELOGIN_INFO,
			global.GenRedisHashName(global.REDIS_HASH_RELOGIN_TOKEN, serverInfoCode),
			reqUserId,
			token)
	}

	c.Ok("success")
}

// @Tags IntercomApi
// @Summary create game record
// @Description 此接口用來創建遊戲紀錄(不做參數檢查)
// @Produce  application/json
// @Param data body model.PlayLogRequest true "代理識別碼, 遊戲紀錄(json), 總變動分數..."
// @Success 200 {object} response.Response{data=model.PlayLogResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/creategamerecord [post]
func (p *IntercomApi) CreateGameRecord(c *ginweb.Context) {
	var req model.PlayLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	gameName := ""
	gameTableObj := global.GameCache.Get(req.GameId)

	if gameTableObj != nil {
		gameName = gameTableObj.Code
	} else {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GAME_ID_NOT_EXIST)
		return
	}

	// 轉小寫
	gameName = strings.ToLower(gameName)
	playLogCommonTableName := "play_log_common"
	userPlayLogTableName := fmt.Sprintf("user_play_log_%s", gameName)
	// playLogTableName := fmt.Sprintf("play_log_%s", gameName)

	betTimeI64 := strconv.FormatInt(req.BetTime, 10)
	betTime := utils.ToTimeUnixMilliUTC(betTimeI64, 0)

	startTimeI64 := strconv.FormatInt(req.StartTime, 10)
	startTime := utils.ToTimeUnixMilliUTC(startTimeI64, 0)

	endTimeI64 := strconv.FormatInt(req.EndTime, 10)
	endTime := utils.ToTimeUnixMilliUTC(endTimeI64, 0)

	var err error

	// 遊戲局紀錄
	err = sqlutils.CreateGamePlayLog(c.DB(), c.Redis(), c.Logger(), playLogCommonTableName, req, betTime, startTime, endTime)
	if err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_CREATE_PLAYLOG_COMMON_FAILED)
		return
	}

	// 用戶個人紀錄
	// 如果個人記錄沒有，就去反查遊戲局記錄的raw data
	err = sqlutils.DispatchUserPlayLogFromGameId(c.DB(), c.Redis(), c.Logger(), userPlayLogTableName, req.Lognumber, req.Playlog,
		req.GameId, req.RoomType, req.DeskId, req.Exchange, betTime)
	if err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_CREATE_USERPLAYLOG_FAILED)
		return
	}

	c.Ok("success")
}

// @Tags IntercomApi
// @Summary get marquee setting
// @Description 此接口供遊戲伺服器取得跑馬燈設定列表
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]model.MarqueeSettingDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/getmarqueesetting [get]
func (p *IntercomApi) GetMarqueeSettingList(c *ginweb.Context) {

	nowTime := time.Now().UTC()
	// column: 11
	query := `SELECT "id", "lang", "type", "order", "freq", "content", "start_time", "end_time"
			FROM marquee 
			WHERE "is_open" = true AND "is_enabled" = true AND "end_time" > $1;`

	rows, err := c.DB().Query(query, nowTime)
	if err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_SELECT_DATABASE_FAILED)
		return
	}

	defer rows.Close()

	resp := make([]model.MarqueeSettingDataResponse, 0)

	for rows.Next() {
		var temp model.MarqueeSettingDataResponse
		// newline every scan 5 column
		if err := rows.Scan(&temp.Id, &temp.Lang, &temp.Type, &temp.Order, &temp.Freq,
			&temp.Content, &temp.StartTime, &temp.EndTime); err != nil {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARSE_ROWS_FAILED)
			return
		}

		resp = append(resp, temp)
	}

	c.Ok(resp)
}

// @Tags IntercomApi
// @Summary get game global switch
// @Description 此接口供遊戲伺服器取得設定遊戲全局開關設定列表
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]model.GetGameServerStateResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/getgameserverstate [get]
func (p *IntercomApi) GetGameServerState(c *ginweb.Context) {

	resp := model.GetGameServerStateResponse{}

	gameServerInfoStorage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO)
	if !ok {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_LOCALMEMORY_FAILED)
		return
	}

	err := utils.ToStruct([]byte(gameServerInfoStorage.Value), &resp)
	if err != nil {
		resp.State = definition.GS_STATE_DEFAULT
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_LOCAL_ANALYZE_FAILED)
		return
	}

	if resp.State > definition.GS_STATE_CLOSE && resp.State < definition.GS_STATE_OPEN {
		resp.State = definition.GS_STATE_DEFAULT
	}

	c.Ok(resp)
}

// @Tags IntercomApi
// @Summary get agent list
// @Description 此接口供遊戲伺服器取得代理列表(僅回傳已啟用的代理資料)
// @Produce  application/json
// @Success 200 {object} response.Response{data=[]model.AgentList,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/getagentlist [get]
func (p *IntercomApi) GetAgentList(c *ginweb.Context) {
	query := `SELECT "id", "name", "md5_key", "aes_key"
            FROM agent 
            WHERE "is_enabled" = 1;`

	rows, err := c.DB().Query(query)
	if err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_SELECT_DATABASE_FAILED)
		return
	}

	defer rows.Close()

	resp := make([]*model.AgentList, 0)

	for rows.Next() {
		temp := new(model.AgentList)
		// newline every scan 5 column
		if err := rows.Scan(&temp.Id, &temp.Name, &temp.Md5Key, &temp.AesKey); err != nil {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARSE_ROWS_FAILED)
			return
		}

		resp = append(resp, temp)
	}

	c.Ok(resp)
}

// @Tags IntercomApi
// @Summary notify backend service, whem game service restart
// @Description 此接口用於遊戲服務重啟時，通知後臺使用
// @Produce  application/json
// @Param data body model.ServiceParams true "service param"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/servicerestart [post]
func (p *IntercomApi) ServiceRestart(c *ginweb.Context) {
	var req model.ServiceParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	code := strings.ToLower(req.Id)

	_, ok := global.ServerInfoCache.Load(code)
	if !ok {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARAMS_PARSE_FAILED)
		return
	}

	// delete InGameUser and ReLoginTokne index
	c.Redis().Del(global.REDIS_IDX_LOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_INGAME_USER, code))
	c.Redis().Del(global.REDIS_IDX_RELOGIN_INFO,
		global.GenRedisHashName(global.REDIS_HASH_RELOGIN_TOKEN, code))
	c.Redis().Del(global.REDIS_IDX_LOGIN_INFO,
		global.REDIS_HASH_INGAME_USER)
	c.Redis().Del(global.REDIS_IDX_RELOGIN_INFO,
		global.REDIS_HASH_RELOGIN_TOKEN)

	c.Ok("success")
}

// @Tags IntercomApi
// @Summary create jackpot record
// @Description 此接口用來創建jackpot紀錄(不做參數檢查)
// @Produce  application/json
// @Param data body model.JackpotLogRequest true "代理識別碼, 遊戲紀錄(json), 總變動分數..."
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/createjackpotrecord [post]
func (p *IntercomApi) CreateJackpotRecord(c *ginweb.Context) {
	var req model.JackpotLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	c.Logger().Info("Receive CreateJackpotRecord, req: %s", utils.ToJSON(req))

	if err := sqlutils.CreateJackpotLog(c.DB(), req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_CREATE_JACKPOTLOG_FAILED)
		return
	}

	c.Logger().Info("Receive CreateJackpotRecord success, lognumber: %s, token id: %s, agent id: %d, user: %s", req.Lognumber, req.TokenId, req.AgentId, req.Username)

	c.Ok("success")
}

// @Tags IntercomApi
// @Summary notify backend service, single wallet get balance
// @Description 此接口用於單一錢包 Api 整合
// @Produce  application/json
// @Param data body model.SingleWalletRequest true "service param"
// @Success 200 {object} response.Response{data=model.SingleWalletResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/singlewallet [post]
func (p *IntercomApi) SingleWallet(c *ginweb.Context) {
	var req model.SingleWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	if !req.CheckParam() {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_PARAMS_PARSE_FAILED)
		return
	}

	// 取得遊戲用戶關聯代理資訊
	obj := global.AgentDataOfGameUserCache.Get(req.UserId)
	if obj == nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_AGENT_GAME_USER_DATA_FAILED)
		return
	}

	// 檢查代理 ID 範圍
	if obj.AgentId <= 0 {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_AGENT_ID_FAILED)
		return
	}

	agentObj := global.AgentCache.Get(obj.AgentId)
	if agentObj == nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_GET_AGENT_GAME_USER_DATA_FAILED)
		return
	}

	// check agent wallet type
	if agentObj.WalletType != definition.AGENT_WALLET_SINGLE {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_AGENT_WALLET_TYPE_FAILED)
		return
	}

	// user lock key format is : [agentId]_[userId]
	lock := p.GetUserLock(fmt.Sprintf("%d_%d", obj.AgentId, req.UserId))

	lock.Lock()
	defer lock.Unlock()

	// get api address from table of agent
	// agentObj.WalletConnInfo.Scheme
	agentStr := strconv.Itoa(agentObj.Id)
	timestamp := time.Now().UnixMilli()
	timestampStr := utils.Int64ToString(timestamp)
	ss := agentStr + timestampStr + agentObj.Md5Key
	hs32 := md5.Hash32bit(ss)

	// create order id
	// create wallet ledger id for group up/down game coin
	kind := 0
	walletLedgerId := req.WalletLedgerId
	switch req.Command {
	case model.SingleWalletRequest_Command_AddScore:
		kind = definition.WALLET_LEDGER_KIND_SINGLE_WALLET_UP
	case model.SingleWalletRequest_Command_MinusScore:
		kind = definition.WALLET_LEDGER_KIND_SINGLE_WALLET_DOWN
	}
	salt := fmt.Sprintf("%d%d%d%v%s_%d", definition.ORDER_TYPE_AGENT_WALLET_LEDGER,
		agentObj.Id, kind, req.Point, req.Username, timestamp%int64(time.Millisecond))
	orderId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_SINGLE_WALLET_LOG, salt, time.Now())

	var paramQuery, paramQueryCancel url.Values
	var b64EncodingParam, b64EncodingParamCancel string
	var u, uc url.URL

	paramQuery, paramQueryCancel = model.CreateSingleWalletParamQueryWithCancel(req.Command, model.SingleWalletRequest_Command_CancelAddScore, req.Username, fmt.Sprintf("%f", req.Point), orderId,
		strconv.Itoa(req.GameId), strconv.Itoa(req.RoomId), agentObj.Currency, req.BetId, walletLedgerId)

	// 參數加密字符串
	b64EncodingParam = model.CreateParamEncpy(paramQuery.Encode(), agentObj.AesKey)
	if b64EncodingParam == "" {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_ENCODE_PARAM_FAILED)
		return
	}

	u = model.CreateSingleWalletCallbackURL(global.DEF_PLATFORM, agentStr, b64EncodingParam, timestampStr, hs32,
		agentObj.WalletConnInfo.Scheme, agentObj.WalletConnInfo.Domain, agentObj.WalletConnInfo.Path)

	if req.Command == model.SingleWalletRequest_Command_AddScore {
		b64EncodingParamCancel = model.CreateParamEncpy(paramQueryCancel.Encode(), agentObj.AesKey)
		if b64EncodingParamCancel == "" {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_ENCODE_PARAM_FAILED)
			return
		}
		uc = model.CreateSingleWalletCallbackURL(global.DEF_PLATFORM, agentStr, b64EncodingParamCancel, timestampStr, hs32,
			agentObj.WalletConnInfo.Scheme, agentObj.WalletConnInfo.Domain, agentObj.WalletConnInfo.Path)
	}

	exchangeData := global.ExchangeDataCache.Get(agentObj.Currency)

	info := ""
	creator := c.Ip
	toCoin := exchangeData.ToCoin

	if req.Command != model.SingleWalletRequest_Command_QueryMinusScore {
		// create wallet record
		code, err := sqlutils.SWCoinInOutStart(c.Logger(), c.DB(), agentObj.Id, agentObj.Cooperation, req.UserId,
			kind, req.Point, toCoin, agentObj.Currency, orderId,
			req.Username, agentObj.LevelCode, info, creator, utils.ToJSON(u), agentObj.Name, walletLedgerId)
		if err != nil || code != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_START_SINGLE_WALLET_LEDGER_FAILED)
			return
		}
	}

	type ResponseData struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	var resBodyData ResponseData
	resBodyData.Code = -1
	apiErrCode := game_model.Response_Success

	// call third wallet api
	body, httpCode, err := utils.GetBasicAuthAPIWithHttpCode(u.String(), agentObj.WalletConnInfo.ApiKey, "")
	// 實際發生error or 平台方回應 httpCode = 500 強制啟動重送
	if err != nil || httpCode == http.StatusInternalServerError {
		log.Printf("SingleWallet GetBasicAuthAPI() httpCode: %v, error: %v", httpCode, err)
		// 沒有正常 response, 就要做例外錯誤處理
		// 如果上分失敗，要送取消上分(平台扣款)，排入重送機制
		if req.Command == model.SingleWalletRequest_Command_AddScore {
			global.InsertReSendData(c.Logger(), c.DB(), agentObj.Id, global.RESEND_TYPE_SINGLEWALLET, agentObj.LevelCode,
				orderId, utils.ToJSON(req), utils.ToJSON(uc))
		} else if req.Command == model.SingleWalletRequest_Command_MinusScore {
			// 如果下分失敗，排入重送機制
			global.InsertReSendData(c.Logger(), c.DB(), agentObj.Id, global.RESEND_TYPE_SINGLEWALLET, agentObj.LevelCode,
				orderId, utils.ToJSON(req), utils.ToJSON(u))
		}

		if os.IsTimeout(err) {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_IO_TIMEOUT_FAILED)
			return
		}

		c.OkWithCode(definition.INTERCOME_ERROR_CODE_CALL_API_FAILED)
		return
	} else {
		// "{\"code\":0,\"msg\":\"操作成功\",\"data\":{\"code\":0,\"s\":101,\"account\":\"kinco\",\"currency\":\"RMB\",\"money\":\"0.0000\",\"point\":\"0.0000\"}}"
		err = json.Unmarshal([]byte(body), &resBodyData)
		if err != nil || resBodyData.Code > 0 {
			// 平台方回應 api code = INTERCOME_ERROR_CODE_SINGLE_WALLET_FORCED_RESENDING(699) 強制啟動重送
			if resBodyData.Code == definition.INTERCOME_ERROR_CODE_SINGLE_WALLET_FORCED_RESENDING {
				if req.Command == model.SingleWalletRequest_Command_AddScore {
					global.InsertReSendData(c.Logger(), c.DB(), agentObj.Id, global.RESEND_TYPE_SINGLEWALLET, agentObj.LevelCode,
						orderId, utils.ToJSON(req), utils.ToJSON(uc))
				} else if req.Command == model.SingleWalletRequest_Command_MinusScore {
					// 如果下分失敗，排入重送機制
					global.InsertReSendData(c.Logger(), c.DB(), agentObj.Id, global.RESEND_TYPE_SINGLEWALLET, agentObj.LevelCode,
						orderId, utils.ToJSON(req), utils.ToJSON(u))
				}

				c.OkWithCode(definition.INTERCOME_ERROR_CODE_SINGLE_WALLET_FORCED_RESENDING)
				return
			} else {
				// 真的 API 錯誤，不重送，結束該筆注單
				apiErrCode = game_model.Response_GameServeDepositFailed_Error
				// 命令非查詢餘額
				if req.Command != model.SingleWalletRequest_Command_QueryMinusScore {
					code, err := sqlutils.SWCoinInOutFinish(c.Logger(), c.DB(), agentObj.Id, agentObj.Cooperation, req.UserId,
						kind, apiErrCode, req.Point, toCoin, agentObj.Currency,
						orderId, req.Username, agentObj.Name, utils.ToMap(body))

					if apiErrCode != definition.ERROR_CODE_SUCCESS ||
						err != nil ||
						code != definition.ERROR_CODE_SUCCESS {
						c.OkWithCode(definition.INTERCOME_ERROR_CODE_FINISH_SINGLE_WALLET_LEDGER_FAILED)
						return
					}
				}
			}

			if resBodyData.Code > 0 {
				c.OkWithCode(definition.INTERCOME_ERROR_CODE_API_RESULT_CODE_FAILED)
				return
			}
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_API_RESULT_UNMARSHAL_FAILED)
			return
		}
	}

	dataMap := make(map[string]interface{}, 0)
	jj := utils.ToJSON(resBodyData.Data)
	_ = json.Unmarshal([]byte(jj), &dataMap)

	money := float64(0)
	moneyStr, ok := dataMap["money"].(string)
	if !ok {
		money = .0
	} else {
		money = utils.StringToFloat64(moneyStr, .0)
	}

	point := float64(0)
	pointStr, ok := dataMap["point"].(string)
	if !ok {
		point = .0
	} else {
		point = utils.StringToFloat64(pointStr, .0)
	}

	if req.Command != model.SingleWalletRequest_Command_QueryMinusScore {
		code, err := sqlutils.SWCoinInOutFinish(c.Logger(), c.DB(), agentObj.Id, agentObj.Cooperation, req.UserId,
			kind, apiErrCode, req.Point, toCoin, agentObj.Currency, orderId, req.Username, agentObj.Name, utils.ToMap(body))

		if apiErrCode != definition.ERROR_CODE_SUCCESS ||
			err != nil ||
			code != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(definition.INTERCOME_ERROR_CODE_FINISH_SINGLE_WALLET_LEDGER_FAILED)
			return
		}
	}

	res := model.NewEmptySingleWalletResponse()
	res.Command = req.Command
	res.Money = money
	res.Point = point
	res.WalletLedgerId = walletLedgerId
	c.Ok(res)
}

// @Tags IntercomApi
// @Summary recalculate game record
// @Description 此接口用來重新計算代理遊戲平台營收統計表(要做參數檢查)
// @Produce  application/json
// @Param data body model.RecalculateGAmeRecordRequest true "代理識別碼, 遊戲紀錄(json), 總變動分數..."
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/recalculategamerecord [post]
func (p *IntercomApi) RecalculateGameRecord(c *ginweb.Context) {

	var req model.RecalculateGAmeRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	query := `SELECT playlog, game_id, room_type, desk_id, bet_time 
			FROM play_log_common 
			WHERE bet_time>=$1 AND bet_time < $2`

	rows, err := c.DB().Query(query, req.StartTime, req.EndTime)
	if err != nil {
		c.Logger().Printf("RecalculateGameRecord() has error to select play_log_common, err  is: %v", err)
	}
	defer rows.Close()

	// map[logTime]map[id]
	bufCalPlayLogData := make(map[string]map[string]*sqlutils.PlayerStatData)

	timeBefore := time.Now()
	for rows.Next() {
		var dbPlaylog string
		var dbGameId, dbRoomType, dbDeskId int
		var betTime time.Time
		if err := rows.Scan(&dbPlaylog, &dbGameId, &dbRoomType, &dbDeskId, &betTime); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		playlogList, err := sqlutils.ParseUserPlayLog(dbPlaylog, dbGameId, dbRoomType, dbDeskId)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_FAIL)
			return
		}

		logTime := utils.Get15MinFormatString(betTime)
		dat, ok := bufCalPlayLogData[logTime]
		if !ok || dat == nil {
			dat = make(map[string]*sqlutils.PlayerStatData)
		}

		for _, playerStatData := range playlogList {
			if playerStatData.AgentId > 0 {
				id := fmt.Sprintf("%d_%d_%d_%d", playerStatData.AgentId, playerStatData.GameId, playerStatData.GameType, playerStatData.RoomType)
				tmp, ok := dat[id]
				if !ok {
					tmp = new(sqlutils.PlayerStatData)
					dat[id] = tmp
					tmp.AgentId = playerStatData.AgentId
					// tmp.BetId = playerStatData.BetId
					tmp.RoomType = playerStatData.RoomType
					// tmp.UserId = playerStatData.UserId
					// tmp.Username = playerStatData.Username
					tmp.GameId = playerStatData.GameId
					tmp.GameType = playerStatData.GameType
					tmp.LevelCode = playerStatData.LevelCode
				}
				tmp.BigWinCount += playerStatData.BigWinCount
				tmp.Bonus += playerStatData.Bonus
				tmp.Tax += playerStatData.Tax
				tmp.DeScore += playerStatData.DeScore
				tmp.YaScore += playerStatData.YaScore
				tmp.VaildYaScore += playerStatData.VaildYaScore
				tmp.OrderCount += playerStatData.OrderCount
				tmp.WinCount += playerStatData.WinCount
				tmp.LoseCount += playerStatData.LoseCount
			}
		}
		if len(dat) > 0 {
			bufCalPlayLogData[logTime] = dat
		}
		// dataCount++
	}
	timeAfter := time.Now()

	dataCount := 0

	for logTime, v := range bufCalPlayLogData {
		finalPlayLogData := make(map[int]*sqlutils.PlayerStatData, 0)
		for _, data := range v {
			finalPlayLogData[dataCount] = data
			dataCount++
		}

		bett, err := utils.GetUnsignedTimeUTCFromStr(logTime, "minute")
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_FAIL)
			return
		}
		sqlutils.CalRPAgentGameRatioStat(c.DB(), finalPlayLogData, bett)
	}

	log.Println("重新計算代理遊戲平台營收統計表中")
	log.Printf("startTime is : %v", req.StartTime)
	log.Printf("endTime is : %v", req.EndTime)
	log.Printf("dataCount is : %d", dataCount)
	log.Printf("spend time is : %v", timeAfter.Sub(timeBefore))

	c.Ok("success")
}

// @Tags IntercomApi
// @Summary create jackpot record
// @Description 此接口用來創建jackpot紀錄(不做參數檢查)
// @Produce  application/json
// @Param data body model.JackpotLogRequest true "代理識別碼, 遊戲紀錄(json), 總變動分數..."
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/intercom/createfriendroomrecord [post]
func (p *IntercomApi) CreateFriendRoomRecord(c *ginweb.Context) {
	var req model.FriendRoomLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_WRONG_FORMAT_FAILED)
		return
	}

	reqStr := utils.ToJSON(req)

	c.Logger().Info("Receive CreateFriendRoomRecord, req: %s", reqStr)

	if err := sqlutils.CreateFriendRoomLog(c.DB(), req); err != nil {
		c.OkWithCode(definition.INTERCOME_ERROR_CODE_CREATE_FRIEND_ROOM_LOG_FAILED)
		return
	}

	c.Logger().Info("Receive CreateFriendRoomRecord success, req: %s", reqStr)

	c.Ok("success")
}
