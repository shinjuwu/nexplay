package health

import (
	"backend/internal/api_cluster"
	"backend/internal/ginweb"

	"github.com/gin-gonic/gin"
)

/*
if you want create new one, copy that and paste on file.
*/

type HealthApi struct {
	BasePath string
}

/** create new api group instance
 * basePath: path of base group router
 */
func NewHealthApi(basePath string) api_cluster.IApiEach {
	return &HealthApi{
		BasePath: basePath,
	}
}

/** get base group path */
func (p *HealthApi) GetGroupPath() string {
	return p.BasePath
}

/** 註冊 api */
func (p *HealthApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.GET("/health", ginHandler.Handle(p.health))
}

// @Tags Health
// @Summary check serice isn't still live
// @Description 檢查伺服器是否活著
// @accept application/json
// @Produce application/json
// @Params
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/health/health [get]
func (p *HealthApi) health(c *ginweb.Context) {
	c.Ok("")
}
