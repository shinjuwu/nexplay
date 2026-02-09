package api_module

import (
	"monitorservice/internal/api"
	"monitorservice/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const (
	filter_all = "all"

	proxyPath = "/monitor"

	// login page
	loginPreApiPath = proxyPath + "/login.api/v1/"
	// account page
	accountPreApiPath = proxyPath + "/account.api/v1/"
	// admin page
	adminPreApiPath = proxyPath + "/admin.api/v1/"
	// monitor page
	monitorPreApiPath = proxyPath + "/monitor.api/v1/"
	// collector page
	collectorPreApiPath = proxyPath + "/collector.api/v1/"

	// 玩家上下分監控設定查詢上/下分指定數值
	// coinInOutMax = 5000
	// 玩家上下分監控資料筆數
	coinInOutRowsLimit = 100
)

type ApiHandler struct {
	routerGroup *gin.RouterGroup
	jwt         *jwt.JwtManager
	// basePath    string
}

func NewApiHandler(routerGroup *gin.RouterGroup, jwt *jwt.JwtManager) api.IApiHandler {
	return &ApiHandler{
		routerGroup: routerGroup,
		jwt:         jwt,
		// basePath:    basePath,
	}
}

func (p *ApiHandler) RouterGroup() *gin.RouterGroup {
	return p.routerGroup
}

// func (p *ApiHandler) BasePsth() string {
// 	return p.basePath
// }

func (p *ApiHandler) ApiHandleRegister(fn api.ApiHandleFunc) {
	// 取得驗證碼
	p.routerGroup.POST(loginPreApiPath+"/captcha", fn(UsersCaptcha))
	// 登入
	p.routerGroup.POST(loginPreApiPath+"/login", fn(UsersLogin))

	//***************************************************

	// 取自己的帳戶資料
	p.routerGroup.POST(accountPreApiPath+"/usersinfo", api.JWT(p.jwt), fn(UsersInfo))
	// 驗證目前登入 token 是否有效
	p.routerGroup.POST(accountPreApiPath+"/verifytoken", api.JWT(p.jwt), fn(VerifyToken))
	// 修改自己的帳戶資料
	p.routerGroup.POST(accountPreApiPath+"/modifyinfo", api.JWT(p.jwt), fn(ModifyInfo))
	// 修改自己的密碼
	p.routerGroup.POST(accountPreApiPath+"/modifypassword", api.JWT(p.jwt), fn(ModifyPassword))

	//***************************************************

	// 註冊(不放外部注冊，此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/register", api.JWT(p.jwt), fn(Register))
	// 封鎖指定用戶(此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/blockusers", api.JWT(p.jwt), fn(BlockUsers))
	// 修改帳戶資料(此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/modifyusersinfo", api.JWT(p.jwt), fn(ModifyUsersInfo))
	// 修改密碼(此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/modifyuserspassword", api.JWT(p.jwt), fn(ModifyUsersPassword))
	// 取所有用戶的帳戶資料(此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/getusersinfolist", api.JWT(p.jwt), fn(GetUsersInfoList))
	// 取指定用戶的帳戶資料(此API只有admin可用)
	p.routerGroup.POST(adminPreApiPath+"/getusersinfo", api.JWT(p.jwt), fn(GetUsersInfo))

	//***************************************************

	// 平台服務在線狀態
	p.routerGroup.POST(monitorPreApiPath+"/servicestatus", api.JWT(p.jwt), fn(ServiceStatus))
	// 取玩家上下分監控狀態
	p.routerGroup.POST(monitorPreApiPath+"/coininoutstatus", api.JWT(p.jwt), fn(CoinInOutStatus))
	// 取異常輸贏監控狀態
	p.routerGroup.POST(monitorPreApiPath+"/abnormalwinandlosestatus", api.JWT(p.jwt), fn(AbnormalWinAndLoseStatus))
	// 取平台RTP監控狀態
	p.routerGroup.POST(monitorPreApiPath+"/platformrtpstatus", api.JWT(p.jwt), fn(PlatformRTPStatus))

	//***************************************************

	// 同步平台遊戲列表
	p.routerGroup.POST(collectorPreApiPath+"/collectorgamelist", fn(CollectorGameList))
	// 玩家上下分資料收集
	p.routerGroup.POST(collectorPreApiPath+"/collectorcoininout", fn(CollectorCoinInOut))
	// 異常輸贏資料收集
	p.routerGroup.POST(collectorPreApiPath+"/collectorabnormalwinandlose", fn(CollectorAbnormalWinAndLose))
	// RTP 統計資料收集
	p.routerGroup.POST(collectorPreApiPath+"/collectorrtpstat", fn(CollectorRTPStat))

}
