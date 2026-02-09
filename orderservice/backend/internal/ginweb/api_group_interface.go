package ginweb

import "github.com/gin-gonic/gin"

type IGinGroupApi interface {
	RegisterApiRouter(g *gin.RouterGroup, ginHandler *GinHanlder)
}
