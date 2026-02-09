package api_module

import (
	"encoding/json"
	"fmt"
	"monitorservice/cmd/baseserver/api_module/data"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/cmd/baseserver/api_module/module_lib"
	"monitorservice/internal/api"
	"monitorservice/pkg/utils"
	"sync"

	sq "github.com/Masterminds/squirrel"
)

// @Tags monitor.api
// @Summary 平台服務在線狀態
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ServiceStatusRequest true "指定查詢全部/單一服務狀態資料"
// @Success 200 {object} model.Response{data=model.ServiceStatusResponse,msg=string,code=int} ""
// @Router /monitor/monitor.api/v1/servicestatus [post]
func ServiceStatus(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	var req model.ServiceStatusRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	var res model.ServiceStatusResponse
	res.StatusList = make([]*model.ServiceStatus, 0)

	// 分執行緒執行
	var wg sync.WaitGroup

	tmps := data.ServiceInfo.GetAll()
	if req.Filter == filter_all {
		for _, v := range tmps {
			if v.IsEnabled {
				apiRes := new(model.ServiceStatus)
				apiRes.Info = v.Info
				apiRes.Name = v.Name
				apiRes.SubName = v.SubName
				apiRes.Status = api.ERROR_CODE_SUCCESS
				res.StatusList = append(res.StatusList, apiRes)

				wg.Add(1)
				go func(d *model.ServiceInfo) {
					defer wg.Done()
					apiRes.Status = module_lib.CallAPICheckServiceStatus(d)
				}(v)
			}
		}
	} else {
		if req.Filter != "" {
			for _, v := range tmps {
				if v.IsEnabled {
					if v.Name == req.Filter {
						apiRes := new(model.ServiceStatus)
						apiRes.Info = v.Info
						apiRes.Name = v.Name
						apiRes.SubName = v.SubName
						apiRes.Status = api.ERROR_CODE_SUCCESS
						res.StatusList = append(res.StatusList, apiRes)
						// apiRes.Status = module_lib.CallAPICheckServiceStatus(v)
						wg.Add(1)
						go func(d *model.ServiceInfo) {
							defer wg.Done()
							apiRes.Status = module_lib.CallAPICheckServiceStatus(d)
						}(v)
					}
				}
			}
		} else {
			return utils.ToJSON(res), fmt.Errorf("platform is not in list")
		}
	}
	wg.Wait()
	return utils.ToJSON(res), nil
}

// @Tags monitor.api
// @Summary 玩家上下分監控(顯示最新100筆資料)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.CoinInOutStatusRequest true "指定查詢單一服務狀態資料(默認 上/下分5000)"
// @Success 200 {object} model.Response{data=model.CoinInOutStatusResponse,msg=string,code=int} ""
// @Router /monitor/monitor.api/v1/coininoutstatus [post]
func CoinInOutStatus(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	var req model.CoinInOutStatusRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if !module_lib.CheckPermissionWithClaim(claims.Permissions, req.Filter) {
		return "", fmt.Errorf("insufficient permission")
	}

	// dbSetTmp, ok := data.DBSet.Load(req.Filter)
	// if !ok {
	// 	return "", fmt.Errorf("filter value is illege")
	// }

	// currectDB := dbSetTmp.(*sql.DB)

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", "agent_name", "user_id",
			"username", "kind", "status", "changeset", "create_time").
		From("wallet_ledger").
		Where(sq.Eq{"platform": req.Filter}).
		OrderBy("create_time DESC").
		Limit(coinInOutRowsLimit).
		ToSql()

	if err != nil {
		return "", err
	}

	var res model.CoinInOutStatusResponse

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		tmp := new(model.CoinInOutStatus)

		if err := rows.Scan(&tmp.Id, &tmp.AgentId, &tmp.AgentName, &tmp.UserId, &tmp.UserName,
			&tmp.Kind, &tmp.Status, &tmp.ChangeSet, &tmp.CreateTime); err != nil {
			return "", err
		}

		res.CoinInOutStatusList = append(res.CoinInOutStatusList, tmp)
	}

	return utils.ToJSON(res), nil
}

