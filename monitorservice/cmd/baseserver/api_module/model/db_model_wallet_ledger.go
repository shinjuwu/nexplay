package model

import (
	"time"
)

type DBWalletLedger struct {
	Platform   string    `json:"platform"`
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
