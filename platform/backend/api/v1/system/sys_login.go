package system

import (
	"backend/api/v1/model"
	"backend/internal/api_cluster"
	"backend/internal/ginweb"
	"backend/internal/ginweb/middleware"
	"backend/pkg/encrypt"
	"backend/pkg/jwt"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

type SystemLoginApi struct {
	BasePath        string
	checkHeader     bool
	accManageHeader string
	accNormalHeader string
}

func NewSystemLoginApi(basePath string, checkHeader bool, accManageHeader, accNormalHeader string) api_cluster.IApiEach {
	return &SystemLoginApi{
		BasePath:        basePath,
		checkHeader:     checkHeader,
		accManageHeader: accManageHeader,
		accNormalHeader: accNormalHeader,
	}
}

func (p *SystemLoginApi) GetGroupPath() string {
	return p.BasePath
}

func (p *SystemLoginApi) RegisterApiRouter(g *gin.RouterGroup, ginHandler *ginweb.GinHanlder) {
	// g.POST("/authtoken", ginHandler.Handle(p.GetAuth))
	db := ginHandler.DB.GetDefaultDB()
	logger := ginHandler.Logger

	g.POST("/login", middleware.LoginLogger(db, logger), ginHandler.Handle(p.Login))
	g.POST("/captcha", ginHandler.Handle(p.Captcha))
	g.POST("/autologin", ginHandler.Handle(p.AutoLogin))
}

// @Tags Login
// @Summary 取得驗證碼
// @accept application/json
// @Produce  application/json
// @Success 200 {object} response.Response{data=model.CaptchaResponse,msg=string} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /api/v1/login/captcha [post]
func (p *SystemLoginApi) Captcha(c *ginweb.Context) {
	id, bs64, captchaValue, expiredTime, err := c.GetCaptcha().GenerateCaptcha()
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_LOCAL, "", nil)
		return
	}

	c.Ok(&model.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       bs64,
		CaptchaLength: len(captchaValue),
		CaptchaValue:  captchaValue,
		ExpiredTime:   expiredTime,
	})
}

