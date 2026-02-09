package api_module

import (
	"encoding/json"
	"fmt"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/cmd/baseserver/api_module/module_lib"
	"monitorservice/internal/api"
	"monitorservice/pkg/encrypt"
	"monitorservice/pkg/encrypt/base64url"
	"monitorservice/pkg/utils"
)

// @Tags account.api
// @Summary 取自己的帳戶資料
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} model.Response{data=model.UserInfoResponse,msg=string,code=int} "回傳用戶資料(在線才有資料)"
// @Router /monitor/account.api/v1/usersinfo [post]
func UsersInfo(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	var res = &model.UserInfoResponse{
		TopCode:     claims.TopCode,
		Username:    c.GetUsername(),
		Nickname:    c.GetNickname(),
		IsAdmin:     claims.IsAdmin,
		LoginTime:   claims.LoginTime,
		ExpiresAt:   claims.ExpiresAt.Unix() * 1000,
		Permissions: claims.Permissions,
	}

	return utils.ToJSON(res), nil
}

// @Tags account.api
// @Summary 驗證目前登入 token 是否有效
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} model.Response{data=string,msg=string,code=int} "回傳用戶資料(在線才有資料)"
// @Router /monitor/account.api/v1/verifytoken [post]
func VerifyToken(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	return "", nil
}

// @Tags account.api
// @Summary 修改帳戶資料
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ModifyInfoRequest true "修改帳戶資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/account.api/v1/modifyinfo [post]
func ModifyInfo(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	usernameByte, _ := base64url.Decode(claims.BaseClaims.Username)
	username := string(usernameByte)

	var req model.ModifyInfoRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	query := `UPDATE "public"."users"
		SET "nickname" = $1
		WHERE "username" = $2;`
	result, _ := c.DB().Exec(query,
		req.Nickname, username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("ModifyInfo() query exec failed,query = %v, err = %v", query, err)
		return "", fmt.Errorf("update rowsAffectedCount is 0")
	}

	return "", nil
}

// @Tags account.api
// @Summary 修改密碼
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ModifyPasswordRequest true "修改密碼，修改後原 token 失效，必須重新登入"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/account.api/v1/modifypassword [post]
func ModifyPassword(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	usernameByte, _ := base64url.Decode(claims.BaseClaims.Username)
	username := string(usernameByte)

	var req model.ModifyPasswordRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if req.Password == "" {
		return "", fmt.Errorf("param is empty")
	}

	topCode := claims.BaseClaims.TopCode

	/*
		密碼處理
		1. 前端使用 base64 加密後傳送至後端進行註冊
		2. 後端 base64 解密後，使用 aes 加密後存入 DB
	*/
	encodePWD := utils.EncodeBase64([]byte(req.Password))

	// 密碼加密
	saltKey := topCode + topCode + topCode + topCode
	var err error
	req.Password, err = encrypt.EncryptSaltToken(encodePWD, saltKey)
	if err != nil {
		return "", err
	}

	query := `UPDATE "public"."users"
		SET "password" = $1
		WHERE "username" = $2;`
	result, _ := c.DB().Exec(query, req.Password, username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("ModifyPassword() query exec failed,query = %v, err = %v", query, err)
		return "", fmt.Errorf("update rowsAffectedCount is 0")
	}

	// 原 token 失效，必須重新登入
	module_lib.LibBlockUsers(c.JWT(), username)

	return "", nil
}
