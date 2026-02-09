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
	"sort"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

/*
風控功能
base path: riskcontrol
*/

// 風控功能
type SystemRiskControlApi struct {
	BasePath string
}

func NewSystemRiskControlApi(basePath string) api_cluster.IApiEach {
	return &SystemRiskControlApi{
		BasePath: basePath,
	}
}

func (p *SystemRiskControlApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemRiskControlApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/getagentincomeratiolist", ginHandler.Handle(p.GetAgentIncomeRatioList))
	g.POST("/getagentincomeratio", ginHandler.Handle(p.GetAgentIncomeRatio))
	g.POST("/setagentincomeratio",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentIncomeRatio))
	g.POST("/getincomeratiolist", ginHandler.Handle(p.GetIncomeRatioList))
	g.POST("/getincomeratio", ginHandler.Handle(p.GetIncomeRatio))
	g.POST("/setincomeratio",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetIncomeRatio))
	g.POST("/setincomeratios",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetIncomeRatios))
	g.POST("/getagentincomeratioandgamedata", ginHandler.Handle(p.GetAgentIncomeRatioAndGameData))
	// Notice: 目前前端沒有頁面需求，所以api先註解起來 by chouyang
	// g.POST("/getplayerincomeratioandgamedata", ginHandler.Handle(p.GetPlayerIncomeRatioAndGameData))
	g.POST("/getagentcustomtagsettinglist", ginHandler.Handle(p.GetAgentCustomTagSettingList))
	g.POST("/setagentcustomtagsettinglist",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAgentCustomTagSettingList))
	g.POST("/getgameuserscustomtaglist", ginHandler.Handle(p.GetGameUsersCustomTagList))
	g.POST("/getgameuserscustomtag", ginHandler.Handle(p.GetGameUsersCustomTag))
	g.POST("/setgameuserscustomtag",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameUsersCustomTag))
	g.POST("/getautoriskcontrolsetting", ginHandler.Handle(p.GetAutoRiskControlSetting))
	g.POST("/setautoriskcontrolsetting",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetAutoRiskControlSetting))
	g.POST("/getgameuserriskcontroltag", ginHandler.Handle(p.GetGameUserRiskControlTag))
	g.POST("/setgameuserriskcontroltag",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameUserRiskControlTag))
	g.POST("/getgamesetting", ginHandler.Handle(p.GetGameSetting))
	g.POST("/setgamesetting",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameSetting))
	g.POST("/getrealtimegameratio", ginHandler.Handle(p.GetRealtimeGameRatio))
}

// @Tags 風控功能/總代理風控設定
// @Summary Get income ratio list by level code of agent
// @Description 此接口用來取得當前總代理風控設定資料列表(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentIncomeRatioListRequest true "取得總代理風控設定資料列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentIncomeRatioListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getagentincomeratiolist [post]
func (p *SystemRiskControlApi) GetAgentIncomeRatioList(c *ginweb.Context) {
	var req model.GetAgentIncomeRatioListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		sqAnd := sq.And{}

		sqAnd = append(sqAnd, sq.Eq{"LENGTH(level_code)": 8})
		sqAnd = append(sqAnd, sq.Eq{"is_enabled": definition.STATE_TYPE_ENABLED})

		if req.AgentId > definition.AGENT_ID_ALL {
			sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})
		}

		if req.StateType > definition.STATE_TYPE_ALL {
			st := false
			if req.StateType == definition.STATE_TYPE_ENABLED {
				st = true
			}
			sqAnd = append(sqAnd, sq.Eq{"kill_switch": st})
		}

		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("id", "name", "kill_switch", "kill_ratio", "kill_info", "kill_update_time").
			From("agent").
			Where(sqAnd).
			ToSql()

		rows, err := c.DB().Query(query, args...)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		defer rows.Close()

		resp := make([]*model.GetAgentIncomeRatioListResponse, 0)
		for rows.Next() {
			tmp := &model.GetAgentIncomeRatioListResponse{}
			if err = rows.Scan(&tmp.AgentId, &tmp.AgentName, &tmp.State, &tmp.Ratio, &tmp.Info,
				&tmp.UpdateTime); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}

			resp = append(resp, tmp)
		}

		c.Ok(resp)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/總代理風控設定
// @Summary Get income ratio by target id
// @Description 此接口用來取得指定id總代理風控設定資料(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentIncomeRatioRequest true "取得指定id總代理風控設定資料參數"
// @Success 200 {object} response.Response{data=model.GetAgentIncomeRatioResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getagentincomeratio [post]
func (p *SystemRiskControlApi) GetAgentIncomeRatio(c *ginweb.Context) {
	var req model.GetAgentIncomeRatioRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		sqAnd := sq.And{}
		sqAnd = append(sqAnd, sq.Eq{"LENGTH(level_code)": 8})
		sqAnd = append(sqAnd, sq.Eq{"is_enabled": definition.STATE_TYPE_ENABLED})
		sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})

		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("id", "name", "kill_switch", "kill_ratio", "kill_info", "kill_update_time").
			From("agent").
			Where(sqAnd).
			ToSql()

		res := &model.GetAgentIncomeRatioListResponse{}
		_ = c.DB().QueryRow(query, args...).Scan(&res.AgentId, &res.AgentName, &res.State, &res.Ratio, &res.Info, &res.UpdateTime)

		c.Ok(res)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/總代理風控設定
