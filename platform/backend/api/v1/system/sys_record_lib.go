package system

import (
	"backend/api/v1/model"
	"backend/server/global"
	table_model "backend/server/table/model"
	"database/sql"
	"definition"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func generateUserPlayLogListTableName(gameId int, startTime, endTime *time.Time) string {
	sqlStr := `
	  SELECT %[1]s.bet_id,
	    %[1]s.lognumber,
	    %[1]s.agent_id,
	    %[1]s.game_id,
	    %[1]s.room_type,
	    %[1]s.desk_id,
	    %[1]s.seat_id,
		%[1]s.exchange,
		%[1]s.de_score,
		%[1]s.ya_score,
		%[1]s.valid_score,
		%[1]s.start_score,
		%[1]s.end_score,
		%[1]s.create_time,
		%[1]s.is_robot,
		%[1]s.is_big_win,
		%[1]s.is_issue,
		%[1]s.bet_time,
		%[1]s.tax,
		%[1]s.level_code,
		%[1]s.username,
		%[1]s.kill_type,
		%[1]s.bonus,
		%[1]s.jp_inject_water_rate,
		%[1]s.jp_inject_water_score,
		%[1]s.wallet_ledger_id,
		%[1]s.kill_prob,
		%[1]s.kill_level,
		%[1]s.real_players,
		%[1]s.room_id
	  FROM user_play_log_%[1]s %[1]s
	  WHERE %[1]s.is_robot = 0 AND %[1]s.bet_time >= '%[2]s' AND %[1]s.bet_time < '%[3]s'
	`

	startTimeStr := startTime.Format(time.RFC3339)
	endTimeStr := endTime.Format(time.RFC3339)

	game := global.GameCache.Get(gameId)
	if game != nil &&
		game.Type != definition.GAME_TYPE_LOBBY {
		return "(" + fmt.Sprintf(sqlStr, game.Code, startTimeStr, endTimeStr) + ") tmp"
	}

	tableName := ""

	for _, game := range global.GameCache.GetAll() {
		if game.Type == definition.GAME_TYPE_LOBBY {
			continue
		}

		if tableName != "" {
			tableName += "UNION ALL "
		}
		tableName += fmt.Sprintf(sqlStr, game.Code, startTimeStr, endTimeStr)
	}

	return "(" + tableName + ") tmp"
}

func getBatchUserPlayLogListSumInfo(db *sql.DB, tableName string, sqAnd *sq.And) (recordsTotal, totalKillGames, totalDiveGames, totalPlayerWinGames int, totalValidScore, totalPlatformWinloseScore, totalJpInjectWaterScore float64, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"COUNT(tmp.bet_id)",
			"tmp.kill_type",
			"SUM(tmp.valid_score)",
			"SUM(tmp.ya_score) + SUM(tmp.tax) - SUM(tmp.de_score) - SUM(tmp.bonus)",
			"SUM(CASE WHEN tmp.de_score + tmp.bonus - tmp.ya_score - tmp.tax > 0 THEN 1 ELSE 0 END)",
			"SUM(tmp.jp_inject_water_score)",
		).
		From(tableName).
		Where(sqAnd).
		GroupBy("kill_type").
		ToSql()
	if err != nil {
		return
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		var count, killType, playerWinGames int
		var validScore, platformWinloseScore, jpInjectWaterScore float64

		if err = rows.Scan(&count, &killType, &validScore, &platformWinloseScore, &playerWinGames, &jpInjectWaterScore); err != nil {
			return
		}

		recordsTotal += count
		totalValidScore += validScore
		totalPlatformWinloseScore += platformWinloseScore
		totalPlayerWinGames += playerWinGames
		totalJpInjectWaterScore += jpInjectWaterScore

		// 殺放狀態統計
		if killType > definition.KILLDIVE_KILL_TYPE_NORMAL {
			if killType == definition.KILLDIVE_KILL_TYPE_RELEASE {
				totalDiveGames += count
			} else {
				totalKillGames += count
			}
		}
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func getBatchUserPlayLogList(db *sql.DB, req *model.GetBatchUserPlayLogListRequest, tableName string, sqAnd *sq.And) (resp []*model.GetUserPlayLogResponse, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	sqBuilder := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("tmp.agent_id", "tmp.game_id", "tmp.room_type", "tmp.desk_id", "tmp.seat_id",
			"tmp.valid_score", "tmp.start_score", "tmp.end_score", "tmp.de_score", "tmp.ya_score",
			"tmp.bet_time", "tmp.username", "tmp.lognumber", "tmp.bet_id", "tmp.kill_type",
			"tmp.tax", "tmp.bonus", "tmp.jp_inject_water_rate", "tmp.jp_inject_water_score", "tmp.wallet_ledger_id",
			"tmp.kill_prob", "tmp.kill_level", "tmp.real_players", "tmp.room_id").
		From(tableName).
		Where(sqAnd)

	sqBuilder = getBatchUserPlayLogListEvaluateOrderBy(sqBuilder, req.SortColumn, req.SortDirection)

	query, args, _ := sqBuilder.
		Limit(uint64(req.Length)).Offset(uint64(req.Start)).
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}

	defer rows.Close()

	tmpAgents := make(map[int]*table_model.Agent)

	for rows.Next() {
		tmp := &model.GetUserPlayLogResponse{}
		if err = rows.Scan(&tmp.AgentId, &tmp.GameId, &tmp.RoomType, &tmp.DeskId, &tmp.SeatId,
			&tmp.ValidScore, &tmp.StartScore, &tmp.EndScore, &tmp.DeScore, &tmp.YaScore,
			&tmp.BetTime, &tmp.UserName, &tmp.LogNumber, &tmp.BetId, &tmp.KillType,
			&tmp.Tax, &tmp.Bonus, &tmp.JpInjectWaterRate, &tmp.JpInjectWaterScore, &tmp.WalletLedgerId,
			&tmp.KillProb, &tmp.KillLevel, &tmp.RealPlayers, &tmp.RoomId); err != nil {
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, tmp)
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func getBatchUserPlayLogListEvaluateOrderBy(sqBuilder sq.SelectBuilder, sortColumn string, sortDirection int) sq.SelectBuilder {
	if sortColumn == "betid" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.bet_id asc")
		} else {
			return sqBuilder.OrderBy("tmp.bet_id desc")
		}
	} else if sortColumn == "agentname" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.agent_id asc")
		} else {
			return sqBuilder.OrderBy("tmp.agent_id desc")
		}
	} else if sortColumn == "username" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.username asc")
		} else {
			return sqBuilder.OrderBy("tmp.username desc")
		}
	} else if sortColumn == "gameid" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.game_id asc")
		} else {
			return sqBuilder.OrderBy("tmp.game_id desc")
		}
	} else if sortColumn == "roomtype" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.room_type asc")
		} else {
			return sqBuilder.OrderBy("tmp.room_type desc")
		}
	} else if sortColumn == "yascore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.ya_score asc")
		} else {
			return sqBuilder.OrderBy("tmp.ya_score desc")
		}
	} else if sortColumn == "validscore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.valid_score asc")
		} else {
			return sqBuilder.OrderBy("tmp.valid_score desc")
		}
	} else if sortColumn == "descore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.de_score asc")
		} else {
			return sqBuilder.OrderBy("tmp.de_score desc")
		}
	} else if sortColumn == "tax" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.tax asc")
		} else {
			return sqBuilder.OrderBy("tmp.tax desc")
		}
	} else if sortColumn == "lognumber" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.lognumber asc")
		} else {
			return sqBuilder.OrderBy("tmp.lognumber desc")
		}
	} else if sortColumn == "killtype" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.kill_type asc")
		} else {
			return sqBuilder.OrderBy("tmp.kill_type desc")
		}
	} else if sortColumn == "winlosescore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("(tmp.de_score + tmp.bonus - tmp.ya_score - tmp.tax) asc")
		} else {
			return sqBuilder.OrderBy("(tmp.de_score + tmp.bonus - tmp.ya_score - tmp.tax) desc")
		}
	} else if sortColumn == "bonus" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.bonus asc")
		} else {
			return sqBuilder.OrderBy("tmp.bonus desc")
		}
	} else if sortColumn == "jpinjectwaterrate" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.jp_inject_water_rate asc")
		} else {
			return sqBuilder.OrderBy("tmp.jp_inject_water_rate desc")
		}
	} else if sortColumn == "jpinjectwaterscore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.jp_inject_water_score asc")
		} else {
			return sqBuilder.OrderBy("tmp.jp_inject_water_score desc")
		}
	} else if sortColumn == "wallet_ledger_id" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.wallet_ledger_id asc")
		} else {
			return sqBuilder.OrderBy("tmp.wallet_ledger_id desc")
		}
	} else if sortColumn == "roomid" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tmp.room_id asc")
		} else {
			return sqBuilder.OrderBy("tmp.room_id desc")
		}
	} else {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("bet_time asc")
		} else {
			return sqBuilder.OrderBy("bet_time desc")
		}
	}
}

