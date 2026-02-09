package module_lib

import (
	"database/sql"
	"sync"
	"time"
)

var (
	uspInsertGameRatioStatLock = new(sync.Mutex)

	GameRatioStatTableName = "game_ratio_stat"
)

/*
	統計代理遊戲機率設定
	usp_insert_agent_game_ratio_stat
*/
func UspInsertGameRatioStat(db *sql.DB, platform string, gameId, gameType, roomType, playCount int, ya, de, validYa, tax, bonus float64, lastBetTime time.Time) {

	uspInsertGameRatioStatLock.Lock()
	defer uspInsertGameRatioStatLock.Unlock()

	query := `CALL "public"."usp_insert_game_ratio_stat"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, _ = db.Exec(query, platform,
		gameId, gameType, roomType, de, ya,
		validYa, tax, bonus, playCount, lastBetTime)
}
