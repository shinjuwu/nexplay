package api

import (
	"chatservice/pkg/config"
	"chatservice/pkg/melody"
	"chatservice/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	// blankPagePath = "/"
	preApiPath = "/chatservice.api/v1/"
	preWsPath  = "/chatservice.ws"
)

type ApiServer struct {
	g               *gin.Engine
	m               *melody.Melody
	db              *sql.DB
	config          config.Config
	sessionManager  SessionManager
	pipeline        *Pipeline
	loginTokenStore *sync.Map
	logger          ILogger
	imBotHandler    *IMBot
}

func NewApiServer(logger ILogger, db *sql.DB, config config.Config, sessionManager SessionManager, pipeline *Pipeline) *ApiServer {

	router := gin.Default()

	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	ginConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

	router.Use(cors.New(ginConfig))

	return &ApiServer{
		g:               router,
		m:               melody.New(),
		db:              db,
		config:          config,
		sessionManager:  sessionManager,
		pipeline:        pipeline,
		loginTokenStore: new(sync.Map),
		logger:          logger,
		imBotHandler:    NewIMBot(),
	}
}

func (p *ApiServer) StartApiServer() error {

	p.ApiHandlerRegister()
	p.WsHandlerRegister()

	return p.g.Run(fmt.Sprintf("%v:%d", p.config.GetSocket().Address, p.config.GetSocket().Port))
}

func (p *ApiServer) ApiHandlerRegister() {

	authorized := p.g.Group("/", gin.BasicAuth(gin.Accounts{
		p.config.GetSocket().ServerKey: "",
	}))

	authorized.GET(preApiPath+"login", func(c *gin.Context) {
		username := utils.ToString(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_USER_NAME), "")
		agentId := utils.ToInt(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_AGENT_ID), 0)
		platform := utils.ToString(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_PLATFORM), "")
		if username == "" || platform == "" || agentId == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": ErrCodeCommonFailed,
				"msg":  "parse param failed.",
				"data": nil,
			})
		} else {
			Login(p.db, c, p.loginTokenStore, username, platform, agentId)
		}
	})

	authorized.POST(preApiPath+"broadcast", func(c *gin.Context) {
		platform := utils.ToString(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_PLATFORM), "")
		msgData := utils.ToString(c.Request.URL.Query().Get("data"), "")
		subject := utils.ToString(c.Request.URL.Query().Get("subject"), "")
		if platform == "" || msgData == "" || subject == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": ErrCodeCommonFailed,
				"msg":  "parse param failed.",
				"data": nil,
			})
		} else {
			Broadcast(p.db, c, p.sessionManager, platform, msgData, subject)
		}
	})

	authorized.POST(preApiPath+"message", func(c *gin.Context) {
		platform := utils.ToString(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_PLATFORM), "")
		msgData := utils.ToString(c.Request.URL.Query().Get("data"), "")
		subject := utils.ToString(c.Request.URL.Query().Get("subject"), "")
		username := utils.ToString(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_USER_NAME), "")
		agentId := utils.ToInt(c.Request.URL.Query().Get(TEQUILA_CTX_RUNTIME_AGENT_ID), 0)
		if platform == "" || msgData == "" || subject == "" || username == "" || agentId == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": ErrCodeCommonFailed,
				"msg":  "parse param failed.",
				"data": nil,
			})
		} else {
			Message(p.db, c, p.sessionManager, agentId, platform, username, msgData, subject)
		}
	})

	authorized.GET(preApiPath+"im", func(c *gin.Context) {
		jsonData := utils.ToString(c.Request.URL.Query().Get("data"), "")
		idx := utils.ToString(c.Request.URL.Query().Get("im"), "")
		imIdx := utils.StringToInt(idx, -1)
		if jsonData == "" || imIdx < 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": ErrCodeCommonFailed,
				"msg":  "parse param failed.",
				"data": nil,
			})
		} else {
			IM(p.db, c, p.imBotHandler, imIdx, jsonData)
		}
	})
}

func (p *ApiServer) WsHandlerRegister() {

	// 註冊 websocket
	p.g.GET(preWsPath, NewSocketWsAcceptor(p.config, p.g, p.m, p.loginTokenStore))

	// 連線事件觸發
	p.m.HandleConnect(func(s *melody.Session) {
		log.Println("HandleConnect")
		p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_CONNECT, s, &Envelope{})
	})

	// 斷線事件觸發
	p.m.HandleDisconnect(func(s *melody.Session) {
		log.Println("HandleDisconnect")
		p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_DISCONNECT, s, &Envelope{})
	})

	// 連線錯誤事件觸發
	p.m.HandleError(func(s *melody.Session, err error) {
		log.Println("HandleError")
		var mm Envelope
		mm.Payload = err.Error()
		p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_ERROR, s, &mm)
	})

	// 連線關閉事件觸發
	p.m.HandleClose(func(s *melody.Session, code int, msg string) error {
		log.Println("HandleClose")
		var mm Envelope
		mm.Code = code
		mm.Payload = msg
		p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_CLOSE, s, &mm)
		return nil
	})

	// 發送訊息廣播
	p.m.HandleMessage(func(s *melody.Session, msg []byte) {
		var mm Envelope
		err := json.Unmarshal(msg, &mm)
		if err != nil {
			mm.Payload = err.Error()
			p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_ERROR, s, &mm)
			return
		}
		mm.Id = strings.ToLower(mm.Id)
		success, err := p.pipeline.ProcessRequest(mm.Id, s, &mm)
		if !success {
			mm.Payload = ErrMessageUnknowRPC
			p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_ERROR, s, &mm)
			return
		}

		if err != nil {
			mm.Payload = err.Error()
			p.pipeline.ProcessInternalRequest(TEQUILA_RUNTIME_EVENT_ERROR, s, &mm)
			return
		}
	})
}

func (p *ApiServer) Stop() {

}

func (p *ApiServer) Healthcheck() (string, error) {

	return "", nil
}

/*
 */
func (p *ApiServer) Listen(address string) {
	p.g.Run(address)
}
