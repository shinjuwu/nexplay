package middleware

import (
	"backend/internal/ginweb"
	"backend/internal/ginweb/response"

	"github.com/gin-gonic/gin"
)

/*
攔截異常 (panic or other)
*/

func SetExceptionHandler(logger ginweb.ILogger) gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Intercept exception: %v", err)

				response.StatusInternalServerError(c)

				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
