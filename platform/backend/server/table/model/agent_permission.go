package model

type AgentPermission struct {
	Id              string          `json:"id"`           // id(uuid)
	AgentId         int             `json:"agent_id"`     // 代理id
	AccountType     int             `json:"account_type"` // 帳號類型
	Name            string          `json:"name"`         // 角色名稱
	Info            string          `json:"info"`         // 備註
	PermissionBytes []byte          `json:"-"`            // DB paring用(權限內容)
	Permission      PermissionSlice `json:"permission"`   // 權限內容
}

type PermissionSlice struct {
	List []int `json:"list"` // 權限
}
