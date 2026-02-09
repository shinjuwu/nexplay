package router

import (
	api_culster "backend/api"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/config"
	"backend/server/global"
	"net/http"

	_ "backend/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(config config.Config, ginHandler *ginweb.GinHanlder) (*gin.Engine, error) {

	// redis object create for store user text verifycode
	// if err := core.CreateRedisObject(core.REDIS_TextVerifyCode, core.REDIS_ADDRESS_DEFAULT); err != nil {
	// 	log.Printf("InitRouter(), Redis connect failed, db idx : REDIS_TextVerifyCode, error: %v", err)
	// }

	global.JobSchedulerSwitch = config.GetApp().LoadJobScheduler
	global.UpdateRedisData = config.GetApp().UpdateRedisData

	// 初始化必要資料
	if err := global.InitGlobalData(ginHandler.Logger, ginHandler.DB.GetDefaultDB(), ginHandler.Redis, ginHandler.JobScheduler); err != nil {
		ginHandler.Logger.Info("InitGlobalData() error: %v", err)
		return nil, err
	}

	if config.GetApp().LoadJobScheduler {
		go global.T_CalKillDiveState(ginHandler.Logger, ginHandler.DB.GetDefaultDB())
		go global.T_CalAutoRiskControl(ginHandler.Logger, ginHandler.DB.GetDefaultDB())
		go global.T_CalRealtimeGameUserStat(ginHandler.Logger, ginHandler.DB.GetDefaultDB(), ginHandler.Redis)
		go global.T_CalDbArchive(ginHandler.Logger, ginHandler.DB.GetDefaultDB())
		go global.T_ResendMonitorServiceData(ginHandler.Logger)
	}
	gin.SetMode(config.GetApp().Env)

	router := gin.Default()

	// 設置受信任代理,如果不設置默認信任所有代理，不安全
	// router.SetTrustedProxies([]string{"172.19.10.12"})

	//跨網域設定
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	router.Use(cors.New(ginConfig))

	// 是否載入本地端網頁資源
	if config.GetApp().LoadFront {
		// 載入靜態資源
		router.Static("/assets", "vue_front/assets")
		// 載入favicon
		router.StaticFile("/vite.svg", "vue_front/vite.svg")
		router.StaticFile("/favicon.ico", "vue_front/vite.svg")
		// 載入頁面
		router.LoadHTMLFiles("vue_front/index.html")
		// use frontend route, not gin route
		router.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "DCC Backend Service",
			})
		})
	}

	// router.Use(middleware.RequestLimit(middleware.MAX_ALLOWED))
	router.Use(middleware.Logger(ginHandler.DB.GetDefaultDB()))
	router.Use(middleware.CheckUserRole(config, &global.PermissionList, global.AgentPermissionCache, ginHandler.Jwt))
	router.Use(middleware.SetExceptionHandler(ginHandler.Logger))

	ApiRouterCluster := api_culster.NewApiCluster(config)
	ApiRouterCluster.RouterGroupRegister(router, ginHandler)

	// enable swagger API doc
	// Swagger 路由導入
	if config.GetApp().LoadSwagger {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router, nil
}
