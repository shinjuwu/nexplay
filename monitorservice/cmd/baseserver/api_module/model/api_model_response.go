package model

// for swagger model
type Response struct {
	Code    int         `json:"code"` // 回傳編碼
	Message string      `json:"msg"`  // 自定義訊息
	Data    interface{} `json:"data"` // 回傳資料
}
