package game

import (
	"backend/internal/api_cluster"
	"backend/internal/ginweb"

	"github.com/gin-gonic/gin"
)

type IntercomOrderApi struct {
	BasePath string
}

/** create new api group instance
 * basePath: path of base group router
 */
func NewIntercomOrderApi(basePath string) api_cluster.IApiEach {
	return &IntercomOrderApi{
		BasePath: basePath,
	}
}

/** get base group path */
func (p *IntercomOrderApi) GetGroupPath() string {
	return p.BasePath
}

/** 註冊 api */
func (p *IntercomOrderApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.POST("/rechargeorder", ginHandler.Handle(p.RechargeOrder))
}

// @Tags IntercomOrderApi
// @Summary order api service of intercom
// @Description 注單服務內部補單使用
// @Produce  application/json
// @Param data body model.PlayLogRequest true "代理識別碼, 遊戲紀錄(json), 總變動分數..."
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /intercomorder/rechargeorder [post]
func (p *IntercomOrderApi) RechargeOrder(c *ginweb.Context) {

	c.Ok("")
}