func getAgentGameRatioStatListSumInfo(db *sql.DB, tablenames []string, queryCond []string, args []interface{}) (recordsTotal int, totalPlatformWinloseScore float64, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	/* tableName = agent_game_ratio_stat_YYYYMM AS s*/
	/*
		SELECT
		    COUNT(DISTINCT s.id), COALESCE(SUM(s.ya) + SUM(s.tax) - SUM(s.de) - SUM(s.bonus), 0)
		FROM (
		    SELECT * FROM agent_game_ratio_stat_202311
		    WHERE level_code LIKE $1 AND log_time >= $2 AND log_time <= $3
			UNION
			SELECT * FROM agent_game_ratio_stat_202310
		    WHERE level_code LIKE $4 AND log_time >= $5 AND log_time <= $6
		) AS s;
	*/

	selectQuery := `SELECT * FROM %s WHERE 1=1 %s`
	unionQuery := " UNION "

	finalQuerys := make([]string, 0)
	finalQuerys = append(finalQuerys, "(")
	finalArgs := make([]interface{}, 0)
	paramIdx := 1
	for i := 0; i < len(tablenames); i++ {
		var combitionQuery string
		for _, v := range queryCond {
			combitionQuery += " AND " + fmt.Sprintf(v, paramIdx)
			paramIdx += 1
		}

		queryTmp := fmt.Sprintf(selectQuery, tablenames[i], combitionQuery)

		finalQuerys = append(finalQuerys, queryTmp)
		if i+1 < len(tablenames) {
			finalQuerys = append(finalQuerys, unionQuery)
		} else {
			finalQuerys = append(finalQuerys, ") as s")
		}
		finalArgs = append(finalArgs, args...)
	}

	selectUnionQuery := ""
	for i := 0; i < len(finalQuerys); i++ {
		selectUnionQuery += finalQuerys[i]
	}

	query, _, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(DISTINCT s.id)", "COALESCE(SUM(s.ya) + SUM(s.tax) - SUM(s.de) - SUM(s.bonus), 0)").
		From(selectUnionQuery).
		ToSql()
	if err != nil {
		return
	}

	err = db.QueryRow(query, finalArgs...).Scan(&recordsTotal, &totalPlatformWinloseScore)
	if err != nil {
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func getAgentGameRatioStatList(db *sql.DB, req *model.GetAgentGameRatioStatListRequest, tablenames []string, queryCond []string, args []interface{}) (resp []interface{}, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	/* tableName = agent_game_ratio_stat_YYYYMM AS s*/
	/*
		SELECT
		    a.name,
		    s.agent_id,
		    s.game_id,
		    s.room_type,
		    SUM(s.play_count) AS play_count,
		    SUM(s.de) AS de,
		    SUM(s.ya) AS ya,
		    SUM(s.tax) AS tax,
		    SUM(s.de) + SUM(s.bonus) - SUM(s.ya) - SUM(s.tax) AS player_winlose,
		    (CASE WHEN SUM(s.ya) > 0 THEN (SUM(s.de) + SUM(s.bonus) - SUM(s.tax)) / SUM(s.ya) ELSE 0 END) AS taxed_rtp,
		    SUM(s.bonus) AS bonus
		FROM
		    (
		        SELECT * FROM agent_game_ratio_stat_202310
		        WHERE level_code LIKE $1 AND log_time >= $2 AND log_time <= $3
		        UNION ALL
		        SELECT * FROM agent_game_ratio_stat_202311
		        WHERE level_code LIKE $1 AND log_time >= $2 AND log_time <= $3
		    ) s
		INNER JOIN
		    agent AS a ON s.agent_id = a.id
		GROUP BY
		    s.agent_id, s.level_code, s.game_id, s.room_type, a.name
		ORDER BY
		    s.agent_id ASC
		LIMIT 10 OFFSET 0;
	*/
	selectUnionQuery := ""

	selectQuery := `SELECT * FROM %s WHERE 1=1 %s`
	unionQuery := " UNION "

	finalQuerys := make([]string, 0)
	finalQuerys = append(finalQuerys, "(")
	finalArgs := make([]interface{}, 0)
	paramIdx := 1
	for i := 0; i < len(tablenames); i++ {
		var combitionQuery string
		for _, v := range queryCond {
			combitionQuery += " AND " + fmt.Sprintf(v, paramIdx)
			paramIdx += 1
		}

		queryTmp := fmt.Sprintf(selectQuery, tablenames[i], combitionQuery)

		finalQuerys = append(finalQuerys, queryTmp)
		if i+1 < len(tablenames) {
			finalQuerys = append(finalQuerys, unionQuery)
		} else {
			finalQuerys = append(finalQuerys, ") as s")
		}
		finalArgs = append(finalArgs, args...)
	}

	for i := 0; i < len(finalQuerys); i++ {
		selectUnionQuery += finalQuerys[i]
	}

	sqBuilder := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select(
			"a.name",
			"s.agent_id",
			"s.game_id",
			"s.room_type",
			"SUM(s.play_count) AS play_count",
			"SUM(s.de) AS de",
			"SUM(s.ya) AS ya",
			"SUM(s.tax) AS tax",
			"SUM(s.de) + SUM(s.bonus) - SUM(s.ya) - SUM(s.tax) AS player_winlose",
			"(CASE WHEN SUM(s.ya) > 0 THEN (SUM(s.de) + SUM(s.bonus) - SUM(s.tax)) / SUM(s.ya) ELSE 0 END) AS taxed_rtp",
			"SUM(s.bonus) AS bonus",
		).
		From(selectUnionQuery).
		InnerJoin("agent AS a ON s.agent_id = a.id").
		GroupBy("s.agent_id", "s.level_code", "s.game_id", "s.room_type", "a.name")

	sqBuilder = getAgentGameRatioStatListEvaluateOrderBy(sqBuilder, req.SortColumn, req.SortDirection)

	query, _, _ := sqBuilder.
		Limit(uint64(req.Length)).Offset(uint64(req.Start)).
		ToSql()

	rows, err := db.Query(query, finalArgs...)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		tmp := &model.GetAgentGameRatioStatResponse{}
		if err = rows.Scan(&tmp.AgentName, &tmp.AgentId, &tmp.GameId, &tmp.RoomType, &tmp.PlayCount,
			&tmp.DeScore, &tmp.YaScore, &tmp.Tax, &tmp.PlayerWinlose, &tmp.TaxedRtp,
			&tmp.Bonus); err != nil {
			return
		}

		resp = append(resp, tmp)
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func getAgentGameRatioStatListEvaluateOrderBy(sqBuilder sq.SelectBuilder, sortColumn string, sortDirection int) sq.SelectBuilder {
	if sortColumn == "gamename" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("s.game_id asc")
		} else {
			return sqBuilder.OrderBy("s.game_id desc")
		}
	} else if sortColumn == "roomtype" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("s.room_type asc")
		} else {
			return sqBuilder.OrderBy("s.room_type desc")
		}
	} else if sortColumn == "playcount" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("play_count asc")
		} else {
			return sqBuilder.OrderBy("play_count desc")
		}
	} else if sortColumn == "yascore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("ya asc")
		} else {
			return sqBuilder.OrderBy("ya desc")
		}
	} else if sortColumn == "descore" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("de asc")
		} else {
			return sqBuilder.OrderBy("de desc")
		}
	} else if sortColumn == "tax" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("tax asc")
		} else {
			return sqBuilder.OrderBy("tax desc")
		}
	} else if sortColumn == "playerwinlose" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("player_winlose asc")
		} else {
			return sqBuilder.OrderBy("player_winlose desc")
		}
	} else if sortColumn == "taxedrtp" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("taxed_rtp asc")
		} else {
			return sqBuilder.OrderBy("taxed_rtp desc")
		}
	} else if sortColumn == "bonus" {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("bonus asc")
		} else {
			return sqBuilder.OrderBy("bonus desc")
		}
	} else {
		if sortDirection == definition.TABLE_SORT_DIRECTION_ASC {
			return sqBuilder.OrderBy("s.agent_id asc")
		} else {
			return sqBuilder.OrderBy("s.agent_id desc")
		}
	}
}

