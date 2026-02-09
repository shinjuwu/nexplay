package api

import (
	"database/sql"
	"monitorservice/pkg/captcha"
	"monitorservice/pkg/config"
	"monitorservice/pkg/encrypt/base64url"
	"monitorservice/pkg/jwt"
	"sync"
)

type MyContext struct {
	db              *sql.DB
	config          config.Config
	loginTokenStore *sync.Map
	logger          ILogger
	jwt             *jwt.JwtManager
	captcha         captcha.ICaptcha
	claims          *jwt.CustomClaims
}

func NewMyContext(db *sql.DB, config config.Config, loginTokenStore *sync.Map, logger ILogger, jwt *jwt.JwtManager, captcha captcha.ICaptcha, claims *jwt.CustomClaims) IContext {
	return &MyContext{
		db:              db,
		config:          config,
		loginTokenStore: loginTokenStore,
		logger:          logger,
		jwt:             jwt,
		captcha:         captcha,
		claims:          claims,
	}
}

func (c *MyContext) DB() *sql.DB {
	return c.db
}

func (c *MyContext) Config() config.Config {
	return c.config
}

func (c *MyContext) LoginTokenStore() *sync.Map {
	return c.loginTokenStore
}

func (c *MyContext) Logger() ILogger {
	return c.logger
}

func (p *MyContext) JWT() *jwt.JwtManager {
	return p.jwt
}

func (p *MyContext) Captcha() captcha.ICaptcha {
	return p.captcha
}

func (c *MyContext) Claims() *jwt.CustomClaims {
	return c.claims
}

func (c *MyContext) GetUsername() string {
	creatorBytes, err := base64url.Decode(c.claims.Username)
	if err != nil {
		return ""
	}
	return string(creatorBytes)
}

func (c *MyContext) GetNickname() string {
	creatorBytes, err := base64url.Decode(c.claims.Nickname)
	if err != nil {
		return ""
	}
	return string(creatorBytes)
}
