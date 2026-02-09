package middleware

import (
	"backend/internal/ginweb"
	"backend/internal/ginweb/response"
	"backend/pkg/encrypt/base64url"
	"backend/pkg/jwt"
	"backend/pkg/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		startTime := time.Now().UTC().UnixMilli()

		//处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()
		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			response := response.Response{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				responseCode = response.Code
				responseMsg = response.Message
				responseData = response.Data
			}
		}

		//结束时间
		endTime := time.Now().UTC().UnixMilli()

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		//日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		// accessLogMap["request_client_ip"] = c.ClientIP()
		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", endTime-startTime)

		accessLogJson := utils.ToJSON(accessLogMap)

		// 失敗才紀錄本地端LOG
		if responseCode != 0 {
			log.Printf("%v", accessLogJson)
		}

		data, isHave := c.Get("claims")
		var username string
		if isHave && data != nil {
			claims := data.(*jwt.CustomClaims)
			byteData, _ := base64url.Decode(claims.Username)
			username = string(byteData)
		}

		if username != "" {
			aual := &AdminUserActionLog{
				Username:  username,
				ErrorCode: responseCode,
				ActionLog: accessLogJson,
				Ip:        c.ClientIP(),
				LogTime:   utils.Get18UnsignedTimeNowUTC(),
			}

			_, err := db.Exec(`INSERT INTO public.admin_user_action_log(
			log_time, username, error_code, action_log, ip, method, request_url)
			VALUES ($1, $2, $3, $4, $5, $6, $7);`,
				aual.LogTime, aual.Username, aual.ErrorCode, aual.ActionLog, aual.Ip, c.Request.Method, c.Request.RequestURI)
			if err != nil {
				log.Printf("Record admin_user_action_log failed, data is %v", aual)
			}
		}
	}
}

func BackendLogger(db *sql.DB, logger ginweb.ILogger, systemPermissionData *sync.Map) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		// 開始時間
		startTime := time.Now().UTC()

		// 處理請求
		c.Next()

		// 结束时间
		endTime := time.Now().UTC()

		responseBody := bodyLogWriter.body.String()
		response := response.Response{}
		if responseBody != "" {
			json.Unmarshal([]byte(responseBody), &response)
		}

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		// http日誌格式
		httpLogMap := make(map[string]interface{})

		httpLogMap["request_time"] = startTime
		httpLogMap["request_method"] = c.Request.Method
		httpLogMap["request_uri"] = c.Request.RequestURI
		httpLogMap["request_proto"] = c.Request.Proto
		httpLogMap["request_ua"] = c.Request.UserAgent()
		httpLogMap["request_referer"] = c.Request.Referer()
		httpLogMap["request_post_data"] = c.Request.PostForm.Encode()
		httpLogMap["response_time"] = endTime
		httpLogMap["response_code"] = response.Code
		httpLogMap["response_msg"] = response.Message
		httpLogMap["response_data"] = response.Data

		httpLogMap["cost_time"] = fmt.Sprintf("%vms", endTime.UnixMilli()-startTime.UnixMilli())

		httpLogJson := utils.ToJSON(httpLogMap)

		permissionVal, ok := systemPermissionData.Load(c.Request.URL.Path)
		if !ok {
			logger.Error("BackendLogger api path:%s undefinded.", c.Request.URL.Path)
			return
		}

		permission := permissionVal.(map[string]interface{})
		featureCode := int(permission["feature_code"].(float64))
		actionType := int(permission["action_type"].(float64))

		// 操作日誌
		actionLog, isHave := c.Get("action_log")
		actionLogJson := "{}"
		if isHave {
			actionLogJson = utils.ToJSON(actionLog)
		}

		// 操作者
		data, _ := c.Get("claims")
		claims := data.(*jwt.CustomClaims)
		byteData, _ := base64url.Decode(claims.Username)
		username := string(byteData)

		aubal := &AdminUserBackendActionLog{
			AgentId:     int(claims.BaseClaims.ID),
			Username:    username,
			ErrorCode:   response.Code,
			Ip:          c.ClientIP(),
			FeatureCode: featureCode,
			ActionType:  actionType,
			HttpLog:     httpLogJson,
			ActionLog:   actionLogJson,
			CreateTime:  startTime,
		}

		_, err := db.Exec(`INSERT INTO public.admin_user_backend_action_log(
			agent_id, username, feature_code, action_type, error_code,
			action_log, http_log, ip, method, request_url, 
			create_time)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`,
			aubal.AgentId, aubal.Username, aubal.FeatureCode, aubal.ActionType, aubal.ErrorCode,
			aubal.ActionLog, aubal.HttpLog, aubal.Ip, c.Request.Method, c.Request.RequestURI,
			aubal.CreateTime)
		if err != nil {
			logger.Error("Record admin_user_backend_action_log failed, data is %v, err is %v", aubal, err)
		}
	}
}

func LoginLogger(db *sql.DB, logger ginweb.ILogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		c.Next()

		userName := c.GetString("username")
		var agentId int = -1

		query := `SELECT agent_id FROM admin_user WHERE username = $1`
		db.QueryRow(query, userName).Scan(&agentId)

		responseBody := bodyLogWriter.body.String()
		response := response.Response{}
		if responseBody != "" {
			json.Unmarshal([]byte(responseBody), &response)
		}

		loginUser := &AdminUserLoginLog{
			AgentId:   agentId,
			Username:  userName,
			Ip:        c.ClientIP(),
			ErrorCode: response.Code,
			LogTime:   time.Now().UTC(),
		}

		_, err := db.Exec(`INSERT INTO backend_login_log (login_time, agent_id, username, ip, error_code)
				VALUES($1, $2, $3, $4, $5);`, loginUser.LogTime, loginUser.AgentId, loginUser.Username, loginUser.Ip, loginUser.ErrorCode)
		if err != nil {
			logger.Error("Record backend_login_log is failed admin_user is:%v", loginUser)
		}

	}
}
