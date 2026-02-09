package model

import (
	"time"
)

/*
INSERT INTO "public"."user_play_log" ("platform", "lognumber", "agent_id", "agent_name", "user_id", "username", "game_id", "room_type", "de_score", "bonus", "bet_id", "bet_time", "create_time", "game_name")
VALUES ('dev', '1001-1684765252727607296', 3, 'test1', 1001, 'simon', 1001, 0, '8856.0200', '0.0000', '420230728031900340361899875', '2023-07-28 03:19:00.361+00', '2023-07-28 03:19:00.368583+00', '百家樂');

*/
type DBUserPlayLog struct {
	Platform   string    `json:"platform"`
	LogNumber  string    `json:"lognumber"`
	AgentID    int       `json:"agent_id"`
	AgentName  string    `json:"agent_name"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
	GameID     int       `json:"game_id"`
	GameName   string    `json:"game_name"`
	RoomType   int       `json:"room_type"`
	DeScore    float64   `json:"de_score"`
	Bonus      float64   `json:"bonus"`
	BetID      string    `json:"bet_id"`
	BetTime    time.Time `json:"bet_time"`
	CreateTime time.Time `json:"create_time"`
}
