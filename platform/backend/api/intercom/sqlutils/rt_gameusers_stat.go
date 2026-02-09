package sqlutils

import (
	"backend/server/global"
	"database/sql"
	"time"
)

type PlayerStatData struct {
	LevelCode    string
	Username     string
	AgentId      int
	GameId       int
	GameType     int
	RoomType     int
	UserId       int
	OrderCount   int // 注單數量(排除機器人)
	WinCount     int
	LoseCount    int
	BigWinCount  int // 中大獎數量
	YaScore      float64
	VaildYaScore float64
	DeScore      float64
	Tax          float64
	Bonus        float64
	BetId        string
}

func CalRTGameUsersStat(db *sql.DB, playerStatDataList map[int]*PlayerStatData, betTime time.Time) error {

	for userId, playerStatData := range playerStatDataList {
		agentId := playerStatData.AgentId
		gameId := playerStatData.GameId
		playCount := playerStatData.OrderCount
		winCount := playerStatData.WinCount
		loseCount := playerStatData.LoseCount
		bigWinCount := playerStatData.BigWinCount

		global.UspGameUsersStat(db, agentId, userId, gameId, playCount, bigWinCount, winCount, loseCount, playerStatData.LevelCode, playerStatData.YaScore,
			playerStatData.DeScore, playerStatData.VaildYaScore, playerStatData.Tax, playerStatData.Bonus, betTime)
	}
	return nil
}
