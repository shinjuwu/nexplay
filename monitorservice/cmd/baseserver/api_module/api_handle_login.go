package api_module

import (
	"encoding/json"
	"fmt"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/internal/api"
	"monitorservice/pkg/encrypt"
	"monitorservice/pkg/jwt"
	"monitorservice/pkg/utils"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// @Tags login.api
// @Summary 取得驗證碼
// @accept application/json
// @Produce  application/json
// @Success 200 {object} model.Response{data=model.CaptchaResponse,msg=string,code=int} "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router /monitor/login.api/v1/captcha [post]
func UsersCaptcha(c api.IContext, params string) (string, error) {

	id, bs64, expiredTime, err := c.Captcha().GenerateCaptcha()
	if err != nil {
		return "", err
	}
	captchaValue := c.Captcha().GetCaptchaValue(id)

	res := &model.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       bs64,
		CaptchaLength: len(captchaValue),
		CaptchaValue:  captchaValue,
		ExpiredTime:   expiredTime,
	}

	return utils.ToJSON(res), nil
}

// @Tags login.api
// @Summary 登入
// @accept application/json
// @Produce  application/json
// @param data body model.UsersLoginRequest true "登入帳號參數"
// @Success 200 {object} model.Response{data=model.UsersLoginResponse,msg=string,code=int} "登入成功回傳帳號資料與驗證token"
// @Router /monitor/login.api/v1/login [post]
func UsersLogin(c api.IContext, params string) (string, error) {

	var req model.UsersLoginRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	// 檢查驗證碼是否過期
	if c.Captcha().CheckExpired(req.CaptchaId) {
		return "", fmt.Errorf("code is expired")
	}
	// 檢查驗證碼
	isVerify := c.Captcha().VerifyCaptcha(req.CaptchaId, req.Captcha)
	if !isVerify {
		return "", fmt.Errorf("verify failed")
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "username", "password", "nickname", "user_metadata",
			"is_enabled", "last_login_time", "info", "create_time", "update_time",
			"disable_time", "top_code", "is_admin", "permission").
		From("users").
		Where(sq.Eq{"username": req.Username}).
		ToSql()

	if err != nil {
		return "", err
	}

	var user model.Users
	var permissionBytes []byte
	err = c.DB().QueryRow(query, args...).
		Scan(&user.Id, &user.Username, &user.Password, &user.Nickname, &user.UserMetadata,
			&user.IsEnabled, &user.LastLoginTime, &user.Info, &user.CreateTime, &user.UpdateTime,
			&user.DisableTime, &user.TopCode, &user.IsAdmin, &permissionBytes)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(permissionBytes, &user.Permissions); err != nil {
		return "", err
	}

	if !user.IsEnabled {
		return "", fmt.Errorf("this account is be block, contact administrator")
	}

	/*
		密碼處理
		1. 前端使用 base64 加密後傳送至後端進行登入
		2. 後端 base64 解密後，使用aes解密後與 DB 內數值比對
	*/
	encodePWD := utils.EncodeBase64([]byte(req.Password))

	// 密碼加密後比對
	saltKey := user.TopCode + user.TopCode + user.TopCode + user.TopCode

	realPwd, err := encrypt.DecryptSaltToken(user.Password, saltKey)
	if err != nil {
		return "", err
	}

	if realPwd != encodePWD {
		return "", fmt.Errorf("password verify is failed")
	}

	// 成功後回傳JWT產生之TOKEN
	loginTime := time.Now().UTC()

	claims := c.JWT().CreateClaims(jwt.BaseClaims{
		ID:          user.Id,
		TopCode:     user.TopCode,
		Username:    user.Username,
		Password:    user.Password,
		Nickname:    user.Nickname,
		LoginTime:   loginTime.String(),
		IsAdmin:     user.IsAdmin,
		Permissions: user.Permissions,
	})

	token, _ := c.JWT().GenerateToken(claims)

	var res = &model.UsersLoginResponse{
		UserData: model.UserLoginData{
			TopCode:       user.TopCode,
			Username:      user.Username,
			Nickname:      user.Nickname,
			UserMetadata:  user.UserMetadata,
			CreateTime:    user.CreateTime,
			LastLoginTime: user.LastLoginTime,
			IsAdmin:       user.IsAdmin,
			Permissions:   user.Permissions,
		},
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix() * 1000,
	}

	// update users login time
	query = `UPDATE "public"."users"
				SET "last_login_time" = $1
				WHERE "id" = $2`
	result, _ := c.DB().Exec(query, loginTime, user.Id)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("query exec failed,query = %v, err = %v", query, err)
	}

	return utils.ToJSON(res), nil
}
