package model

import "time"

type PermissionList struct {
	FeatureCode int       `json:"feature_code"` // 功能代碼(PK)
	Name        string    `json:"name"`         // 名稱
	ApiPath     string    `json:"api_path"`     // api 路徑
	IsEnabled   bool      `json:"is_enabled"`   // 是否開放
	IsRequired  bool      `json:"is_required"`  // 是否開放
	Remark      string    `json:"remark"`       // 備註
	CreateTime  time.Time `json:"create_time"`  // 創建時間
	UpdateTime  time.Time `json:"update_time"`  // 更新時間
	ActionType  int       `json:"action_type"`  // 操作類型
}

func NewEmptyPermissionList() *PermissionList {
	return &PermissionList{}
}

func NewPermissionList(featureCode, actionType int, name, apiPath, remark string, isEnabled, isRequired bool, createTimem, updateTime time.Time) *PermissionList {
	return &PermissionList{
		FeatureCode: featureCode,
		Name:        name,
		ApiPath:     apiPath,
		IsEnabled:   isEnabled,
		IsRequired:  isRequired,
		Remark:      remark,
		CreateTime:  createTimem,
		UpdateTime:  updateTime,
		ActionType:  actionType,
	}
}
