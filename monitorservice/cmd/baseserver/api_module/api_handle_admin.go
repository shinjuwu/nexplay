package api_module

import (
	"encoding/json"
	"fmt"
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/cmd/baseserver/api_module/module_lib"
	"monitorservice/internal/api"
	"monitorservice/pkg/encrypt"
	"monitorservice/pkg/utils"
	"time"

	md5 "monitorservice/pkg/encrypt/md5hash"

	sq "github.com/Masterminds/squirrel"
)

// @Tags admin.api
// @Summary 註冊(不開放外部注冊，此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.UsersRegisterRequest true "註冊帳號參數，密碼部分轉成 base64 再送出"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/admin.api/v1/register [post]
func Register(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	var req model.UsersRegisterRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if !utils.LowercaseEnglishAndNumber4To16.MatchString(req.Username) ||
		!utils.EnglishAndNumber8To16.MatchString(req.Password) {
		return "", fmt.Errorf("username or password format check failed")
	}

	// 權限檢查，比對檢查claim 裡面的權限就可以
	// 避免重複與不可超出admin 的權限
	fliterPermissions := make([]string, 0)
	permissionMap := make(map[string]int)
	for i := 0; i < len(claims.Permissions); i++ {
		permissionMap[claims.Permissions[i]] = 1
	}

	for i := 0; i < len(req.Permissions); i++ {
		permissionMap[req.Permissions[i]] += 1
	}

	for permissionKey, count := range permissionMap {
		if count > 1 {
			fliterPermissions = append(fliterPermissions, permissionKey)
		}
	}

	userID := ""
	timeNowStr := time.Now().String()
	topCode := md5.Hash8bit(req.Username + timeNowStr)

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

	err = c.DB().QueryRow(`INSERT INTO public.users(
		"username", "password", "nickname", "is_enabled", "top_code",
		"permission")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`,
		req.Username, req.Password, req.Nickname, true, topCode,
		utils.ToJSON(fliterPermissions)).Scan(&userID)
	if err != nil || userID == "" {
		return "", fmt.Errorf("Register() failed, err is : %v", err)
	}

	return "", nil
}

// @Tags admin.api
// @Summary 封鎖指定用戶(此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.BlockUsersRequest true "指定用戶資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/admin.api/v1/blockusers [post]
func BlockUsers(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	var req model.BlockUsersRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	query := `UPDATE "public"."users"
		SET "is_enabled" = $1
		WHERE "id" = $2;`
	result, _ := c.DB().Exec(query, req.IsEnabled, req.ID)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("BlockUsers() query exec failed,query = %v, err = %v", query, err)
		return "", fmt.Errorf("update rowsAffectedCount is 0")
	}

	// 原 token 失效
	module_lib.LibBlockUsers(c.JWT(), req.Username)

	return "", nil
}

// @Tags admin.api
// @Summary 修改帳戶資料(此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ModifyUsersInfoRequest true "修改帳戶資料"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/admin.api/v1/modifyusersinfo [post]
func ModifyUsersInfo(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	var req model.ModifyUsersInfoRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	query := `UPDATE "public"."users"
		SET "nickname" = $1, "is_enabled" = $2, "info" = $3, "permission" = $4
		WHERE "username" = $5;`
	result, _ := c.DB().Exec(query,
		req.Nickname, req.IsEnabled, req.Info, utils.ToJSON(req.Permissions), req.Username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("ModifyUsersInfo() query exec failed,query = %v, err = %v", query, err)
		return "", fmt.Errorf("update rowsAffectedCount is 0")
	}

	return "", nil
}

// @Tags admin.api
// @Summary 修改密碼(此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.ModifyUsersPasswordRequest true "修改密碼，修改後原 token 失效，必須重新登入"
// @Success 200 {object} model.Response{data=string,msg=string,code=int} ""
// @Router /monitor/admin.api/v1/modifyuserspassword [post]
func ModifyUsersPassword(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	var req model.ModifyUsersPasswordRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	if req.Username == "" || req.Password == "" {
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
	result, _ := c.DB().Exec(query, req.Password, req.Username)
	if count, err := result.RowsAffected(); count != 1 {
		c.Logger().Printf("ModifyUsersInfo() query exec failed,query = %v, err = %v", query, err)
		return "", fmt.Errorf("update rowsAffectedCount is 0")
	}

	// 原 token 失效，必須重新登入
	module_lib.LibBlockUsers(c.JWT(), req.Username)

	return "", nil
}

// @Tags admin.api
// @Summary 取所有用戶的帳戶資料(此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @Success 200 {object} model.Response{data=model.GetUserInfoListResponse,msg=string,code=int} "回傳用戶資料"
// @Router /monitor/admin.api/v1/getusersinfolist [post]
func GetUsersInfoList(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("username", "nickname", "user_metadata", "is_enabled", "last_login_time",
			"info", "create_time", "permission").
		From("users").
		ToSql()

	if err != nil {
		return "", err
	}

	var res model.GetUserInfoListResponse
	res.Data = make([]*model.GetUserInfoResponse, 0)

	rows, err := c.DB().Query(query, args...)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		user := new(model.GetUserInfoResponse)
		var permissionBytes []byte
		if err := rows.Scan(&user.Username, &user.Nickname, &user.UserMetadata, &user.IsEnabled, &user.LastLoginTime,
			&user.Info, &user.CreateTime, &permissionBytes); err != nil {
			return "", err
		}

		if err = json.Unmarshal(permissionBytes, &user.Permissions); err != nil {
			return "", err
		}

		res.Data = append(res.Data, user)
	}

	return utils.ToJSON(res), nil
}

// @Tags admin.api
// @Summary 取指定用戶的帳戶資料(此API只有admin可用)
// @accept application/json
// @Produce  application/json
// @Security BearerAuth
// @param data body model.GetUserInfoRequest true "輸入用戶帳號"
// @Success 200 {object} model.Response{data=model.GetUserInfoResponse,msg=string,code=int} "回傳用戶資料"
// @Router /monitor/admin.api/v1/getusersinfo [post]
func GetUsersInfo(c api.IContext, params string) (string, error) {

	claims := c.Claims()
	if claims == nil {
		return "", fmt.Errorf("jwt verify failed")
	}

	if !claims.IsAdmin {
		return "", fmt.Errorf("insufficient permission")
	}

	var req model.GetUserInfoRequest
	if err := json.Unmarshal([]byte(params), &req); err != nil {
		return "", err
	}

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("username", "nickname", "user_metadata", "is_enabled", "last_login_time",
			"info", "create_time", "permission").
		From("users").
		Where(sq.Eq{"username": req.Username}).
		ToSql()

	if err != nil {
		return "", err
	}

	res := new(model.GetUserInfoResponse)
	var permissionBytes []byte
	err = c.DB().QueryRow(query, args...).Scan(&res.Username, &res.Nickname, &res.UserMetadata, &res.IsEnabled, &res.LastLoginTime,
		&res.Info, &res.CreateTime, &permissionBytes)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(permissionBytes, &res.Permissions); err != nil {
		return "", err
	}

	return utils.ToJSON(res), nil
}
