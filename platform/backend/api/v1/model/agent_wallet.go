package model

import (
	"backend/pkg/utils"
	"definition"
	"time"
)

type GetAgentWalletListRequest struct {
	AgentId int `json:"agent_id"` // 代理id
}

func (r *GetAgentWalletListRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type GetAgentWalletListResponse struct {
	AdminUserUsername string    `json:"admin_user_username"` // 代理帳號
	Id                int       `json:"id"`                  // 代理id
	Name              string    `json:"name"`                // 代理名稱
	LevelCode         string    `json:"level_code"`          // 代理層級代碼
	Commission        int       `json:"commission"`          // 分成
	Balance           float64   `json:"balance"`             // 餘額
	IsEnabled         int       `json:"is_enabled"`          // 狀態(1:啟用,0:關閉)
	CreateTime        time.Time `json:"create_time"`         // 創建時間
}

type SetAgentWalletRequest struct {
	AgentId     int     `json:"agent_id"`      // 代理id
	ChangeScore float64 `json:"change_amount"` // 變更點數
	Info        string  `json:"info"`          // 備註
}

func (r *SetAgentWalletRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL ||
		r.ChangeScore == 0 ||
		r.Info == "" || utils.WordLength(r.Info) > 100 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}

	return definition.ERROR_CODE_SUCCESS
}
