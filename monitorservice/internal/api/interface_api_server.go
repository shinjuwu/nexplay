package api

import (
	"monitorservice/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type IApiServer interface {
	JWT() *jwt.JwtManager
	GetApiRouter() *gin.Engine
	StartApiServer(apiHandler ...IApiHandler) error
	ApiHandlerRegister(apiHandler ...IApiHandler)
	WsHandlerRegister()
	Stop()
	Healthcheck() (string, error)
	Listen(string)
	ApiHandle(h ApiFunc) gin.HandlerFunc
}
