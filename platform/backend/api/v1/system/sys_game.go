package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	"definition"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemGameApi struct {
	BasePath string
}

func NewSystemGameApi(basePath string) api_cluster.IApiEach {
	return &SystemGameApi{
		BasePath: basePath,
	}
}

func (p *SystemGameApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemGameApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.GET("/getgamelist", ginHandler.Handle(p.GetGameList))
	g.POST("/setgamestate",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameState))
	g.GET("/getgameserverstate", ginHandler.Handle(p.GetGameServerState))
	g.POST("/setgameserverstate",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameServerState))
	g.POST("/notifygameserver",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.NotifyGameServer))
	g.GET("/getgameiconlist", ginHandler.Handle(p.GetGameIconList))
	g.POST("/setgameiconlist",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetGameIconList))
	g.GET("/getplinkoballmaxodds", ginHandler.Handle(p.GetPlinkoballMaxOdds))
	g.POST("/setplinkoballmaxodds",
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetPlinkoballMaxOdds))

}

// @Tags 遊戲設置/遊戲設置
// @Summary 取得遊戲列表
// @Produce application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]model.GetGameResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/getgamelist [get]
func (p *SystemGameApi) GetGameList(c *ginweb.Context) {
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "code", "state").
		From("game").
		Where(sq.NotEq{"type": definition.GAME_TYPE_LOBBY}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer rows.Close()

	resp := make([]*model.GetGameResponse, 0)
	for rows.Next() {
		var temp model.GetGameResponse
		if err := rows.Scan(&temp.Id, &temp.Code, &temp.State); err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
		resp = append(resp, &temp)
	}

	c.Ok(resp)
}

// @Tags 遊戲設置/遊戲設置
// @Summary 修改遊戲狀態
// @Produce  application/json
// @Security BearerAuth
// @param data body model.SetGameStateRequest true "修改遊戲狀態參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/setgamestate [post]
func (p *SystemGameApi) SetGameState(c *ginweb.Context) {
	var req model.SetGameStateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if code := req.CheckParams(); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	game := global.GameCache.Get(req.GameId)
	if game == nil || game.Type == definition.GAME_TYPE_LOBBY {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if game.State != int16(req.State) {
		if code := setGameState(c.DB(), &req); code != definition.ERROR_CODE_SUCCESS {
			c.OkWithCode(code)
			return
		}

		gamelists := make([]map[string]interface{}, 1)
		gamelists = append(gamelists, getGameList(game.Id, game.Code, req.State))

		if code, err := notification.SendSetGameList(gamelists); err != nil {
			c.Logger().Info("notification.SendSetGameList fail, gamelists=%v, resp code=%d", gamelists, code)

			// 通知失敗要重設DB資料
			req.State = game.State
			if code := setGameState(c.DB(), &req); code != definition.ERROR_CODE_SUCCESS {
				c.Logger().Warn("notification.SendSetGameList fail and db rollback fail, req=%v", req)
				c.OkWithCode(code)
				return
			}

			// 重設DB資料完成才傳通知送失敗
			c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
			return
		}
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["game_id"] = game.Id
	if game.State != int16(req.State) {
		actionLog["state"] = createBancendActionLogDetail(game.State, req.State)
	}
	c.Set("action_log", actionLog)

	// 全部成功才更新cache
	game.State = int16(req.State)

	c.Ok(nil)
}

// @Tags 遊戲設置/遊戲設置
// @Summary 取得遊戲server狀態
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/getgameserverstate [get]
func (p *SystemGameApi) GetGameServerState(c *ginweb.Context) {
	gameServerInfoStorage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO)
	if !ok {
		c.Logger().Error("GetGameServerState() STORAGE_KEY_GAMESERVERINFO is empty")
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
		return
	}
	resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
	if len(resp) <= 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS)
		return
	}
	c.Ok(resp)
}

type setgameserverstateRequest struct {
	State int `json:"state"`
}

