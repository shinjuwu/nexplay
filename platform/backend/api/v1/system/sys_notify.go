package system

import (
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/internal/notification"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"

	"github.com/gin-gonic/gin"
)

/***/

type SystemNotifyApi struct {
	BasePath string
}

func NewSystemNotifyApi(basePath string) api_cluster.IApiEach {
	return &SystemNotifyApi{
		BasePath: basePath,
	}
}

func (p *SystemNotifyApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemNotifyApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	g.Use(middleware.JWT(ginHandler.Jwt, global.AgentCache))
	g.GET("/getchatserviceconnInfo", ginHandler.Handle(p.GetChatServiceConnInfo))
	g.POST("/notifybroadcastmessage", ginHandler.Handle(p.NotifyBroadcastMessage))
}

// @Tags 即時通訊功能
// @Summary Get chat service connect information
// @Description 此接口用來取得chat service 連線資訊
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/notify/getchatserviceconnInfo [get]
func (p *SystemNotifyApi) GetChatServiceConnInfo(c *ginweb.Context) {
	// message := c.PostForm("message")
	var address []byte
	isEnabled := false
	query := "SELECT addresses, is_enabled FROM server_info WHERE code = $1 "
	err := c.DB().QueryRowContext(c.Request.Context(), query, "chat").Scan(&address, &isEnabled)
	if err != nil {
		if err != sql.ErrNoRows {
			c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
			return
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE_NO_ROWS)
			return
		}
	} else if !isEnabled {
		c.OkWithCode(definition.ERROR_CODE_ERROR_FEATURE_DISABLED)
		return
	}

	connInfo := utils.ToMap(address)

	c.Ok(connInfo)
}

// @Tags 即時通訊功能
// @Summary Notify user real time
// @Description 此接口用來即時通知後台前端新訊息發佈
// @Produce  application/json
// @Security BearerAuth
// @Param message formData string true "通知訊息"
// @Success 200 {object} response.Response{data=string,msg=string} "返回成功或失敗訊息"
// @Router /api/v1/notify/notifybroadcastmessage [post]
func (p *SystemNotifyApi) NotifyBroadcastMessage(c *ginweb.Context) {
	message := c.PostForm("message")

	if message == "" {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// check permission
	claims := c.GetUserClaims()
	if claims == nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	if claims.AccountType != definition.ACCOUNT_TYPE_ADMIN {
		c.OkWithCode(definition.ERROR_CODE_ERROR_PERMISSION)
		return
	}

	connInfo, code := getChatConnInfo(c.DB(), c.Request.Context())
	if code != definition.ERROR_CODE_SUCCESS {
		c.OkWithCode(code)
		return
	}

	channelId, ok := connInfo["channel"].(string)
	if !ok {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHANNEL_ID)
		return
	}

	success := notification.SendNotifyToFrontend(channelId, message, notification.ChatNotification_broadcast, connInfo, notification.API_CHAT_MESSAGE_SUBJECT_MESSAGE)
	if !success {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CHAT_SEND_MESSAGE_FAILED)
		return
	}
	c.Ok("")
}
