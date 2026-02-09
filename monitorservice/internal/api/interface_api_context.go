package api

import (
	"database/sql"
	"monitorservice/pkg/captcha"
	"monitorservice/pkg/config"
	"monitorservice/pkg/jwt"
	"sync"
)

type IContext interface {
	DB() *sql.DB
	Config() config.Config
	LoginTokenStore() *sync.Map
	Logger() ILogger
	JWT() *jwt.JwtManager
	Captcha() captcha.ICaptcha
	Claims() *jwt.CustomClaims
	GetUsername() string
	GetNickname() string
}
