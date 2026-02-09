package global

import (
	"backend/internal/ginweb"
	"backend/pkg/database"
	"backend/pkg/job"
	"backend/pkg/redis"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
)

/*
全域參數定義
*/

const (
	DB_IDX_ORDER = 0
	DB_IDX_GAME  = 1

	/*
		redis feature idx define
	*/

	REDIS_IDX_VERIFY_INFO             = 0 // 驗證 game user login token
	REDIS_IDX_LOGIN_INFO              = 1 // 儲存 game user in game info
	REDIS_IDX_RELOGIN_INFO            = 2 // 儲存 game user relogin token
	REDIS_IDX_REALTIME_DATA_STAT_INFO = 3 // 儲存資訊總覽 (未使用)
	REDIS_IDX_AGENT_DATA              = 4 // 儲存 agent 資料

	// 遊戲用戶註冊 token 在 redis 存活時間
	REDIS_LOGIN_TOKEN_LIFETIME = 30 * time.Minute

	// 遊戲用戶 relogin token 在 redis 存活時間
	REDIS_RELOGIN_TOKEN_LIFETIME = 1 * time.Minute

	// [username, token], 後蓋前
	REDIS_HASH_LOGIN_TOKEN = "LoginToken"
	// [username, 登入資訊(json string)]
	REDIS_HASH_INGAME_USER = "InGameUser"
	// [username, token], 後蓋前
	REDIS_HASH_RELOGIN_TOKEN = "ReLoginToken"
	// [代理資料(後台更新&初始化，多服務讀取)]
	REDIS_HASH_AGENT_DATA = "AgentData"
	// [遊戲用戶關聯代理資料(後台更新&初始化，多服務讀取)]
	REDIS_HASH_AGEN_OF_GAMEUSER_DATA = "AgentOfGameUserData"
	// [遊戲幣種設定資料(後台更新&初始化，多服務讀取)]
	REDIS_HASH_EXCHANGE_DATA = "ExchangeData"

	SERVER_INFO_DEFAULT = 0 // mapping with lobby of game
)

var (
	/*
		DB 資料表資料轉存 local 端定義
	*/

	AgentCache      *GlobalAgentCache
	ServerInfoCache sync.Map

	// [game user id(string), AgentDataOfGameUser(struct)]
	AgentDataOfGameUserCache *GlobalAgentDataOfGameUserCache

	// [ScheduleToBackup id(string), ScheduleToBackup(struct)]
	ScheduleToBackupCache *GlobalScheduleToBackupCache
	/*
		key-value syncmap cache
	*/
	GlobalStorage *storageDatabaseCache

	RELOADDATA_COOLTIME_SEC       = 60 * time.Second
	RELOADDATA_COOLTIME_LAST_TIME = time.Unix(0, 0).UTC()

	// uspGameUsersStatLock            = new(sync.Mutex)
	// uspInsertAgentGameRatioStatLock = new(sync.Mutex)

)

type initFunc func(ginweb.ILogger, *sql.DB, redis.IRedisCliect) error
type cronJobFunc func(context.Context, ginweb.ILogger, *sql.DB, redis.IRedisCliect, *job.JobScheduler) error

// start_func must after init_func exec.
type globalCacheInitHandler struct {
	name         string
	db_idx       int
	init_func    initFunc
	cronjob_func cronJobFunc
}

func InitGlobalData(logger ginweb.ILogger, idb database.IDatabase, rdb redis.IRedisCliect, job *job.JobScheduler) error {

	// 設定短時間內不可重複更新
	utcTimeNow := time.Now().UTC()
	if utcTimeNow.After(RELOADDATA_COOLTIME_LAST_TIME.Add(RELOADDATA_COOLTIME_SEC)) {
		RELOADDATA_COOLTIME_LAST_TIME = time.Now().UTC()
	} else {
		return fmt.Errorf("the update time is too short, still have to wait %d seconds", (RELOADDATA_COOLTIME_SEC-utcTimeNow.Sub(RELOADDATA_COOLTIME_LAST_TIME))/time.Second)
	}

	initFuncs := []*globalCacheInitHandler{
		{name: "InitGlobalAgent", db_idx: DB_IDX_GAME, init_func: InitGlobalAgent},
		{name: "InitGlobalServerInfo", db_idx: DB_IDX_GAME, init_func: InitGlobalServerInfo},
		{name: "InitGlobalAgentIdOfGameUser", db_idx: DB_IDX_GAME, init_func: InitGlobalAgentIdOfGameUser},
		{name: "InitGlobalScheduleToBackup", db_idx: DB_IDX_ORDER, init_func: InitGlobalScheduleToBackup},
	}

	for _, handler := range initFuncs {
		if handler.init_func != nil {
			db := idb.GetDB(handler.db_idx)
			if db == nil {
				return fmt.Errorf("idb.GetDB failed, idx is :%d", handler.db_idx)
			}
			if err := handler.init_func(logger, db, rdb); err != nil {
				logger.Printf("%s() failed, err: %v", handler.name, err)
				return err
			}
			logger.Printf("%s() init success", handler.name)
		}
	}

	ctx := context.Background()

	// start func must after init func exec.
	for _, handler := range initFuncs {
		if handler.cronjob_func != nil {
			db := idb.GetDB(handler.db_idx)
			if db == nil {
				return fmt.Errorf("idb.GetDB failed, idx is :%d", handler.db_idx)
			}
			if err := handler.cronjob_func(ctx, logger, db, rdb, job); err != nil {
				logger.Printf("%s() failed, err: %v", handler.name, err)
				return err
			}
			logger.Printf("%s() init success", handler.name)
		}
	}

	return nil
}

