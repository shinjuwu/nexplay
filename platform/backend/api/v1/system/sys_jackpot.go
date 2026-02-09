package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemJackpotApi struct {
	BasePath string
}

func NewSystemJackpotApi(basePath string) api_cluster.IApiEach {
	return &SystemJackpotApi{
		BasePath: basePath,
	}
}

func (p *SystemJackpotApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemJackpotApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))

	g.POST("/getagentjackpotlist", ginHandler.Handle(p.GetAgentJackpotList))
	g.POST("/getagentjackpot", ginHandler.Handle(p.GetAgentJackpot))
	g.POST("/setagentjackpot",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentJackpot),
	)
	g.GET("/getjackpotsetting", ginHandler.Handle(p.GetJackpotSetting))
	g.POST("/setjackpotsetting",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetJackpotSetting),
	)
	g.POST("/getjackpottokenlist", ginHandler.Handle(p.GetJackpotTokenList))
	g.POST("/createjackpottoken",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.CreateJackpotToken),
	)
	g.POST("/getjackpotlist", ginHandler.Handle(p.GetJackpotList))
	g.GET("/getjackpotpooldata", ginHandler.Handle(p.GetJackpotPoolData))
	g.POST("/notifygameserveragentjackpotinfo",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.NotifyGameServerAgentJackpotInfo),
	)
	g.POST("/getjackpotleaderboard", ginHandler.Handle(p.GetJackpotLeaderboard))
}

// @Tags JP功能管理
// @Summary get gernal agent jackpot list
// @Description 取得總代理jackpot設定列表(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentJackpotListRequest true "取得總代理jackpot列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentJackpotListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getagentjackpotlist [post]
func (p *SystemJackpotApi) GetAgentJackpotList(c *ginweb.Context) {
	var req model.GetAgentJackpotListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"LENGTH(level_code)": 8})
	sqAnd = append(sqAnd, sq.Eq{"is_enabled": definition.STATE_TYPE_ENABLED})
	if req.AgentId > definition.AGENT_ID_ALL {
		sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"id", "name", "child_agent_count", "jackpot_status", "jackpot_start_time",
			"jackpot_end_time",
		).
		From("agent").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetAgentJackpotListResponse, 0)
	for rows.Next() {
		temp := new(model.GetAgentJackpotListResponse)

		if err := rows.Scan(&temp.Id, &temp.Name, &temp.ChildAgentCount, &temp.JackpotStatus, &temp.JackpotStartTime,
			&temp.JackpotEndTime); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		resp = append(resp, temp)
	}

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary get gernal agent jackpot
// @Description 取得指定總代理jackpot設定(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetAgentJackpotListRequest true "取得指定總代理jackpot參數"
// @Success 200 {object} response.Response{data=model.GetAgentJackpotListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getagentjackpot [post]
func (p *SystemJackpotApi) GetAgentJackpot(c *ginweb.Context) {
	var req model.GetAgentJackpotListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"LENGTH(level_code)": 8})
	sqAnd = append(sqAnd, sq.Eq{"is_enabled": definition.STATE_TYPE_ENABLED})
	sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"id", "name", "child_agent_count", "jackpot_status", "jackpot_start_time",
			"jackpot_end_time",
		).
		From("agent").
		Where(sqAnd).
		ToSql()

	var resp model.GetAgentJackpotListResponse
	err := c.DB().QueryRow(query, args...).Scan(&resp.Id, &resp.Name, &resp.ChildAgentCount, &resp.JackpotStatus, &resp.JackpotStartTime, &resp.JackpotEndTime)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary set gernal agent jackpot
