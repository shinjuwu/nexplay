package api_cluster

import (
	"backend/internal/ginweb"

	"github.com/gin-gonic/gin"
)

type IApiCluster interface {
	RouterGroupRegister(router *gin.Engine, ginHandler *ginweb.GinHanlder)
}

type IApiGroupCluster interface {
	RegisterApiGroup(basePath, groupPath, fn func(c *ginweb.Context), middleware ...gin.HandlerFunc)
	RegisterApi(relativePath, method string)
}

type ApiGroupCluster struct {
	BasePath     string
	GroupPath    string
	RelativePath string
}
