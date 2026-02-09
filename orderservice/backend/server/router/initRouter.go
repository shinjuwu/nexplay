package router

import (
	api_culster "backend/api"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/config"
	"backend/server/global"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(config config.Config, ginHandler *ginweb.GinHanlder) (*gin.Engine, error) {

	// 初始化必要資料
	if err := global.InitGlobalData(ginHandler.Logger, ginHandler.DB, ginHandler.Redis, ginHandler.JobScheduler); err != nil {
		ginHandler.Logger.Info("InitGlobalData() error: %v", err)
		return nil, err
	}

	go global.T_execTimeForGetPlaylogCommon(ginHandler.Logger, ginHandler.DB)
	scheduleToBackup := global.ScheduleToBackupCache.GetAll()
	if scheduleToBackup != nil {
		go global.T_execScheduleToBackup(ginHandler.Logger, ginHandler.DB, scheduleToBackup)
	}

	gin.SetMode(config.GetApp().Env)

	router := gin.Default()

	// 設置受信任代理,如果不設置默認信任所有代理，不安全
	// router.SetTrustedProxies([]string{"172.19.10.12"})

	//跨網域設定
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	router.Use(cors.New(ginConfig))

	router.Use(middleware.RequestLimit(middleware.MAX_ALLOWED))
	router.Use(middleware.Logger())
	// router.Use(middleware.CheckUserRole(config, &global.PermissionList, global.AgentPermissionCache, ginHandler.Jwt))
	router.Use(middleware.SetExceptionHandler())

	ApiRouterCluster := api_culster.NewApiCluster(config)
	ApiRouterCluster.RouterGroupRegister(router, ginHandler)

	// enable swagger API doc
	// Swagger 路由導入
	if config.GetApp().LoadSwagger {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router, nil
}
