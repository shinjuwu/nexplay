package squirrel

import (
	"backend/pkg/utils"
)

/*
TODO:
定義 table 所屬model
使用定義好的 model 操作 TABLE
*/

const (
	default_tag_name = "db"
)

type SquirrelModel struct {
	table       string             // 資料表名稱
	Scanner     func() interface{} // target table model
	tableFileds []string           // 指定結構的 tag 資料(對應資料庫 table)
}

func NewSquirrelModel(table string, destTable interface{}) *SquirrelModel {

	if table == "" {
		return nil
	}

	tableFileds := utils.GetFieldsName(default_tag_name, destTable)

	return &SquirrelModel{
		table: table,
		Scanner: func() interface{} {
			return destTable
		},
		tableFileds: tableFileds,
	}
}

func (p *SquirrelModel) Table() string {
	return p.table
}

func (p *SquirrelModel) TableFileds() []string {
	return p.tableFileds
}

func (p *SquirrelModel) BuildInsertSql() string {
	return ""
}

func (p *SquirrelModel) BuildSelectSql() string {
	return ""
}

func (p *SquirrelModel) BuildUpdateSql() string {
	return ""
}

func (p *SquirrelModel) BuildDeleteSql() string {
	return ""
}

func (p *SquirrelModel) BuildCountSql() string {
	return ""
}

func (p *SquirrelModel) BuilddExistSql() string {
	return ""
}
