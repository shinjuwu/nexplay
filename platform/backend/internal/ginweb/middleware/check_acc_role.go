package middleware

import (
	"backend/internal/ginweb/response"
	"backend/pkg/config"
	"backend/pkg/jwt"
	"backend/server/global"
	"log"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

/*
角色權限設置定義
*/

func CheckUserRole(config config.Config, systemPermissionData *sync.Map, agentPermissionCache *global.GlobalAgentPermissionCache, jwt *jwt.JwtManager) gin.HandlerFunc {

	isLoadSwagger := config.GetApp().LoadSwagger
	checkHeader := config.GetApp().CheckHeader

	accManageHeader := config.GetApp().AccManageHeader
	accNormalHeader := config.GetApp().AccNormalHeader

	return func(c *gin.Context) {

		uri := c.Request.RequestURI

		checkIsSwagger := false
		// 如果 config 有設定載入 swagger 那就不阻擋 swagger uri
		if isLoadSwagger {
			referer := c.Request.Header.Get("Referer")
			if referer != "" {
				pathArray := strings.Split(referer, "/")[1:]
				if pathArray[2] == "swagger" {
					checkIsSwagger = true
				}
			}
		}

		if !checkIsSwagger {
			idx := strings.Index(uri, "?")
			if idx >= 0 {
				uri = uri[:idx]
			}
			// check api path exist
			val, ok := systemPermissionData.Load(uri)
			if ok {
				// log.Println(val)

				systemPermission := val.(map[string]interface{})

				// 檢查該 API 是否已被關閉
				if !systemPermission["is_enabled"].(bool) {
					response.StatusBadRequest(c)
					c.Abort()
					return
				}

				//不用驗證直接跳過不檢查
				if systemPermission["is_required"].(bool) {

					token := c.Request.Header.Get("Dcc-Token")

					// 沒有token 代表沒有登入
					if token != "" {
						claims := jwt.GetClaims(token)
						if claims != nil {

							if checkHeader {
								// 固定角色只能使用固定 header
								// config 內設定之隱藏碼，只有在經過nginx 之後才會加上
								// # 加入自定義標頭(不固定位數長度)
								// proxy_set_header X-Dcc-Header "XXXXXXXX";
								customerHeader := c.Request.Header.Get("X-Dcc-Header")

								if len(claims.BaseClaims.LevelCode) == 4 {
									if customerHeader != accManageHeader {
										response.StatusBadRequest(c)
										c.Abort()
										return
									}
								} else if len(claims.BaseClaims.LevelCode) > 4 {
									if customerHeader != accNormalHeader {
										response.StatusBadRequest(c)
										c.Abort()
										return
									}
								} else {
									// 有帶token 的狀態下，要比對TOKEN
									response.StatusBadRequest(c)
									c.Abort()
									return
								}
							}

							fcCode := int(systemPermission["feature_code"].(float64))

							agentPermission := agentPermissionCache.Get(claims.PermissionId)
							if !slices.Contains(agentPermission.Permission.List, fcCode) {
								response.StatusUnauthorized(c)
								c.Abort()
								return
							}
						} else {
							// 有token 本地端沒有憑證，直接回傳失敗
							response.StatusUnauthorized(c)
							c.Abort()
							return
						}
					} else {
						response.StatusBadRequest(c)
						c.Abort()
						return
					}
				}
			} else {
				log.Printf("CheckUserRole api not in list, uri: %s", uri)
				// response.StatusBadRequest(c)
				// c.Abort()
				// return
			}
		}

		c.Next()
	}
}
