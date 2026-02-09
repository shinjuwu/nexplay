package global

import (
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func UdfCreateNewAgent(db *sql.DB, agentName, agentCode, agentLevelCode, agentInfo, agentSecretKey,
	agentAesKey, agentMd5Key, agentCurrency string, agentIpWhitelist []*table_model.AgentIPWhitelistObj, agentCreator string,
	agentCommission, agentCooperation, agentTopAgentId int, agentIsTopAgent bool, agentWalletType int,
	agentWalletConninfo *table_model.WalletConnInfo, agentLobbySwitchInfo int, adminUserUsername, adminUserPassword, adminUserNickname, adminUserRole,
	adminUserInfo string, adminUserAccountType, adminUserIsReadonly int) (agent *table_model.Agent, adminUser *table_model.AdminUser, agentGames []*table_model.AgentGame, agentGameRooms []*table_model.AgentGameRoom, err error) {
	/*
		CREATE FUNCTION "public"."udf_create_agent"("_agent_name" varchar, "_agent_code" varchar,
			"_agent_level_code" varchar, "_agent_info" varchar, "_agent_secret_key" varchar,
			"_agent_aes_key" varchar, "_agent_md5_key" varchar, "_agent_currency" varchar,
			"_agent_ip_whitelist" varchar, "_agent_creator" varchar, "_agent_commission" int4,
			"_agent_cooperation" int4, "_agent_top_agent_id" int4, "_agent_is_top_agent" bool,
			"_admin_user_username" varchar,	"_admin_user_password" varchar, "_admin_user_nickname" varchar,
			"_admin_user_role" uuid, "_admin_user_info" varchar, "_admin_user_account_type" int4,
			"_admin_user_is_readonly" int4)
			RETURNS json AS $$
	*/
	// 此 sql function 內含
	// 1. 創建代理
	// 2. 創建代理後台帳號
	// 3. 回傳 json 格式，包含創建的agent、adminUser、agentGames及agentGameRooms資料

	ipWhitelistStr := utils.ToJSON(agentIpWhitelist)
	walletConninfo := utils.ToJSON(agentWalletConninfo)
	jsonResult := ""
	query := `SELECT "public"."udf_create_agent"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
		$11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)`

	err = db.QueryRow(query,
		agentName, agentCode, agentLevelCode, agentInfo, agentSecretKey,
		agentAesKey, agentMd5Key, agentCurrency, ipWhitelistStr, agentCreator,
		agentCommission, agentCooperation, agentTopAgentId, agentIsTopAgent, agentWalletType,
		walletConninfo, agentLobbySwitchInfo, adminUserUsername, adminUserPassword, adminUserNickname,
		adminUserRole, adminUserInfo, adminUserAccountType, adminUserIsReadonly).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))

	agent, adminUser, agentGames, agentGameRooms = UdfCreateAgentReturningDataToAgentAndAdminUser(result)

	return
}

func UdfCreateAgentReturningDataToAgentAndAdminUser(data map[string]interface{}) (agent *table_model.Agent, adminUser *table_model.AdminUser, agentGames []*table_model.AgentGame, agentGameRooms []*table_model.AgentGameRoom) {
	adminUserData := data["admin_user"].(map[string]interface{})
	adminUser = UdfCreateAdminUserReturningDataToAdminUser(adminUserData)

	agentData := data["agent"].(map[string]interface{})
	agent = &table_model.Agent{
		Id:               int(agentData["id"].(float64)),
		Name:             agentData["name"].(string),
		Code:             agentData["code"].(string),
		LevelCode:        agentData["level_code"].(string),
		SecretKey:        agentData["secret_key"].(string),
		AesKey:           agentData["aes_key"].(string),
		Md5Key:           agentData["md5_key"].(string),
		Commission:       int(agentData["commission"].(float64)),
		Info:             agentData["info"].(string),
		IsEnabled:        int(agentData["is_enabled"].(float64)),
		DisableTime:      utils.MicroSecondsToUtcTime(int64(agentData["disable_time"].(float64))),
		UpdateTime:       utils.MicroSecondsToUtcTime(int64(agentData["update_time"].(float64))),
		CreateTime:       utils.MicroSecondsToUtcTime(int64(agentData["create_time"].(float64))),
		IsTopAgent:       agentData["is_top_agent"].(bool),
		TopAgentId:       int(agentData["top_agent_id"].(float64)),
		Cooperation:      int(agentData["cooperation"].(float64)),
		CoinLimit:        agentData["coin_limit"].(float64),
		CoinUse:          agentData["coin_use"].(float64),
		Creator:          agentData["creator"].(string),
		IPWhitelist:      UdfDBAgentIpWhitelistResultToAgentWhitelist([]byte(utils.ToJSON(agentData["ip_whitelist"]))),
		AdminUsername:    adminUser.Username,
		Currency:         agentData["currency"].(string),
		WalletType:       int(agentData["wallet_type"].(float64)),
		WalletConnInfo:   UdfDBAgentIpWhitelistResultToAgentWalletConnInfo([]byte(utils.ToJSON(agentData["wallet_conninfo"]))),
		JackpotStatus:    int(agentData["jackpot_status"].(float64)),
		JackpotStartTime: utils.MicroSecondsToUtcTime(int64(agentData["jackpot_start_time"].(float64))),
		JackpotEndTime:   utils.MicroSecondsToUtcTime(int64(agentData["jackpot_end_time"].(float64))),
		LobbySwitchInfo:  int(agentData["lobby_switch_info"].(float64)),
	}

	agentGamesData := data["agent_games"].([]interface{})
	agentGames = make([]*table_model.AgentGame, 0)
	for _, agentGameDataI := range agentGamesData {
		agentGameData := agentGameDataI.(map[string]interface{})
		agentGames = append(agentGames, &table_model.AgentGame{
			AgentId: int(agentGameData["agent_id"].(float64)),
			GameId:  int(agentGameData["game_id"].(float64)),
			State:   int16(agentGameData["state"].(float64)),
		})
	}

	agentGameRoomsData := data["agent_game_rooms"].([]interface{})
	agentGameRooms = make([]*table_model.AgentGameRoom, 0)
	for _, agentGameRoomDataI := range agentGameRoomsData {
		agentGameRoomData := agentGameRoomDataI.(map[string]interface{})
		agentGameRooms = append(agentGameRooms, &table_model.AgentGameRoom{
			AgentId:    int(agentGameRoomData["agent_id"].(float64)),
			GameRoomId: int(agentGameRoomData["game_room_id"].(float64)),
			State:      int16(agentGameRoomData["state"].(float64)),
		})
	}

	return
}

