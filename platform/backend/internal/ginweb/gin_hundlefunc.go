package ginweb

import (
	"backend/pkg/captcha"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/job"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"backend/pkg/redis"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *Context)

type GinHanlder struct {
	Logger       *logger.RuntimeGoLogger
	DB           database.IDatabase // DB 物件
	Redis        redis.IRedisCliect // redis DB 物件
	Jwt          *jwt.JwtManager
	Captcha      captcha.ICaptcha
	JobScheduler *job.JobScheduler
	Config       config.Config
}

func NewGinHanlder(logger *logger.RuntimeGoLogger, db database.IDatabase, rdb redis.IRedisCliect, jwt *jwt.JwtManager, captcha captcha.ICaptcha, jobScheduler *job.JobScheduler, config config.Config) *GinHanlder {
	return &GinHanlder{
		Logger:       logger,
		DB:           db,
		Redis:        rdb,
		Jwt:          jwt,
		Captcha:      captcha,
		JobScheduler: jobScheduler,
		Config:       config,
	}
}

func (g *GinHanlder) Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		ctx := NewContext(g.Logger, c, g.DB, g.Redis, g.Jwt, g.Captcha, g.JobScheduler, ip, g.Config.GetApp().Env == "debug", g.Config.GetBackend().Address)

		h(ctx)
	}
}
