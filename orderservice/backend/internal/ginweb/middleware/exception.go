package middleware

import (
	"backend/internal/ginweb/response"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
攔截異常 (panic or other)
*/

func SetExceptionHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				response.StatusInternalServerError(c)

				c.Abort()
				return
			}
		}()
		c.Next()
	}
}
