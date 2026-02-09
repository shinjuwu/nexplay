package api

import (
	"database/sql"
)

const (
	IB_HANDLE_IDX_TELEGRAM = 1
	IB_HANDLE_IDX_SLACK    = 2
	IB_HANDLE_IDX_LINE     = 3
)

type IMBotHandleFunc func(db *sql.DB, payload string) error

type IMBot struct {
	IMBotHandle map[int]IMBotHandleFunc
}

func NewIMBot() *IMBot {
	return &IMBot{
		IMBotHandle: map[int]IMBotHandleFunc{
			IB_HANDLE_IDX_TELEGRAM: TelegramSend,
		},
	}
}
