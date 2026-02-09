package api_cluster

import (
	"backend/internal/ginweb"

	"github.com/gin-gonic/gin"
)

type IApiEach interface {
	GetGroupPath() string
	RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder)
}