// @Summary Get income ratio by target id
// @Description 此接口用來設定指定id總代理風控設定資料(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetAgentIncomeRatioRequest true "設定總代理風控設定資料參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setagentincomeratio [post]
func (p *SystemRiskControlApi) SetAgentIncomeRatio(c *ginweb.Context) {
	var req model.SetAgentIncomeRatioRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		dbTargetLevelCode := ""
		isEnabled := 0
		if req.State {
			isEnabled = 1
		}

		sqAnd := sq.And{}

		sqAnd = append(sqAnd, sq.Eq{"id": req.AgentId})

		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("id", "level_code", "name", "kill_switch", "kill_ratio", "kill_info", "kill_update_time").
			From("agent").
			Where(sqAnd).
			ToSql()

		res := &model.GetAgentIncomeRatioListResponse{}
		_ = c.DB().QueryRow(query, args...).Scan(&res.AgentId, &dbTargetLevelCode, &res.AgentName, &res.State, &res.Ratio, &res.Info, &res.UpdateTime)

		if res.State != req.State || res.Ratio != req.Ratio {
			// 狀態改變、殺率改變,通知遊戲端,通知失敗就直接回傳錯誤
			if _, _, err := notification.SendSetKillAdminInfo(dbTargetLevelCode, req.Ratio, isEnabled); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
				return
			}
		}

		query = `UPDATE agent
				SET kill_switch = $1, kill_ratio = $2, kill_info = $3, kill_update_time= now()
				WHERE LENGTH(level_code)= 8 AND is_enabled = 1 AND id = $4`
		result, err := c.DB().Exec(query, req.State, req.Ratio, req.Info, req.AgentId)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_EXEC)
			return
		}
		if count, err := result.RowsAffected(); count != 1 {
			c.Logger().Printf("query exec failed,query = %v, err = %v", query, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
			return
		}

		agentObj := global.AgentCache.Get(req.AgentId)
		agentObj.KillSwitch = req.State
		agentObj.KillRatio = req.Ratio
		global.AgentCache.Add(agentObj)

		// 操作紀錄
		actionLog := createBackendActionLogWithTitle(res.AgentName)
		if res.State != req.State {
			actionLog["bool_status"] = createBancendActionLogDetail(res.State, req.State)
		}
		if res.Ratio != req.Ratio {
			actionLog["ratio"] = createBancendActionLogDetail(res.Ratio, req.Ratio)
		}
		if res.Info != req.Info {
			actionLog["info"] = createBancendActionLogDetail(res.Info, req.Info)
		}
		c.Set("action_log", actionLog)

		c.Ok("")
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/遊戲風控設定
// @Summary Get income ratio list by level code of agent
// @Description 此接口用來取得當前平台機率設定列表(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetIncomeRatioListRequest true "取得平台風控設定列表參數"
// @Success 200 {object} response.Response{data=[]model.GetIncomeRatioListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getincomeratiolist [post]
func (p *SystemRiskControlApi) GetIncomeRatioList(c *ginweb.Context) {
	var req model.GetIncomeRatioListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		agentGameRatioCache, err := global.AgentGameRatioCache.Select()
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		checkGameId := !(req.GameId == definition.GAME_ID_ALL)
		checkRoomType := !(req.RoomType == definition.ROOM_TYPE_ALL)
		checkAgentId := !(req.AgentId == definition.AGENT_ID_ALL)

		if req.GameType == definition.GAME_TYPE_BAIREN {
			checkAgentId = false
		}

		tmps := make([]*global.AgentGameRatio, 0)
		for _, tmp := range agentGameRatioCache {

			if tmp.GameType != req.GameType {
				continue
			}
			if checkGameId {
				if req.GameId != tmp.GameId {
					continue
				}
			}
			if checkRoomType {
				if req.RoomType != tmp.RoomType {
					continue
				}
			}
			if checkAgentId {
				if req.AgentId != tmp.AgentId {
					continue
				}
			}
			tmps = append(tmps, tmp)
		}

		resp := make([]*model.GetIncomeRatioListResponse, 0)

		for _, val := range tmps {
			res := &model.GetIncomeRatioListResponse{}
			res.Id = val.Id
			res.AgentId = val.AgentId
			res.RoomType = val.RoomType
			res.GameId = val.GameId
			res.GameType = val.GameType
			res.KillRatio = val.KillRatio
			res.NewKillRatio = val.NewKillRatio
			res.ActiveNum = val.ActiveNum
			res.LastUpdateTime = val.UpdateTime
			res.Info = val.Info

			gameCache := global.GameCache.Get(val.GameId)
			if gameCache == nil {
				continue
			}
			res.GameCode = gameCache.Code

			if req.GameType != definition.GAME_TYPE_BAIREN {
				agentCache := global.AgentCache.Get(val.AgentId)
				if agentCache == nil {
					continue
				}

				generalAgent := agentCache
				if len(agentCache.LevelCode) > 8 {
					generalAgentId, _ := strconv.ParseInt(agentCache.LevelCode[4:8], 16, 32)
					generalAgent = global.AgentCache.Get(int(generalAgentId))
					if generalAgent == nil {
						continue
					}
				}

				res.IsParent = generalAgent.KillSwitch
				res.AgnetName = agentCache.Name
			}

			resp = append(resp, res)
		}

		c.Ok(resp)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/遊戲風控設定
// @Summary Get income ratio by target id
// @Description 此接口用來取得指定id平台機率設定資料(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetIncomeRatioRequest true "取得平台機率設定列表參數"
// @Success 200 {object} response.Response{data=model.GetIncomeRatioResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getincomeratio [post]
func (p *SystemRiskControlApi) GetIncomeRatio(c *ginweb.Context) {
	var req model.GetIncomeRatioRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		key := global.AgentGameRatioCache.GetKey(req.AgentId, req.GameId, req.GameType, req.RoomType)

		agentGameRatioCache, ok := global.AgentGameRatioCache.SelectOne(key)
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		// trans values
		res := &model.GetIncomeRatioResponse{}
		res.Id = agentGameRatioCache.Id
		res.AgentId = agentGameRatioCache.AgentId
		res.KillRatio = agentGameRatioCache.KillRatio
		res.NewKillRatio = agentGameRatioCache.NewKillRatio
		res.ActiveNum = agentGameRatioCache.ActiveNum
		res.RoomType = agentGameRatioCache.RoomType
		res.GameId = agentGameRatioCache.GameId
		res.GameType = agentGameRatioCache.GameType
		res.LastUpdateTime = agentGameRatioCache.UpdateTime
		res.Info = agentGameRatioCache.Info

		gameCache := global.GameCache.Get(agentGameRatioCache.GameId)
		if gameCache != nil {
			res.GameCode = gameCache.Code
		}

		agentCache := global.AgentCache.Get(agentGameRatioCache.AgentId)
		if agentCache != nil {
			res.AgnetName = agentCache.Name

			generalAgent := agentCache
			if len(agentCache.LevelCode) > 8 {
				generalAgentId, _ := strconv.ParseInt(agentCache.LevelCode[4:8], 16, 32)
				generalAgent = global.AgentCache.Get(int(generalAgentId))
			}
			if generalAgent != nil {
				res.IsParent = generalAgent.KillSwitch
			}
		}

		c.Ok(res)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/遊戲風控設定
// @Summary Get income ratio by target id
// @Description 此接口用來設定指定id平台機率設定資料(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetIncomeRatioRequest true "設定平台機率設定列表參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setincomeratio [post]
func (p *SystemRiskControlApi) SetIncomeRatio(c *ginweb.Context) {
	var req model.SetIncomeRatioRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有管理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		if !isAdmin {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		agentGameRatioCache, ok := global.AgentGameRatioCache.SelectOne(req.Id)
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		/*
			當百人場時，AgentId請送-1，
			以代表設定不分代理

			當對戰場及電子遊戲時，上水線和下水線請送0，
			以代表此類型無法設定

			當Server接收後並設定成功，
			會回傳所有的殺放設定(類似setlobbyinfo)
		*/
		sendAgentId := agentGameRatioCache.AgentId
		if agentGameRatioCache.GameType == definition.GAME_TYPE_BAIREN {
			sendAgentId = definition.AGENT_ID_UNKNOW
		}
		backupKillRatio := agentGameRatioCache.KillRatio
		backupNewKillRatio := agentGameRatioCache.NewKillRatio
		backupActiveNum := agentGameRatioCache.ActiveNum
		backupInfo := agentGameRatioCache.Info

		// TO DO: 追蹤Bug問題加入的Log，問題解決後要移除
		roomId, _ := getRoomId(agentGameRatioCache.GameId, agentGameRatioCache.RoomType)
		c.Logger().Info("[代理殺放追蹤]遊戲同步代理設置殺放，殺放內容: %v", getKillDiveInfo(
			sendAgentId,
			agentGameRatioCache.GameId,
			roomId,
			req.ActiveNum,
			req.KillRatio,
			req.NewKillRatio,
		))
		apiResult, code, err := notification.SendSetKillDive(
			int64(sendAgentId),
			int64(agentGameRatioCache.GameId),
			int64(agentGameRatioCache.RoomType),
			req.KillRatio,
			req.NewKillRatio,
			req.ActiveNum)
		if err != nil {
			c.Logger().Printf("SetIncomeRatio() has error in notification.SendSetKillDive():result is %s, code is %d, err is %v", apiResult, code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_API_SERVER_REQUEST_FAILED)
			return
		}

		if err := global.AgentGameRatioCache.Update(
			req.Id,
			req.KillRatio,
			req.NewKillRatio,
			req.ActiveNum,
			req.Info); err != nil {
			c.Logger().Printf("SetIncomeRatio() has error in AgentGameRatioCache.Update():err is %v", err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_API_SERVER_REQUEST_FAILED)
			return
		}

		sendAgent := global.AgentCache.Get(sendAgentId)

		actionLog := make(map[string]interface{})
		actionLog["game_id"] = agentGameRatioCache.GameId
		actionLog["room_type"] = agentGameRatioCache.RoomType
		if sendAgent != nil {
			actionLog["agent_name"] = sendAgent.Name
		}
		if backupKillRatio != req.KillRatio {
			actionLog["ratio"] = createBancendActionLogDetail(backupKillRatio, req.KillRatio)
		}
		if backupNewKillRatio != req.NewKillRatio {
			actionLog["new_ratio"] = createBancendActionLogDetail(backupNewKillRatio, req.NewKillRatio)
		}
		if backupActiveNum != req.ActiveNum {
			actionLog["active_num"] = createBancendActionLogDetail(backupActiveNum, req.ActiveNum)
		}
		if backupInfo != req.Info {
			actionLog["info"] = createBancendActionLogDetail(backupInfo, req.Info)
		}
		c.Set("action_log", actionLog)

		c.Ok("")
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/遊戲風控設定
// @Summary Get income ratio by target id
// @Description 此接口用來批次設定殺數設定資料(只有管理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetIncomeRatiosRequest true "設定平台機率設定列表參數"
// @Success 200 {object} response.Response{data=model.SetIncomeRatiosResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setincomeratios [post]
func (p *SystemRiskControlApi) SetIncomeRatios(c *ginweb.Context) {
	var req model.SetIncomeRatiosRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 只有管理可以修改
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) > 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	loopAgentGameRatio := func(loopFunc func(agentId, gameId, roomType int, agentGameRatio *global.AgentGameRatio)) {
		for _, agentId := range req.AgentIdList {
			for _, gameId := range req.GameIdList {
				for _, roomType := range req.RoomTypeList {
					gameType := gameId / 1000

					agentGameRatioKey := global.AgentGameRatioCache.GetKey(agentId, gameId, gameType, roomType)
					agentGameRatio, _ := global.AgentGameRatioCache.SelectOne(agentGameRatioKey)

					loopFunc(agentId, gameId, roomType, agentGameRatio)
				}
			}
		}
	}

	resp := new(model.SetIncomeRatiosResponse)
	resp.ErrorList = make([]map[string]int, 0)

	sendKillDiveInfo := make([]map[string]interface{}, 0)
	actionLogAgentGameRatios := make([]map[string]interface{}, 0)
	updateDbKeys := make([]string, 0)

	loopAgentGameRatio(func(agentId, gameId, roomType int, agentGameRatio *global.AgentGameRatio) {
		if agentGameRatio == nil {
			resp.ErrorList = append(resp.ErrorList, map[string]int{
				"agent_id":  agentId,
				"game_id":   gameId,
				"room_type": roomType,
			})
			return
		}

		needSendGameServer := agentGameRatio.KillRatio != req.KillRatio ||
			agentGameRatio.NewKillRatio != req.NewKillRatio ||
			agentGameRatio.ActiveNum != req.ActiveNum
		if needSendGameServer {
			roomId, _ := getRoomId(agentGameRatio.GameId, agentGameRatio.RoomType)
			sendKillDiveInfo = append(sendKillDiveInfo, getKillDiveInfo(agentGameRatio.AgentId, agentGameRatio.GameId, roomId, req.ActiveNum, req.KillRatio, req.NewKillRatio))
		}

		needUpdateDB := needSendGameServer ||
			agentGameRatio.Info != req.Info
		if needUpdateDB {
			updateDbKeys = append(updateDbKeys, agentGameRatio.Id)

			actionLogAgentGameRatio := make(map[string]interface{}, 0)
			actionLogAgentGameRatio["game_id"] = gameId
			actionLogAgentGameRatio["room_type"] = roomType

			if agent := global.AgentCache.Get(agentId); agent != nil {
				actionLogAgentGameRatio["agent_name"] = agent.Name
			}
			if agentGameRatio.KillRatio != req.KillRatio {
				actionLogAgentGameRatio["ratio"] = createBancendActionLogDetail(agentGameRatio.KillRatio, req.KillRatio)
			}
			if agentGameRatio.NewKillRatio != req.NewKillRatio {
				actionLogAgentGameRatio["new_ratio"] = createBancendActionLogDetail(agentGameRatio.NewKillRatio, req.NewKillRatio)
			}
			if agentGameRatio.ActiveNum != req.ActiveNum {
				actionLogAgentGameRatio["active_num"] = createBancendActionLogDetail(agentGameRatio.ActiveNum, req.ActiveNum)
			}
			if agentGameRatio.Info != req.Info {
				actionLogAgentGameRatio["info"] = createBancendActionLogDetail(agentGameRatio.Info, req.Info)
			}

			actionLogAgentGameRatios = append(actionLogAgentGameRatios, actionLogAgentGameRatio)
		}
	})

	if len(resp.ErrorList) > 0 {
		c.OkWithCodeAndData(definition.ERROR_CODE_ERROR_INCOME_RATIO_DATA, resp.ErrorList)
		return
	}

	if len(updateDbKeys) == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_NO_CHANGE_IN_DATA)
		return
	}

	if len(sendKillDiveInfo) > 0 {
		batch := 5000
		curSendStart := 0
		curSendEnd := batch
		for curSendStart < len(sendKillDiveInfo) {
			if curSendEnd > len(sendKillDiveInfo) {
				curSendEnd = len(sendKillDiveInfo)
			}

			// TO DO: 追蹤Bug問題加入的Log，問題解決後要移除
			c.Logger().Info("[代理殺放追蹤]遊戲風控設定批次設置殺放，殺放內容: %v", sendKillDiveInfo[curSendStart:curSendEnd])

			_, errorCode, err := notification.SendSetKillDives(sendKillDiveInfo[curSendStart:curSendEnd])
			if err != nil {
				c.Logger().Info("notification.SendSetKillDives fail, killDiveInfo=%v, resp code=%d, err=%v", sendKillDiveInfo[curSendStart:curSendEnd], errorCode, err)
				c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
				return
			}

			curSendStart += batch
			curSendEnd += batch
		}
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("agent_game_ratio").
		SetMap(sq.Eq{
			"kill_ratio":     req.KillRatio,
			"new_kill_ratio": req.NewKillRatio,
			"active_num":     req.ActiveNum,
			"info":           req.Info,
			"update_time":    sq.Expr("now()"),
		}).
		Where(sq.Eq{"id": updateDbKeys}).
		ToSql()
	_, err := c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["agent_game_ratios"] = actionLogAgentGameRatios
	c.Set("action_log", actionLog)

	loopAgentGameRatio(func(agentId, gameId, roomType int, agentGameRatio *global.AgentGameRatio) {
		if agentGameRatio == nil {
			return
		}

		agentGameRatio.KillRatio = req.KillRatio
		agentGameRatio.NewKillRatio = req.NewKillRatio
		agentGameRatio.ActiveNum = req.ActiveNum
		agentGameRatio.Info = req.Info
	})

	c.Ok(resp.ErrorList)
}

// @Tags 風控功能/代理風控統計
// @Summary Get the agent sets the probability & the outcome of the game
// @Description 此接口用來取得代理設定機率&遊戲輸贏結果
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetAgentIncomeRatioAndGameRequest true "取得平台機率設定列表參數"
// @Success 200 {object} response.Response{data=[]model.GetAgentIncomeRatioAndGameResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getagentincomeratioandgamedata [post]
func (p *SystemRiskControlApi) GetAgentIncomeRatioAndGameData(c *ginweb.Context) {
	var req model.GetAgentIncomeRatioAndGameRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode
	isAdmin := len(myLevelCode) == 4
	if !isAdmin {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	// 產生欲查詢 tablename 並檢查 table 是否存在
	// 取開始到結束時間所有的格式化時間字串
	months := utils.GetTimeIntervalList("month", req.StartTime, req.EndTime)

	checkTablenames := make([]string, 0)
	for i := 0; i < len(months); i++ {
		checkTablenames = append(checkTablenames, "agent_game_ratio_stat"+"_"+months[i])
	}

	// 因為資料表是自動產生，故在此檢查資料表是否存在
	tablenames := make([]string, 0)
	for i := 0; i < len(checkTablenames); i++ {
		var tableExist bool
		query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`

		err := c.DB().QueryRow(query, checkTablenames[i]).Scan(&tableExist)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		if tableExist {
			tablenames = append(tablenames, checkTablenames[i])
		}
	}

	// 沒有資料表，直接送空資料
	if len(tablenames) == 0 {
		c.Ok(nil)
		return
	}
	sort.Strings(tablenames)

	startTime := utils.TransUnsignedTimeUTCFormat("15min", req.StartTime)
	endTime := utils.TransUnsignedTimeUTCFormat("15min", req.EndTime)

	selectQuery := `SELECT log_time, id, level_code, agent_id, game_id, game_type, room_type, de, ya, vaild_ya, tax, bonus, play_count 
			FROM %s
			WHERE 1=1 %s`
	unionQuery := " UNION "

	args := make([]interface{}, 0)
	queryAnd := make([]string, 0)

	if req.AgentId > definition.AGENT_ID_ALL {
		queryAnd = append(queryAnd, "agent_id = $%d")
		args = append(args, req.AgentId)
	}

	if req.GameId > definition.GAME_ID_ALL {
		queryAnd = append(queryAnd, "game_id = $%d")
		args = append(args, req.GameId)
	}

	if req.RoomType > definition.ROOM_TYPE_ALL {
		queryAnd = append(queryAnd, "room_type = $%d")
		args = append(args, req.RoomType)
	}

	queryAnd = append(queryAnd, "log_time >= $%d")
	args = append(args, startTime)

	queryAnd = append(queryAnd, "log_time <= $%d")
	args = append(args, endTime)

	finalQuerys := make([]string, 0)
	finalArgs := make([]interface{}, 0)
	paramIdx := 1
	for i := 0; i < len(tablenames); i++ {
		var combitionQuery string
		for _, v := range queryAnd {
			combitionQuery += " AND " + fmt.Sprintf(v, paramIdx)
			paramIdx += 1
		}

		queryTmp := fmt.Sprintf(selectQuery, tablenames[i], combitionQuery)
		finalQuerys = append(finalQuerys, queryTmp)
		finalArgs = append(finalArgs, args...)
	}

	finalQuery := ""
	for i := 0; i < len(finalQuerys); i++ {
		if i < len(finalQuerys)-1 {
			finalQuery += finalQuerys[i] + unionQuery
		} else {
			finalQuery += finalQuerys[i] + ";"
		}
	}

	rows, err := c.DB().Query(finalQuery, finalArgs...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	tmps := make(map[string]*model.GetAgentIncomeRatioAndGameResponse, 0)

	for rows.Next() {
		tmp := &model.GetAgentIncomeRatioAndGameResponse{}

		// newline every scan 5 column
		if err := rows.Scan(&tmp.LogTime, &tmp.Id, &tmp.LevelCode, &tmp.AgentId, &tmp.GameId, &tmp.GameType,
			&tmp.RoomType, &tmp.DeScore, &tmp.YaScore, &tmp.VaildYaScore, &tmp.Tax, &tmp.Bonus, &tmp.BetCount); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		agentObj := global.AgentCache.Get(tmp.AgentId)

		if agentObj.IsNotKillDiveCal {
			continue
		}

		tmp.AgnetName = agentObj.Name

		generalAgentObj := agentObj
		if len(agentObj.LevelCode) > 8 {
			generalAgentId, _ := strconv.ParseInt(agentObj.LevelCode[4:8], 16, 32)
			generalAgentObj = global.AgentCache.Get(int(generalAgentId))
		}

		if generalAgentObj.KillSwitch && tmp.GameType != definition.GAME_TYPE_BAIREN {
			tmp.KillRatio = generalAgentObj.KillRatio
		} else {
			pKey := global.AgentGameRatioCache.GetKey(tmp.AgentId, tmp.GameId, tmp.GameType, tmp.RoomType)
			agrc, ok := global.AgentGameRatioCache.SelectOne(pKey)
			if ok {
				tmp.KillRatio = agrc.KillRatio
			}
		}

		if _, ok := tmps[tmp.Id]; ok {
			tmps[tmp.Id].BetCount += tmp.BetCount
			tmps[tmp.Id].DeScore += tmp.DeScore
			tmps[tmp.Id].YaScore += tmp.YaScore
			tmps[tmp.Id].VaildYaScore += tmp.VaildYaScore
			tmps[tmp.Id].Tax += tmp.Tax
			tmps[tmp.Id].Bonus += tmp.Bonus
		} else {
			tmps[tmp.Id] = tmp
		}
	}

	c.Ok(tmps)
}

// Notice: 目前前端沒有頁面需求，所以api先不調整，未來使用調整時router註冊的註解要取消 by chouyang，
// @Tags 風控功能/代理機率統計
// @Summary Get players set the probability & the outcome of the game
// @Description 此接口用來取得玩家設定機率&遊戲輸贏結果
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetPlayerIncomeRatioAndGameRequest true "取得平台機率設定列表參數"
// @Success 200 {object} response.Response{data=[]model.GetPlayerIncomeRatioAndGameResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getplayerincomeratioandgamedata [post]
func (p *SystemRiskControlApi) GetPlayerIncomeRatioAndGameData(c *ginweb.Context) {
	var req model.GetPlayerIncomeRatioAndGameRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myLevelCode := claims.BaseClaims.LevelCode
	isAdmin := len(myLevelCode) == 4
	if !isAdmin {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	gameUserIds := make([]int, 0)

	if req.Username != "" {
		tmps, err := global.GetGameUserIdsByUsername(c.DB(), req.Username)
		if len(tmps) == 0 || err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_GAME_USERS_NOT_EXIST)
			return
		}

		gameUserIds = tmps
	}

	selectQuery := `SELECT gus.agent_id, gus.game_users_id, gu.original_username as game_users_username, gus.de, gus.ya, gus.vaild_ya, (gus.ya-gus.de) as betresult
	FROM game_users_stat gus, game_users gu
	WHERE gus.agent_id=gu.agent_id AND gus.game_users_id=gu.id %s
	ORDER BY gus.ya-de DESC;`

	args := make([]interface{}, 0)
	queryAnd := make([]string, 0)

	if req.AgentId > definition.AGENT_ID_ALL {
		queryAnd = append(queryAnd, "gus.agent_id = $%d")
		args = append(args, req.AgentId)
	}

	combitionQuery := ""
	for i, v := range queryAnd {
		combitionQuery += " AND " + fmt.Sprintf(v, i+1)
	}

	if len(gameUserIds) > 0 {
		// queryAnd = append(queryAnd, "gus.game_users_id IN ($%d)")
		queryTmp := "AND gus.game_users_id IN ("
		for idx, gameUserId := range gameUserIds {
			// args = append(args, gameUserId)
			queryTmp += utils.IntToString(gameUserId)
			if idx < len(gameUserIds)-1 {
				queryTmp += ", "
			}
		}

		queryTmp += ")"

		combitionQuery += queryTmp
	}

	query := fmt.Sprintf(selectQuery, combitionQuery)

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	tmps := make([]*model.GetPlayerIncomeRatioAndGameResponse, 0)

	for rows.Next() {
		tmp := &model.GetPlayerIncomeRatioAndGameResponse{}

		// newline every scan 5 column
		if err := rows.Scan(&tmp.AgentId, &tmp.GameUsersId, &tmp.GameUsername, &tmp.DeScore, &tmp.YaScore, &tmp.VaildYaScore,
			&tmp.BetResult); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		agentObj := global.AgentCache.Get(tmp.AgentId)
		tmp.AgnetName = agentObj.Name

		tmps = append(tmps, tmp)
	}

	// if len(tmps) <= 0 {
	// 	c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS)
	// 	return
	// }

	c.Ok(tmps)
}

// @Tags 風控功能/玩家標示設定
// @Summary Get agent custom tag setting list
// @Description 此接口用來取得玩家標示設定資料(只有總代理、子代理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.GetAgentCustomTagSettingResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getagentcustomtagsettinglist [post]
func (p *SystemRiskControlApi) GetAgentCustomTagSettingList(c *ginweb.Context) {

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isLevelOK := len(myLevelCode) >= 8

		if !isLevelOK {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(int(claims.BaseClaims.ID))
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		res := &model.GetAgentCustomTagSettingResponse{
			AgentId:       agentGameRatioCache.AgentId,
			CustomTagInfo: agentGameRatioCache.CustomTagInfo,
			UpdateTime:    agentGameRatioCache.UpdateTime,
		}

		c.Ok(res)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/玩家標示設定
// @Summary Set agent custom tag setting list
// @Description 此接口用來設定玩家標示設定資料(只有總代理、子代理可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetAgentCustomTagSettingListRequest true "設定玩家標示設定資料參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setagentcustomtagsettinglist [post]
func (p *SystemRiskControlApi) SetAgentCustomTagSettingList(c *ginweb.Context) {
	var req model.SetAgentCustomTagSettingListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isLevelOK := len(myLevelCode) >= 8

		if !isLevelOK {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(int(claims.BaseClaims.ID))
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		for idx, val := range req.CustomTagInfo {
			agentGameRatioCache.CustomTagInfo[idx] = val
		}

		if err := global.AgentCustomTagInfoCache.Update(agentGameRatioCache.AgentId, agentGameRatioCache.CustomTagInfo); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		c.Ok("")
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/玩家標示
// @Summary Get tag setting list
// @Description 此接口用來取得玩家標示資料列表(開發商（營運商）可看到全部代理商資訊；總代理、子代理可看到自身&下級)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUsersCustomTagListRequest true "取得玩家標示資料參數"
// @Success 200 {object} response.Response{data=model.GetGameUsersCustomTagListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getgameuserscustomtaglist [post]
func (p *SystemRiskControlApi) GetGameUsersCustomTagList(c *ginweb.Context) {
	var req model.GetGameUsersCustomTagListRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		// isAdmin := len(myLevelCode) == 4

		// 開發商（營運商）可看到全部代理商資訊；總代理、子代理可看到自身&下級
		sqlLevelCodeParam := ""
		if req.AgentId == definition.AGENT_ID_ALL {
			sqlLevelCodeParam = fmt.Sprintf("AND gush.level_code LIKE '%s%%'", myLevelCode)
		} else {
			targetAgent := global.AgentCache.Get(req.AgentId)
			sqlLevelCodeParam = fmt.Sprintf("AND gush.level_code='%s'", targetAgent.LevelCode)
		}

		// TODO: 先找出封線條件設定的玩家，再對應代理資訊，找出代理自定義的標籤

		startDate := utils.TransUnsignedTimeUTCFormat("hour", req.StartTime)
		endDate := utils.TransUnsignedTimeUTCFormat("hour", req.EndTime)
		// 開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號
		query := fmt.Sprintf(`SELECT gu.agent_id, gu.original_username, gush.game_users_id, gu.is_enabled, gu.is_risk,
		  gu.kill_dive_state, gu.kill_dive_value, gu.custom_status, SUM(gush.ya) as y, SUM(gush.vaild_ya) as v_y,
		  SUM(gush.de) as d, SUM(gush.tax) as t, SUM(play_count) as p_c, SUM(win_count) as w_c, SUM(big_win_count) as b_w_c,
		  SUM(gush.bonus) as b
		FROM game_users_stat_hour gush, game_users gu 
		WHERE gush.game_users_id = gu.id AND gush.log_time >= $1 AND gush.log_time <= $2 %s
		GROUP BY gu.agent_id, gu.original_username, gush.game_users_id, gu.is_enabled, gu.is_risk, gu.kill_dive_state, gu.kill_dive_value, gu.custom_status
		ORDER BY gush.game_users_id ASC;`, sqlLevelCodeParam)

		rows, err := c.DB().Query(query, startDate, endDate)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		defer rows.Close()

		reps := make(map[int]*model.GetGameUsersCustomTagListResponse, 0)

		for rows.Next() {
			var tmp model.GameUsersCustomTagList
			// newline every scan 5 column
			if err := rows.Scan(&tmp.AgentId, &tmp.GameUsersName, &tmp.GameUserId, &tmp.IsEnabled, &tmp.HighRisk,
				&tmp.KillDiveState, &tmp.KillDiveValue, &tmp.TagList, &tmp.Ya, &tmp.ValidYa,
				&tmp.De, &tmp.Tax, &tmp.PlayCount, &tmp.WinCount, &tmp.BgiWinCount,
				&tmp.Bonus); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}

			// get agent name from local memory
			agentObj := global.AgentCache.Get(tmp.AgentId)
			if agentObj != nil {
				tmp.AgentName = agentObj.Name
			}

			if val, ok := reps[tmp.AgentId]; ok {
				val.DataList = append(val.DataList, tmp)
			} else {
				reps[tmp.AgentId] = new(model.GetGameUsersCustomTagListResponse)
				reps[tmp.AgentId].DataList = append(reps[tmp.AgentId].DataList, tmp)
				agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(tmp.AgentId)
				if !ok {
					c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
					return
				}
				reps[tmp.AgentId].CustomTagInfo = agentGameRatioCache.CustomTagInfo
			}
		}

		c.Ok(reps)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/玩家標示
// @Summary Get tag from game users
// @Description 此接口用來取得單一標示玩家資料(開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUsersCustomTagRequest true "取得玩家標示資料參數"
// @Success 200 {object} response.Response{data=model.GetGameUsersCustomTagResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getgameuserscustomtag [post]
func (p *SystemRiskControlApi) GetGameUsersCustomTag(c *ginweb.Context) {
	var req model.GetGameUsersCustomTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isLevelOK := len(myLevelCode) >= 8

		if !isLevelOK {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		var dbId int
		var dbUsername string
		var dbIsRisk bool
		var dbKillDiveState int
		var dbKillDiveValue float64
		var dbCustomStatus string
		// 開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號
		query := `SELECT id, original_username, is_risk, kill_dive_state, kill_dive_value, custom_status
		 FROM game_users 
		 WHERE id=$1 AND original_username=$2 AND level_code=$3`

		err := c.DB().QueryRow(query, req.GameUserId, req.GameUsersName, myLevelCode).
			Scan(&dbId, &dbUsername, &dbIsRisk, &dbKillDiveState, &dbKillDiveValue,
				&dbCustomStatus)
		if err != nil && err != sql.ErrNoRows {
			return
		}

		agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(int(claims.BaseClaims.ID))
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		res := &model.GetGameUsersCustomTagResponse{
			GameUsersId:   dbId,
			GameUsersname: dbUsername,
			HighRisk:      dbIsRisk,
			KillDiveState: dbKillDiveState,
			KillDiveValue: dbKillDiveValue,
			TagList:       dbCustomStatus,
			CustomTagInfo: agentGameRatioCache.CustomTagInfo,
		}

		c.Ok(res)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/玩家標示
// @Summary Set tag to game users
// @Description 此接口用來設定標示玩家資料(開發商（營運商）不可使用；總代理、子代理只可設定自身的玩家帳號)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetGameUsersCustomTagRequest true "設定玩家標示資料參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setgameuserscustomtag [post]
func (p *SystemRiskControlApi) SetGameUsersCustomTag(c *ginweb.Context) {
	var req model.SetGameUsersCustomTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if success := req.CheckParams(); !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isLevelOK := len(myLevelCode) >= 8

		if !isLevelOK {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		if !global.AgentCustomTagInfoCache.CheckCustomTagInfoFormat(req.CustomTagInfo) {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		// 狀態改變通知遊戲
		if msg, code, err := notification.SendSetPlayerStatus([]int{req.GameUsersId}, []int{req.KillDiveState}); err != nil {
			c.Logger().Printf("UpdateGameUserInfo() SendSetPlayerStatus has error, msg: %s, code: %d, err: %v", msg, code, err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}

		// 定點用戶資料新增/刪除
		if req.KillDiveState == definition.GAMEUSERS_STATUS_KILLDIVE_CONFIGKILL {
			kdMap := make(map[string]float64)
			kdMap["need_update"] = float64(0)
			kdMap["kill_dive_value"] = req.KillDiveValue
			global.GameUsersKillDiveValueList.Store(req.GameUsersId, kdMap)
		} else {
			global.GameUsersKillDiveValueList.Delete(req.GameUsersId)
		}

		query := `SELECT is_risk, kill_dive_state, kill_dive_value, custom_status
		FROM game_users 
		WHERE id=$1 AND original_username=$2 AND level_code=$3`

		var originalIsRisk bool
		var originalKillDiveState int
		var originalKillDiveValue float64
		var originalCustomStatus string

		err := c.DB().QueryRow(query, req.GameUsersId, req.GameUsername, myLevelCode).Scan(&originalIsRisk, &originalKillDiveState, &originalKillDiveValue, &originalCustomStatus)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		query = `UPDATE game_users
		SET is_risk=$1, kill_dive_state=$2, kill_dive_value=$3, custom_status=$4 ,"update_time" = now()
		WHERE id=$5 AND original_username=$6 AND level_code=$7`
		_, err = c.DB().Exec(query,
			req.HighRisk, req.KillDiveState, req.KillDiveValue, req.TagList, req.GameUsersId,
			req.GameUsername, myLevelCode)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		// 更新自定義玩家標示 start
		agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(int(claims.BaseClaims.ID))
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}

		agentGameRatioCache.CustomTagInfo = req.CustomTagInfo
		global.AgentCustomTagInfoCache.Update(agentGameRatioCache.AgentId, agentGameRatioCache.CustomTagInfo)
		// 更新自定義玩家標示 end

		// 操作紀錄 start
		agentName := "unknow"
		agent := global.AgentCache.Get(int(claims.BaseClaims.ID))

		if agent != nil {
			agentName = agent.Name
		}
		actionLog := make(map[string]interface{})
		actionLog["agent_name"] = agentName
		actionLog["username"] = req.GameUsername
		if originalIsRisk != req.HighRisk {
			actionLog["is_risk"] = createBancendActionLogDetail(originalIsRisk, req.HighRisk)
		}
		if originalKillDiveState != req.KillDiveState {
			actionLog["kill_dive_state"] = createBancendActionLogDetail(originalKillDiveState, req.KillDiveState)
		}
		if originalKillDiveValue != float64(req.KillDiveValue) {
			actionLog["kill_dive_value"] = createBancendActionLogDetail(originalKillDiveValue, req.KillDiveValue)
		}
		if originalCustomStatus != req.TagList {
			actionLog["custom_status"] = createBancendActionLogDetail(originalCustomStatus, req.TagList)
			actionLog["custom_tag_info"] = req.CustomTagInfo
		}
		c.Set("action_log", actionLog)
		// 操作紀錄 end

		c.Ok("")
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags 風控功能/自動風控設定
// @Summary Get auto risk control setting
// @Description 此接口用來取得自動風控設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.GetAutoRiskControlSetting,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getautoriskcontrolsetting [post]
func (p *SystemRiskControlApi) GetAutoRiskControlSetting(c *ginweb.Context) {
	// 只有開發商（營運商）可使用
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	current := global.AutoRiskControlSettingCache.Get()
	if current == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
		return
	}

	c.Ok(&model.GetAutoRiskControlSetting{
		Current: current,
		Default: table_model.NewAutoRiskControlSetting(),
	})
}

// @Tags 風控功能/自動風控設定
// @Summary Set auto risk control setting
// @Description 此接口用來設定自動風控設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body table_model.AutoRiskControlSetting true "設定自動風控設定參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setautoriskcontrolsetting [post]
func (p *SystemRiskControlApi) SetAutoRiskControlSetting(c *ginweb.Context) {
	var req table_model.AutoRiskControlSetting
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.GameUserCoinInAndOutRequestPerMinuteLimit <= 0 ||
		req.GameUserApiRequestPerSecondLimit <= 0 ||
		req.GameUserCoinInAndOutDiffLimit <= 0 ||
		req.GameUserWinRateLimit <= 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 只有開發商（營運商）可使用
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	s, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_AUTO_RISK_CONTROL_SETTING)
	if !ok {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	backup := new(table_model.AutoRiskControlSetting)
	if err := json.Unmarshal([]byte(s.Value), backup); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	if backup.IsEnabled != req.IsEnabled && !req.IsEnabled {
		if err := global.AutoRiskControlStatCache.ClearAllStat(); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
			return
		}

		global.GameUsersAutoRiskWinRateList.Range(func(key, value interface{}) bool {
			global.GameUsersAutoRiskWinRateList.Delete(key.(int))
			return true
		})
	}

	if err := global.GlobalStorage.Update(definition.STORAGE_KEY_AUTO_RISK_CONTROL_SETTING, utils.ToJSON(req)); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	if err := global.AutoRiskControlSettingCache.Add(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
		return
	}

	actionLog := make(map[string]interface{})
	if backup.GameUserCoinInAndOutRequestPerMinuteLimit != req.GameUserCoinInAndOutRequestPerMinuteLimit {
		actionLog["game_user_coin_in_and_out_request_per_minute_limit"] = createBancendActionLogDetail(backup.GameUserCoinInAndOutRequestPerMinuteLimit, req.GameUserCoinInAndOutRequestPerMinuteLimit)
	}
	if backup.GameUserApiRequestPerSecondLimit != req.GameUserApiRequestPerSecondLimit {
		actionLog["game_user_api_request_per_second_limit"] = createBancendActionLogDetail(backup.GameUserApiRequestPerSecondLimit, req.GameUserApiRequestPerSecondLimit)
	}
	if backup.GameUserCoinInAndOutDiffLimit != req.GameUserCoinInAndOutDiffLimit {
		actionLog["game_user_coin_in_and_out_diff_limit"] = createBancendActionLogDetail(backup.GameUserCoinInAndOutDiffLimit, req.GameUserCoinInAndOutDiffLimit)
	}
	if backup.GameUserWinRateLimit != req.GameUserWinRateLimit {
		actionLog["game_user_win_rate_limit"] = createBancendActionLogDetail(backup.GameUserWinRateLimit, req.GameUserWinRateLimit)
	}
	if backup.IsEnabled != req.IsEnabled {
		actionLog["is_enabled"] = createBancendActionLogDetail(backup.IsEnabled, req.IsEnabled)
	}

	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 風控功能/玩家處置設定
// @Summary Set auto risk control setting
// @Description 此接口用來取得指定玩家的處置設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetGameUserRiskControlTagRequest true "取得指定玩家的處置設定參數"
// @Success 200 {object} response.Response{data=model.GetGameUserRiskControlTagResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getgameuserriskcontroltag [post]
func (p *SystemRiskControlApi) GetGameUserRiskControlTag(c *ginweb.Context) {
	var req model.GetGameUserRiskControlTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.Id <= 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 只有開發商（營運商）可使用
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	var resp model.GetGameUserRiskControlTagResponse

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(`gu."id"`, "gu.original_username", "gu.agent_id", `aa."name"`, "gu.risk_control_status").
		From("game_users AS gu").
		InnerJoin(`agent AS aa ON gu.agent_id = aa."id"`).
		Where(sq.Eq{`gu."id"`: req.Id}).
		ToSql()

	err := c.DB().QueryRow(query, args...).Scan(&resp.Id, &resp.Username, &resp.AgentId, &resp.AgentName, &resp.RiskControlTag)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	c.Ok(resp)
}

// @Tags 風控功能/玩家處置設定
// @Summary Set auto risk control setting
// @Description 此接口用來設定指定玩家的處置設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetGameUserRiskControlTagRequest true "設定指定玩家的處置設定參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setgameuserriskcontroltag [post]
func (p *SystemRiskControlApi) SetGameUserRiskControlTag(c *ginweb.Context) {
	var req model.SetGameUserRiskControlTagRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.Id <= 0 || len(req.RiskControlTag) != 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	for _, r := range req.RiskControlTag {
		if string(r) != definition.RISK_CONTROL_STATUS_DISABLED && string(r) != definition.RISK_CONTROL_STATUS_ENABLED {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
	}

	// 只有開發商（營運商）可使用
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	var username string
	var agentName string
	var originalRiskControlTag string
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("gu.original_username", `aa."name"`, "gu.risk_control_status").
		From("game_users AS gu").
		InnerJoin(`agent AS aa ON gu.agent_id = aa."id"`).
		Where(sq.Eq{`gu."id"`: req.Id}).
		ToSql()
	err := c.DB().QueryRow(query, args...).Scan(&username, &agentName, &originalRiskControlTag)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	originalRiskControlBetStatus := string(originalRiskControlTag[definition.RISK_CONTROL_BET_IDX])
	reqRiskControlBetStatus := string(req.RiskControlTag[definition.RISK_CONTROL_BET_IDX])
	// 禁止下注狀態有變更要通知game server
	if originalRiskControlBetStatus != reqRiskControlBetStatus {
		if _, _, err := notification.SendSoftBlockPlayer([]int{req.Id}, []int{utils.ToInt(reqRiskControlBetStatus)}); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}

		// 解封要清除redis中的資料
		if reqRiskControlBetStatus == definition.RISK_CONTROL_STATUS_DISABLED {
			if err := global.AutoRiskControlStatCache.ClearGameUserTotalGame(req.Id); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
				return
			}
			if err := global.AutoRiskControlStatCache.ClearGameUserTotalWin(req.Id); err != nil {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
				return
			}
		}
	}

	originalRiskControlCoinOutStatus := string(originalRiskControlTag[definition.RISK_CONTROL_COIN_OUT_IDX])
	reqRiskControlCoinOutStatus := string(req.RiskControlTag[definition.RISK_CONTROL_COIN_OUT_IDX])
	// 解封禁止下分要清除redis中的資料，並加上目前的餘額
	if originalRiskControlCoinOutStatus != reqRiskControlCoinOutStatus &&
		reqRiskControlCoinOutStatus == definition.RISK_CONTROL_STATUS_DISABLED {
		if err := global.AutoRiskControlStatCache.ClearGameUserTotalCoinIn(req.Id); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
			return
		}
		if err := global.AutoRiskControlStatCache.ClearGameUserTotalCoinOut(req.Id); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
			return
		}

		gold, _, err := notification.GetGameServerGold(int64(req.Id))
		if err != nil {
			c.Logger().Printf("SetGameUserRiskControlTag exec notification.GetGameServerGold() has error: %v", err)
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
		if _, err := global.AutoRiskControlStatCache.IncrGameUserTotalCoinIn(req.Id, gold); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REDIS)
			return
		}
	}

	query, args, _ = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("game_users").
		Set("risk_control_status", req.RiskControlTag).
		Where(sq.Eq{`"id"`: req.Id}).
		ToSql()
	_, err = c.DB().Exec(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	actionLog := make(map[string]interface{})
	actionLog["agent_name"] = agentName
	actionLog["username"] = username
	if originalRiskControlTag != req.RiskControlTag {
		actionLog["risk_control_status"] = createBancendActionLogDetail(originalRiskControlTag, req.RiskControlTag)
	}
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 風控功能/遊戲基礎設定
// @Summary Get game setting
// @Description 此接口用來取得遊戲基礎設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.GetGameSettingResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getgamesetting [post]
func (p *SystemRiskControlApi) GetGameSetting(c *ginweb.Context) {
	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	val, _, err := notification.Getdefaultgamesetting()
	if err != nil {
		c.Logger().Error("notification.Getdefaultgamesetting() has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	var sourceGameSetting []model.GameSetting
	err = json.Unmarshal([]byte(val), &sourceGameSetting)
	if err != nil {
		c.Logger().Error("GetGameSetting() sourceGameSetting json.Unmarshal has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	storage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESETTINGSUPPORTINFO)
	if !ok || storage.Value == "" {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	var supportGameSettings []int
	err = json.Unmarshal([]byte(storage.Value), &supportGameSettings)
	if err != nil {
		c.Logger().Error("GetGameSetting() supportGameSettings json.Unmarshal has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	supportMap := make(map[int]struct{})
	for i := 0; i < len(supportGameSettings); i++ {
		supportMap[supportGameSettings[i]] = struct{}{}
	}

	resp := make(model.GetGameSettingResponse, 0)
	for i := 0; i < len(sourceGameSetting); i++ {
		if _, find := supportMap[sourceGameSetting[i].GameId]; !find {
			continue
		}

		resp = append(resp, &sourceGameSetting[i])
	}

	c.Ok(resp)
}

// @Tags 風控功能/遊戲基礎設定
// @Summary Set game setting
// @Description 此接口用來設定遊戲基礎設定(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetGameSettingRequest true "設定遊戲基礎設定參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/setgamesetting [post]
func (p *SystemRiskControlApi) SetGameSetting(c *ginweb.Context) {
	var req model.SetGameSettingRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	} else if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	val, _, err := notification.Getdefaultgamesetting()
	if err != nil {
		c.Logger().Error("notification.Getdefaultgamesetting() has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	var sourceGameSetting []model.GameSetting
	err = json.Unmarshal([]byte(val), &sourceGameSetting)
	if err != nil {
		c.Logger().Error("SetGameSetting() sourceGameSetting json.Unmarshal has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	sourceGameSettingMap := make(map[int]*model.GameSetting)
	for i := 0; i < len(sourceGameSetting); i++ {
		sourceGameSettingMap[sourceGameSetting[i].GameId] = &sourceGameSetting[i]
	}

	storage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESETTINGSUPPORTINFO)
	if !ok || storage.Value == "" {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	var supportGameSettings []int
	err = json.Unmarshal([]byte(storage.Value), &supportGameSettings)
	if err != nil {
		c.Logger().Error("SetGameSetting() supportGameSettings json.Unmarshal has error: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	supportGameSettingMap := make(map[int]struct{})
	for i := 0; i < len(supportGameSettings); i++ {
		supportGameSettingMap[supportGameSettings[i]] = struct{}{}
	}

	gameIds := make([]int, 0)
	matchGameRTPs := make([]float64, 0)
	matchGameKillRates := make([]float64, 0)
	matchGames := make([]int, 0)
	normalMatchGameRtps := make([]float64, 0)
	normalMatchGameKillRates := make([]float64, 0)
	lowBoundRtps := make([]float64, 0)
	limitOdds := make([]float64, 0)
	limitMoneys := make([]float64, 0)
	actionLogGameSettings := make([]map[string]interface{}, 0)

	for i := 0; i < len(req); i++ {
		gameId := req[i].GameId

		if _, support := supportGameSettingMap[gameId]; !support {
			c.Logger().Error("SetGameSetting game not support gameSetting, game id: %d", gameId)
			continue
		}

		origin := sourceGameSettingMap[gameId]
		if origin == nil {
			c.Logger().Error("SetGameSetting game not find in origin gameSetting, game id: %d", gameId)
			continue
		}

		game := global.GameCache.Get(gameId)
		if game == nil {
			c.Logger().Error("SetGameSetting game not find in cache, game id: %d", gameId)
			continue
		}

		if origin.IsEqual(&req[i]) {
			continue
		}

		before := utils.ToJSON(origin)
		after := utils.ToJSON(req[i])

		origin.MatchGameRTP = req[i].MatchGameRTP
		origin.MatchGameKillRate = req[i].MatchGameKillRate
		origin.MatchGames = req[i].MatchGames
		origin.NormalMatchGameRTP = req[i].NormalMatchGameRTP
		origin.NormalMatchGameKillRate = req[i].NormalMatchGameKillRate
		origin.LowBoundRTP = req[i].LowBoundRTP
		origin.LimitOdds = req[i].LimitOdds
		origin.LimitMoney = req[i].LimitMoney

		gameIds = append(gameIds, gameId)
		matchGameRTPs = append(matchGameRTPs, req[i].MatchGameRTP)
		matchGameKillRates = append(matchGameKillRates, req[i].MatchGameKillRate)
		matchGames = append(matchGames, req[i].MatchGames)
		normalMatchGameRtps = append(normalMatchGameRtps, req[i].NormalMatchGameRTP)
		normalMatchGameKillRates = append(normalMatchGameKillRates, req[i].NormalMatchGameKillRate)
		lowBoundRtps = append(lowBoundRtps, req[i].LowBoundRTP)
		limitOdds = append(limitOdds, req[i].LimitOdds)
		limitMoneys = append(limitMoneys, req[i].LimitMoney)
		actionLogGameSettings = append(actionLogGameSettings, map[string]interface{}{
			"game_id": gameId,
			"before":  before,
			"after":   after,
		})
	}

	if len(actionLogGameSettings) == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_NO_CHANGE_IN_DATA)
		return
	}

	apiResult, code, err := notification.SendGameSetting(
		gameIds,
		matchGameRTPs,
		matchGameKillRates,
		matchGames,
		normalMatchGameRtps,
		normalMatchGameKillRates,
		lowBoundRtps,
		limitOdds,
		limitMoneys,
	)
	if err != nil {
		c.Logger().Printf("SetGameSetting() has error in notification.SendGameSetting():result is %s, code is %d, err is %v", apiResult, code, err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_API_SERVER_REQUEST_FAILED)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["game_settings"] = actionLogGameSettings
	c.Set("action_log", actionLog)

	c.Ok("")
}

// @Tags 風控功能/遊戲即時殺率資訊
// @Summary Get realtime game ratio
// @Description 此接口用來取得遊戲即時殺率資訊(只有開發商（營運商）可使用)
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetRealtimeGameRatioRequest true "設定參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/riskcontrol/getrealtimegameratio [post]
func (p *SystemRiskControlApi) GetRealtimeGameRatio(c *ginweb.Context) {
	var req model.GetRealtimeGameRatioRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	} else if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	claims := c.GetUserClaims()
	if claims == nil || len(claims.BaseClaims.LevelCode) >= 8 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	apiBody, apiCode, err := notification.GetRealtimeGameRatio(req.GameId)
	code := notification.TransformApiCodeToModelErrorCode(apiCode)
	if err != nil {
		c.Logger().Error("GetRealtimeGameRatio() sourceGameSetting json.Unmarshal has error: %v", err)
		c.OkWithCode(code)
		return
	}

	c.Ok(apiBody)
}
