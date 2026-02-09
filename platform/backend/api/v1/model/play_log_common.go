package model

import (
	"definition"
	"time"
)

// 遊戲日誌解析
type GetPlayLogCommonRequest struct {
	LogNumber string `json:"lognumber"` // 局號
}

func (p *GetPlayLogCommonRequest) CheckParams() int {
	if p.LogNumber == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetPlayLogCommonResponse struct {
	LogNumber    string                 `json:"lognumber"`  // 單號(局號)
	GameId       int                    `json:"game_id"`    // 遊戲編號
	GameCode     string                 `json:"game_code"`  // 遊戲代碼
	RoomType     int                    `json:"room_type"`  // 房間類型
	PlayLog      map[string]interface{} `json:"play_log"`   // 遊戲詳情
	PlayLogBytes []byte                 `json:"-"`          // DB parsing用
	BetTime      time.Time              `json:"bet_time"`   // 遊戲結算時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	StartTime    time.Time              `json:"start_time"` // 遊戲開始時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	EndTime      time.Time              `json:"end_time"`   // 遊戲結束時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
}
