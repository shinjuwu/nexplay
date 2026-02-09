package global

import (
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"fmt"
	"log"
	"time"
)

/* combination sql query to cal performance report
盈虧報表計算
*/
func CalPerformanceReportRecord(db *sql.DB, dataType string, startTime, endTime time.Time) (int, []table_model.RpRecordData) {
	tmps := make([]table_model.RpRecordData, 0)

	getTmpKey := func(logTime string, agentId int) string {
		return fmt.Sprintf("%s_%d", logTime, agentId)
	}

	code, dataFromUserPlayLog := getDataFromUserPlayLog(db, dataType, startTime, endTime, getTmpKey)
	if code != definition.ERROR_CODE_SUCCESS {
		return code, tmps
	}

	code, dataFromJackpotLog := getDataFromJackpotLog(db, dataType, startTime, endTime, getTmpKey)
	if code != definition.ERROR_CODE_SUCCESS {
		return code, tmps
	}

	for key, value := range dataFromUserPlayLog {
		jackpotLogData, ok := dataFromJackpotLog[key]
		if ok {
			value.JackpotUser = jackpotLogData.JackpotUser
			value.JackpotCount = jackpotLogData.JackpotCount
			value.SumJpPrizeScore = jackpotLogData.SumJpPrizeScore
		}

		tmps = append(tmps, *value)
	}

	for key, value := range dataFromJackpotLog {
		// 找到表示已經處理過可跳過
		if _, ok := dataFromUserPlayLog[key]; ok {
			continue
		}

		tmps = append(tmps, *value)
	}

	return 0, tmps
}

func getDataFromUserPlayLog(db *sql.DB, dataType string, startTime, endTime time.Time, getTmpKey func(string, int) string) (int, map[string]*table_model.RpRecordData) {
	tmps := make(map[string]*table_model.RpRecordData)

	// get user play log table name
	userPlayLogTableNames := make([]string, 0)
	gGameData := GameCache.GetAll()
	for _, v := range gGameData {
		if v.Type != definition.GAME_TYPE_LOBBY && v.CalState == 1 {
			userPlayLogTableNames = append(userPlayLogTableNames, "user_play_log_"+v.Code)
		}
	}

	dateQueryPart := ""

	switch dataType {
	case "15min":
		dateQueryPart = `date_trunc('hour'::text, tmp.bet_time)+ date_part('minute', tmp.bet_time)::int / 15 * interval '15 min'`
	case "hour":
		fallthrough
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		dateQueryPart = fmt.Sprintf(`date_trunc('%s'::text, tmp.bet_time)`, dataType)
	default:
		return definition.ERROR_CODE_ERROR_REQUEST_DATA, tmps
	}

	// combination sql query start
	preQueryPart :=
		`SELECT %s AS log_time,
			tmp.agent_id,
			count(DISTINCT tmp.user_id) AS bet_user,
			count(tmp.lognumber) AS bet_count,
			sum(tmp.ya_score) AS sum_ya,
			sum(tmp.valid_score) AS sum_valid_ya,
			sum(tmp.de_score) AS sum_de,
			sum(tmp.bonus) AS sum_bonus,
			sum(tmp.tax) AS sum_tax,
			sum(tmp.jp_inject_water_score) AS sum_jp_inject_water_score
		FROM ( %s ) tmp
		GROUP BY log_time, tmp.agent_id;`

	tableQueryPart := `SELECT lognumber,
						agent_id,
						user_id,
						game_id,
						ya_score,
						valid_score,
						de_score,
						bonus,
						tax,
						bet_time,
						jp_inject_water_score
					FROM %s
					WHERE is_robot = 0 AND bet_time >= $%d AND bet_time < $%d`
	unionQueryPart := " UNION ALL "
	fromQueryPart := ""
	dateTimeCount := 1
	arges := make([]any, 0)
	for k, v := range userPlayLogTableNames {

		fromQueryPart = fromQueryPart + fmt.Sprintf(tableQueryPart, v, dateTimeCount, dateTimeCount+1)
		dateTimeCount += 2
		arges = append(arges, startTime, endTime)

		if len(userPlayLogTableNames)-1 > k {
			fromQueryPart = fromQueryPart + unionQueryPart
		}

	}

	query := fmt.Sprintf(preQueryPart, dateQueryPart, fromQueryPart)
	// combination sql query end

	rows, err := db.Query(query, arges...)
	if err != nil {
		log.Printf("err is: %v", err)
		return definition.ERROR_CODE_ERROR_DATABASE, tmps
	}

	defer rows.Close()

	for rows.Next() {

		var tmp table_model.RpRecordData
		var logTime time.Time
		// newline every scan 5 column
		if err := rows.Scan(&logTime, &tmp.AgentId, &tmp.BetUser, &tmp.BetCount, &tmp.SumYa,
			&tmp.SumValidYa, &tmp.SumDe, &tmp.SumBonus, &tmp.SumTax, &tmp.SumJpInjectWaterScore); err != nil {
			log.Printf("CalPerformanceReportRecord has error: %v", err)
			return definition.ERROR_CODE_ERROR_DATABASE, tmps
		}

		// 轉換時間格式 time -> YYYYMMDDhh
		tmp.LogTime = utils.TransUnsignedTimeUTCFormat(dataType, logTime)

		agentTmp := AgentCache.Get(tmp.AgentId)

		if agentTmp == nil {
			continue
		}
		// must
		tmp.LevelCode = agentTmp.LevelCode

		tmps[getTmpKey(tmp.LogTime, tmp.AgentId)] = &tmp
	}

	return 0, tmps
}

