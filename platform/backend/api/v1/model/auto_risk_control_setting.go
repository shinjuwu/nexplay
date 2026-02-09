package model

import (
	table_model "backend/server/table/model"
)

type GetAutoRiskControlSetting struct {
	Current *table_model.AutoRiskControlSetting `json:"current"` // 目前風控設定
	Default *table_model.AutoRiskControlSetting `json:"default"` // 預設風控設定
}
