package model

// play log 有兩層

// 第一層 playlog
type PlayLogRequest struct {
	Lognumber  string  `json:"lognumber"`   //單號
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
}

// 第二層 playlog
type PlayerStatData struct {
	LevelCode    string
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
}
