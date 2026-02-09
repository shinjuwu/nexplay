package model

import "time"

type UserPlayLog struct {
	Lognumber  string    `json:"lognumber"`   //單號
	AgentId    int       `json:"agent_id"`    //代理id
	LevelCode  string    `json:"level_code"`  //層級碼
	UserId     int       `json:"user_id"`     //玩家id
	UserName   string    `json:"username"`    //玩家名稱
	GameId     int       `json:"game_id"`     //遊戲id
	RoomType   int       `json:"room_type"`   //房間類型
	DeskId     int       `json:"desk_id"`     //桌子id
	SeatId     int       `json:"seat_id"`     //座位id
	Exchange   float64   `json:"exchange"`    //一幣分值(沒有就填1)
	DeScore    float64   `json:"de_score"`    //總得遊戲分(float64)
	YaScore    float64   `json:"ya_score"`    //總壓遊戲分(float64)
	ValidScore float64   `json:"valid_score"` //有效投注(float64)
	Tax        float64   `json:"tax"`         //抽水(float64)
	StartScore float64   `json:"start_score"` // 玩家壓住前遊戲分
	EndScore   float64   `json:"end_score"`   // 玩家結算後遊戲分
	CreateTime time.Time `json:"create_time"` // 建立時間
	BetTime    time.Time `json:"bet_time"`    //遊戲結算時間(UTC+0, format: 2006-01-02T15:04:05.000Z)
	IsRobot    int       `json:"is_robot"`    //是否為機器人
	IsBigWin   bool      `json:"is_big_win"`  //是否為大獎
	IsIssue    bool      `json:"is_issue"`    //是否為問題單
}

func NewUserPlayLog(lognumber string, agentId, userId, gameId, roomType,
	deskId, seatId int, exchange, deScore, yaScore,
	validScore, tax, startScore, endScore float64, createTime,
	betTime time.Time, isRobot int, isBigWin, isIssue bool, userName, levelCode string) *UserPlayLog {
	return &UserPlayLog{
		Lognumber:  lognumber,
		AgentId:    agentId,
		LevelCode:  levelCode,
		UserId:     userId,
		GameId:     gameId,
		RoomType:   roomType,
		DeskId:     deskId,
		SeatId:     seatId,
		Exchange:   exchange,
		DeScore:    deScore,
		YaScore:    yaScore,
		ValidScore: validScore,
		Tax:        tax,
		StartScore: startScore,
		EndScore:   endScore,
		CreateTime: createTime,
		BetTime:    betTime,
		IsRobot:    isRobot,
		IsBigWin:   isBigWin,
		IsIssue:    isIssue,
		UserName:   userName,
	}
}