// @Tags monitor.api
// @Summary 異常輸贏監控(顯示最新100筆資料)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.AbnormalWinAndLoseStatusRequest true "指定查詢單一服務狀態資料"
// @Success 200 {object} model.Response{data=model.AbnormalWinAndLoseStatusResponse,msg=string,code=int} ""
// @Router /monitor/monitor.api/v1/abnormalwinandlosestatus [post]
func AbnormalWinAndLoseStatus(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	var req model.AbnormalWinAndLoseStatusRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if !module_lib.CheckPermissionWithClaim(claims.Permissions, req.Filter) {
		return "", fmt.Errorf("insufficient permission")
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("lognumber", "agent_id", "agent_name", "user_id", "username",
			"game_id", "game_name", "room_type", "room_name", "de_score",
			"bonus", "bet_id", "bet_time", "create_time").
		From("user_play_log").
		Where(sq.Eq{"platform": req.Filter}).
		OrderBy("create_time DESC").
		Limit(coinInOutRowsLimit).
		ToSql()

	if err != nil {
		return "", err
	}

	var res model.AbnormalWinAndLoseStatusResponse

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		tmp := new(model.AbnormalWinAndLoseStatus)

		if err := rows.Scan(&tmp.LogNumber, &tmp.AgentID, &tmp.AgentName, &tmp.UserID, &tmp.Username,
			&tmp.GameID, &tmp.GameName, &tmp.RoomType, &tmp.RoomName, &tmp.DeScore,
			&tmp.Bonus, &tmp.BetID, &tmp.BetTime, &tmp.CreateTime); err != nil {
			return "", err
		}

		res.AbnormalWinAndLoseStatusList = append(res.AbnormalWinAndLoseStatusList, tmp)
	}

	return utils.ToJSON(res), nil
}

