package global

import (
	"backend/pkg/encrypt"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
	"time"
)

func UdfCreateNewAdminUser(db *sql.DB, agentId int, username, password, nickname string, accountType,
	isReadonly int, role, info string) (adminUser *table_model.AdminUser, err error) {
	/*
		CREATE FUNCTION "public"."udf_create_admin_user"("_agent_id" int4, "_username" varchar,
		    "_password" varchar, "_nickname" varchar, "_account_type" int4, "_is_readonly" int4,
		    "_is_added" bool, "_role" uuid, "_info" varchar)
		    RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 創建代理後台帳號
	// 2. 回傳 json 格式，包含創建的adminUser資料

	jsonResult := ""
	query := `SELECT "public"."udf_create_admin_user"($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	err = db.QueryRow(query, agentId, username, password, nickname, accountType,
		isReadonly, true, role, info).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))

	adminUser = UdfCreateAdminUserReturningDataToAdminUser(result)

	return
}

func UdfCreateAdminUserReturningDataToAdminUser(data map[string]interface{}) *table_model.AdminUser {
	return &table_model.AdminUser{
		AgentId:      int(data["agent_id"].(float64)),
		Username:     data["username"].(string),
		Password:     data["password"].(string),
		Nickname:     data["nickname"].(string),
		AccountType:  int(data["account_type"].(float64)),
		IsReadonly:   int(data["is_readonly"].(float64)),
		IsEnabled:    int(data["is_enabled"].(float64)),
		UpdateTime:   utils.MicroSecondsToUtcTime(int64(data["update_time"].(float64))),
		CreateTime:   utils.MicroSecondsToUtcTime(int64(data["create_time"].(float64))),
		IsAdded:      data["is_added"].(bool),
		LoginTime:    utils.MicroSecondsToUtcTime(int64(data["login_time"].(float64))),
		PermissionId: data["role"].(string),
		Info:         data["info"].(string),
	}
}

func UdfUpdateAdminUser(db *sql.DB, agentId int, username, role, info string,
	isEnabled int) (updateTime time.Time, err error) {
	/*
		CREATE FUNCTION "public"."udf_update_admin_user"("_agent_id" int4, "_username" varchar,
			"_role" uuid, "_info" varchar, "_is_enabled")
			RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 修改代理
	// 2. 修改代理後台帳號
	// 3. 回傳 json 格式，目前僅提供更新時間

	jsonResult := ""
	query := `SELECT "public"."udf_update_admin_user"($1, $2, $3, $4, $5)`

	err = db.QueryRow(query, agentId, username, role, info, isEnabled).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))

	usec := int64(result["update_time"].(float64))
	updateTime = utils.MicroSecondsToUtcTime(usec)

	return
}

func CreateNewAdminUser(db *sql.DB, agentId, accountType, isReadonly, topAgentId int, role, username, password, nickname, secretKey, info string, isAdded bool) error {

	user := table_model.NewEmptyAdminUser()
	user.Username = username
	user.AgentId = agentId
	user.AccountType = accountType
	user.IsReadonly = isReadonly
	user.Nickname = nickname

	// 將密碼做加密
	decryptPwd, err := encrypt.EncryptSaltToken(password, secretKey)
	if err != nil {
		return err
	}

	user.Password = decryptPwd

	query := `
		INSERT INTO "public"."admin_user"(
			"agent_id", "username", "password", "nickname", "account_type",
			"is_readonly", "role", "info")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`

	_, err = db.Exec(query,
		user.AgentId, user.Username, decryptPwd, user.Nickname, user.AccountType,
		user.IsReadonly, role, info)
	if err != nil {
		return err
	}

	return nil
}

func CheckAdminUserIsExist(db *sql.DB, username string) bool {
	query := `SELECT username FROM admin_user WHERE username = $1`

	queryUsername := ""

	err := db.QueryRow(query, username).Scan(&queryUsername)
	if err != nil {
		return err != sql.ErrNoRows
	}
	// 帳號存在
	return true
}
