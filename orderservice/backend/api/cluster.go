package api

import (
	"backend/api/game"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/pkg/config"

	"github.com/gin-gonic/gin"
)

/*
TIP:
1. API 路徑的前墜在此設定一個全域的設定值
2. ApiCluster 內每個物件都視為一個 gin Group 各自在內設定自己的 group path
3. 在 ApiCluster 的 RouterRegister 內執行每個物件的 RegisterApiRouter 即可完成個別的 api 註冊

* api path format:
	BasePath + CheckApi.GroupApth + relativePath(個別 api 設定)
*/

var (
	BasePath    string = "/api/v1/" // inside service call out
	OutSidePath string = "/"        // outside service call out
	// IntercomPath string = "/intercom/v1/" // intercom service call out
)

type ApiCluster struct {
	cluster []api_cluster.IApiEach
}

func NewApiCluster(config config.Config) api_cluster.IApiCluster {
	apiCluster := &ApiCluster{}

	apiCluster.cluster = append(apiCluster.cluster,
		game.NewIntercomOrderApi(OutSidePath+"intercomorder"),
		game.NewGetGameRecordApi(OutSidePath+"record"),
	)

	return apiCluster
}

func (p *ApiCluster) RouterGroupRegister(router *gin.Engine, ginHandler *ginweb.GinHanlder) {
	// Don't modifiable code
	// middleware 在個別的 IApiEach 裡面加
	for idx := range p.cluster {
		tempGroup := router.Group(p.cluster[idx].GetGroupPath())
		{
			p.cluster[idx].RegisterApiRouter(tempGroup, ginHandler)
		}
	}
}