// @Tags monitor.api
// @Summary 各平台RTP監控
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.PlatformRTPStatusRequest true "指定查詢單一服務狀態資料"
// @Success 200 {object} model.Response{data=model.PlatformRTPStatusResponse,msg=string,code=int} ""
// @Router /monitor/monitor.api/v1/platformrtpstatus [post]
func PlatformRTPStatus(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	var req model.PlatformRTPStatusRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if !module_lib.CheckPermissionWithClaim(claims.Permissions, req.Filter) {
		return "", fmt.Errorf("insufficient permission")
	}

	// 時區計算
	timZoneMin := req.TimeZone

	// table需要有資料才會產生，所以要確認table存在
	tableNameMonths := module_lib.GenerateTableNameMonth(timZoneMin)

	queryCreateTable := `CALL "public"."usp_check_game_ratio_stat"($1, $2)`
	for i := 0; i < len(tableNameMonths); i++ {
		// 沒table 就直接創建
		_, err := c.DB().Exec(queryCreateTable, tableNameMonths[i], req.Filter)
		if err != nil {
			return "", err
		}
	}

	platformTableNameMonths := []string{
		req.Filter + "_game_ratio_stat_" + tableNameMonths[0],
		req.Filter + "_game_ratio_stat_" + tableNameMonths[1],
	}

	// platformTableNameMonths := module_lib.GenerateTableNameMonth(req.Filter+"_"+module_lib.GameRatioStatTableName, timZoneMin)
	selectPeriodOfTimeMonths := module_lib.GenerateMonthPeriodOfTime("15min", timZoneMin)
	selectPeriodOfTimeWeeks := module_lib.GenerateWeekPeriodOfTime("15min", timZoneMin)
	selectPeriodOfTimeDays := module_lib.GenerateDayPeriodOfTime("15min", timZoneMin)

	combitionQuery := ""

	selectQuery := `SELECT * 
			FROM %s
			WHERE log_time >= '%s' AND log_time < '%s'`
	for i := 0; i < len(platformTableNameMonths); i++ {
		if i == 0 {
			combitionQuery += fmt.Sprintf(selectQuery, platformTableNameMonths[i], selectPeriodOfTimeMonths[0], selectPeriodOfTimeMonths[1])
		} else {
			combitionQuery += " UNION "
			combitionQuery += fmt.Sprintf(selectQuery, platformTableNameMonths[i], selectPeriodOfTimeMonths[0], selectPeriodOfTimeMonths[1])
		}
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("log_time", "game_id", "game_type", `sum(de) as "de"`, `sum(ya) as "ya"`,
			`sum(tax) as "tax"`, `sum(bonus) as "bonus"`, `sum(play_count) as "play_count"`).
		From("(" + combitionQuery + ") as combined_data").
		GroupBy("log_time, game_id, game_type").
		OrderBy("log_time, game_id ASC").
		ToSql()

	if err != nil {
		return "", err
	}

	var gameStatData model.PlatformGameRatioStat

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		tmp := new(model.GameRatioStat)

		if err := rows.Scan(&tmp.LogTime, &tmp.GameID, &tmp.GameType, &tmp.DE, &tmp.YA,
			&tmp.Tax, &tmp.Bonus, &tmp.PlayCount); err != nil {
			return "", err
		}

		gameStatData.GameRatioStatList = append(gameStatData.GameRatioStatList, tmp)
	}

	var res model.PlatformRTPStatusResponse

	todayRTPStat := model.NewRTPStatus("day", "日", -1, selectPeriodOfTimeDays[0], selectPeriodOfTimeDays[1])
	weekRTPStat := model.NewRTPStatus("week", "週", -1, selectPeriodOfTimeWeeks[0], selectPeriodOfTimeWeeks[1])
	monthRTPStat := model.NewRTPStatus("month", "月", -1, selectPeriodOfTimeMonths[0], selectPeriodOfTimeMonths[1])

	gameRTPDayStat := make(map[int]*model.RTPStatus, 0)
	gameRTPWeekStat := make(map[int]*model.RTPStatus, 0)
	gameRTPMonthStat := make(map[int]*model.RTPStatus, 0)

	game := data.Game.Get(req.Filter)
	for _, v := range game.GameList {
		gameRTPDayStat[v.ID] = model.NewRTPStatus(v.Code, v.Name, v.Type, selectPeriodOfTimeDays[0], selectPeriodOfTimeDays[1])
		gameRTPWeekStat[v.ID] = model.NewRTPStatus(v.Code, v.Name, v.Type, selectPeriodOfTimeWeeks[0], selectPeriodOfTimeWeeks[1])
		gameRTPMonthStat[v.ID] = model.NewRTPStatus(v.Code, v.Name, v.Type, selectPeriodOfTimeMonths[0], selectPeriodOfTimeMonths[1])
	}

	for _, gameStat := range gameStatData.GameRatioStatList {
		// 同一天的所有資料
		if gameStat.LogTime >= todayRTPStat.CalST && gameStat.LogTime < todayRTPStat.CalET {
			todayRTPStat.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)

			if val, ok := gameRTPDayStat[gameStat.GameID]; ok {
				val.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)
			}
		}
		// 本周的資料
		if gameStat.LogTime >= weekRTPStat.CalST && gameStat.LogTime < weekRTPStat.CalET {
			weekRTPStat.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)

			if val, ok := gameRTPWeekStat[gameStat.GameID]; ok {
				val.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)
			}
		}
		// 整個月的資料
		monthRTPStat.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)

		if val, ok := gameRTPMonthStat[gameStat.GameID]; ok {
			val.CalProcess(gameStat.DE, gameStat.YA, gameStat.Tax, gameStat.Bonus, gameStat.PlayCount)
		}
	}

	res.RTPStatusList = append(res.RTPStatusList, todayRTPStat)
	res.RTPStatusList = append(res.RTPStatusList, weekRTPStat)
	res.RTPStatusList = append(res.RTPStatusList, monthRTPStat)

	for _, v := range gameRTPDayStat {
		res.RTPStatusDayList = append(res.RTPStatusDayList, v)
	}

	for _, v := range gameRTPWeekStat {
		res.RTPStatusWeekList = append(res.RTPStatusWeekList, v)
	}

	for _, v := range gameRTPMonthStat {
		res.RTPStatusMonthList = append(res.RTPStatusMonthList, v)
	}

	return utils.ToJSON(res), nil
}