func getUserPlayLogList(db *sql.DB, tableName string, sqAnd *sq.And) (resp []*model.GetUserPlayLogResponse, errorCode int) {
	errorCode = definition.ERROR_CODE_ERROR_DATABASE

	sqBuilder := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("tmp.agent_id", "tmp.game_id", "tmp.room_type", "tmp.desk_id", "tmp.seat_id",
			"tmp.valid_score", "tmp.start_score", "tmp.end_score", "tmp.de_score", "tmp.ya_score",
			"tmp.bet_time", "tmp.username", "tmp.lognumber", "tmp.bet_id", "tmp.kill_type",
			"tmp.tax", "tmp.bonus", "tmp.jp_inject_water_rate", "tmp.jp_inject_water_score", "tmp.wallet_ledger_id",
			"tmp.kill_prob", "tmp.kill_level", "tmp.real_players").
		From(tableName).
		Where(sqAnd)

	query, args, _ := sqBuilder.ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return
	}

	defer rows.Close()

	tmpAgents := make(map[int]*table_model.Agent)

	for rows.Next() {
		tmp := &model.GetUserPlayLogResponse{}
		if err = rows.Scan(&tmp.AgentId, &tmp.GameId, &tmp.RoomType, &tmp.DeskId, &tmp.SeatId,
			&tmp.ValidScore, &tmp.StartScore, &tmp.EndScore, &tmp.DeScore, &tmp.YaScore,
			&tmp.BetTime, &tmp.UserName, &tmp.LogNumber, &tmp.BetId, &tmp.KillType,
			&tmp.Tax, &tmp.Bonus, &tmp.JpInjectWaterRate, &tmp.JpInjectWaterScore, &tmp.WalletLedgerId,
			&tmp.KillProb, &tmp.KillLevel, &tmp.RealPlayers); err != nil {
			return
		}

		if _, find := tmpAgents[tmp.AgentId]; !find {
			tmpAgents[tmp.AgentId] = global.AgentCache.Get(tmp.AgentId)
		}

		tmpAgent := tmpAgents[tmp.AgentId]
		if tmpAgent != nil {
			tmp.AgentName = tmpAgent.Name
		}

		resp = append(resp, tmp)
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}
