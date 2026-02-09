package global

import (
	"backend/internal/statistical"
	table_model "backend/server/table/model"
	"encoding/json"

	"backend/pkg/logger"
	"backend/pkg/utils"
	"database/sql"
	"definition"
	"time"
)

func T_CalDbArchive(logger *logger.RuntimeGoLogger, db *sql.DB) {
	var next time.Duration
	var c *time.Timer
	for {
		next = _execDbArchive(logger, db)
		c = time.NewTimer(next)
		<-c.C
	}
}

func _execDbArchive(logger *logger.RuntimeGoLogger, db *sql.DB) time.Duration {
	// 固定檢測時間
	now := time.Now()
	ftTimeNow := now.UnixNano() / int64(time.Millisecond)

	logger.Printf("_execDbArchive run, ftTimeNow: %v", ftTimeNow)

	storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_ARCHIVE_SETTING)
	if !ok {
		logger.Error("db storage %s not exist", definition.STORAGE_KEY_ARCHIVE_SETTING)
		return time.Duration(0)
	}

	var archiveSetting table_model.ArchiveSetting
	json.Unmarshal([]byte(storage.Value), &archiveSetting)

	today := utils.GetTimeNowUTCTodayTime()
	todayWorkTime := today.Add(time.Duration(archiveSetting.ArchiveFireTime) * time.Hour)
	nextWorkTime := todayWorkTime.AddDate(0, 0, 1)

	// 今日備份時間還沒到
	if !now.After(todayWorkTime) {
		return time.Until(todayWorkTime)
	}

	// 今日已經備份完成不再處理(避免server重開多次執行)
	if !archiveSetting.LastArchiveTime.Before(today) {
		return time.Until(nextWorkTime)
	}

	archiveResult := true
	archiveResult = archiveResult && _execDbArchiveAgentGameUsersStatMin(logger, db, archiveSetting.AgentGameUsersStatMinRetainDataDays, archiveSetting.DeleteRetainDataBatchCount)

	// 備份過程發生失敗則延遲一小時後再進行，避免一直失敗
	if !archiveResult {
		return time.Until(time.Now().Add(time.Hour))
	}

	archiveSetting.LastArchiveTime = today
	if err := GlobalStorage.Update(definition.STORAGE_KEY_ARCHIVE_SETTING, utils.ToJSON(archiveSetting)); err != nil {
		logger.Error("db storage %s update has error: %v", definition.STORAGE_KEY_ARCHIVE_SETTING, err)
		return time.Until(time.Now().Add(time.Hour))
	}

	return time.Until(nextWorkTime)
}

func _execDbArchiveAgentGameUsersStatMin(logger *logger.RuntimeGoLogger, db *sql.DB, retainDataDays int, batchCount int) bool {
	archiveTime := utils.GetTimeNowUTCTodayTime().AddDate(0, 0, -retainDataDays)
	lastArchiveLogTime := utils.TransUnsignedTimeUTCFormat(statistical.SAVE_PERIOD_MINUTE, archiveTime)

	query := `DELETE FROM "public"."agent_game_users_stat_min"
		WHERE "ctid" IN (
			SELECT "ctid"
				FROM "public"."agent_game_users_stat_min"
				WHERE "log_time" < $1
				ORDER BY "log_time", "agent_id"
				LIMIT $2
		)`

	var result sql.Result
	var err error
	var rowsAffetedCount int64
	for {
		result, err = db.Exec(query, lastArchiveLogTime, batchCount)
		if err != nil {
			return false
		}

		rowsAffetedCount, err = result.RowsAffected()
		if err != nil {
			return false
		}

		if rowsAffetedCount < int64(batchCount) {
			return true
		}
	}
}
