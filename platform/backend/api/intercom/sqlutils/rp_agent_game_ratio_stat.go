package sqlutils

import (
	"backend/server/global"
	"database/sql"
	"fmt"
	"time"
)

func CalRPAgentGameRatioStat(db *sql.DB, playerStatDataList map[int]*PlayerStatData, betTime time.Time) error {

	for _, playerStatData := range playerStatDataList {
		if playerStatData.AgentId > 0 {
			id := fmt.Sprintf("%d_%d_%d_%d", playerStatData.AgentId, playerStatData.GameId, playerStatData.GameType, playerStatData.RoomType)
			global.UspInsertAgentGameRatioStat(db,
				id, playerStatData.LevelCode, playerStatData.AgentId, playerStatData.GameId, playerStatData.GameType,
				playerStatData.RoomType, playerStatData.OrderCount, playerStatData.YaScore, playerStatData.DeScore, playerStatData.VaildYaScore,
				playerStatData.Tax, playerStatData.Bonus, betTime)
		}
	}

	return nil
}
