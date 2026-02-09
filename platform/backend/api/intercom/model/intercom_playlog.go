package model

import "time"

type PlayLogRequest struct {
	Lognumber string `json:"lognumber"` //單號
	// UserId     int     `json:"user_id"`     //用戶id
	GameId     int     `json:"game_id"`     //遊戲id
	RoomType   int     `json:"room_type"`   //房間類型
	DeskId     int     `json:"desk_id"`     //桌子id
	Exchange   float64 `json:"exchange"`    //一幣分值(沒有就填1)
	Playlog    string  `json:"playlog"`     //遊戲記錄(json format)
	DeScore    float64 `json:"de_score"`    //總得遊戲分(float64)
	YaScore    float64 `json:"ya_score"`    //總壓遊戲分(float64)
	ValidScore float64 `json:"valid_score"` //有效投注(float64)
	Tax        float64 `json:"tax"`         //抽水(float64)
	Bonus      float64 `json:"bonus"`       //紅利(float64)
	BetTime    int64   `json:"bet_time"`    //遊戲結算時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	IsBigWin   bool    `json:"is_big_win"`  //是否為大獎
	IsIssue    bool    `json:"is_issue"`    //是否為問題單
	StartTime  int64   `json:"start_time"`  //遊戲開始時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	EndTime    int64   `json:"end_time"`    //遊戲結束時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
}

type PlayLogResponse struct {
}

// 遊戲LOG 轉換定義 start
type IPlayLog interface {
}

type PlayLogBaccarat struct {
}

func NewPlayLogBaccarat() IPlayLog {
	return &PlayLogBaccarat{}
}

func ConvertPlayLog(playlog string, gameId int) IPlayLog {

	switch gameId {
	case 1001:
		//
	default:
		//

	}
	return NewPlayLogBaccarat()
}

// 遊戲LOG 轉換定義 end

type RecalculateGAmeRecordRequest struct {
	StartTime time.Time `json:"start_time"` //開始時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	EndTime   time.Time `json:"end_time"`   //結束時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
}
