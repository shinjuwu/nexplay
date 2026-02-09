package model

import (
	"definition"
	"time"
)

type ServerSideTableRequest struct {
	Draw          int    `json:"draw"`           // 計數器 default(0)
	Start         int    `json:"start"`          // 查詢起始筆數(0表示第1筆資料)" default(0)
	Length        int    `json:"length"`         // 查詢筆數 default(10)
	SortColumn    string `json:"sort_column"`    // 排序欄位
	SortDirection int    `json:"sort_direction"` // 排序方向 0:升冪 1:降冪
}

func (r *ServerSideTableRequest) CheckServerSideTableRequest() int {
	if r.Start < 0 || r.Length < 0 || r.Draw < 0 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type DateTimeTableRequest struct {
	StartTime time.Time `json:"start_time"` // 查詢範圍起始時間
	EndTime   time.Time `json:"end_time"`   // 查詢範圍結束時間
}

func (r *DateTimeTableRequest) CheckDateTimeTableRequest() int {
	if r.EndTime.Sub(r.StartTime) < 0 {
		return definition.ERROR_CODE_ERROR_REQUEST_DATA
	}
	return definition.ERROR_CODE_SUCCESS
}

type DataTablesResponse struct {
	Draw         int           `json:"draw" form:"draw"` // 計數器
	Data         []interface{} `json:"data"`             // 查詢結果
	RecordsTotal int           `json:"recordsTotal"`     // 查詢範圍資料總筆數
}

type DefinitionTableConstant struct {
	DefaultLength     int   // 預設查詢筆數
	DefaultLengthMenu []int // 預設查詢筆數選單
}