func InitGlobalAgent(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {

	if AgentCache == nil {
		AgentCache = NewGlobalAgentCache(rdb, REDIS_IDX_AGENT_DATA, REDIS_HASH_AGENT_DATA)
	}

	AgentCache.RemoveAll()

	query := `SELECT aa."id", aa."name", aa."code", aa."secret_key", aa."aes_key",
				aa."md5_key", aa."commission", aa."info", aa."is_enabled", aa."disable_time",
				aa."update_time", aa."create_time", aa."is_top_agent", aa."top_agent_id", aa."cooperation",
				aa."coin_limit", aa."coin_use", aa."level_code", aa."member_count", aa."creator",
				au."username" as "admin_username", aa."ip_whitelist", aa."kill_switch", aa."kill_ratio",
				aa."api_ip_whitelist", aa."currency"
			FROM "agent" aa, "admin_user" au
			WHERE aa."id" = au."agent_id" AND au."is_added" = false;`

	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	tmps := make([]*table_model.Agent, 0)
	for rows.Next() {
		tmp := table_model.NewEmptyAgent()

		if err := rows.Scan(&tmp.Id, &tmp.Name, &tmp.Code, &tmp.SecretKey, &tmp.AesKey,
			&tmp.Md5Key, &tmp.Commission, &tmp.Info, &tmp.IsEnabled, &tmp.DisableTime,
			&tmp.UpdateTime, &tmp.CreateTime, &tmp.IsTopAgent, &tmp.TopAgentId, &tmp.Cooperation,
			&tmp.CoinLimit, &tmp.CoinUse, &tmp.LevelCode, &tmp.MemberCount,
			&tmp.Creator, &tmp.AdminUsername, &tmp.IPWhitelistBytes, &tmp.KillSwitch, &tmp.KillRatio,
			&tmp.ApiIPWhitelistBytes, &tmp.Currency); err != nil {
			return err
		}

		tmp.TranslateWhiteIPList()

		// AgentCache.Add(tmp)
		tmps = append(tmps, tmp)
	}

	if len(tmps) > 0 {
		AgentCache.Adds(tmps)
	}

	return nil
}

func InitGlobalServerInfo(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("code", "ip", "addresses", "is_enabled").
		From("server_info").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	// clear map
	ServerInfoCache.Range(func(key interface{}, value interface{}) bool {
		ServerInfoCache.Delete(key)
		return true
	})

	for rows.Next() {
		var temp table_model.ServerInfo

		if err := rows.Scan(&temp.Code, &temp.Ip, &temp.AddressesBytes, &temp.IsEnabled); err != nil {
			return err
		}

		utils.ToStruct(temp.AddressesBytes, &temp.Addresses)

		ServerInfoCache.Store(temp.Code, temp)
	}

	return nil
}

func InitGlobalAgentIdOfGameUser(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {

	if AgentDataOfGameUserCache == nil {
		AgentDataOfGameUserCache = NewGlobalGameUserCache(rdb, REDIS_IDX_AGENT_DATA, REDIS_HASH_AGEN_OF_GAMEUSER_DATA)
	}

	AgentDataOfGameUserCache.RemoveAll()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", "level_code", "kill_dive_state", "kill_dive_value").
		From("game_users").
		Where(sq.Eq{"is_enabled": true}).
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	tmps := make([]*AgentDataOfGameUser, 0)

	for rows.Next() {
		tmp := new(AgentDataOfGameUser)
		var kds int
		var kdv float64

		if err := rows.Scan(&tmp.GameUserId, &tmp.AgentId, &tmp.LevelCode, &kds, &kdv); err != nil {
			return err
		}

		tmps = append(tmps, tmp)
	}

	AgentDataOfGameUserCache.Adds(tmps)

	return nil
}

func InitGlobalScheduleToBackup(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {

	ScheduleToBackupCache = NewGlobalScheduleToBackupCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "data_keeping_table", "data_keeping_day", "is_enabled", "create_time",
			"update_time", "disable_time", "last_exec_time").
		From("schedule_to_backup").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.ScheduleToBackup

		if err := rows.Scan(&temp.Id, &temp.DataKeepingTable, &temp.DataKeepingDay, &temp.IsEnabled, &temp.CreateTime,
			&temp.UpdateTime, &temp.DisableTime, &temp.LastExecTime); err != nil {
			return err
		}

		ScheduleToBackupCache.cache.Add(temp.Id, &temp)
	}

	return nil
}
