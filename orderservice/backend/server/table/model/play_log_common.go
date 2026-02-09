package model

import "time"

type PlayLogCommon struct {
	Lognumber  string    `json:"lognumber"`   //單號
	GameId     int       `json:"game_id"`     //遊戲id
	RoomType   int       `json:"room_type"`   //房間類型
	DeskId     int       `json:"desk_id"`     //桌子id
	Exchange   float64   `json:"exchange"`    //一幣分值(沒有就填1)
	Playlog    string    `json:"playlog"`     //詳情
	DeScore    float64   `json:"de_score"`    //總得遊戲分(float64)
	YaScore    float64   `json:"ya_score"`    //總壓遊戲分(float64)
	ValidScore float64   `json:"valid_score"` //有效投注(float64)
	Tax        float64   `json:"tax"`         //抽水(float64)
	CreateTime time.Time `json:"create_time"` //建立時間
	BetTime    time.Time `json:"bet_time"`    //遊戲結算時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	IsBigWin   bool      `json:"is_big_win"`  //是否為大獎
	IsIssue    bool      `json:"is_issue"`    //是否為問題單
}

func NewPlayLogCommon(lognumber string, gameId, roomType, deskId int, exchange float64,
	playlog string, deScore, yaScore, validScore, tax float64,
	createTime, betTime time.Time, isBigWin, isIssue bool) *PlayLogCommon {
	return &PlayLogCommon{
		Lognumber:  lognumber,
		GameId:     gameId,
		RoomType:   roomType,
		DeskId:     deskId,
		Exchange:   exchange,
		Playlog:    playlog,
		DeScore:    deScore,
		YaScore:    yaScore,
		ValidScore: validScore,
		Tax:        tax,
		CreateTime: createTime,
		BetTime:    betTime,
		IsBigWin:   isBigWin,
		IsIssue:    isIssue,
	}
}