func UdfDBAgentIpWhitelistResultToAgentWhitelist(bytes []byte) []*table_model.AgentIPWhitelistObj {
	ipWhitelist := make([]*table_model.AgentIPWhitelistObj, 0)
	for _, result := range utils.ToArrayMap(bytes) {
		ipWhitelist = append(ipWhitelist, &table_model.AgentIPWhitelistObj{
			CreateTime: int64(result["create_time"].(float64)),
			IPAddress:  result["ip_address"].(string),
			Info:       result["info"].(string),
			Creator:    result["creator"].(string),
		})
	}
	return ipWhitelist
}

func UdfDBAgentIpWhitelistResultToAgentWalletConnInfo(bytes []byte) *table_model.WalletConnInfo {
	tmp := new(table_model.WalletConnInfo)
	mapTmp := utils.ToMap(bytes)
	tmp.Path = utils.ToString(mapTmp["path"], "")
	tmp.Domain = utils.ToString(mapTmp["domain"], "")
	tmp.Scheme = utils.ToString(mapTmp["scheme"], "")
	tmp.ApiKey = utils.ToString(mapTmp["api_key"], "")
	return tmp
}

func UdfUpdateAgent(db *sql.DB, agentId int, agentName, agentInfo string, agentCommission int,
	adminUserUsername, adminUserRole, adminUserInfo string, adminUserIsEnabled int, isAdminUserRoleChanged bool,
	walletConninfoJson string, agentLobbySwitchInfo int) (agentUpdateTime time.Time, adminUserUpdateTime time.Time, err error) {
	/*
		CREATE FUNCTION "public"."udf_update_agent"("_agent_id" int4, "_agent_name" varchar,
			"_agent_info" varchar, "_agent_commission" int4, "_agent_cooperation" int4,
			"_agent_coin_supply_setting" jsonb, "_admin_user_username" varchar, "_admin_user_role" uuid,
		    "_admin_user_info" varchar, "_admin_user_is_enabled" int4, "_is_admin_user_role_changed" boolean)
			RETURNS int AS $$
	*/
	// 此 sql function 內含
	// 1. 修改代理
	// 2. 修改代理後台帳號
	// 3. 回傳 json 格式，目前僅回傳agent跟adminUser更新時間

	jsonResult := ""
	query := `SELECT "public"."udf_update_agent"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	err = db.QueryRow(query, agentId, agentName, agentInfo, agentCommission, adminUserUsername,
		adminUserRole, adminUserInfo, adminUserIsEnabled, isAdminUserRoleChanged, walletConninfoJson,
		agentLobbySwitchInfo).Scan(&jsonResult)
	if err != nil {
		return
	}

	result := utils.ToMap([]byte(jsonResult))

	agent := result["agent"].(map[string]interface{})
	agentUpdateTime = utils.MicroSecondsToUtcTime(int64(agent["update_time"].(float64)))

	adminUser := result["admin_user"].(map[string]interface{})
	adminUserUpdateTime = utils.MicroSecondsToUtcTime(int64(adminUser["update_time"].(float64)))

	return
}

func UpdateAgentCoinSupplySetting(db *sql.DB, agentId, commission, cooperation int, name, info string) error {

	query, args, err := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Update("agent").
		Set("commission", commission).
		Set("cooperation", cooperation).
		Set("name", name).
		Set("info", info).
		Where(sq.Eq{"id": agentId}).
		ToSql()
	if err != nil {
		return err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}
