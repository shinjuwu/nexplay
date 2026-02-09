package game

import (
	"backend/api/game/controller"
	"backend/api/game/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type CasinoHandleFunc func(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.CasinoHandleRequest, lock *sync.Mutex) (bool, *model.CasinoHandleResponse, error)

type CasinoChannelApi struct {
	BasePath         string
	ControllerHandle map[int]CasinoHandleFunc
	userLockMap      *sync.Map
}

func NewCasinoChannelApi(basePath string) api_cluster.IApiEach {
	return &CasinoChannelApi{
		BasePath: basePath,
		// define API func
		ControllerHandle: map[int]CasinoHandleFunc{
			model.CasinoHandle_GameLogin: controller.CasinoGameLogin,
			model.CasinoHandle_Register:  controller.CasinoRegister,
			model.CasinoHandle_Credit:    controller.CasinoCredit,
		},
		userLockMap: new(sync.Map),
	}
}

func (p *CasinoChannelApi) GetGroupPath() string {
	return p.BasePath
}

// if get nil, create new one
func (p *CasinoChannelApi) GetUserLock(thirdAccount string) *sync.Mutex {
	tmp, _ := p.userLockMap.LoadOrStore(thirdAccount, new(sync.Mutex))
	return tmp.(*sync.Mutex)
}

func (p *CasinoChannelApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.GET("/casinoHandle", ginHandler.Handle(p.CasinoHandle))
}

// @Tags Casion
// @Summary 娛樂城遊戲API
// @Produce  application/json
// @param s query string true "命令"
// @param account query string true "用戶帳號"
// @param password query string true "用戶密碼"
// @param lang query string false "語系" default(en-us)
// @param coin query string false "充值分數" default(0)
// @Success 200 {object} model.CasinoHandleResponse "返回成功或失敗訊息"
// @Router /casino/casinoHandle [get]
func (p *CasinoChannelApi) CasinoHandle(c *ginweb.Context) {
	/*
	   固定代理編號
	   用戶輸入帳號、密碼(加密後傳送)
	*/
	if gameServerInfoStorage, ok := global.GlobalStorage.SelectFromDB(definition.STORAGE_KEY_GAMESERVERINFO); ok {
		resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
		cs, ok := resp["state"].(float64)
		if !ok {
			c.Logger().Error("getGameServerState code:%d", definition.ERROR_CODE_ERROR_DATABASE)
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		curState := int(cs)

		if curState != definition.GS_STATE_OPEN {
			c.OkWithCode(definition.ERROR_CODE_ERROR_GAME_SERVER_IN_MAINTENANCE)
			return
		}
	}

	var request model.CasinoHandleRequest

	agentIdStr := model.Casino_Agent_IdStr
	commandStr := c.Request.URL.Query().Get("s")
	account := c.Request.URL.Query().Get("account")
	password := c.Request.URL.Query().Get("password")
	// keyStr := c.Request.URL.Query().Get("key")
	langStr := c.Request.URL.Query().Get("lang")
	coin := c.Request.URL.Query().Get("coin")

	// 檢查命令
	if len(commandStr) <= 0 || len(account) <= 0 || len(password) <= 0 || !utils.EnglishAndNumber5To20.MatchString(account) {
		c.Fail("assign value failed")
		return
	}

	command := utils.ToInt(commandStr)

	// 一律轉小寫
	account = strings.ToLower(account)

	if langStr == "" {
		langStr = definition.LANG_TYPE_EN
	}

	// check agent id is exist
	// TODO: agent data save in redis
	agent := global.AgentCache.Get(utils.ToInt(agentIdStr))
	// 20220928 add. 代理是開發者，下面不能有玩家
	if agent == nil || agent.TopAgentId == -1 {
		c.Fail("agent id failed")
		return
	}

	// // check api ip whitelist
	// apiIpWhiteList := agent.ApiIPWhitelist
	// clientIp := c.ClientIP()
	// backendIp := c.GetBackendIp()

	// if clientIp != LOCALIP && clientIp != backendIp {
	// 	if len(apiIpWhiteList) == 0 || !middleware.CheckAgentIpWhitelist(clientIp, apiIpWhiteList) {
	// 		c.Fail("ip is not allow")
	// 		return
	// 	}
	// }

	request.Agent = agentIdStr
	request.AgentCode = agent.Code
	request.LevelCode = agent.LevelCode
	request.Currency = agent.Currency
	request.Lang = langStr
	request.Account = account
	request.Password = password
	request.ToCoin = 1
	request.Coin = coin

	lock := p.GetUserLock(account)

	if fn, ok := p.ControllerHandle[command]; ok {
		isSuccess, reponse, err := fn(c.Logger(), c.DB(), c.Redis(), request, lock)
		if err != nil {
			c.Fail(err.Error())
			return
		} else {
			if isSuccess {
				c.Ok(reponse)
				return
			} else {
				c.Fail(reponse)
			}
		}
	} else {
		c.Fail("invalid command")
		return
	}
}
