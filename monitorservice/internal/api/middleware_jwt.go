package api

import (
	"monitorservice/pkg/jwt"
	"time"

	"github.com/gin-gonic/gin"
)

// const (
// 	localIp = "127.0.0.1"
// )

func JWT(jwtManager *jwt.JwtManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		// var data interface{}
		var claims *jwt.CustomClaims
		code = ERROR_CODE_SUCCESS
		token := ""
		if len(c.Request.Header["Token"]) > 0 {
			token = c.Request.Header["Token"][0]
		}

		if token == "" {
			code = ERROR_CODE_ERROR_JWT
		} else {
			if isHasBlack := jwtManager.CheckBlackTokenExist(token); isHasBlack {
				code = ERROR_CODE_ERROR_JWT
			} else {
				// 解析token
				parseResult, err := jwtManager.ParseToken(token)
				claims = parseResult
				if err != nil {
					code = ERROR_CODE_ERROR_JWT
				} else if time.Now().UTC().Unix() > claims.ExpiresAt.Unix() {
					code = ERROR_CODE_ERROR_JWT
				}
			}
		}
		if code != ERROR_CODE_SUCCESS {
			StatusUnauthorized(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
