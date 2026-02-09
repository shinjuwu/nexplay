package middleware

import (
	"backend/internal/ginweb/response"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SAME_USER_ALLOWED_TIME_INTERVAL_SEC = 10 // 限制同一用戶連線等待時間
)

var (
	userConnections = make(map[string]time.Time)
	mutex           = &sync.Mutex{}
)

func ConnectionsLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		agentId := c.Request.URL.Query().Get("agent") // 假设用户ID作为查询参数传递

		mutex.Lock()
		defer mutex.Unlock()

		if lastConnectionTime, ok := userConnections[agentId]; ok {
			elapsedTime := time.Since(lastConnectionTime)
			if elapsedTime.Seconds() < SAME_USER_ALLOWED_TIME_INTERVAL_SEC {
				// within 10 seconds, the same user can only connect once
				response.StatusTooManyRequestsError(c, agentId)
				c.Abort()
				return
			}
		}

		// update user last connect time
		userConnections[agentId] = time.Now()

		c.Next()
	}
}
