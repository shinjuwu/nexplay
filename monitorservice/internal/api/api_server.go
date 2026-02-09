package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"monitorservice/pkg/captcha"
	"monitorservice/pkg/config"
	"monitorservice/pkg/jwt"
	"monitorservice/pkg/melody"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "monitorservice/cmd/baseserver/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ApiFunc func(IContext, string) (string, error)

type ApiHandleFunc func(ApiFunc) gin.HandlerFunc

var (
	// // login page
	// loginPreApiPath = "/login.api/v1/"

	// monitor service
	msPreWsPath = "/monitorservice.ws"
)

type ApiServer struct {
	g               *gin.Engine // api
	s               *gin.Engine
	m               *melody.Melody
	db              *sql.DB
	config          config.Config
	sessionManager  SessionManager
	pipeline        *Pipeline
	loginTokenStore *sync.Map
	logger          ILogger
	apiHandles      []IApiHandler
	jwt             *jwt.JwtManager
	captcha         captcha.ICaptcha
}

func NewApiServer(logger ILogger, db *sql.DB, config config.Config, sessionManager SessionManager, pipeline *Pipeline, jwt *jwt.JwtManager, captcha captcha.ICaptcha) IApiServer {

	router := gin.Default()

	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	ginConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}

	router.Use(cors.New(ginConfig))

	// enable swagger API doc
	// Swagger 路由導入
	if config.GetApp().LoadSwagger {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srouter := gin.Default()

	srouter.Use(cors.New(ginConfig))

	return &ApiServer{
		g:               router,
		s:               srouter,
		m:               melody.New(),
		db:              db,
		config:          config,
		sessionManager:  sessionManager,
		pipeline:        pipeline,
		loginTokenStore: new(sync.Map),
		logger:          logger,
		apiHandles:      make([]IApiHandler, 0),
		jwt:             jwt,
		captcha:         captcha,
	}
}

func (p *ApiServer) JWT() *jwt.JwtManager {
	return p.jwt
}

func (p *ApiServer) GetApiRouter() *gin.Engine {
	return p.g
}

func (p *ApiServer) StartApiServer(apiHandler ...IApiHandler) error {

	// global middleware
	p.g.Use(SetExceptionHandler(p.logger))

	p.ApiHandlerRegister(apiHandler...)

	go p.g.Run(fmt.Sprintf(":%d", p.config.GetApp().Addr))

	p.s.Use(SetExceptionHandler(p.logger))

	p.WsHandlerRegister()

	go p.s.Run(fmt.Sprintf("%v:%d", p.config.GetSocket().Address, p.config.GetSocket().Port))

	return nil
}

func (p *ApiServer) ApiHandlerRegister(apiHandler ...IApiHandler) {

	p.apiHandles = append(p.apiHandles, apiHandler...)

	if len(p.apiHandles) > 0 {
		for _, apiHandle := range p.apiHandles {
			apiHandle.ApiHandleRegister(p.ApiHandle)
		}
	}
}

func (p *ApiServer) WsHandlerRegister() {

	// 註冊 websocket
	p.s.GET(msPreWsPath, NewSocketWsAcceptor(p.config, p.g, p.m, p.loginTokenStore))

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

func (p *ApiServer) ApiHandle(h ApiFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonMap := make(map[string]interface{}, 0)
		c.ShouldBind(&jsonMap)

		jsonData, err := json.Marshal(jsonMap)
		if err != nil {
			Fail(c, err.Error())
			return
		}

		claims := new(jwt.CustomClaims)

		tmp, isExits := c.Get("claims")
		if isExits {
			claims = tmp.(*jwt.CustomClaims)
		}

		processST := time.Now()

		myContect := NewMyContext(p.db, p.config, p.loginTokenStore, p.logger, p.jwt, p.captcha, claims)
		resJson, err := h(myContect, string(jsonData))

		processET := time.Now()
		resultCode := http.StatusOK
		if err != nil {
			resultCode = http.StatusBadRequest
		}
		p.logger.Printf("%d | %v | %s | %s %s", resultCode, processET.Sub(processST), c.ClientIP(), c.Request.Method, c.FullPath())
		if err == nil {
			if resJson == "" {
				resJson = "{}"
			}
			Ok(c, resJson)
			return
		} else {
			Fail(c, err.Error())
			return
		}
	}
}
