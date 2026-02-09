package system

import (
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/ginweb/response"
	"backend/server/global"
	"definition"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemSystemApi struct {
	BasePath string
}

func NewSystemSystemApi(basePath string) api_cluster.IApiEach {
	return &SystemSystemApi{
		BasePath: basePath,
	}
}

func (p *SystemSystemApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemSystemApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.POST("/getserversetting", ginHandler.Handle(p.GetServerSetting))
	g.POST("/getexchangedatalist",
		middleware.JWT(ginHandler.Jwt, global.AgentCache),
		ginHandler.Handle(p.GetExchangeDataList))
	g.POST("/setexchangedatalist",
		middleware.JWT(ginHandler.Jwt, global.AgentCache),
		middleware.BackendLogger(db, logger, &global.PermissionList),
		ginHandler.Handle(p.SetExchangeDataList))
}

// @Tags 系統/取得伺服器設定
// @Summary 取得伺服器設定
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=model.ServerInfoSetting,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/system/getserversetting [post]
func (p *SystemSystemApi) GetServerSetting(c *ginweb.Context) {
	serverInfoSetting := global.GetServerInfo()
	if serverInfoSetting == nil {
		response.StatusInternalServerError(c.Context)
		return
	}
	c.Ok(serverInfoSetting)
}

// @Tags 系統/匯率設定
// @Summary 取得匯率設定
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]global.ExchangeData,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/system/getexchangedatalist [post]
func (p *SystemSystemApi) GetExchangeDataList(c *ginweb.Context) {
	// 只有管理後台可以用
	userClaims := c.GetUserClaims()
	if userClaims == nil || len(userClaims.LevelCode) > 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	tmps := global.ExchangeDataCache.GetAll()
	if len(tmps) == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}
	c.Ok(tmps)
}

// @Tags 系統/匯率設定
// @Summary 設定匯率設定
// @Produce  application/json
// @Security BearerAuth
// @Param data body []global.ExchangeData true "匯率設定資料"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/system/setexchangedatalist [post]
func (p *SystemSystemApi) SetExchangeDataList(c *ginweb.Context) {
	var req []*global.ExchangeData
	if err := c.ShouldBindJSON(&req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 只有管理後台可以用
	userClaims := c.GetUserClaims()
	if userClaims == nil || len(userClaims.LevelCode) > 4 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	tmps := global.ExchangeDataCache.GetAll()
	if len(tmps) == 0 {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	generateKey := func(exchangeData *global.ExchangeData) string {
		return fmt.Sprintf("%d_%s", exchangeData.Id, exchangeData.Currency)
	}

	backupExchangeData := make(map[string]*global.ExchangeData, len(tmps))
	for _, tmp := range tmps {
		key := generateKey(tmp)
		backupExchangeData[key] = tmp
	}

	tx, err := c.DB().Begin()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	defer tx.Rollback()

	actionLogExchangeData := make([]map[string]interface{}, 0)
	for _, tmp := range req {
		key := generateKey(tmp)
		if _, find := backupExchangeData[key]; !find {
			c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
			return
		}

		backup := backupExchangeData[key]

		if backup.ToCoin == tmp.ToCoin {
			continue
		}

		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Update("exchange_data").
			Set("to_coin", tmp.ToCoin).
			Where(sq.And{
				sq.Eq{"id": backup.Id},
				sq.Eq{"currency": backup.Currency},
			}).
			ToSql()

		_, err = tx.Exec(query, args...)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}

		actionLogExchangeData = append(actionLogExchangeData, map[string]interface{}{
			"id":       backup.Id,
			"currency": backup.Currency,
			"to_coin":  createBancendActionLogDetail(backup.ToCoin, tmp.ToCoin),
		})
	}

	err = tx.Commit()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	global.ExchangeDataCache.Adds(req)

	// 操作紀錄
	actionLog := make(map[string]interface{})
	if len(actionLogExchangeData) > 0 {
		actionLog["exchange_data"] = actionLogExchangeData
	}
	c.Set("action_log", actionLog)

	c.Ok(nil)
}
