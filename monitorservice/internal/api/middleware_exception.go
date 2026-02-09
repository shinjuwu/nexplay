package api

import "github.com/gin-gonic/gin"

/*
攔截異常 (panic or other)
*/

func SetExceptionHandler(logger ILogger) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Intercept exception: %v", err)

				StatusInternalServerError(c)

				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
