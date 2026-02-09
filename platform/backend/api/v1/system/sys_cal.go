package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/utils"
	"backend/server/global"
	"definition"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

/*
計算報表使用API
*/

type SystemCalApi struct {
	BasePath string
}

func NewSystemCalApi(basePath string) api_cluster.IApiEach {
	return &SystemCalApi{
		BasePath: basePath,
	}
}

func (p *SystemCalApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemCalApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.POST("/calperformancereport", ginHandler.Handle(p.CalPerformanceReport))
	g.POST("/getperformancereportlist", ginHandler.Handle(p.GetPerformanceReportList))
	g.POST("/getperformancereport", ginHandler.Handle(p.GetPerformanceReport))
}

// @Tags 報表管理/業績報表
// @Summary Cal performance report
// @Description 此接口用來重新計算業績報表
// @Produce  application/json
// @Security BearerAuth
// @Param startDate formData string true "指定計算開始時間，格式: 2006-01-02T15:04:05Z"
// @Param endDate formData string true "指定計算結束時間，格式: 2006-01-02T15:04:05Z"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/cal/calperformancereport [post]
func (p *SystemCalApi) CalPerformanceReport(c *ginweb.Context) {
	// dateType := c.PostForm("dateType")
	startDate := c.PostForm("startDate")
	endDate := c.PostForm("endDate")
	dateType := "15min"
	// if dateType == "" {
	// 	c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
	// 	return
	// }

	startTime, err := time.ParseInLocation("2006-01-02T15:04:05Z", startDate, time.UTC)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}
	endTime, err := time.ParseInLocation("2006-01-02T15:04:05Z", endDate, time.UTC)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 計算盈虧報表
	errCode, tmps := global.CalPerformanceReportRecord(c.DB(), dateType, startTime, endTime)
	if errCode != 0 {
		c.OkWithCode(errCode)
		return
	}

	dataCount := len(tmps)

	//如果有資料
	if dataCount > 0 {
		// TODO: 刪除原數據
		err = global.DelPerformanceReportRecord(c.DB(), dateType, startTime, endTime)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
		// TODO:  insert into 新數據
		err = global.InsertPerformanceReportRecord(c.DB(), dateType, tmps)
		if err != nil {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
	}

	c.Ok(dataCount)
}

