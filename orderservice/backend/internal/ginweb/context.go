package ginweb

import (
	"backend/internal/ginweb/response"
	"backend/pkg/captcha"
	"backend/pkg/database"
	"backend/pkg/encrypt/base64url"
	"backend/pkg/job"
	"backend/pkg/jwt"
	"backend/pkg/logger"
	"backend/pkg/redis"
	"database/sql"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Context struct {
	logger       ILogger
	idb          database.IDatabase // DB 物件
	redis        redis.IRedisCliect // redis DB 物件
	Ip           string             // 用戶 ip
	Jwt          *jwt.JwtManager
	captcha      captcha.ICaptcha
	job          *job.JobScheduler
	claims       *jwt.CustomClaims
	isDebugMode  bool
	*gin.Context // gin object to be executed every time a api call is executed
}

func NewContext(_logger *logger.RuntimeGoLogger, _c *gin.Context, _idb database.IDatabase, _rdb redis.IRedisCliect, _jwt *jwt.JwtManager,
	_captcha captcha.ICaptcha, _jobScheduler *job.JobScheduler, _ip string, _isDebugMode bool) *Context {

	claims := new(jwt.CustomClaims)

	tmp, isExits := _c.Get("claims")
	if isExits {
		claims = tmp.(*jwt.CustomClaims)
	}

	return &Context{
		logger:      _logger,
		Context:     _c,
		idb:         _idb,
		redis:       _rdb,
		Jwt:         _jwt,
		captcha:     _captcha,
		job:         _jobScheduler,
		Ip:          _ip,
		claims:      claims,
		isDebugMode: _isDebugMode,
	}
}

func (c *Context) Logger() ILogger {
	return c.logger
}

// get database.IDatabase object to find select db connect
func (c *Context) IDB() database.IDatabase {
	return c.idb
}

// get default db object
func (c *Context) DB() *sql.DB {
	return c.idb.GetDefaultDB()
}

// get db object by idx (idx define in config)
func (c *Context) GetDB(idx int) *sql.DB {
	return c.idb.GetDB(idx)
}

func (c *Context) Redis() redis.IRedisCliect {
	return c.redis
}

func (c *Context) GetJwt() *jwt.JwtManager {
	return c.Jwt
}

func (c *Context) GetCaptcha() captcha.ICaptcha {
	return c.captcha
}

func (c *Context) GetJob() *job.JobScheduler {
	return c.job
}

func (c *Context) GetIsDebugMode() bool {
	return c.isDebugMode
}

func (c *Context) GetUserClaims() *jwt.CustomClaims {
	return c.claims
}

func (c *Context) GetUsername() string {
	creatorBytes, err := base64url.Decode(c.claims.Username)
	if err != nil {
		return ""
	}
	return string(creatorBytes)
}

func (c *Context) GetNickname() string {
	creatorBytes, err := base64url.Decode(c.claims.Nickname)
	if err != nil {
		return ""
	}
	return string(creatorBytes)
}

// 客戶端傳來的當下時間
func (c *Context) CheckTimeout(timeStamp string) bool {
	unixTime, _ := strconv.ParseInt(timeStamp, 10, 64)
	expiredTime := time.Unix(unixTime, 0)

	return !(time.Now().UTC().Sub(expiredTime.UTC()).Seconds() < 30 && time.Now().UTC().Sub(expiredTime.UTC()).Seconds() > -30)
}

func (c *Context) HTMLSuccess(pageName string, data map[string]interface{}) {
	c.HTML(http.StatusOK, pageName, data)
}

func (c *Context) HTMLFailed(pageName string, data map[string]interface{}) {
	c.HTML(http.StatusBadRequest, pageName, data)
}

func (c *Context) Result(code int, msg string, data interface{}) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	c.logger.Info("RESULT %s %s%s, RawQuery: %s, body: %s, code: %d, msg: %s, data: %v",
		c.Context.Request.Method,
		c.Context.Request.Host,
		c.Context.Request.RequestURI,
		c.Context.Request.URL.RawQuery,
		jsonData,
		code, msg, data)
	response.Result(c.Context, code, msg, data)
}

func (c *Context) OkWithCode(code int) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	c.logger.Info("CODE %s %s%s, RawQuery: %s, body: %s, code is %v",
		c.Context.Request.Method,
		c.Context.Request.Host,
		c.Context.Request.RequestURI,
		c.Context.Request.URL.RawQuery,
		jsonData,
		code)
	response.Result(c.Context, code, "", nil)
}

func (c *Context) Ok(data interface{}) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	c.logger.Info("OK %s %s%s, RawQuery: %s, body: %s, response is %v",
		c.Context.Request.Method,
		c.Context.Request.Host,
		c.Context.Request.RequestURI,
		c.Context.Request.URL.RawQuery,
		jsonData,
		data)
	response.Ok(c.Context, data)
}

func (c *Context) Fail(data interface{}) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	c.logger.Info("FAIL %s %s%s, RawQuery: %s, body: %s, response is %v",
		c.Context.Request.Method,
		c.Context.Request.Host,
		c.Context.Request.RequestURI,
		c.Context.Request.URL.RawQuery,
		jsonData,
		data)
	response.Fail(c.Context, data)
}

func (c *Context) FailWithLogError(err error, data interface{}) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	c.logger.Error("ERROR %s %s%s, RawQuery: %s, body: %s, error is %s",
		c.Context.Request.Method,
		c.Context.Request.Host,
		c.Context.Request.RequestURI,
		c.Context.Request.URL.RawQuery,
		jsonData,
		err.Error())
	response.Fail(c.Context, data)
}
