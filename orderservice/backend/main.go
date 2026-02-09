package main

import (
	"backend/internal/ginweb"
	"backend/pkg/captcha"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/job"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"backend/pkg/redis"
	"backend/server/router"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	// _ "backend/docs"
)

// @EnableOpenApi
// @title Swagger DCC backend API
// @version 0.0.1
// @description This is a backend service
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Dcc-Token
// @BasePath /
func main() {

	// 初始化 logger 物件
	tmpLogger := logger.NewJSONLogger(os.Stdout, zapcore.InfoLevel, logger.JSONFormat)

	// 讀取設定檔
	yamlConfig := config.ReadYamlFile(tmpLogger)

	// setting log object
	newLogger, startupLogger := logger.SetupLogging(tmpLogger, yamlConfig)

	startupLogger.Info("Backend service starting")
	startupLogger.Info("Name", zap.String("name", yamlConfig.GetName()))
	startupLogger.Info("Data directory", zap.String("path", yamlConfig.GetDataDir()))
	startupLogger.Info("Number of databases", zap.Int("number", len(yamlConfig.GetDatabase().ConnInfo)))
	startupLogger.Info("Database connections", zap.Strings("dsns", yamlConfig.GetDatabase().ConnInfo))
	startupLogger.Info("Redis address", zap.String("dsns", yamlConfig.GetRedis().Address))
	startupLogger.Info("Env mode", zap.String("is open", yamlConfig.GetApp().Env))

	if yamlConfig.GetApp().LoadDBMigration {
		// db migrate
		m, err := database.NewSMigration(yamlConfig)
		if err != nil {
			startupLogger.Error("Database migrate", zap.String("sourceURL", m.SourceURL()), zap.String("databaseURL", m.DatabaseURL()), zap.Error(err))
			return
		}

		ver, dirty, err := m.Up()
		if err != nil {
			if err != database.ErrNoChange {
				startupLogger.Error("Database migrate up", zap.Uint("version", ver), zap.Bool("dirty", dirty), zap.Error(err))
				return
			} else {
				startupLogger.Info("Database migrate up no change")
			}
		}
		startupLogger.Info("Database migrate information", zap.Uint("version", ver), zap.Bool("dirty", dirty))
	}

	// 初始化 DB 物件
	db := database.DBConnectDispatcher(startupLogger, yamlConfig)
	startupLogger.Info("Database information", zap.String("version", db.Version()))

	// 初始化 redis DB 物件
	rdb := redis.NewRedisCliect(startupLogger, yamlConfig)
	startupLogger.Info("Redis successfully started")

	jwt := jwt.NewJwtManager(yamlConfig)

	runtimeLogger := logger.NewRuntimeGoLogger(newLogger)
	captcha := captcha.NewCaptcha(yamlConfig)
	jobScheduler := job.NewJobScheduler()
	ginHandler := ginweb.NewGinHanlder(runtimeLogger, db, rdb, jwt, captcha, jobScheduler, yamlConfig)

	router, err := router.InitRouter(yamlConfig, ginHandler)
	if router == nil || err != nil {
		startupLogger.Error("InitRouter() failed, router is nil", zap.Error(err))
		return
	}

	addr := fmt.Sprintf(":%d", yamlConfig.GetApp().Addr)

	if err := router.Run(addr); err != nil {
		panic("Router run failed.")
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

	startupLogger.Info("Shutdown complete")

	os.Exit(0)
}
