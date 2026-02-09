package main

import (
	"monitorservice/cmd/baseserver/api_module"
	"monitorservice/cmd/baseserver/api_module/data"
	"monitorservice/internal/api"
	"monitorservice/internal/module"
	"monitorservice/migrate"
	"monitorservice/pkg/captcha"
	"monitorservice/pkg/config"
	"monitorservice/pkg/database"
	"monitorservice/pkg/jwt"
	"monitorservice/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "monitorservice/cmd/baseserver/docs"
)

// @EnableOpenApi
// @title Swagger Monitor Service API
// @version 0.0.1
// @description This is a backend service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Token
// @BasePath /
func main() {

	// 初始化 logger 物件
	tmpLogger := logger.NewJSONLogger(os.Stdout, zapcore.InfoLevel, logger.JSONFormat)

	// 讀取設定檔
	yamlConfig := config.ReadYamlFile(tmpLogger)

	// setting log object
	newLogger, startupLogger := logger.SetupLogging(tmpLogger, yamlConfig)

	startupLogger.Info("monitorservice service starting")
	startupLogger.Info("Name", zap.String("name", yamlConfig.GetName()))
	startupLogger.Info("Data directory", zap.String("path", yamlConfig.GetDataDir()))
	startupLogger.Info("Number of databases", zap.Int("number", len(yamlConfig.GetDatabase().ConnInfo)))
	startupLogger.Info("Database connections", zap.Strings("dsns", yamlConfig.GetDatabase().ConnInfo))
	// startupLogger.Info("Redis address", zap.String("dsns", yamlConfig.GetRedis().Address))
	startupLogger.Info("Env mode", zap.String("is open", yamlConfig.GetApp().Env))

	// 初始化 DB 物件
	db := database.DBConnectDispatcher(startupLogger, yamlConfig)
	startupLogger.Info("Database information", zap.String("version", db.Version()))

	// default migrate action is up
	migrate.StartUpCheck(startupLogger, db.GetDefaultDB())
	if len(os.Args) >= 2 {
		migrate.Parse(os.Args[1:], startupLogger, db.GetDefaultDB())
	}

	sessionManager := api.NewLocalSessionManager()

	runtimeLogger := logger.NewRuntimeGoLogger(newLogger)

	tequila := api.NewTequila(runtimeLogger, db.GetDefaultDB(), sessionManager)

	// 自定義 socket 模組.start
	mod := module.NewModule(tequila)
	err := mod.InitModule(runtimeLogger, db.GetDefaultDB(), tequila)
	if err != nil {
		runtimeLogger.Fatal("Error looking up InitModule function in tequila module", zap.Error(err))
		return
	}
	// 自定義 socket 模組.end
	jwt := jwt.NewJwtManager(yamlConfig)

	captcha := captcha.NewCaptcha(yamlConfig)

	pipeline := api.NewPipeline(runtimeLogger, db.GetDefaultDB(), yamlConfig, sessionManager, tequila)

	apiServer := api.NewApiServer(runtimeLogger, db.GetDefaultDB(), yamlConfig, sessionManager, pipeline, jwt, captcha)

	err = data.InitData(runtimeLogger, db.GetDefaultDB())
	if err != nil {
		runtimeLogger.Fatal("Error looking up InitData function in data module", zap.Error(err))
		return
	}

	// 自定義 api 模組.start
	apiRouter := apiServer.GetApiRouter()
	if apiRouter == nil {
		runtimeLogger.Fatal("Error looking up GetApiRouter function in ApiServer module")
		return
	}
	routerGroup := apiRouter.Group("/")

	loginHandler := api_module.NewApiHandler(routerGroup, apiServer.JWT())
	// 自定義 api 模組.end

	err = apiServer.StartApiServer(loginHandler)
	if err != nil {
		// panic("Router run failed.")
		runtimeLogger.Fatal("Router run failed.", zap.Error(err))
		return
	}

	// Respect OS stop signals.
	quit := make(chan os.Signal, 2)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	startupLogger.Info("Startup done")

	// Wait for a termination signal.
	<-quit

	graceSeconds := yamlConfig.GetShutdownGraceSec()

	// If a shutdown grace period is allowed, prepare a timer.
	var timer *time.Timer
	timerCh := make(<-chan time.Time, 1)
	if graceSeconds != 0 {
		timer = time.NewTimer(time.Duration(graceSeconds) * time.Second)
		timerCh = timer.C
		startupLogger.Info("Shutdown started - use CTRL^C to force stop server", zap.Int("grace_period_sec", graceSeconds))
	} else {
		// No grace period.
		startupLogger.Info("Shutdown started")
	}

	// Stop any running authoritative matches and do not accept any new ones.
	select {
	case <-timerCh:
		// Timer has expired, terminate matches immediately.
		startupLogger.Info("Shutdown grace period expired")
	case <-quit:
		// A second interrupt has been received.
		startupLogger.Info("Skipping graceful shutdown")
	}
	if timer != nil {
		timer.Stop()
	}

	// Gracefully stop remaining server components.
	// TIP: any program you want to end can be put here.
	apiServer.Stop()

	startupLogger.Info("Shutdown complete")

	os.Exit(0)
}
