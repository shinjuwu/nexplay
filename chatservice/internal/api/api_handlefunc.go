package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func Login(db *sql.DB, c *gin.Context, loginTokneStore *sync.Map, username, platform string, agentId int) {

	returnCode := ErrCodeCommonSuccessed
	returnMessage := ErrMessageCommonSuccessed

	// username := utils.ToString(c.Request.URL.Query().Get("username"), "")
	// agentId := utils.ToInt(c.Request.URL.Query().Get("agent_id"), 0)
	// platform := utils.ToString(c.Request.URL.Query().Get("platform"), "")

	// check user account isn't exist, create new when user not exist
	userId := ""
	isDisabled := false
	createNew := false
	query := "SELECT id, is_disabled FROM users WHERE platform = $1 AND username = $2 AND agent_id = $3 "
	err := db.QueryRowContext(c.Request.Context(), query, platform, username, agentId).Scan(&userId, &isDisabled)
	if err != nil {
		if err != sql.ErrNoRows {
			// 	returnCode = ErrCodeCommonFailed
			// 	returnMessage = "User account not found."
			// } else {
			returnCode = ErrCodeDatabaseFailed
			returnMessage = "Error finding user account."
		} else {
			createNew = true
		}
	} else if isDisabled {
		returnCode = ErrCodeCommonFailed
		returnMessage = "User account banned."
	}

	if createNew {
		// create new account
		userId = uuid.Must(uuid.NewV4()).String()
		// Setup System user
		query = `INSERT INTO users (id, platform, username, agent_id)
			VALUES ($1, $2, $3, $4)`
		result, err := db.ExecContext(c.Request.Context(), query, userId, platform, username, agentId)
		if err != nil {
			returnCode = ErrCodeDatabaseFailed
			returnMessage = "Error finding or creating user account."
		} else {
			if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
				returnCode = ErrCodeDidNotInsertNew
				returnMessage = "Error finding or creating user account."
			}
		}
	}

	tmp := make(map[string]interface{}, 0)
	token := ""
	if returnCode == ErrCodeCommonSuccessed {
		// create login token
		token = uuid.Must(uuid.NewV4()).String()

		tmp["user_id"] = userId
		tmp["platform"] = platform
		tmp["token"] = token
		tmp["username"] = username
		tmp["agent_id"] = agentId

		loginTokneStore.Store(token, tmp)
	}

	httpCode := http.StatusOK
	if returnCode != ErrCodeCommonSuccessed {
		httpCode = http.StatusBadRequest
	}

	ret := make(map[string]interface{}, 0)

	ret["user_id"] = userId
	ret["platform"] = platform
	ret["token"] = token

	c.JSON(httpCode, gin.H{
		"code": returnCode,
		"msg":  returnMessage,
		"data": ret,
	})
}

func Broadcast(db *sql.DB, c *gin.Context, sessionManager SessionManager, platform, msgData, subject string) {

	returnCode := ErrCodeCommonSuccessed
	returnMessage := ErrMessageCommonSuccessed

	sendData := NewEnvelope("notifybroadcastmessage", msgData, subject, 0)

	send, err := json.Marshal(sendData)
	if err != nil {
		returnCode = ErrCodeCommonFailed
		returnMessage = ErrMessageCommonFailed
	} else {
		sessions := sessionManager.GetAll()
		for _, v := range sessions {
			if platform == v.Platform() {
				v.Sendb(send)
			}
		}
	}

	query := `INSERT INTO notification (content, platform)
			VALUES ($1, $2)`
	result, err := db.ExecContext(c.Request.Context(), query, send, platform)
	if err != nil {
		returnCode = ErrCodeDatabaseFailed
		returnMessage = "Error finding or creating user account."
	} else {
		if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
			returnCode = ErrCodeDidNotInsertNew
			returnMessage = "Error insert notification content."
		}
	}

	query = `UPDATE users SET unread=unread+1 WHERE platform = $1;`
	_, err = db.ExecContext(c.Request.Context(), query, platform)
	if err != nil {
		returnCode = ErrCodeDatabaseFailed
		returnMessage = "Error update users unread count."
	}

	c.JSON(http.StatusOK, gin.H{
		"code": returnCode,
		"msg":  returnMessage,
		"data": nil,
	})
}

// 針對某個用戶發出警告
// 之後可能新增 其他第三方通訊軟體的警告
func Message(db *sql.DB, c *gin.Context, sessionManager SessionManager, agentId int, platform, username, msgData, subject string) {
	returnCode := ErrCodeCommonSuccessed
	returnMessage := ErrMessageCommonSuccessed

	sendData := NewEnvelope("notifymessage", msgData, subject, 0)

	send, err := json.Marshal(sendData)
	if err != nil {
		returnCode = ErrCodeCommonFailed
		returnMessage = ErrMessageCommonFailed
	} else {
		sessions := sessionManager.GetAll()
		for _, v := range sessions {
			if platform == v.Platform() && username == v.Username() {
				v.Sendb(send)
			}
		}
	}

	// TODO: maybe need add send log?

	c.JSON(http.StatusOK, gin.H{
		"code": returnCode,
		"msg":  returnMessage,
		"data": nil,
	})
}

func IM(db *sql.DB, c *gin.Context, fns *IMBot, imIdx int, jsonData string) {
	returnCode := ErrCodeCommonSuccessed
	returnMessage := ErrMessageCommonSuccessed

	if fn, ok := fns.IMBotHandle[imIdx]; ok {
		err := fn(db, jsonData)
		if err != nil {
			returnCode = ErrCodeIMHandleError
			returnMessage = err.Error()
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": returnCode,
		"msg":  returnMessage,
		"data": nil,
	})
}
