package model

import (
	"backend/pkg/utils"
	"definition"
)

type GetAgentPermissionListRequest struct {
	Name string `json:"name"` // 角色名稱
}

type GetAgentPermissionListResponse struct {
	Id          string `json:"id"`           // uuid(唯一)
	AgentId     int    `json:"agent_id"`     // 代理id
	Name        string `json:"name"`         // 角色名稱
	Info        string `json:"info"`         // 備註
	AccountType int    `json:"account_type"` // 角色類型
}

type GetAgentPermissionRequest struct {
	Id string `json:"id"` // uuid(唯一)
}

func (r *GetAgentPermissionRequest) CheckParams() int {
	if r.Id == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type GetAgentPermissionResponse struct {
	GetAgentPermissionListResponse
	Permissions     []int  `json:"permissions"` // 權限列表
	PermissionBytes []byte `json:"-"`           // 權限列表(DB parsing用)
}

type CreateAgentPermissionRequest struct {
	Name        string `json:"name"`         // 角色名稱(長度16)
	Info        string `json:"info"`         // 備註(長度100)
	AccountType int    `json:"account_type"` // 角色類型(0:admin 1:總代理 2:子代)
	Permissions []int  `json:"permissions"`  // 權限列表
}

func (r *CreateAgentPermissionRequest) CheckParams() int {
	if r.Name == "" || utils.WordLength(r.Name) > 16 ||
		utils.WordLength(r.Info) > 100 ||
		r.AccountType < definition.ACCOUNT_TYPE_ADMIN || r.AccountType > definition.ACCOUNT_TYPE_NORMAL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type SetAgentPermissionRequest struct {
	CreateAgentPermissionRequest
	Id string `json:"id"` // uuid(唯一)
}

func (r *SetAgentPermissionRequest) CheckParams() int {
	if r.Id == "" ||
		r.Name == "" || utils.WordLength(r.Name) > 16 ||
		utils.WordLength(r.Info) > 100 ||
		r.AccountType < definition.ACCOUNT_TYPE_ADMIN || r.AccountType > definition.ACCOUNT_TYPE_NORMAL {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type DeleteAgentPermissionRequest struct {
	Id string `json:"id"` // uuid(唯一)
}

func (r *DeleteAgentPermissionRequest) CheckParams() int {
	if r.Id == "" {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}
