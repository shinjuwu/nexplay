package api

import (
	"backend/api/game"
	"backend/api/health"
	"backend/api/intercom"
	"backend/api/v1/system"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
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

	// TIP: add customize api group after here

	checkHeader := config.GetApp().CheckHeader

	accManageHeader := config.GetApp().AccManageHeader
	accNormalHeader := config.GetApp().AccNormalHeader

	if config.GetApp().LoadBackendChannel {
		apiCluster.cluster = append(apiCluster.cluster,
			// // for example use
			health.NewHealthApi(BasePath+"health"),
			// for backend use
			system.NewSystemLoginApi(BasePath+"login", checkHeader, accManageHeader, accNormalHeader),
			system.NewSystemManageApi(BasePath+"manage"), // 運營管理
			system.NewSystemUserApi(BasePath+"user"),
			system.NewSystemGameApi(BasePath+"game"),
			system.NewSystemRecordApi(BasePath+"record"),
			system.NewSystemAgentApi(BasePath+"agent"),
			system.NewSystemGlobalApi(BasePath+"global"),
			system.NewSystemCalApi(BasePath+"cal"),
			system.NewSystemNotifyApi(BasePath+"notify"),
			system.NewSystemServiceStatusApi(BasePath+"servicestatus"),
			system.NewSystemRiskControlApi(BasePath+"riskcontrol"),
			system.NewSystemSystemApi(BasePath+"system"),
			system.NewSystemJackpotApi(BasePath+"jackpot"),
			// for intercom use
			intercom.NewIntercomApi(BasePath+"intercom"),
		)
	}

	if config.GetApp().LoadApiChannel {
		apiCluster.cluster = append(apiCluster.cluster,
			// for game use
			game.NewGameChannelApi(OutSidePath+"channel"),
		)
	}

	if config.GetApp().LoadCasinoChannel {
		apiCluster.cluster = append(apiCluster.cluster,
			// for game use
			game.NewCasinoChannelApi(OutSidePath+"casino"),
		)
	}

	// if config.GetApp().LoadRecordChannel {
	// 	apiCluster.cluster = append(apiCluster.cluster,
	// 		game.NewGetGameRecordApi(OutSidePath+"record"),
	// 	)
	// }

	return apiCluster
}

func (p *ApiCluster) RouterGroupRegister(router *gin.Engine, ginHandler *ginweb.GinHanlder) {
	// Don't modifiable code
	// middleware 在個別的 IApiEach 裡面加
	for idx := range p.cluster {
		tempGroup := router.Group(p.cluster[idx].GetGroupPath())
		{
			if p.cluster[idx].GetGroupPath() != BasePath+"intercom" {
				tempGroup.Use(middleware.RequestLimit(middleware.MAX_ALLOWED))
			}
			p.cluster[idx].RegisterApiRouter(tempGroup, ginHandler)
		}
	}
}
