package model

import "time"

/*
通用查詢條件設定
*/

/*
通用查詢條件
*/
type CommonQueryConditionsRequest struct {
	Param1 string `json:"param1"` // 單一搜尋數字
	Param2 string `json:"Param2"` // 單一搜尋字串
	// QueryCondition []string  `json:"query_condition"` // 查詢條件(字串)
	QueryCount   int       `json:"query_count"`    // 查詢筆數(預設20)
	PageCount    int       `json:"page_count"`     // 查詢頁數
	IsFuzzyQuery bool      `json:"is_fuzzy_query"` // 是否模糊查詢
	StartTime    time.Time `json:"start_time"`     // 查詢開始時間
	EndTime      time.Time `json:"end_time"`       // 查詢結束時間
}

/*
通用查詢回覆
*/
type CommonQueryConditionsResponse struct {
	Result         []interface{} `json:"rt"` // 查詢結果
	DataCount      int           `json:"dc"` // 每頁顯示資料筆數
	PageCount      int           `json:"pc"` // 總頁數
	TotalDataCount int           `json:"tc"` // 資料總筆數
	PageNow        int           `json:"pn"` // 目前頁數
}
