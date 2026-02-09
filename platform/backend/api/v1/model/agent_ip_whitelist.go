package model

import (
	"backend/pkg/utils"
	"backend/server/table/model"
	"definition"
)

type GetAgentIpWhitelistListRequest struct {
	AgentId int `json:"agent_id"` // 代理id
}

func (r *GetAgentIpWhitelistListRequest) CheckParams() int {
	if r.AgentId < definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type GetAgentIpWhitelistListResponse struct {
	Id                int    `json:"id"`                   // 代理id
	Name              string `json:"name"`                 // 代理名稱
	LevelCode         string `json:"level_code"`           // 代理level代碼
	IpAddressCount    int    `json:"ip_address_count"`     // 代理白名單數量
	ApiIpAddressCount int    `json:"api_ip_address_count"` // 代理白名單數量
	AdminUsername     string `json:"admin_username"`       // 代理帳號名稱
}

type GetAgentIpWhitelistRequest struct {
	AgentId int `json:"agent_id"` // 代理id
}

func (r *GetAgentIpWhitelistRequest) CheckParams() int {
	if r.AgentId <= definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type SetAgentIpWhitelistRequest struct {
	AgentId     int                         `json:"agent_id"`     // 代理id
	IpWhitelist []model.AgentIPWhitelistObj `json:"ip_whitelist"` // 代理ip白名單
}

func (r *SetAgentIpWhitelistRequest) CheckParams() int {
	checkIpWhitelistInfo := func(objs []model.AgentIPWhitelistObj) bool {
		for _, obj := range objs {
			if utils.WordLength(obj.Info) > 30 {
				return false
			}
		}
		return true
	}

	checkIpWhitelistAddress := func(objs []model.AgentIPWhitelistObj) bool {
		cache := make(map[string]struct{})
		for _, obj := range objs {
			if _, find := cache[obj.IPAddress]; find {
				return false
			}
			cache[obj.IPAddress] = struct{}{}
		}

		return true
	}

	if r.AgentId <= definition.AGENT_ID_ALL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	} else if len(r.IpWhitelist) <= 0 {
		return definition.ERROR_CODE_ERROR_AGENT_IP_WHITELIST_COUNT
	} else if len(r.IpWhitelist) > 30 ||
		!checkIpWhitelistInfo(r.IpWhitelist) ||
		!checkIpWhitelistAddress(r.IpWhitelist) {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}
