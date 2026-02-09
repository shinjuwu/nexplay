package api

import (
	"chatservice/pkg/config"
	"chatservice/pkg/melody"
	"chatservice/pkg/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// user 連線產生 websocket 使用
func NewSocketWsAcceptor(config config.Config, g *gin.Engine, m *melody.Melody, loginTokneStore *sync.Map) func(c *gin.Context) {
	return func(c *gin.Context) {

		// TODO: parse token to check authentication
		token := c.Request.URL.Query().Get("token")
		if token == "" {
			http.Error(c.Writer, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		// verify login token
		tmp, ok := loginTokneStore.Load(token)
		if !ok {
			http.Error(c.Writer, "Invalid token", http.StatusNotFound)
			return
		} else {
			// after verify, delete login token
			loginTokneStore.Delete(token)
		}

		// 帶入個人資料
		userInfo := tmp.(map[string]interface{})

		userId := utils.ToString(userInfo["user_id"], "")
		agentId := utils.ToInt(userInfo["agent_id"], 0)
		username := utils.ToString(userInfo["username"], "")
		platform := utils.ToString(userInfo["platform"], "")

		if userId == "" || agentId == 0 || username == "" {
			http.Error(c.Writer, "Error to get user info", http.StatusNotFound)
			return
		}

		// create user info
		userInfo[TEQUILA_CTX_RUNTIME_USER_ID] = userId
		userInfo[TEQUILA_CTX_RUNTIME_AGENT_ID] = agentId
		userInfo[TEQUILA_CTX_RUNTIME_USER_NAME] = username
		userInfo[TEQUILA_CTX_RUNTIME_PLATFORM] = platform

		m.HandleRequestWithKeys(c.Writer, c.Request, userInfo)
	}
}