// @Tags Login
// @Summary 用戶登錄
// @Produce  application/json
// @Param data body model.Login true "用戶名, 密碼, 驗證碼"
// @Success 200 {object} response.Response{data=model.LoginResponse,msg=string} "返回包括用戶信息,token,過期時間"
// @Router /api/v1/login/login [post]
func (p *SystemLoginApi) Login(c *ginweb.Context) {
	var l model.Login
	_ = c.ShouldBindJSON(&l)
	c.Set("username", l.Username)

	// 檢查參數是否合法
	if !l.CheckParams() {
		c.OkWithCode(definition.ERROR_CODE_ERROR_REQUEST_DATA)
		return
	}

	// 檢查驗證碼是否過期
	if c.GetCaptcha().CheckExpired(l.CaptchaId) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CAPTCHA_EXPIRED)
		return
	}
	// 檢查驗證碼
	isVerify := c.GetCaptcha().VerifyCaptcha(l.CaptchaId, l.Captcha)
	if !isVerify {
		c.OkWithCode(definition.ERROR_CODE_ERROR_CAPTCHA)
		return
	}

	l.Username = strings.ToLower(l.Username)

	// check user relogin
	if isHad := c.GetJwt().CheckUsernameExist(l.Username); isHad {
		// 如果前面有登入，直接把前面的 token 拉黑
		c.GetJwt().AddBlackTokenByUsername(l.Username)
	}

	loginTime := time.Now().UTC()

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "username", "password", "nickname", "google_auth",
			"google_key", "allow_ip", "account_type", "is_readonly", "is_enabled",
			"permission", "is_added", "role").
		From("view_admin_user").
		Where(sq.Eq{"username": l.Username}).
		ToSql()

	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	var user model.LoginAdminUser
	err = c.DB().QueryRow(query, args...).
		Scan(&user.AgentId, &user.Username, &user.Password, &user.Nickname, &user.GoogleAuth,
			&user.GoogleKey, &user.AllowIp, &user.AccountType, &user.IsReadonly, &user.IsEnabled,
			&user.PermissionBytes, &user.IsAdded, &user.PermissionId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
			return
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
			return
		}
	}

	// check admin user custome header
	if p.checkHeader {
		customerHeader := c.Request.Header.Get("X-Dcc-Header")
		if customerHeader != "" {
			myHeader := ""
			if user.AccountType == definition.ACCOUNT_TYPE_ADMIN {
				myHeader = p.accManageHeader
			} else {
				myHeader = p.accNormalHeader
			}

			if myHeader != customerHeader {
				c.OkWithCode(definition.ERROR_CODE_ERROR_ADMIN_USER_TYPE_ILLEGAL)
				return
			}
		} else {
			c.OkWithCode(definition.ERROR_CODE_ERROR_ADMIN_USER_TYPE_ILLEGAL)
			return
		}
	}

	utils.ToStruct(user.PermissionBytes, &user.Permission)
	if user.IsEnabled == definition.ACCOUNT_STATUS_DISABLE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_DISABLED)
		return
	}

	// 檢查代理是否存在
	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("is_enabled", "secret_key", "level_code", "cooperation", "ip_whitelist",
			"currency", "jackpot_start_time", "jackpot_end_time", "wallet_type").
		From("agent").
		Where(sq.Eq{"id": user.AgentId}).
		ToSql()
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	isEnable := 0
	secretKey := ""
	levelCode := ""
	cooperation := 0
	ipWhitelistBytes := make([]byte, 0)
	currency := ""
	jackpotStartTime := time.Time{}
	jackpotEndTime := time.Time{}
	walletType := 0

	err = c.DB().QueryRow(query, args...).Scan(&isEnable, &secretKey, &levelCode, &cooperation, &ipWhitelistBytes,
		&currency, &jackpotStartTime, &jackpotEndTime, &walletType)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_DATABASE)
		return
	}

	if isEnable == definition.ACCOUNT_STATUS_DISABLE {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT_DISABLED)
		return
	}

	if secretKey == "" {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	// 驗證密碼
	realPwd, err := encrypt.DecryptSaltToken(user.Password, secretKey)
	if err != nil {
		c.OkWithCode(definition.ERROR_CODE_ERROR_LOCAL)
		return
	}

	if realPwd != l.Password {
		c.OkWithCode(definition.ERROR_CODE_ERROR_ACCOUNT)
		return
	}

	// 檢查ip
	clientIp := c.ClientIP()
	ipWhitelist := global.UdfDBAgentIpWhitelistResultToAgentWhitelist(ipWhitelistBytes)
	if !middleware.CheckAgentIpWhitelist(clientIp, ipWhitelist) {
		c.OkWithCode(definition.ERROR_CODE_ERROR_AGENT_IP_WHITELIST)
		return
	}

	// 總代理及子代理檢查是否正在jackpot期間，有的話要加入jackpot紀錄的權限
	if user.AccountType == definition.ACCOUNT_TYPE_GENERAL || user.AccountType == definition.ACCOUNT_TYPE_NORMAL {
		now := time.Now().UTC()
		if !jackpotStartTime.After(now) && !jackpotEndTime.Before(now) {
			user.Permission.List = append(user.Permission.List, definition.PERMISSION_LIST_FEATURE_CODE_JACKPOT_GET_JACKPOT_LIST)
		}
	}

	claims := c.GetJwt().CreateClaims(jwt.BaseClaims{
		ID:           uint(user.AgentId),
		LevelCode:    levelCode,
		Username:     user.Username,
		Password:     user.Password,
		Nickname:     user.Nickname,
		AccountType:  user.AccountType,
		Cooperation:  cooperation,
		Currency:     currency,
		PermissionId: user.PermissionId,
		IsAdded:      user.IsAdded,
		LoginTime:    loginTime.String(),
		Ip:           c.ClientIP(), // token 綁IP
		WalletType:   walletType,
	})

	token, _ := c.GetJwt().GenerateToken(claims)

	var returnValue = &model.LoginResponse{
		UserData: model.AdminUserResponse{
			AgentId:   user.AgentId,
			LevelCode: levelCode,
			Username:  user.Username,
			// Password: user.Password,
			Nickname:         user.Nickname,
			Permission:       user.Permission.List,
			AccountType:      (len(levelCode) / 4),
			IsAdded:          user.IsAdded,
			Cooperation:      cooperation,
			Currency:         currency,
			JackpotStartTime: jackpotStartTime,
			JackpotEndTime:   jackpotEndTime,
			WalletType:       walletType,
		},
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix() * 1000,
	}

	// update admin_user login time
	query = `UPDATE "public"."admin_user"
				SET "login_time" = $1
				WHERE "agent_id" = $2 AND "username" = $3`
	result, _ := c.DB().Exec(query, loginTime, user.AgentId, user.Username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("query exec failed,query = %v, err = %v", query, err)
	}

	c.Ok(returnValue)
}

