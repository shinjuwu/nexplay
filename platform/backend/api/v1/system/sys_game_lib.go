package system

import (
	"backend/api/v1/model"
	"backend/internal/ginweb"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"

	sq "github.com/Masterminds/squirrel"
)

func setGameState(db *sql.DB, req *model.SetGameStateRequest) int {
	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("game").
		Set("state", req.State).
		Where(sq.Eq{"id": req.GameId}).
		ToSql()

	if err != nil {
		return definition.ERROR_CODE_ERROR_DATABASE
	}

	_, err = db.Exec(query, args...)
	if err != nil {
		return definition.ERROR_CODE_ERROR_DATABASE
	}

	return definition.ERROR_CODE_SUCCESS
}

func notifyGameList(logger ginweb.ILogger, db *sql.DB, gameId int, gameCode string) (errorCode int, err error) {
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("state").
		From("game").
		Where(sq.Eq{"id": gameId}).
		ToSql()

	state := int16(-1)

	err = db.QueryRow(query, args...).Scan(&state)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	gamelists := make([]map[string]interface{}, 0)
	gamelists = append(
		gamelists,
		getGameList(
			gameId,
			gameCode,
			state,
		),
	)

	errorCode, err = notification.SendSetGameList(gamelists)
	if err != nil {
		logger.Info("notification.SendSetGameList fail, gamelists=%v, resp code=%d", gamelists, errorCode)
		errorCode = definition.ERROR_CODE_ERROR_NOTIFICATION
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func notifyGameInfo(logger ginweb.ILogger, db *sql.DB, gameId int, gameCode string) (errorCode int, err error) {
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(agent_id)").
		From("agent_game").
		Where(sq.Eq{"game_id": gameId}).
		ToSql()

	count := 0

	err = db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	start := 0
	length := 1000

	gameInfos := make([]map[string]interface{}, 0)
	for start < count {
		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("agent_id", "state").
			From("agent_game").
			Where(sq.Eq{"game_id": gameId}).
			Offset(uint64(start)).
			Limit(uint64(length)).
			ToSql()

		var rows *sql.Rows
		rows, err = db.Query(query, args...)
		if err != nil {
			errorCode = definition.ERROR_CODE_ERROR_DATABASE
			return
		}

		for rows.Next() {
			agentId := 0
			state := int16(0)
			if err = rows.Scan(&agentId, &state); err != nil {
				errorCode = definition.ERROR_CODE_ERROR_DATABASE
				rows.Close()
				return
			}

			gameInfos = append(
				gameInfos,
				getGameInfo(
					agentId,
					gameId,
					gameCode,
					state,
				),
			)
		}

		start += length

		rows.Close()
	}

	errorCode, err = notification.SendSetGameInfo(gameInfos)
	if err != nil {
		logger.Info("notification.SendSetGameInfo fail, gameInfos=%v, resp code=%d", gameInfos, errorCode)
		errorCode = definition.ERROR_CODE_ERROR_NOTIFICATION
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func notifyLobbyInfo(logger ginweb.ILogger, db *sql.DB, gameId int) (errorCode int, err error) {
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(agr.agent_id)").
		From("agent_game_room AS agr").
		InnerJoin("game_room AS gr ON agr.game_room_id = gr.id").
		Where(sq.Eq{"gr.game_id": gameId}).
		ToSql()

	count := 0

	err = db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	start := 0
	length := 1000

	lobbyInfos := make([]map[string]interface{}, 0)
	for start < count {
		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("agr.agent_id", "agr.game_room_id", "agr.state").
			From("agent_game_room AS agr").
			InnerJoin("game_room AS gr ON agr.game_room_id = gr.id").
			Where(sq.Eq{"gr.game_id": gameId}).
			Offset(uint64(start)).
			Limit(uint64(length)).
			ToSql()

		var rows *sql.Rows
		rows, err = db.Query(query, args...)
		if err != nil {
			errorCode = definition.ERROR_CODE_ERROR_DATABASE
			return
		}

		for rows.Next() {
			agentId := 0
			tableId := 0
			state := int16(0)

			if err = rows.Scan(&agentId, &tableId, &state); err != nil {
				errorCode = definition.ERROR_CODE_ERROR_DATABASE
				rows.Close()
				return
			}

			lobbyInfos = append(
				lobbyInfos,
				getLobbyInfo(
					agentId,
					gameId,
					tableId,
					state,
				),
			)
		}

		start += length

		rows.Close()
	}

	errorCode, err = notification.SendSetLobbyInfo(lobbyInfos)
	if err != nil {
		logger.Info("notification.SendSetLobbyInfo fail, lobbyInfos=%v, resp code=%d", lobbyInfos, errorCode)
		errorCode = definition.ERROR_CODE_ERROR_NOTIFICATION
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func updateKillDiveInfo(logger ginweb.ILogger, db *sql.DB, gameId int) (errorCode int, err error) {
	val, apiCode, apiErr := notification.Getdefaultkilldiveinfo()
	if apiErr != nil {
		logger.Error("Getdefaultkilldiveinfo code:%d, err: %v", apiCode, apiErr)
		errorCode = definition.ERROR_CODE_ERROR_NOTIFICATION
		err = apiErr
		return
	}

	curKillDiveInfo := utils.ToJSON(val)

	err = global.GlobalStorage.Update(definition.STORAGE_KEY_GAMEKILLDIVEINFO, curKillDiveInfo)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	isExistsDefaultGameKillDiveInfo := false

	for _, curDefaultKillDiveInfo := range utils.ToArrayMap([]byte(curKillDiveInfo)) {
		curGameIdF, ok := curDefaultKillDiveInfo["GameId"].(float64)
		if !ok {
			continue
		}

		curGameId := int(curGameIdF)
		if curGameId == gameId {
			isExistsDefaultGameKillDiveInfo = true
			break
		}
	}

	// 遊戲預設殺放不存在視為成功，不需要進行新增代理遊戲殺放，下次重新同步時再確認即可
	if !isExistsDefaultGameKillDiveInfo {
		errorCode = definition.ERROR_CODE_SUCCESS
		return
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("1").
		Prefix("SELECT EXISTS (").
		From("agent_game_ratio").
		Where(sq.Eq{"game_id": gameId}).
		Suffix(")").
		ToSql()

	var isExistsAgentGameKillDiveInfo bool
	err = db.QueryRow(query, args...).Scan(&isExistsAgentGameKillDiveInfo)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	// 代理遊戲殺放只需要新增一次，後續修改透過遊戲風控設定進行處理
	if isExistsAgentGameKillDiveInfo {
		errorCode = definition.ERROR_CODE_SUCCESS
		return
	}

	err = global.AgentGameRatioCache.CreateNewGameToCacheFromJson(gameId, curKillDiveInfo)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}

func notifyKillDiveInfo(logger ginweb.ILogger, db *sql.DB, gameId int) (errorCode int, err error) {
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("COUNT(id)").
		From("agent_game_ratio").
		Where(sq.Eq{"game_id": gameId}).
		ToSql()

	count := 0

	err = db.QueryRow(query, args...).Scan(&count)
	if err != nil {
		errorCode = definition.ERROR_CODE_ERROR_DATABASE
		return
	}

	start := 0
	length := 1000

	killDiveInfo := make([]map[string]interface{}, 0)
	for start < count {
		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select(
				"agent_id", "game_id", "room_type", "kill_ratio", "new_kill_ratio",
				"active_num",
			).
			From("agent_game_ratio").
			Where(sq.Eq{"game_id": gameId}).
			OrderBy("id asc").
			Offset(uint64(start)).
			Limit(uint64(length)).
			ToSql()

		var rows *sql.Rows
		rows, err = db.Query(query, args...)
		if err != nil {
			errorCode = definition.ERROR_CODE_ERROR_DATABASE
			return
		}

		for rows.Next() {
			agentId := 0
			gameId := 0
			roomType := 0
			killRatio := float64(0)
			newKilRatio := float64(0)
			activeNum := 0

			if err = rows.Scan(
				&agentId, &gameId, &roomType, &killRatio, &newKilRatio,
				&activeNum,
			); err != nil {
				errorCode = definition.ERROR_CODE_ERROR_DATABASE
				rows.Close()
				return
			}

			roomId, convErr := getRoomId(gameId, roomType)
			if convErr != nil {
				errorCode = definition.ERROR_CODE_ERROR_DATABASE
				err = convErr
				rows.Close()
				return
			}

			killDiveInfo = append(
				killDiveInfo,
				getKillDiveInfo(
					agentId,
					gameId,
					roomId,
					activeNum,
					killRatio,
					newKilRatio,
				),
			)
		}

		start += length

		rows.Close()
	}

	if len(killDiveInfo) > 0 {
		// TO DO: 追蹤Bug問題加入的Log，問題解決後要移除
		logger.Info("[代理殺放追蹤]遊戲同步代理設置殺放，殺放內容: %v", killDiveInfo)
		_, errorCode, err = notification.SendSetKillDives(killDiveInfo)
		if err != nil {
			logger.Info("notification.SendSetKillDives fail, killDiveInfo=%v, resp code=%d, err=%v", killDiveInfo, errorCode, err)
			errorCode = definition.ERROR_CODE_ERROR_NOTIFICATION
			return
		}
	}

	errorCode = definition.ERROR_CODE_SUCCESS

	return
}
