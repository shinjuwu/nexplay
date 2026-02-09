package model

import "time"

//*****************************************************************************

type CollectorGameListRequest struct {
	ID   string      `json:"id"`
	Data []*GameList `json:"data"`
}

//*****************************************************************************

type WalletLedger struct {
	ID         string    `json:"id"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	Kind       int       `json:"kind"`
	Status     int       `json:"status"`
	ChangeSet  string    `json:"changeset"` // You can use map[string]interface{} for JSON
	CreateTime time.Time `json:"create_time"`
}

type CollectorCoinInOutRequest struct {
	ID   string          `json:"id"`
	Data []*WalletLedger `json:"data"`
}

// type CollectorCoinInOutResponse struct {
// }

//*****************************************************************************

type UserPlayLog struct {
	LogNumber  string    `json:"lognumber"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	GameID     int       `json:"game_id"`
	GameName   string    `json:"game_name"`
	RoomType   int       `json:"room_type"`
	RoomName   string    `json:"room_name"`
	DeScore    float64   `json:"de_score"`
	Bonus      float64   `json:"bonus"`
	BetID      string    `json:"bet_id"`
	BetTime    time.Time `json:"bet_time"`
	CreateTime time.Time `json:"create_time"`
}

type CollectorAbnormalWinAndLoseRequest struct {
	ID   string         `json:"id"`
	Data []*UserPlayLog `json:"data"`
}

// type CollectorAbnormalWinAndLoseResponse struct {
// }

//*****************************************************************************

type CollectorRTPStat struct {
	GameId     int       `json:"game_id"`
	GameType   int       `json:"game_type"`
	RoomType   int       `json:"room_type"`
	Ya         float64   `json:"ya"`
	VaildYa    float64   `json:"vaild_ya"`
	De         float64   `json:"de"`
	Tax        float64   `json:"tax"`
	Bonus      float64   `json:"bonus"`
	PlayCount  int       `json:"play_count"`
	UpdateTime time.Time `json:"update_time"`
}

type CollectorRTPStatRequest struct {
	ID   string              `json:"id"`
	Data []*CollectorRTPStat `json:"data"`
}

// type CollectorRTPStatResponse struct {
// }