func getDataFromJackpotLog(db *sql.DB, dataType string, startTime, endTime time.Time, getTmpKey func(string, int) string) (int, map[string]*table_model.RpRecordData) {
	tmps := make(map[string]*table_model.RpRecordData)

	dateQueryPart := ""

	switch dataType {
	case "15min":
		dateQueryPart = `date_trunc('hour'::text, winning_time)+ date_part('minute', winning_time)::int / 15 * interval '15 min'`
	case "hour":
		fallthrough
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		dateQueryPart = fmt.Sprintf(`date_trunc('%s'::text, winning_time)`, dataType)
	default:
		return definition.ERROR_CODE_ERROR_REQUEST_DATA, tmps
	}

	// combination sql query start
	preQueryPart :=
		`SELECT %s AS log_time,
			agent_id,
			count(DISTINCT user_id) AS jackpot_user,
			count(lognumber) AS jackpot_count,
			sum(prize_score) AS sum_jackpot_prize_score
		FROM jackpot_log
		WHERE is_robot = 0 AND winning_time >= $1 AND winning_time < $2
		GROUP BY log_time, agent_id;`

	query := fmt.Sprintf(preQueryPart, dateQueryPart)
	// combination sql query end

	rows, err := db.Query(query, startTime, endTime)
	if err != nil {
		log.Printf("err is: %v", err)
		return definition.ERROR_CODE_ERROR_DATABASE, tmps
	}

	defer rows.Close()

	for rows.Next() {

		var tmp table_model.RpRecordData
		var logTime time.Time
		// newline every scan 5 column
		if err := rows.Scan(&logTime, &tmp.AgentId, &tmp.JackpotUser, &tmp.JackpotCount, &tmp.SumJpPrizeScore); err != nil {
			log.Printf("CalPerformanceReportRecord has error: %v", err)
			return definition.ERROR_CODE_ERROR_DATABASE, tmps
		}

		// 轉換時間格式 time -> YYYYMMDDhh
		tmp.LogTime = utils.TransUnsignedTimeUTCFormat(dataType, logTime)

		agentTmp := AgentCache.Get(tmp.AgentId)

		if agentTmp == nil {
			continue
		}
		// must
		tmp.LevelCode = agentTmp.LevelCode

		tmps[getTmpKey(tmp.LogTime, tmp.AgentId)] = &tmp
	}

	return 0, tmps
}

/* combination sql query to cal performance report
刪除指定報表數據盈虧報表計算
*/
func DelPerformanceReportRecord(db *sql.DB, dataType string, startTime, endTime time.Time) error {

	args := utils.GetTimeIntervalList(dataType, startTime, endTime)

	if len(args) <= 0 {
		return fmt.Errorf("DelPerformanceReportRecord() sql param is empty, dataType: %v, startTime: %v, endTime: %v", dataType, startTime, endTime)
	}

	inPartQuery := ""
	for _, v := range args {
		inPartQuery = fmt.Sprintf(inPartQuery+"'%s',", v)
	}

	apRecordtableName := fmt.Sprintf("rp_agent_stat_%s", dataType)

	inPartQuery = inPartQuery[0 : len(inPartQuery)-1]

	// preQuery := `DELETE FROM %s WHERE log_time IN (%s);`

	query := fmt.Sprintf(`DELETE FROM %s WHERE log_time IN (%s);`, apRecordtableName, inPartQuery)

	rowResult, err := db.Exec(query)
	if err != nil {
		log.Printf("DelPerformanceReportRecord has error is: %v", err)
		return err
	}
	effectRows, _ := rowResult.RowsAffected()

	log.Printf("DelPerformanceReportRecord has RowsAffected is: %d", effectRows)

	return nil
}

/* combination sql query to cal performance report
插入指定報表數據盈虧報表計算
*/
func InsertPerformanceReportRecord(db *sql.DB, dataType string, records []table_model.RpRecordData) error {

	now := time.Now().UTC()
	log.Printf("InsertPerformanceReportRecord start time is %v", now)

	rpRecordTableName := fmt.Sprintf("rp_agent_stat_%s", dataType)

	sqlStr := fmt.Sprintf(`INSERT INTO
		"public"."%s" ("log_time", "agent_id", "level_code", "bet_user", "bet_count",
		"sum_ya", "sum_vaild_ya", "sum_de", "sum_bonus", "sum_tax",
		"jackpot_user", "jackpot_count", "sum_jp_inject_water_score", "sum_jp_prize_score") VALUES `, rpRecordTableName)

	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),` //14

	vals := []interface{}{}

	idx := 1
	for _, val := range records {
		sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4,
			idx+5, idx+6, idx+7, idx+8, idx+9,
			idx+10, idx+11, idx+12, idx+13)
		idx += 14

		vals = append(vals,
			val.LogTime, val.AgentId, val.LevelCode, val.BetUser, val.BetCount,
			val.SumYa, val.SumValidYa, val.SumDe, val.SumBonus, val.SumTax,
			val.JackpotUser, val.JackpotCount, val.SumJpInjectWaterScore, val.SumJpPrizeScore)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return err
		}
	}

	return nil
}
