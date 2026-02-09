package game

import (
	"backend/api/game/controller"
	"backend/api/game/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/redis"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	CONST_ACCOUNT_LENGTH_MIN = 6
	CONST_ACCOUNT_LENGTH_MAX = 64

	LOCALIP = "127.0.0.1"
)

// type HandleCommand int

type ChannelHandleFunc func(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest, lock *sync.Mutex) (bool, *model.ChannelHandleResponse, error)

type ChannelHandleApi struct {
	BasePath         string
	ControllerHandle map[int]ChannelHandleFunc
	userLockMap      *sync.Map
}

func NewGameChannelApi(basePath string) api_cluster.IApiEach {
	return &ChannelHandleApi{
		BasePath: basePath,
		// define API func
		ControllerHandle: map[int]ChannelHandleFunc{
			model.ChannelHandle_GameLogin:             controller.GameLogin,
			model.ChannelHandle_CheckCoinOutLimit:     controller.CheckCoinOutLimit,
			model.ChannelHandle_CoinIn:                controller.CoinIn,
			model.ChannelHandle_CoinOut:               controller.CoinOut,
			model.ChannelHandle_CheckUserOrder:        controller.CheckUserOrder,
			model.ChannelHandle_CheckUserOnlineStatus: controller.CheckUserOnlineStatus,
			model.ChannelHandle_CheckUserCoin:         controller.CheckUserCoin,
			model.ChannelHandle_KickUser:              controller.KickUser,
			model.ChannelHandle_CheckAgentCoin:        controller.CheckAgentCoin,
		},
		userLockMap: new(sync.Map),
	}
}

func (p *ChannelHandleApi) GetGroupPath() string {
	return p.BasePath
}

// func (p *ChannelHandleApi) CreateLock(thirdAccount string) {
// 	p.userLockMap.Store(thirdAccount, new(sync.Mutex))
// }

// if get nil, create new one
func (p *ChannelHandleApi) GetUserLock(thirdAccount string) *sync.Mutex {
	tmp, _ := p.userLockMap.LoadOrStore(thirdAccount, new(sync.Mutex))
	return tmp.(*sync.Mutex)
}

// func (p *ChannelHandleApi) DelLock(thirdAccount string) {
// 	p.userLockMap.Delete(thirdAccount)
// }

func (p *ChannelHandleApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.GET("/channelHandle", ginHandler.Handle(p.ChannelHandle))
}

// @Tags Channel
// @Summary 遊戲API
// @Produce  application/json
// @param agent query string true "代理編號（平台提供）"
// @param timestamp query string true "時間戳(Unix 時間戳帶上毫秒)" default(1234567890123)
// @param param query string true "參數加密字符串(先用AES加密, 再使用base64加密成字串)"
// @param key query string true "Md5 校驗字符串(Encrypt.MD5(agent+timestamp+MD5Key))" default(hongkong3345678)
// @param lang query string false "語系" default(en-us)
// @Success 200 {object} model.ChannelHandleResponse "返回成功或失敗訊息"
// @Router /channel/channelHandle [get]
func (p *ChannelHandleApi) ChannelHandle(c *ginweb.Context) {

	if gameServerInfoStorage, ok := global.GlobalStorage.SelectFromDB(definition.STORAGE_KEY_GAMESERVERINFO); ok {
		resp := utils.ToMap([]byte(gameServerInfoStorage.Value))
		cs, ok := resp["state"].(float64)
		if !ok {
			c.Logger().Error("getGameServerState code:%d", model.Response_QueryRow_Error)
			c.OkWithCode(model.Response_QueryRow_Error)
			return
		}

		curState := int(cs)

		if curState != definition.GS_STATE_OPEN {
			c.OkWithCode(model.Response_GameNotOpen_Error)
			return
		}
	}

	var request model.ChannelHandleRequest

	agentIdStr := c.Request.URL.Query().Get("agent")
	timestampStr := c.Request.URL.Query().Get("timestamp")
	paramStr := c.Request.URL.Query().Get("param")
	keyStr := c.Request.URL.Query().Get("key")
	langStr := c.Request.URL.Query().Get("lang")

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

	// check api ip whitelist
	apiIpWhiteList := agent.ApiIPWhitelist
	clientIp := c.ClientIP()
	backendIp := c.GetBackendIp()

	if clientIp != LOCALIP && clientIp != backendIp {
		if len(apiIpWhiteList) == 0 || !middleware.CheckAgentIpWhitelist(clientIp, apiIpWhiteList) {
			c.Fail("ip is not allow")
			return
		}
	}

	// assign and check
	if isSuccess, msg := request.Assign(
		agentIdStr,
		agent.Code,
		agent.LevelCode,
		timestampStr,
		paramStr,
		keyStr,
		agent.AesKey,
		agent.Md5Key,
		agent.Currency,
		langStr); !isSuccess {
		c.Fail(msg)
		return
	}

	// 檢查命令
	if _, ok := request.ParamMap["s"]; !ok {
		c.Fail("assign value failed")
		return
	} else if len(request.ParamMap["s"][0]) == 0 {
		c.Fail("command is illegal")
		return
	}

	command := utils.ToInt(request.ParamMap["s"][0], -1)
	account := ""
	if command == model.ChannelHandle_CheckAgentCoin ||
		command == model.ChannelHandle_CheckUserOrder {
		account = agentIdStr
	} else {
		// 檢查帳號參數
		if _, ok := request.ParamMap["account"]; !ok {
			c.Fail("assign value failed")
			return
		} else if len(request.ParamMap["account"][0]) == 0 {
			c.Fail("account is illegal")
			return
		} else if account = utils.ToString(request.ParamMap["account"][0], ""); !utils.EnglishAndNumber5To20.MatchString(account) {
			// 檢查帳號是否有中文
			c.Fail("account is illegal")
			return
		}

		// 一律轉小寫
		request.ParamMap["account"][0] = strings.ToLower(account)
	}

	if (command == model.ChannelHandle_CoinIn ||
		command == model.ChannelHandle_CoinOut) &&
		agent.WalletType == definition.AGENT_WALLET_SINGLE {
		c.Fail("wallet type is illegal")
		return
	}

	switch command {
	case model.ChannelHandle_GameLogin:
		fallthrough
	case model.ChannelHandle_CoinIn:
		fallthrough
	case model.ChannelHandle_CoinOut:
		strMoney := request.ParamMap["money"][0]
		if !controller.CheckFloatPlaces(strMoney, 2) {
			c.Fail("no more than 2 decimal places")
			return
		}

		// get exchange ToCoin value
		exchangeData := global.ExchangeDataCache.Get(agent.Currency)
		request.Currency = exchangeData.Currency
		request.ToCoin = exchangeData.ToCoin
	}

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
