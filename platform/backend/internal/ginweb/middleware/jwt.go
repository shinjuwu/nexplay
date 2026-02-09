package middleware

import (
	"definition"
	"time"

	"backend/internal/ginweb/response"
	"backend/pkg/jwt"
	"backend/server/global"
	"backend/server/table/model"

	"github.com/gin-gonic/gin"
)

const (
	localIp = "127.0.0.1"
)

func JWT(jwtManager *jwt.JwtManager, agentCache *global.GlobalAgentCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var claims *jwt.CustomClaims
		code = definition.ERROR_CODE_SUCCESS

		token := c.Request.Header.Get("Dcc-Token")

		if token == "" {
			code = definition.ERROR_CODE_ERROR_JWT
		} else {
			if isHasBlack := jwtManager.CheckBlackTokenExist(token); isHasBlack {
				code = definition.ERROR_CODE_ERROR_JWT
			} else {
				// 解析token
				parseResult, err := jwtManager.ParseToken(token)
				claims = parseResult
				if err != nil {
					code = definition.ERROR_CODE_ERROR_JWT
				} else if time.Now().UTC().Unix() > claims.ExpiresAt.Unix() {
					code = definition.ERROR_CODE_ERROR_JWT
				} else {
					clientIp := c.ClientIP()
					if clientIp != localIp {
						clientAgent := agentCache.Get(int(claims.BaseClaims.ID))
						if clientIp != claims.Ip {
							code = definition.ERROR_CODE_ERROR_JWT
						} else if !CheckAgentIpWhitelist(clientIp, clientAgent.IPWhitelist) {
							code = definition.ERROR_CODE_ERROR_JWT
						}
					}
				}

			}
		}
		if code != definition.ERROR_CODE_SUCCESS {
			response.StatusUnauthorized(c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func CheckAgentIpWhitelist(clientIp string, ipWhitelist []*model.AgentIPWhitelistObj) bool {
	// Notice: 後白IP白名單必須設定
	if len(ipWhitelist) <= 0 {
		return false
	}

	for _, ipWhitelistObj := range ipWhitelist {
		ip := ipWhitelistObj.IPAddress

		if ip[len(ip)-1:] == "*" {
			ip = ip[:len(ip)-1]
		}

		if len(clientIp) < len(ip) {
			continue
		}

		if clientIp[:len(ip)] == ip {
			return true
		}
	}

	return false
}