// @Tags 遊戲設置/遊戲設置
// @Summary 設置遊戲server狀態
// @Produce  application/json
// @Security BearerAuth
// @param state body setgameserverstateRequest true "設置遊戲server狀態"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/setgameserverstate [post]
func (p *SystemGameApi) SetGameServerState(c *ginweb.Context) {
	req := setgameserverstateRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	reqState := req.State
	if reqState != definition.GS_STATE_OPEN && reqState != definition.GS_STATE_CLOSE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	curState := 0
	gameServerInfoStorage, isNew := global.GlobalStorage.SelectOrInsertOne(definition.STORAGE_KEY_GAMESERVERINFO, utils.ToJSON(req), false)
	if gameServerInfoStorage != nil {
		resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
		cs, ok := resp["state"].(float64)
		if !ok {
			c.Logger().Error("getGameServerState code:%d", definition.ERROR_CODE_ERROR_DATABASE)
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		curState = int(cs)

		if curState != reqState {
			if err := global.GlobalStorage.Update(definition.STORAGE_KEY_GAMESERVERINFO, utils.ToJSON(req)); err != nil {
				c.Logger().Error("getGameServerState update has error: %v", err)
				c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
				return
			}
		}
	}

	code, err := notification.SetGameServerState(reqState)
	if code != definition.ERROR_CODE_SUCCESS || err != nil {
		c.Logger().Error("setGameServerState notification code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	if curState != reqState || isNew {
		if isNew {
			curState = 0
		}
		actionLog["state"] = createBancendActionLogDetail(curState, reqState)
	}
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

type notifygameserverRequest struct {
	GameId int `json:"game_id"`
	State  int `json:"state"`
}

// @Tags 遊戲設置/遊戲設置
// @Summary 創建更新遊戲相關設定(遊戲server維護中才可以使用)
// @Produce  application/json
// @Security BearerAuth
// @param game_id body notifygameserverRequest true "送通知給遊戲伺服器"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/notifygameserver [post]
func (p *SystemGameApi) NotifyGameServer(c *ginweb.Context) {
	req := notifygameserverRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	gameId := req.GameId
	game := global.GameCache.Get(gameId)
	if game == nil || game.Type == definition.GAME_TYPE_LOBBY {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	gameServerInfoStorage, ok := global.GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO)
	if !ok {
		c.Logger().Error("NotifyGameServer() STORAGE_KEY_GAMESERVERINFO is empty")
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
		return
	}
	resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
	cs, ok := resp["state"].(float64)
	if !ok {
		c.Logger().Error("getGameServerState code:%d", definition.ERROR_CODE_ERROR_DATABASE)
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	state := int(cs)

	if state != definition.GAME_STATE_MAINTAIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_GAME_SERVER_NOT_IN_MAINTENANCE)
		return
	}

	code, err := notifyGameList(c.Logger(), c.DB(), game.Id, game.Code)
	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("notifyGameList code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	code, err = notifyGameInfo(c.Logger(), c.DB(), game.Id, game.Code)
	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("notifyGameInfo code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	code, err = notifyLobbyInfo(c.Logger(), c.DB(), game.Id)
	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("notifyLobbyInfo code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	// 殺放資料由game server提供，所以要先跟game server取得預設殺放進行更新
	code, err = updateKillDiveInfo(c.Logger(), c.DB(), game.Id)
	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("updateKillDiveInfo code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	code, err = notifyKillDiveInfo(c.Logger(), c.DB(), game.Id)
	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("notifyKillDiveInfo code:%d, err: %v", code, err)
		c.OkWithCode(code)
		return
	}

	// 操作紀錄
	actionLog := make(map[string]interface{})
	actionLog["game_id"] = game.Id
	c.Set("action_log", actionLog)

	c.Ok(nil)
}

type getGameIconListResponse struct {
	IsAdmin                 bool               `json:"is_admin"`                   // 是否為總代理預設值(只能有一個)
	IsDefault               bool               `json:"is_default"`                 // 是否讀取預設值
	UpdateTime              time.Time          `json:"update_time"`                // 更新時間
	CreateTime              time.Time          `json:"create_time"`                // 創建時間
	GameIconListData        []*getGameIconData `json:"gameicon_list_data"`         // 遊戲ICON 列表
	DefaultGameIconListData []*getGameIconData `json:"default_gameicon_list_data"` // 預設遊戲ICON 列表(上級)
}

type getGameIconData struct {
	Id     int    `json:"id"`     // 遊戲id(PK)
	Name   string `json:"name"`   // 遊戲名稱
	Code   string `json:"code"`   // 遊戲代碼
	Type   int    `json:"type"`   // 遊戲類型
	Rank   int    `json:"rank"`   // 排行 (since from 0)
	Hot    int    `json:"hot"`    // 熱門 (0:none, 1:熱門)
	Newest int    `json:"newest"` // 最新 (0:none, 1:最新)
	Push   int    `json:"push"`   // 推廣大圖 (0:none, 1:大圖1, 2:大圖2...)
}

// @Tags 遊戲設置/遊戲設置
// @Summary 取得遊戲icon list
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=getGameIconListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/getgameiconlist [get]
func (p *SystemGameApi) GetGameIconList(c *ginweb.Context) {
	// 只查詢自己的資料
	claims := c.GetUserClaims()
	if claims != nil {
		myLevelCode := claims.BaseClaims.LevelCode

		myIconSetting, ok := global.AgentGameIconListCache.SelectOne(myLevelCode)
		if ok {
			// search up level setting
			var upIconSetting *global.AgentGameIconList
			for i := len(myLevelCode) - 4; i >= 4; i -= 4 {
				tmp, ok := global.AgentGameIconListCache.SelectOne(myLevelCode[0:i])
				if ok {
					// 只要上層使用 default 設定就再往上找
					// 最後一層直接給值
					if !tmp.IsDefault || i == 4 {
						upIconSetting = tmp
						break
					}
				}
			}

			resp := &getGameIconListResponse{
				IsAdmin:                 myIconSetting.IsAdmin,
				IsDefault:               myIconSetting.IsDefault,
				UpdateTime:              myIconSetting.UpdateTime,
				CreateTime:              myIconSetting.CreateTime,
				GameIconListData:        make([]*getGameIconData, 0),
				DefaultGameIconListData: make([]*getGameIconData, 0),
			}

			gameCacheTmps := global.GameCache.GetAll()
			for _, gameIcon := range myIconSetting.GameIconList {
				for _, game := range gameCacheTmps {
					if gameIcon.GameId == game.Id {
						res := &getGameIconData{
							Id:     game.Id,
							Name:   game.Name,
							Code:   game.Code,
							Type:   game.Type,
							Rank:   gameIcon.Rank,
							Hot:    gameIcon.Hot,
							Newest: gameIcon.Newest,
							Push:   gameIcon.Push,
						}
						resp.GameIconListData = append(resp.GameIconListData, res)
					}
				}
			}

			if upIconSetting == nil {
				upIconSetting = myIconSetting
			}

			for _, gameIcon := range upIconSetting.GameIconList {
				for _, game := range gameCacheTmps {
					if gameIcon.GameId == game.Id {
						res := &getGameIconData{
							Id:     game.Id,
							Name:   game.Name,
							Code:   game.Code,
							Type:   game.Type,
							Rank:   gameIcon.Rank,
							Hot:    gameIcon.Hot,
							Newest: gameIcon.Newest,
							Push:   gameIcon.Push,
						}
						resp.DefaultGameIconListData = append(resp.DefaultGameIconListData, res)
					}
				}
			}

			c.Ok(resp)
			return
		} else {
			c.Logger().Printf("GetGameIconList() AgentGameIconListCache is nil, key is: %v", myLevelCode)
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}
}

type setGameIconListRequest struct {
	IsDefault    bool               `json:"is_default"`     // 是否讀取預設值
	GameIconList []*global.GameIcon `json:"game_icon_list"` // 遊戲 icon 設定
}

// @Tags 遊戲設置/遊戲設置
// @Summary 設置遊戲icon list
// @Produce  application/json
// @Security BearerAuth
// @param state body setGameIconListRequest true "設置遊戲icon list參數(只能改自己)"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/setgameiconlist [post]
func (p *SystemGameApi) SetGameIconList(c *ginweb.Context) {
	req := setGameIconListRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	claims := c.GetUserClaims()
	myLevelCode := claims.BaseClaims.LevelCode
	myAgentId := int(claims.BaseClaims.ID)
	myAccountType := claims.AccountType
	if myAccountType == definition.ACCOUNT_TYPE_ADMIN {
		// 如果是管理者，IsDefault 必定要是 false
		if req.IsDefault {
			req.IsDefault = false
		}
	}

	if !req.IsDefault {
		if len(req.GameIconList) <= 0 {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}
	}

	if len(req.GameIconList) > 0 {
		for _, val := range req.GameIconList {
			if val.Hot > 1 {
				val.Hot = 1
			} else if val.Hot < 0 {
				val.Hot = 0
			}

			if val.Newest > 1 {
				val.Newest = 1
			} else if val.Newest < 0 {
				val.Newest = 0
			}

			if val.Push < 0 {
				val.Push = 0
			}

			if val.Rank < 0 {
				val.Rank = 0
			} else if val.Rank > 9999 {
				val.Rank = 9999
			}

			// 互斥
			if val.Hot == 1 && val.Newest == 1 {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}
		}
	}

	if claims != nil {
		myIconSetting, ok := global.AgentGameIconListCache.SelectOne(myLevelCode)
		if !ok {
			c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
			return
		}
		beforeIsDefault := myIconSetting.IsDefault
		beforeIconSetting := append(make([]*global.GameIcon, 0, len(myIconSetting.GameIconList)), myIconSetting.GameIconList...)

		if req.IsDefault {
			var upIconSetting *global.AgentGameIconList
			for i := len(myLevelCode) - 4; i >= 4; i -= 4 {
				tmp, ok := global.AgentGameIconListCache.SelectOne(myLevelCode[0:i])
				if ok {
					// 只要上層使用 default 設定就再往上找
					// 最後一層直接給值
					if !tmp.IsDefault || i == 4 {
						upIconSetting = tmp
						break
					}
				}
			}

			global.AgentGameIconListCache.Update(myLevelCode, req.IsDefault, utils.ToJSON(upIconSetting.GameIconList))
		} else {
			global.AgentGameIconListCache.Update(myLevelCode, req.IsDefault, utils.ToJSON(req.GameIconList))
		}

		// 操作紀錄
		actionLog := make(map[string]interface{})
		agentObj := global.AgentCache.Get(myAgentId)
		actionLog["agent_name"] = agentObj.Name
		actionLog["is_default"] = createBancendActionLogDetail(beforeIsDefault, req.IsDefault)
		actionLog["icon_list"] = createBancendActionLogDetail(beforeIconSetting, req.GameIconList)

		c.Set("action_log", actionLog)

		c.Ok(nil)
		return
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}
}

// @Tags 遊戲設置/遊戲設置  
// @Summary 取得Plinko球倍率上限
// @Produce application/json
// @Security BearerAuth
// @Param agentName query string true "代理商名称"
// @Param gameId query int true "游戏ID"
// @Success 200 {object} response.Response{data=model.GetPlinkoballMaxOddsResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/getplinkoballmaxodds [get]
func (p *SystemGameApi) GetPlinkoballMaxOdds(c *ginweb.Context) {
	req := model.GetPlinkoballMaxOddsRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.CheckParams() != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 验证游戏ID是否为Plinko
	if req.GameId != definition.GAME_ID_PLINKO {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 调用game server获取倍率上限
	data, code, err := notification.GetPlinktoBallMaxOdds(req.AgentName)
	if err != nil {
		c.Logger().Error("GetPlinkoballMaxOdds() failed, err: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("GetPlinkoballMaxOdds() game server error, code: %d", code)
		c.OkWithCode(code)
		return
	}

	// 解析返回数据
	result := utils.ToMap([]byte(data))
	maxOdds := float64(1) // 默认值
	if val, ok := result["max_odds"].(float64); ok {
		maxOdds = val
	}

	resp := model.GetPlinkoballMaxOddsResponse{
		AgentName: req.AgentName,
		GameId:    req.GameId,
		MaxOdds:   maxOdds,
	}

	c.Ok(resp)
}

// @Tags 遊戲設置/遊戲設置
// @Summary 設定Plinko球倍率上限
// @Produce application/json  
// @Security BearerAuth
// @Param data body model.SetPlinkoballMaxOddsRequest true "設定Plinko球倍率上限參數"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/game/setplinkoballmaxodds [post]
func (p *SystemGameApi) SetPlinkoballMaxOdds(c *ginweb.Context) {
	req := model.SetPlinkoballMaxOddsRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	if req.CheckParams() != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 验证游戏ID是否为Plinko
	if req.GameId != definition.GAME_ID_PLINKO {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 调用game server设定倍率上限
	data, code, err := notification.SetPlinktoBallMaxOdds(req.AgentName, req.MaxOdds)
	if err != nil {
		c.Logger().Error("SetPlinkoballMaxOdds() failed, err: %v", err)
		c.OkWithCode(definition.ERROR_CODE_ERROR_NOTIFICATION)
		return
	}

	if code != definition.ERROR_CODE_SUCCESS {
		c.Logger().Error("SetPlinkoballMaxOdds() game server error, code: %d", code)
		c.OkWithCode(code)
		return
	}

	// 操作記錄
	actionLog := make(map[string]interface{})
	actionLog["agent_name"] = req.AgentName
	actionLog["game_id"] = req.GameId
	actionLog["max_odds"] = req.MaxOdds

	c.Set("action_log", actionLog)
	c.Ok(data)
}
