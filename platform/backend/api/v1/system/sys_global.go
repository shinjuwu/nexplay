package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/utils"
	"backend/server/global"
	"definition"

	"github.com/gin-gonic/gin"
)

type SystemGlobalApi struct {
	BasePath string
}

func NewSystemGlobalApi(basePath string) api_cluster.IApiEach {
	return &SystemGlobalApi{
		BasePath: basePath,
	}
}

func (p *SystemGlobalApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemGlobalApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/roloadglobaldata", ginHandler.Handle(p.RoloadGlobalData))
	// g.POST("/checkgameroomsetting", ginHandler.Handle(p.CheckGameRoomSetting))
	g.GET("/getagentlist", ginHandler.Handle(p.GetAgentList))
	g.GET("/getallgamelist", ginHandler.Handle(p.GetAllGameList))
	g.GET("/getgamelist", ginHandler.Handle(p.GetGameList))
	g.GET("/getroomtypelist", ginHandler.Handle(p.GetRoomTypeList))
	g.GET("/getagentpermissionlist", ginHandler.Handle(p.GetAgentPermissionList))
	g.GET("/getagentcustomtagsettinglist", ginHandler.Handle(p.GetAgentCustomTagSettingList))
	g.GET("/getagentadminuserpermissionlist", ginHandler.Handle(p.GetAgentAdminuserPermissionList))
}

// @Tags Global
// @Summary Roload global date
// @Description 此接口用來重新載入本地端資料
// @Produce  application/json
// @Security BearerAuth
// @Param data body model.RoloadGlobalDateRequest true "指定資料 id (0: all)"
// @Success 200 {object} response.Response{data=model.AgentGame,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/roloadglobaldata [post]
func (p *SystemGlobalApi) RoloadGlobalData(c *ginweb.Context) {

	var req model.RoloadGlobalDateRequest
	_ = c.ShouldBindJSON(&req)

	if req.Id < 0 {
		c.Fail("id range is illegal")
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims != nil {

		myLevelCode := claims.BaseClaims.LevelCode

		if len(myLevelCode) != 4 {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		}

		if req.Id == 0 {
			err := global.InitGlobalData(c.Logger(), c.DB(), c.Redis(), c.GetJob())
			if err != nil {
				c.Fail(err.Error())
				return
			}
		}

	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	c.Ok("success")
}

// @Tags Global
// @Summary 取得全部代理商list(供下拉選單使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getagentlist [get]
func (p *SystemGlobalApi) GetAgentList(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	c.Ok(GetAgentInfoList(int(userClaims.BaseClaims.ID)))
}

// @Tags Global
// @Summary 取得全部遊戲list(供下拉選單使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]int,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getallgamelist [get]
func (p *SystemGlobalApi) GetAllGameList(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	accountType := userClaims.AccountType
	agentId := int(userClaims.BaseClaims.ID)

	c.Ok(GetGameIdList(accountType, agentId, definition.GAME_STATES_ONLINE_MAINTAIN_OFFLINE))
}

// @Tags Global
// @Summary 取得上線及維護中的遊戲list(供下拉選單使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]int,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getgamelist [get]
func (p *SystemGlobalApi) GetGameList(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	accountType := userClaims.AccountType
	agentId := int(userClaims.BaseClaims.ID)

	c.Ok(GetGameIdList(accountType, agentId, definition.GAME_STATES_ONLINE_MAINTAIN))
}

// @Tags Global
// @Summary 取得房間類型list(供下拉選單使用)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]int,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getroomtypelist [get]
func (p *SystemGlobalApi) GetRoomTypeList(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	accountType := userClaims.AccountType
	agentId := int(userClaims.BaseClaims.ID)

	c.Ok(GetRoomTypeList(accountType, agentId))
}

// @Tags Global
// @Summary 取得權限群組層級list(供下拉選單使用)
// @Produce  application/json
// @Security BearerAuth
// @Param account_type query int true "層級"
// @Success 200 {object} response.Response{data=[]map[string]interface{},msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getagentpermissionlist [get]
func (p *SystemGlobalApi) GetAgentPermissionList(c *ginweb.Context) {
	userClaims := c.GetUserClaims()
	if userClaims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	accountType := userClaims.AccountType
	agentId := int(userClaims.BaseClaims.ID)
	permissionId := userClaims.PermissionId

	reqAccountType := utils.ToInt(c.Request.URL.Query().Get("account_type"))
	if reqAccountType <= 0 ||
		reqAccountType < accountType ||
		reqAccountType > accountType+1 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	c.Ok(GetAgentPermissionList(agentId, reqAccountType, permissionId))
}

// @Tags Global
// @Summary Get agent custom tag setting list
// @Description 此接口用來取得玩家標示設定下拉選單資料list(供下拉選單使用,Drop Down Menu)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=global.AgentTagInfoList,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getagentcustomtagsettinglist [get]
func (p *SystemGlobalApi) GetAgentCustomTagSettingList(c *ginweb.Context) {

	// 只能取得自己的設定,所以不用輸入對應參數

	claims := c.GetUserClaims()
	if claims != nil {
		// 只有總代理、子代理可以查詢
		myLevelCode := claims.BaseClaims.LevelCode
		isAdmin := len(myLevelCode) == 4

		var res global.AgentTagInfoList
		res.DefaultTagList = global.NewDefaultTagInfoOutput()
		if isAdmin {
			res.CustomTagInfo = nil
		} else {
			agentGameRatioCache, ok := global.AgentCustomTagInfoCache.SelectOne(int(claims.BaseClaims.ID))
			if !ok {
				c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL_CACHE_NOT_EXIST)
				return
			}

			res.CustomTagInfo = agentGameRatioCache.CustomTagInfo
		}

		c.Ok(res)
	} else {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}
}

// @Tags Global
// @Summary Get agent parent admin user permission list
// @Description 此接口用來取得代理父帳號權限list(供下拉選單使用,Drop Down Menu)
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]int,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/global/getagentadminuserpermissionlist [get]
func (p *SystemGlobalApi) GetAgentAdminuserPermissionList(c *ginweb.Context) {
	// 只能取得自己代理的設定,所以不用輸入對應參數
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	agent := global.AgentCache.Get(int(claims.BaseClaims.ID))
	adminUser := global.AdminUserCache.Get(agent.Id, agent.AdminUsername)
	agentPermission := global.AgentPermissionCache.Get(adminUser.PermissionId)

	c.Ok(agentPermission.Permission.List)
}
