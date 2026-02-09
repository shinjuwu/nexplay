package api

import (
	"github.com/gin-gonic/gin"
)

type IApiHandler interface {
	RouterGroup() *gin.RouterGroup
	ApiHandleRegister(ApiHandleFunc)
}
