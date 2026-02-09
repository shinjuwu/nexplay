package middleware

import (
	"github.com/gin-gonic/gin"
)

/*
限制每個呼叫的用戶在單位時間內呼叫的次數
做法:
1. 設定單位時間內可呼叫的次數，所有用戶共用
	* 可防止server崩潰
	* 但是被攻擊的話用戶也無法呼叫
2. 對每個呼叫的用戶做紀錄，設定單位時間內可呼叫的次數或是每呼叫一次要間隔多少時間
	* 不知道呼叫用戶的資訊，無法主動針對單一用戶，必須讓用戶自帶可識別資訊
*/

const (
	// API rate limit param setting.
	MAX_ALLOWED           = 20 //單位時間最大用戶連接數(可依照 cpu 核數做修改)
	SAME_IP_ALLOWED_COUNT = 10 //同IP單位時間最大用戶連接數
)

func RequestLimit(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