// @Description 設定總代理jackpot(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.SetAgentJackpotRequest true "設定總代理jackpot參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/setagentjackpot [post]
func (p *SystemJackpotApi) SetAgentJackpot(c *ginweb.Context) {
	// 檢查參數
	var req model.SetAgentJackpotRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// 檢查權限
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 取得DB原始資料
	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})
	sqAnd = append(sqAnd, sq.Eq{"LENGTH(level_code)": 8})

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "name", "jackpot_status", "jackpot_start_time", "jackpot_end_time").
		From("agent").
		Where(sqAnd).
		ToSql()

	res := &model.GetAgentJackpotListResponse{}
	err := c.DB().QueryRow(query, args...).Scan(&res.Id, &res.Name, &res.JackpotStatus, &res.JackpotStartTime, &res.JackpotEndTime)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	// 資料無異動直接回傳成功
	if res.JackpotStatus == req.JackpotStatus &&
		res.JackpotStartTime.Equal(req.JackpotStartTime) &&
		res.JackpotEndTime.Equal(req.JackpotEndTime) {
		c.Ok("")
		return
	}

	// 時間有異動才通知game server
	if !res.JackpotStartTime.Equal(req.JackpotStartTime) ||
		!res.JackpotEndTime.Equal(req.JackpotEndTime) {
		start := req.JackpotStartTime.Unix()
		end := req.JackpotEndTime.Unix()

		agentIds := make([]int, 0)
		agentIds = append(agentIds, req.AgentId)

		starts := make([]int64, 0)
		starts = append(starts, start)

		ends := make([]int64, 0)
		ends = append(ends, end)

		for _, childAgent := range global.AgentCache.GetChildAgents(req.AgentId) {
			agentIds = append(agentIds, childAgent.Id)
			starts = append(starts, start)
			ends = append(ends, end)
		}

		// 通知game server設定代理jackpot設定
		if apiResult, errCode, err := notification.SendJapotInfo(agentIds, starts, ends); errCode != definition.ERROR_CODE_SUCCESS {
			c.Logger().Info("set agent jackpot notification.SendJapotInfo fail, agentIds=%v, starts=%v, ends=%v, resp=%s, resp code=%d, err=%v", agentIds, starts, ends, apiResult, errCode, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	// 修改DB
	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("agent").
		Set("jackpot_status", req.JackpotStatus).
		Set("jackpot_start_time", req.JackpotStartTime).
		Set("jackpot_end_time", req.JackpotEndTime).
		Where(sq.Or{
			sq.Eq{"id": req.AgentId},
			sq.Eq{"top_agent_id": req.AgentId},
		}).
		ToSql()
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_EXEC)
		return
	}

	// 修改cache
	agentObjs := make([]*table_model.Agent, 0)
	agentObjs = append(agentObjs, global.AgentCache.Get(req.AgentId))

	childObjs := global.AgentCache.GetChildAgents(req.AgentId)
	if len(childObjs) > 0 {
		agentObjs = append(agentObjs, childObjs...)
	}

	for i := 0; i < len(agentObjs); i++ {
		agentObjs[i].JackpotStatus = req.JackpotStatus
		agentObjs[i].JackpotStartTime = req.JackpotStartTime
		agentObjs[i].JackpotEndTime = req.JackpotEndTime

		// jackpot時間只要有修改就要踢除目前在線的使用者
		for _, adminUser := range global.AdminUserCache.GetAgentAdminUsers(agentObjs[i].Id) {
			c.Jwt.AddBlackTokenByUsername(adminUser.Username)
		}
	}

	global.AgentCache.Adds(agentObjs)

	// 操作紀錄
	actionLog := createBackendActionLogWithTitle(res.Name)
	if res.JackpotStatus != req.JackpotStatus {
		actionLog["jackpot_status"] = createBancendActionLogDetail(res.JackpotStatus, req.JackpotStatus)
	}
	if res.JackpotStartTime != req.JackpotStartTime {
		actionLog["jackpot_start_time"] = createBancendActionLogDetail(res.JackpotStartTime, req.JackpotStartTime)
	}
	if res.JackpotEndTime != req.JackpotEndTime {
		actionLog["jackpot_end_time"] = createBancendActionLogDetail(res.JackpotEndTime, req.JackpotEndTime)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags JP功能管理
// @Summary get systen jackpot setting
// @Description 取得平台jackpot設定(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=table_model.JackpotSetting,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getjackpotsetting [get]
func (p *SystemJackpotApi) GetJackpotSetting(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	storage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_JACKPOT_SETTING)
	if !ok {
		c.Logger().Error("db storage %s not exist", definition.STORAGE_KEY_JACKPOT_SETTING)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	var resp table_model.JackpotSetting
	json.Unmarshal([]byte(storage.Value), &resp)

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary set systen jackpot setting
// @Description 設定平台jackpot設定(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body table_model.JackpotSetting true "設定jackpot總開關參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/setjackpotsetting [post]
func (p *SystemJackpotApi) SetJackpotSetting(c *ginweb.Context) {
	// 檢查參數
	var req table_model.JackpotSetting
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查權限
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 取得DB原始資料
	storage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_JACKPOT_SETTING)
	if !ok {
		c.Logger().Error("db storage %s not exist", definition.STORAGE_KEY_JACKPOT_SETTING)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	var res table_model.JackpotSetting
	json.Unmarshal([]byte(storage.Value), &res)

	// 資料無異動直接回傳成功
	if res.JackpotSwitch == req.JackpotSwitch {
		c.Ok("")
		return
	}

	if req.JackpotSwitch {
		if code, err := notification.GetJpOn(); err != nil {
			c.Logger().Info("SetJackpotSetting GetJpOn failed. code: %d, err: %v", code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	} else {
		if code, err := notification.GetJpOff(); err != nil {
			c.Logger().Info("SetJackpotSetting GetJpOff failed. code: %d, err: %v", code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	// 修改DB
	if err := global.GlobalStorage.Update(definition.STORAGE_KEY_JACKPOT_SETTING, utils.ToJSON(req)); err != nil {
		c.Logger().Error("SetJackpotSetting STORAGE_KEY_JACKPOT_SETTING update has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	if res.JackpotSwitch != req.JackpotSwitch {
		actionLog["jackpot_switch"] = createBancendActionLogDetail(res.JackpotSwitch, req.JackpotSwitch)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags JP功能管理
// @Summary get jackpot token list
// @Description 取得jackpot代幣紀錄列表(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetJackpotTokenListRequest true "取得jackpot代幣紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetJackpotTokenListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getjackpottokenlist [post]
func (p *SystemJackpotApi) GetJackpotTokenList(c *ginweb.Context) {
	// 檢查參數
	var req model.GetJackpotTokenListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(
		serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange*24,
		serverInfoSetting.CommonReportTimeBeforeDays,
	); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// 檢查權限
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))
	startTime := today.Add(-time.Duration(serverInfoSetting.CommonReportTimeBeforeDays) * 24 * time.Hour)
	endTime := time.Now().UTC()

	// 必要條件檢查
	checkExist_tokenId := (req.TokenId != "")
	checkExist_lognumber := (req.Lognumber != "")
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_userName := (req.Username != "")

	sqAnd := sq.And{}
	if checkExist_tokenId {
		sqAnd = append(sqAnd, sq.Eq{"token_id": req.TokenId})
	} else if checkExist_lognumber {
		sqAnd = append(sqAnd, sq.Eq{"source_lognumber": req.Lognumber})
	} else {
		if checkExist_agentId {
			sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
		}
		if checkExist_userName {
			sqAnd = append(sqAnd, sq.Eq{"username": req.Username})
		}

		startTime = req.StartTime
		endTime = req.EndTime
	}
	sqAnd = append(sqAnd, sq.GtOrEq{"token_create_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"token_create_time": endTime})

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"id", "token_id", "agent_id", "username", "source_lognumber",
			"source_bet_id", "jp_bet", "usage_count", "creator", "info",
			"token_create_time", "status", "error_code",
		).
		From("jackpot_token_log").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetJackpotTokenListResponse, 0)

	tmpAgents := make(map[int]*table_model.Agent)
	for rows.Next() {
		var tmp model.GetJackpotTokenListResponse

		if err := rows.Scan(&tmp.Id, &tmp.TokenId, &tmp.AgentId, &tmp.Username, &tmp.SourceLognumber,
			&tmp.SourceBetId, &tmp.JpBet, &tmp.UsageCount, &tmp.Creator, &tmp.Info,
			&tmp.TokenCreateTime, &tmp.Status, &tmp.ErrorCode); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, &tmp)
	}

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary create jackpot token
// @Description 建立jackpot代幣(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.CreateJackpotTokenRequest true "建立jackpot代幣參數"
// @Success 200 {object} response.Response{data=[]model.GetJackpotTokenListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/createjackpottoken [post]
func (p *SystemJackpotApi) CreateJackpotToken(c *ginweb.Context) {
	// 檢查參數
	var req model.CreateJackpotTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	agent := global.AgentCache.Get(req.AgentId)
	if agent == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查權限
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id").
		From("game_users").
		Where(sq.And{
			sq.Eq{"agent_id": req.AgentId},
			sq.Eq{"original_username": req.Username},
		}).
		ToSql()

	userId := 0
	err := c.DB().QueryRow(query, args...).Scan(&userId)
	if err != nil {
		code := definition.ERROR_CODE_ERROR_DATABASE
		if err == sql.ErrNoRows {
			code = definition.ERROR_CODE_ERROR_ACCOUNT_NOT_EXIST
		}

		c.OkWithCode(code)
		return
	}

	tokenCreateTime := time.Now().UTC()
	creator := c.GetUsername()
	status := definition.JACKPOT_TOKEN_LOG_STATUS_CREATED

	// agentId, username, jpBet, info, creator, time.Now().UnixMilli()
	salt := fmt.Sprintf("%d_%s_%f_%s_%s_%d", req.AgentId, req.Username, req.JpBet, req.Info, creator, tokenCreateTime.UnixMilli())
	tokenId := utils.CreatreOrderIdByOrderTypeAndSalt(definition.ORDER_TYPE_JACKPOT_TOKEN_LOG, salt, tokenCreateTime)

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("jackpot_token_log").
		Columns(
			"token_id", "agent_id", "level_code", "user_id", "username",
			"jp_bet", "token_create_time", "status", "creator", "info",
		).
		Values(
			tokenId, req.AgentId, agent.LevelCode, userId, req.Username,
			req.JpBet, tokenCreateTime, status, creator, req.Info,
		).
		ToSql()
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 通知game server建立token
	var apiResult string
	var apiErrorCode int
	apiResult, apiErrorCode, err = notification.SendJapotAddCoin([]string{utils.IntToString(userId)}, []string{tokenId}, []float64{req.JpBet}, []int64{tokenCreateTime.Unix()})
	if err != nil {
		c.Logger().Info("CreateJackpotToken notification.SendJapotAddCoin fail, userId=%d, tokenId=%s, jpBet=%f, createTime=%d resp=%s, resp code=%d, err=%v", userId, tokenId, req.JpBet, tokenCreateTime.Unix(), apiResult, apiErrorCode, err)
	}

	errorCode := notification.TransformApiCodeToModelErrorCode(apiErrorCode)
	status = definition.JACKPOT_TOKEN_LOG_STATUS_SUCCESS

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("jackpot_token_log").
		SetMap(sq.Eq{
			"status":      status,
			"error_code":  errorCode,
			"update_time": sq.Expr("now()"),
		}).
		Where(sq.Eq{"token_id": tokenId}).
		ToSql()
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	c.Set("action_log", map[string]interface{}{
		"agent_name": agent.Name,
		"username":   req.Username,
		"token_id":   tokenId,
	})

	c.Ok("")
}

// @Tags JP功能管理
// @Summary get jackpot list
// @Description 取得jackpot紀錄列表
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetJackpotListRequest true "取得jackpot紀錄列表參數"
// @Success 200 {object} response.Response{data=[]model.GetJackpotListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getjackpotlist [post]
func (p *SystemJackpotApi) GetJackpotList(c *ginweb.Context) {
	// 檢查權限
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
	userAgent := global.AgentCache.Get(int(userClaims.BaseClaims.ID))

	// 營運商開發商依照權限檢查，總代理及子代理則看是否再jackpot開放期間
	if userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN {
		userAgentPermission := global.AgentPermissionCache.Get(userClaims.PermissionId)

		hasPermission := false
		for _, permission := range userAgentPermission.Permission.List {
			if permission == definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LIST {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}
	} else {
		now := time.Now().UTC()
		if userAgent.JackpotStartTime.After(now) || userAgent.JackpotEndTime.Before(now) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}
	}

	// 檢查參數
	var req model.GetJackpotListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(
		serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange*24,
		serverInfoSetting.CommonReportTimeBeforeDays,
	); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	today := utils.GetTimeNowUTCTodayTime().Add(time.Duration(req.TimezoneOffset * int(time.Minute)))
	startTime := today.Add(-time.Duration(serverInfoSetting.CommonReportTimeBeforeDays) * 24 * time.Hour)
	endTime := time.Now().UTC()

	// 必要條件檢查
	checkExist_tokenId := (req.TokenId != "")
	checkExist_lognumber := (req.Lognumber != "")
	checkExist_agentId := (req.AgentId > definition.AGENT_ID_ALL)
	checkExist_username := (req.Username != "")
	isAdmin := (userClaims.AccountType == definition.ACCOUNT_TYPE_ADMIN)

	sqAnd := sq.And{}
	if checkExist_tokenId {
		sqAnd = append(sqAnd, sq.Eq{"token_id": req.TokenId})
		if !isAdmin {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else if checkExist_lognumber {
		sqAnd = append(sqAnd, sq.Eq{"lognumber": req.Lognumber})
		if !isAdmin {
			sqAnd = append(sqAnd, sq.Or{
				sq.Like{"level_code": userAgent.LevelCode + "%"},
			})
		}
	} else {
		// 非自己或下層級的代理不能查詢
		if checkExist_agentId {
			agent := global.AgentCache.Get(req.AgentId)
			if agent == nil || !strings.HasPrefix(agent.LevelCode, userAgent.LevelCode) {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}
			sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
		} else if !isAdmin {
			sqAnd = append(sqAnd, sq.Like{"level_code": userAgent.LevelCode + "%"})
		}

		if checkExist_username {
			sqAnd = append(sqAnd, sq.Eq{"username": req.Username})
		}

		startTime = req.StartTime
		endTime = req.EndTime
	}
	sqAnd = append(sqAnd, sq.GtOrEq{"winning_time": startTime})
	sqAnd = append(sqAnd, sq.Lt{"winning_time": endTime})

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"bet_id", "lognumber", "token_id", "agent_id", "username",
			"jp_bet", "token_create_time", "prize_score", "prize_item", "winning_time",
			"show_pool", "real_pool", "is_robot",
		).
		From("jackpot_log").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	resp := make([]*model.GetJackpotListResponse, 0)

	tmpAgents := make(map[int]*table_model.Agent)
	for rows.Next() {
		var tmp model.GetJackpotListResponse

		if err := rows.Scan(&tmp.BetId, &tmp.Lognumber, &tmp.TokenId, &tmp.AgentId, &tmp.Username,
			&tmp.JpBet, &tmp.TokenCreateTime, &tmp.PrizeScore, &tmp.PrizeItem, &tmp.WinningTime,
			&tmp.ShowPool, &tmp.RealPool, &tmp.IsRobot); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, &tmp)
	}

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary get jackpot pool data
// @Description 取得jackpot獎池資訊(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.GetJackpotPoolDataResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getjackpotpooldata [get]
func (p *SystemJackpotApi) GetJackpotPoolData(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	realPool, showPool, reservePool, jpShowPool, jpInjectWaterRate, code, err := notification.GetJpPool()
	if err != nil {
		c.Logger().Info("GetJackpotPoolData GetJpPool failed. code: %d, err: %v", code, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	resp := &model.GetJackpotPoolDataResponse{}
	resp.RealPool = realPool
	resp.ShowPool = showPool
	resp.ReservePool = reservePool
	resp.JpShowPool = jpShowPool
	resp.JpInjectWaterRate = jpInjectWaterRate

	c.Ok(resp)
}

// @Tags JP功能管理
// @Summary send current all agent jackpot info to game server
// @Description server同步jackpot資訊(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/notifygameserveragentjackpotinfo [post]
func (p *SystemJackpotApi) NotifyGameServerAgentJackpotInfo(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd,
		sq.GtOrEq{"LENGTH(level_code)": 8},
		sq.Eq{"is_enabled": definition.STATE_TYPE_ENABLED},
	)

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "jackpot_start_time", "jackpot_end_time").
		From("agent").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	agentIds := make([]int, 0)
	starts := make([]int64, 0)
	ends := make([]int64, 0)

	for rows.Next() {
		var id int
		var start time.Time
		var end time.Time

		if err := rows.Scan(&id, &start, &end); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		agentIds = append(agentIds, id)
		starts = append(starts, start.Unix())
		ends = append(ends, end.Unix())
	}

	// 通知game server設定代理jackpot設定
	if apiResult, errCode, err := notification.SendJapotInfo(agentIds, starts, ends); errCode != definition.ERROR_CODE_SUCCESS {
		c.Logger().Info("set agent jackpot notification.SendJapotInfo fail, agentIds=%v, starts=%v, ends=%v, resp=%s, resp code=%d, err=%v", agentIds, starts, ends, apiResult, errCode, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	c.Ok("")
}

// @Tags JP功能管理
// @Summary get jackpot game user leaderboard
// @Description 取得jackpot玩家貢獻度(只有開發商(營運商)可使用)
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.GetJackpotLeaderboardRequest true "取得jackpot玩家貢獻度參數"
// @Success 200 {object} response.Response{data=[]model.GetJackpotLeaderboardResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/jackpot/getjackpotleaderboard [post]
func (p *SystemJackpotApi) GetJackpotLeaderboard(c *ginweb.Context) {
	// 檢查參數
	var req model.GetJackpotLeaderboardRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// 檢查權限
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	leaderboard, code, err := notification.GetCRank()
	if err != nil {
		c.Logger().Info("GetJackpotLeaderboard GetCRank failed. code: %d, err: %v", code, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	userIds := make([]int, 0)
	leaderboardMap := make(map[int]map[string]interface{})
	for _, data := range leaderboard {
		userId := utils.ToInt(data["UserID"].(string))
		userIds = append(userIds, userId)
		leaderboardMap[userId] = data
	}

	resp := make([]*model.GetJackpotLeaderboardResponse, 0)
	if len(userIds) == 0 {
		c.Ok(resp)
		return
	}

	sqAnd := sq.And{}
	sqAnd = append(sqAnd, sq.Eq{"gu.id": userIds})
	if req.AgentId > definition.AGENT_ID_ALL {
		sqAnd = append(sqAnd, sq.Eq{"gu.agent_id": req.AgentId})
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("a.id", "a.name", "gu.id", "gu.original_username").
		From("game_users AS gu").
		InnerJoin("agent AS a ON gu.agent_id = a.id").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tmp model.GetJackpotLeaderboardResponse

		if err := rows.Scan(&tmp.AgentId, &tmp.AgentName, &tmp.UserId, &tmp.Username); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		data := leaderboardMap[tmp.UserId]
		tmp.PlayNum = int(data["PlayNum"].(float64))
		tmp.TotalBet = data["TotalBet"].(float64)
		tmp.Win = int(data["Win"].(float64))
		tmp.Lose = int(data["Lose"].(float64))
		tmp.Score = data["Score"].(float64)

		resp = append(resp, &tmp)
	}

	c.Ok(resp)
}