// @Tags Login
// @Summary 自動登錄(測試用, login role is admin)
// @Produce  application/json
// @Success 200 {object} response.Response{data=model.LoginResponse,msg=string} "返回包括用戶信息,token,過期時間"
// @Router /api/v1/login/autologin [post]
func (p *SystemLoginApi) AutoLogin(c *ginweb.Context) {
	l := model.Login{
		Username: "admin",
		Password: "12345678",
	}
	// _ = c.ShouldBindJSON(&l)

	// // 檢查參數是否合法
	// if !l.CheckParams() {
	// 	c.Result(definition.ERROR_CODE_ERROR_REQUEST_DATA, "", nil)
	// 	return
	// }

	// 檢查驗證碼
	// isVerify := c.GetCaptcha().VerifyCaptcha(l.CaptchaId, l.Captcha)
	// if !isVerify {
	// 	c.Result(definition.ERROR_CODE_ERROR_CAPTCHA, "", nil)
	// 	return
	// }

	// check user relogin
	if isHad := c.GetJwt().CheckUsernameExist(l.Username); isHad {
		// 如果前面有登入，直接把前面的 token 拉黑
		c.GetJwt().AddBlackTokenByUsername(l.Username)
	}

	loginTime := time.Now().UTC()

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "username", "password", "nickname", "google_auth",
			"google_key", "allow_ip", "account_type", "is_readonly", "is_enabled",
			"permission", "is_added", "role").
		From("view_admin_user").
		Where(sq.Eq{"username": l.Username}).
		ToSql()
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_ACCOUNT, "", nil)
		return
	}

	var user model.LoginAdminUser
	err = c.DB().QueryRow(query, args...).
		Scan(&user.AgentId, &user.Username, &user.Password, &user.Nickname, &user.GoogleAuth,
			&user.GoogleKey, &user.AllowIp, &user.AccountType, &user.IsReadonly, &user.IsEnabled,
			&user.PermissionBytes, &user.IsAdded, &user.PermissionId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Result(definition.ERROR_CODE_ERROR_ACCOUNT, "", nil)
			return
		} else {
			c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
			return
		}
	}

	utils.ToStruct(user.PermissionBytes, &user.Permission)
	if user.IsEnabled == definition.ACCOUNT_STATUS_DISABLE {
		c.Result(definition.ERROR_CODE_ERROR_ACCOUNT_DISABLED, "", nil)
		return
	}

	// 檢查代理是否存在
	query, args, err = sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("is_enabled", "secret_key", "level_code", "cooperation", "currency").
		From("agent").
		Where(sq.Eq{"id": user.AgentId}).
		ToSql()
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	isEnable := 0
	secretKey := ""
	levelCode := ""
	cooperation := 0
	currency := ""
	err = c.DB().QueryRow(query, args...).Scan(&isEnable, &secretKey, &levelCode, &cooperation, &currency)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_DATABASE, "", nil)
		return
	}

	if isEnable == definition.ACCOUNT_STATUS_DISABLE {
		c.Result(definition.ERROR_CODE_ERROR_ACCOUNT_DISABLED, "", nil)
		return
	}

	if secretKey == "" {
		c.Result(definition.ERROR_CODE_ERROR_LOCAL, "", nil)
		return
	}

	// 驗證密碼
	realPwd, err := encrypt.DecryptSaltToken(user.Password, secretKey)
	if err != nil {
		c.Result(definition.ERROR_CODE_ERROR_LOCAL, "", nil)
		return
	}

	if realPwd != l.Password {
		c.Result(definition.ERROR_CODE_ERROR_ACCOUNT, "", nil)
		return
	}

	jackpotStartTime := time.Time{}
	jackpotEndTime := time.Time{}
	walletType := 0

	claims := c.GetJwt().CreateClaims(jwt.BaseClaims{
		ID:           uint(user.AgentId),
		LevelCode:    levelCode,
		Username:     user.Username,
		Password:     user.Password,
		Nickname:     user.Nickname,
		AccountType:  user.AccountType,
		Cooperation:  cooperation,
		Currency:     currency,
		PermissionId: user.PermissionId,
		IsAdded:      user.IsAdded,
		LoginTime:    loginTime.String(),
		Ip:           c.ClientIP(), // token 綁IP
		WalletType:   walletType,
	})

	token, _ := c.GetJwt().GenerateToken(claims)

	var returnValue = &model.LoginResponse{
		UserData: model.AdminUserResponse{
			AgentId:   user.AgentId,
			LevelCode: levelCode,
			Username:  user.Username,
			// Password: user.Password,
			Nickname:         user.Nickname,
			Permission:       user.Permission.List,
			AccountType:      (len(levelCode) / 4),
			IsAdded:          user.IsAdded,
			Cooperation:      cooperation,
			Currency:         currency,
			JackpotStartTime: jackpotStartTime,
			JackpotEndTime:   jackpotEndTime,
			WalletType:       walletType,
		},
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix() * 1000,
	}

	// update admin_user login time
	query = `UPDATE "public"."admin_user"
				SET "login_time" = $1
				WHERE "agent_id" = $2 AND "username" = $3`
	result, _ := c.DB().Exec(query, loginTime, user.AgentId, user.Username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("query exec failed,query = %v, err = %v", query, err)
	}

	c.Ok(returnValue)
}
