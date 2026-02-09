package global

import (
	"database/sql"
	"errors"
	"time"

	md5 "backend/pkg/encrypt/md5hash"
	"backend/pkg/utils"

	sq "github.com/Masterminds/squirrel"
)

func CreateNewGameUser(db *sql.DB, originalUsername, transUsername, levelCode, userMetadata string, agentId int, coin float64) (userId int, err error) {

	if userMetadata == "" {
		userMetadata = "{}"
	}
	// create new user
	// set coin only at neew accont
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Insert("game_users").
		Columns("agent_id", "original_username", "username", "user_metadata", "sum_coin_in", "sum_coin_out", "level_code").
		Values(agentId, originalUsername, transUsername, userMetadata, coin, .0, levelCode).
		ToSql()

	query = query + " RETURNING id"

	lastInsertId := 0
	db.QueryRow(query, args...).Scan(&lastInsertId)

	if lastInsertId > 0 {
		tmp := new(AgentDataOfGameUser)
		tmp.GameUserId = lastInsertId
		tmp.AgentId = agentId
		tmp.LevelCode = levelCode
		AgentDataOfGameUserCache.Add(tmp)
	} else {
		err = errors.New("create game user failed")
	}

	return
}

/*
	檢查遊戲用戶帳號是否存在，不存在即創建
	udf_check_game_users_data
*/
func UdfCheckGameUserData(db *sql.DB, originalUsername, transUsername, defaultUserMetadata, levelCode string, agentId int, coin float64) (
	userId int, username, userMetadata, riskControlStatus string, isNew, isEnabled bool, killDiveState int, err error) {
	/*
		CREATE OR REPLACE FUNCTION "public"."udf_check_game_users_data"("_original_username" varchar, "_trans_username" varchar, "_agent_id" int4, "_coin" numeric, "_level_code" varchar)
			RETURNS TABLE("gs_id" int4, "gs_usernam" varchar, "gs_is_new" bool) AS $BODY$
	*/
	// 此 sql function 內含
	// 1. 檢查帳號是否存在，不存在即創建
	// 2. 創建帳號後更新 agent 內 member_count 欄位
	// 3. 回傳 json string {"id" : 1005, "username" : "dyg_d1dddec0e53379d8", "is_new" : false} (用戶id, 用戶原始帳號, 是否新用戶)

	jsonResult := ""
	query := `SELECT "public"."udf_check_game_users_data"($1, $2, $3, $4, $5, $6)`

	err = db.QueryRow(query, originalUsername, transUsername, agentId, coin, levelCode, defaultUserMetadata).Scan(&jsonResult)
	if err != nil {
		return
	}
	result := utils.ToMap([]byte(jsonResult))

	userId = utils.ToInt(result["id"], 0)
	username = utils.ToString(result["username"], "")
	isNew = utils.ToBool(result["is_new"], false)
	isEnabled = utils.ToBool(result["is_enabled"], false)
	userMetadata = utils.ToJSON(result["user_metadata"])
	killDiveState = utils.ToInt(result["kill_dive_state"], 0)
	riskControlStatus = utils.ToString(result["risk_control_status"], "0000")

	if isNew {
		tmp := new(AgentDataOfGameUser)
		tmp.GameUserId = userId
		tmp.AgentId = agentId
		tmp.LevelCode = levelCode
		AgentDataOfGameUserCache.Add(tmp)
	}

	return
}

/*
	統計玩家下注資料
	usp_game_users_stat
*/
func UspGameUsersStat(db *sql.DB, agentId, game_user_id, game_id, playCount, bigWinCount, winCount, loseCount int, levelCode string, ya, de, validYa, tax, bonus float64, lastBetTime time.Time) {
	/*"public"."usp_game_users_stat"(
	"_agent_id" int4,
	"_level_code" varchar,
	"_game_users_id" int4,
	"_game_id" int4,
	"_de" float8,
	"_ya" float8,
	"_vaild_ya" float8,
	"_tax" float8,
	"_bonus" float8,
	"_play_count" int4,
	"_big_win_count" int4,
	"_win_count" int4,
	"_lose_count" int4,
	"_last_bet_time" timestamptz)
	*/

	uspGameUsersStatLock.Lock()
	defer uspGameUsersStatLock.Unlock()
	// example
	// CALL "public"."usp_game_users_stat"('1001', '0001', '1007', '10', '20', '20', '0', '0', '1', '0', '0', '0', '2022-12-30T15:00:00Z')
	query := `CALL "public"."usp_game_users_stat"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	_, _ = db.Exec(query,
		agentId, levelCode, game_user_id, game_id, de,
		ya, validYa, tax, bonus, playCount,
		bigWinCount, winCount, loseCount, lastBetTime)
}

/*
	統計代理遊戲機率設定
	usp_insert_agent_game_ratio_stat
*/
func UspInsertAgentGameRatioStat(db *sql.DB, id, levelCode string, agentId, gameId, gameType, roomType, playCount int, ya, de, validYa, tax, bonus float64, lastBetTime time.Time) {
	/*"public"."usp_insert_agent_game_ratio_stat"
	"_id" varchar,
	"_level_code" varchar,
	"_agent_id" int4,
	"_game_id" int4,
	"_game_type" int4,
	"_room_type" int4,
	"_de" numeric,
	"_ya" numeric,
	"_vaild_ya" numeric,
	"_tax" numeric,
	"_bonus" numeric,
	"_play_count" int4,
	"_bet_time" timestamptz)
	*/

	uspInsertAgentGameRatioStatLock.Lock()
	defer uspInsertAgentGameRatioStatLock.Unlock()
	// example
	// CALL "public"."usp_insert_agent_game_ratio_stat"('132132', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '2023-02-20T00:00:00Z')
	query := `CALL "public"."usp_insert_agent_game_ratio_stat"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, _ = db.Exec(query,
		id, levelCode, agentId, gameId, gameType,
		roomType, de, ya, validYa, tax,
		bonus, playCount, lastBetTime)
}

/*
create game user account

原始帳號, hash成16碼字串

dyg_xxxxxxxxxxxxxxxx (加上前墜字串拼成20碼字串)
*/
func CreateNewThirdUsername(agentCode string, thirdAccount string) string {
	return "dyg_" + agentCode + "_" + md5.Hash16bit(thirdAccount)
}

func GetGameUserId(db *sql.DB, agentId int, originalUsername string) (userId int) {

	query := `SELECT id FROM game_users WHERE agent_id=$1 AND original_username=$2`
	db.QueryRow(query, agentId, originalUsername).Scan(&userId)

	return
}

func GetGameUserIdsByUsername(db *sql.DB, originalUsername string) (userIds []int, err error) {

	query := `SELECT id FROM game_users WHERE original_username=$1`
	db.Query(query, originalUsername)

	rows, err := db.Query(query, originalUsername)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var userId int
		if err = rows.Scan(&userId); err != nil {
			return
		}
		userIds = append(userIds, userId)
	}

	return
}
