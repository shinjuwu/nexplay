package system

import (
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/server/global"

	"github.com/gin-gonic/gin"
)

/*
服務狀態
*/

type SystemServiceStatusApi struct {
	BasePath string
}

func NewSystemServiceStatusApi(basePath string) api_cluster.IApiEach {
	return &SystemServiceStatusApi{
		BasePath: basePath,
	}
}

func (p *SystemServiceStatusApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemServiceStatusApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.GET("/getjobshedulerlist", ginHandler.Handle(p.GetJobShedulerList))
}

// @Tags 服務狀態
// @Summary Get job scheduler rutnime info list
// @Description 此接口用來取得當前 job 的資訊清單
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/servicestatus/getjobshedulerlist [get]
func (p *SystemServiceStatusApi) GetJobShedulerList(c *ginweb.Context) {

	tmps := c.GetJob().Lists()
	c.Ok(tmps)
}
