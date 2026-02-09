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

	"github.com/gin-gonic/gin"
)

type GameRecodrHandleFunc func(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, data model.ChannelHandleRequest) (bool, *model.ChannelHandleResponse, error)

type GetGameRecordApi struct {
	BasePath         string
	ControllerHandle map[int]GameRecodrHandleFunc
}

func NewGetGameRecordApi(basePath string) api_cluster.IApiEach {
	return &GetGameRecordApi{
		BasePath: basePath,
		// define API func
		ControllerHandle: map[int]GameRecodrHandleFunc{
			model.GetRecordHandle_CheckGameRecord: controller.CheckGameRecord,
		},
	}
}

func (p *GetGameRecordApi) GetGroupPath() string {
	return p.BasePath
}

func (p *GetGameRecordApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.GET("/getRecordHandle", ginHandler.Handle(p.ChannelHandle))
}

// @Tags Channel
// @Summary 第三方取遊戲記錄
// @Produce  application/json
// @param agent query string true "代理編號（平台提供）"
// @param timestamp query string true "時間戳(Unix 時間戳帶上毫秒)" default(1234567890123)
// @param param query string true "參數加密字符串(aes 加密)"
// @param key query string true "Md5 校驗字符串(Encrypt.MD5(agent+timestamp+MD5Key))"
// @Success 200 {object} model.ChannelHandleResponse "返回成功或失敗訊息"
// @Router /record/getRecordHandle [get]
func (p *GetGameRecordApi) ChannelHandle(c *ginweb.Context) {

	var request model.ChannelHandleRequest

	agentIdStr := c.Request.URL.Query().Get("agent")
	timestampStr := c.Request.URL.Query().Get("timestamp")
	paramStr := c.Request.URL.Query().Get("param")
	keyStr := c.Request.URL.Query().Get("key")

	// check agent id is exist
	agent := global.AgentCache.Get(utils.ToInt(agentIdStr))
	if agent == nil {
		c.Fail("agent id failed")
		return
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
		"en-us"); !isSuccess {
		c.Fail(msg)
		return
	}

	command := utils.ToInt(request.ParamMap["s"][0], -1)

	if fn, ok := p.ControllerHandle[command]; ok {
		isSuccess, reponse, err := fn(c.Logger(), c.DB(), c.Redis(), request)
		if err != nil {
			c.Fail(err)
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
