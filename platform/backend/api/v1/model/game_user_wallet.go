package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetGameUserWalletListRequest struct {
	Username string `json:"username"` // 玩家帳號
	ServerSideTableRequest
}

func (r *GetGameUserWalletListRequest) CheckParams() int {
	if r.Length > 100 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	if code := r.CheckServerSideTableRequest(); code != definition.ERROR_CODE_SUCCESS {
		return code
	}

	return definition.ERROR_CODE_SUCCESS
}

type GetGameUserWalletResponse struct {
	UserId     int       `json:"user_id"`     // 玩家id
	Username   string    `json:"username"`    // 玩家帳號
	AgentName  string    `json:"agent_name"`  // 代理名稱
	Gold       float64   `json:"gold"`        // 餘額
	LockGold   float64   `json:"lock_gold"`   // 遊戲鎖定分數
	IsEnabled  bool      `json:"is_enabled"`  // 狀態
	CreateTime time.Time `json:"create_time"` // 創建時間
}

type GetGameUserWalletListResponse struct {
	AgentBalance float64 `json:"agent_balance"` // 代理錢包餘額
	DataTablesResponse
}

type SetGameUserWalletRequest struct {
	UserId      int     `json:"user_id"`       // 玩家id
	ChangeScore float64 `json:"change_amount"` // 變更點數
	Info        string  `json:"info"`          // 備註
}

func (r *SetGameUserWalletRequest) CheckParams() int {
	if r.ChangeScore == 0 ||
		r.Info == "" || utils.WordLength(r.Info) > 100 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}