// @Tags 報表管理/業績報表
// @Summary 指定時間區段下層代理總和的資料
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.GetPerformanceReportListRequest true "指定時間區段下層代理總和的資料"
// @Success 200 {object} response.Response{data=[]model.GetPerformanceReportListResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/cal/getperformancereportlist [post]
func (p *SystemCalApi) GetPerformanceReportList(c *ginweb.Context) {
	req := model.NewGetPerformanceReportListRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(
		serverInfoSetting.EarningReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	sqAnd := sq.And{
		sq.GtOrEq{"log_time": utils.TransUnsignedTimeUTCFormat("minute", req.StartTime)},
		sq.Lt{"log_time": utils.TransUnsignedTimeUTCFormat("minute", req.EndTime)},
	}

	isAdmin := len(claims.LevelCode) == 4
	myAgentId := int(claims.BaseClaims.ID)
	searchOneAgent := (req.AgentId > definition.AGENT_ID_ALL)
	searchLevelCode := claims.LevelCode
	if isAdmin {
		if searchOneAgent {
			reqAgent := global.AgentCache.Get(req.AgentId)
			if len(reqAgent.LevelCode) == 8 {
				searchLevelCode = reqAgent.LevelCode
			}
		}
		sqAnd = append(sqAnd, sq.Like{"level_code": searchLevelCode + "%"})
	} else {
		// 只能查自己跟下一層不能查超過兩層
		reqAgent := global.AgentCache.Get(req.AgentId)
		if searchOneAgent {
			if reqAgent == nil || (reqAgent.Id != myAgentId && reqAgent.TopAgentId != myAgentId) {
				c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
				return
			}

			sqAnd = append(sqAnd, sq.Eq{"agent_id": req.AgentId})
		} else {
			sqAnd = append(sqAnd, sq.LtOrEq{"LENGTH(level_code)": (claims.AccountType + 1) * 4})
			sqAnd = append(sqAnd, sq.Like{"level_code": searchLevelCode + "%"})
		}
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "level_code", "SUM(bet_user)", "SUM(bet_count)", "SUM(sum_ya)",
			"SUM(sum_vaild_ya)", "SUM(sum_de)", "SUM(sum_bonus)", "SUM(sum_tax)", "SUM(sum_jp_inject_water_score)",
			"SUM(sum_jp_prize_score)", "SUM(jackpot_user)", "SUM(jackpot_count)").
		From("rp_agent_stat_15min").
		GroupBy("agent_id", "level_code").
		Where(sqAnd).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	defer rows.Close()

	resp := model.NewGetPerformanceReportListResponseSlice()

	for rows.Next() {
		tmp := model.NewGetPerformanceReportListResponse()
		// newline every scan 5 column
		if err := rows.Scan(&tmp.AgentId, &tmp.LevelCode, &tmp.BetUser, &tmp.BetCount, &tmp.SumYa,
			&tmp.SumValidYa, &tmp.SumDe, &tmp.SumBonus, &tmp.SumTax, &tmp.SumJpInjectWaterScore,
			&tmp.SumJpPrizeScore, &tmp.JackpotUser, &tmp.JackpotCount); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		agentObj := global.AgentCache.Get(tmp.AgentId)
		if agentObj != nil {
			exchangeObj := global.ExchangeDataCache.Get(agentObj.Currency)
			if exchangeObj != nil {
				tmp.AgentName = agentObj.Name
				tmp.AgentCommission = agentObj.Commission
				tmp.Currency = agentObj.Currency
				tmp.ToCoin = exchangeObj.ToCoin
			}
		}

		resp = append(resp, tmp)
	}

	if isAdmin {
		respMap := make(map[string]*model.GetPerformanceReportListResponse)

		for _, res := range resp {
			tmp, ok := respMap[res.LevelCode[:8]]
			if ok {
				tmp.BetUser += res.BetUser
				tmp.BetCount += res.BetCount
				tmp.JackpotUser += res.JackpotUser
				tmp.JackpotCount += res.JackpotCount
				tmp.SumYa = utils.DecimalAdd(tmp.SumYa, res.SumYa)
				tmp.SumValidYa = utils.DecimalAdd(tmp.SumValidYa, res.SumValidYa)
				tmp.SumDe = utils.DecimalAdd(tmp.SumDe, res.SumDe)
				tmp.SumBonus = utils.DecimalAdd(tmp.SumBonus, res.SumBonus)
				tmp.SumTax = utils.DecimalAdd(tmp.SumTax, res.SumTax)
				tmp.Currency = res.Currency
				tmp.ToCoin = res.ToCoin
				tmp.SumJpInjectWaterScore = utils.DecimalAdd(tmp.SumJpInjectWaterScore, res.SumJpInjectWaterScore)
				tmp.SumJpPrizeScore = utils.DecimalAdd(tmp.SumJpPrizeScore, res.SumJpPrizeScore)
			} else {
				// 轉換顯示的level code(下層通通要加總到上層)
				res.LevelCode = res.LevelCode[:8]
				respMap[res.LevelCode[:8]] = res
			}
		}

		agentObjs := global.AgentCache.GetAll()

		for levelCode, respTmp := range respMap {
			for _, agentObj := range agentObjs {
				if levelCode == agentObj.LevelCode {
					respTmp.AgentCommission = agentObj.Commission
					respTmp.AgentName = agentObj.Name
					respTmp.AgentId = agentObj.Id
					break
				}
			}
		}

		resp = model.NewGetPerformanceReportListResponseSlice()
		for _, respTmp := range respMap {
			if len(respTmp.LevelCode) == 8 {
				resp = append(resp, respTmp)
			}
		}
	}

	c.Ok(resp)
}

// @Tags 報表管理/業績報表
// @Summary 指定時間區段下層代理時間區段的統計資料
// @Produce  multipart/form-data
// @Security BearerAuth
// @Param data body model.GetPerformanceReportRequest true "指定時間區段下層代理時間區段的統計資料"
// @Success 200 {object} response.Response{data=[]model.GetPerformanceReportResponse,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/cal/getperformancereport [post]
func (p *SystemCalApi) GetPerformanceReport(c *ginweb.Context) {
	req := model.NewGetPerformanceReportRequest()

	if err := c.ShouldBind(req); err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	serverInfoSetting := global.GetServerInfo()
	if code := req.CheckParams(serverInfoSetting.CommonReportTimeMinuteIncrement,
		serverInfoSetting.CommonReportTimeRange,
		serverInfoSetting.CommonReportTimeBeforeDays); code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	myAgentId := int(claims.BaseClaims.ID)
	reqAgent := global.AgentCache.Get(req.AgentId)
	// 只能查自己跟下一層不能查超過兩層
	if reqAgent == nil || (reqAgent.Id != myAgentId && reqAgent.TopAgentId != myAgentId) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("log_time", "agent_id", "bet_user", "bet_count", "sum_ya",
			"sum_vaild_ya", "sum_de", "sum_bonus", "sum_tax", "sum_jp_inject_water_score",
			"sum_jp_prize_score", "jackpot_user", "jackpot_count").
		From("rp_agent_stat_15min").
		Where(sq.And{
			sq.GtOrEq{"log_time": utils.TransUnsignedTimeUTCFormat("minute", req.StartTime)},
			sq.Lt{"log_time": utils.TransUnsignedTimeUTCFormat("minute", req.EndTime)},
			sq.Eq{"agent_id": req.AgentId},
		}).
		ToSql()

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	defer rows.Close()

	resp := model.NewGetPerformanceReportResponseSlice()

	for rows.Next() {
		tmp := model.NewGetPerformanceReportResponse()
		// newline every scan 5 column
		if err := rows.Scan(&tmp.LogTime, &tmp.AgentId, &tmp.BetUser, &tmp.BetCount, &tmp.SumYa,
			&tmp.SumValidYa, &tmp.SumDe, &tmp.SumBonus, &tmp.SumTax, &tmp.SumJpInjectWaterScore,
			&tmp.SumJpPrizeScore, &tmp.JackpotUser, &tmp.JackpotCount); err != nil {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}

		agentObj := global.AgentCache.Get(tmp.AgentId)
		if agentObj != nil {
			tmp.AgentName = agentObj.Name
		}

		logTime, err := utils.GetUnsignedTimeUTCFromStr(tmp.LogTime, "minute")
		if err != nil {
			c.Logger().Printf("GetPerformanceReport(), trans logtime format has error: %v", err)
		} else {
			tmp.LogTime = logTime.String()
			resp = append(resp, tmp)
		}
	}

	c.Ok(resp)
}
