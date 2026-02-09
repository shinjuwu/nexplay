package global

import (
	"backend/internal/ginweb"
	"backend/internal/notification"
	"regexp"

	"backend/pkg/cache"
	"backend/pkg/job"
	"backend/pkg/redis"
	"backend/pkg/utils"

	table_model "backend/server/table/model"
	"context"
	"database/sql"
	"definition"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	sq "github.com/Masterminds/squirrel"
)

/*
全域參數定義
*/

const (
	/*
		redis feature idx define
	*/

	REDIS_IDX_VERIFY_INFO             = 0 // 驗證 game user login token
	REDIS_IDX_LOGIN_INFO              = 1 // 儲存 game user in game info
	REDIS_IDX_RELOGIN_INFO            = 2 // 儲存 game user relogin token
	REDIS_IDX_REALTIME_DATA_STAT_INFO = 3 // 儲存資訊總覽 (未使用)
	REDIS_IDX_AGENT_DATA              = 4 // 儲存 agent 資料
	REDIS_IDX_AUTO_RISK_CONTROL       = 5 // 儲存自動風控資料
	REDIS_IDX_MONITOR_SERVICE         = 5 // 儲存監控服務資料

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
	// [自動風控統計資料(多服務更新及讀取)]
	REDIS_KEY_AUTO_RISK_CONTROL_PREFIX = "AutoRiskControl"
	// [自動風控設定資料(後台更新&初始化，多服務讀取)]
	REDIS_KEY_AUTO_RISK_CONTROL_SETTING = REDIS_KEY_AUTO_RISK_CONTROL_PREFIX + "Setting"
	// [監控服務資料(後台更新&初始化)]
	REDIS_KEY_MONITOR_SERVICE_PREFIX = "MonitorService"

	SERVER_INFO_DEFAULT = 0 // mapping with lobby of game
)

var (
	/*
		DB 資料表資料轉存 local 端定義
	*/

	AgentCache                  *GlobalAgentCache
	GameCache                   *GlobalGameCache
	GameRoomCache               *GlobalGameRoomCache
	AgentGameCache              *GlobalAgentGameCache
	AgentGameRoomCache          *GlobalAgentGameRoomCache
	AgentPermissionCache        *GlobalAgentPermissionCache
	AdminUserCache              *GlobalAdminUserCache
	AutoRiskControlSettingCache *GlobalAutoRiskControlSettingCache
	AutoRiskControlStatCache    *GlobalAutoRiskControlStatCache
	MonitorServiceCache         *GlobalMonitorServiceDataCache

	// [ServerInfo.code(string) , ServerInfo(struct)]
	ServerInfoCache sync.Map

	// [game user id(string), AgentDataOfGameUser(struct)]
	AgentDataOfGameUserCache *GlobalAgentDataOfGameUserCache
	// [exchange data currency(string), ExchangeData(struct)]
	ExchangeDataCache *GlobalExchangeDataCache
	// [apiPath(string), PermissionList(map[string]interface{})]
	PermissionList sync.Map
	// [userId(int), kill dive value(map[string]float64)]
	GameUsersKillDiveValueList sync.Map
	// [userId(int), win rate(map[string]interface{})]
	GameUsersAutoRiskWinRateList sync.Map

	/*
		key-value syncmap cache
	*/
	GlobalStorage *storageDatabaseCache

	AgentGameRatioCache *agentGameRatioDatabaseCache

	AgentCustomTagInfoCache *agentCustomTagInfoDatabaseCache

	AgentGameIconListCache *agentGameIconListDatabaseCache

	RELOADDATA_COOLTIME_SEC       = 60 * time.Second
	RELOADDATA_COOLTIME_LAST_TIME = time.Unix(0, 0).UTC()

	uspGameUsersStatLock            = new(sync.Mutex)
	uspInsertAgentGameRatioStatLock = new(sync.Mutex)

	DEF_PLATFORM = "dev"

	// redis資料更新開關
	UpdateRedisData = false
)

type initFunc func(ginweb.ILogger, *sql.DB, redis.IRedisCliect) error
type cronJobFunc func(context.Context, ginweb.ILogger, *sql.DB, redis.IRedisCliect, *job.JobScheduler) error

// start_func must after init_func exec.
type globalCacheInitHandler struct {
	name         string
	init_func    initFunc
	cronjob_func cronJobFunc
}

func InitGlobalData(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, job *job.JobScheduler) error {

	// 設定短時間內不可重複更新
	utcTimeNow := time.Now().UTC()
	if utcTimeNow.After(RELOADDATA_COOLTIME_LAST_TIME.Add(RELOADDATA_COOLTIME_SEC)) {
		RELOADDATA_COOLTIME_LAST_TIME = time.Now().UTC()
	} else {
		return fmt.Errorf("the update time is too short, still have to wait %d seconds", (RELOADDATA_COOLTIME_SEC-utcTimeNow.Sub(RELOADDATA_COOLTIME_LAST_TIME))/time.Second)
	}

	initFuncs := []*globalCacheInitHandler{
		{name: "InitGlobalAgent", init_func: InitGlobalAgent},
		{name: "InitGlobalGame", init_func: InitGlobalGame},
		{name: "InitGlobalGameRoom", init_func: InitGlobalGameRoom},
		{name: "InitGlobalAgentGame", init_func: InitGlobalAgentGame},
		{name: "InitGlobalAgentGameRoom", init_func: InitGlobalAgentGameRoom},
		{name: "InitGlobalServerInfo", init_func: InitGlobalServerInfo},
		{name: "InitGlobalAgentIdOfGameUser", init_func: InitGlobalAgentIdOfGameUser},
		{name: "InitGlobalExchangeData", init_func: InitGlobalExchangeData},
		{name: "InitGlobalPermissionList", init_func: InitGlobalPermissionList},
		{name: "InitGlobalAgentPermission", init_func: InitGlobalAgentPermission},
		{name: "InitGlobalJobScheduler", init_func: InitGlobalJobScheduler},
		{name: "StartGlobalJobScheduler", init_func: InitGlobalJobfunc, cronjob_func: StartGlobalJobScheduler},
		{name: "InitGlobalAdminUserCache", init_func: InitGlobalAdminUserCache},
		{name: "InitGlobalDBTableObject", init_func: InitGlobalDBTableObject}, // must be after InitGlobalGame & InitGlobalServerInfo
		{name: "InitAutoRiskControlStatCache", init_func: InitAutoRiskControlStatCache},
		{name: "SyncRTPMonitorServiceGameList", init_func: SyncRTPMonitorServiceGameList}, // must be after InitGlobalGame & InitGlobalServerInfo
		{name: "InitMonitorServiceCache", init_func: InitMonitorServiceCache},

		// {name: "InitGlobalWSClient", init_func: InitGlobalWSClient},
	}

	for _, handler := range initFuncs {
		if handler.init_func != nil {
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

	if !UpdateRedisData {
		return nil
	}

	AgentCache.RemoveAll()

	query := `SELECT aa."id", aa."name", aa."code", aa."secret_key", aa."aes_key",
				aa."md5_key", aa."commission", aa."info", aa."is_enabled", aa."disable_time",
				aa."update_time", aa."create_time", aa."is_top_agent", aa."top_agent_id", aa."cooperation",
				aa."coin_limit", aa."coin_use", aa."level_code", aa."member_count", aa."creator",
				au."username" as "admin_username", aa."ip_whitelist", aa."kill_switch", aa."kill_ratio",
				aa."api_ip_whitelist", aa."currency", aa."is_not_kill_dive_cal", aa."child_agent_count", aa."jackpot_status",
				aa."jackpot_start_time", aa."jackpot_end_time", aa."wallet_type", aa."wallet_conninfo", aa."lobby_switch_info"
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
			&tmp.ApiIPWhitelistBytes, &tmp.Currency, &tmp.IsNotKillDiveCal, &tmp.ChildAgentCount, &tmp.JackpotStatus,
			&tmp.JackpotStartTime, &tmp.JackpotEndTime, &tmp.WalletType, &tmp.WalletConnInfoBytes, &tmp.LobbySwitchInfo); err != nil {
			return err
		}

		tmp.IPWhitelist = UdfDBAgentIpWhitelistResultToAgentWhitelist(tmp.IPWhitelistBytes)
		tmp.ApiIPWhitelist = UdfDBAgentIpWhitelistResultToAgentWhitelist(tmp.ApiIPWhitelistBytes)

		err = json.Unmarshal(tmp.WalletConnInfoBytes, &tmp.WalletConnInfo)
		if err != nil {
			return err
		}

		tmps = append(tmps, tmp)
	}

	if len(tmps) > 0 {
		AgentCache.Adds(tmps)
	}

	return nil
}

func InitGlobalGame(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	GameCache = NewGlobalGameCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "server_info_code", "name", "code", "state",
			"image", "h5_link", "create_time", "update_time", "type",
			"room_number", "table_number", "cal_state").
		From("game").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.Game

		if err := rows.Scan(&temp.Id, &temp.ServerInfoCode, &temp.Name, &temp.Code, &temp.State,
			&temp.Image, &temp.H5Link, &temp.CreateTime, &temp.UpdateTime, &temp.Type,
			&temp.RoomNumber, &temp.TableNumber, &temp.CalState); err != nil {
			return err
		}

		GameCache.cache.Add(temp.Id, &temp)
	}

	return nil
}

func InitGlobalGameRoom(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	GameRoomCache = NewGlobalGameRoomCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "name", "state", "game_id",
			"create_time", "update_time", "room_type").
		From("game_room").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.GameRoom

		if err := rows.Scan(&temp.Id, &temp.Name, &temp.State, &temp.GameId,
			&temp.CreateTime, &temp.UpdateTime, &temp.RoomType); err != nil {
			return err
		}

		GameRoomCache.cache.Add(temp.Id, &temp)
	}

	return nil
}

func InitGlobalAgentGame(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	AgentGameCache = NewGlobalAgentGameCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "game_id", "state").
		From("agent_game").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.AgentGame

		if err := rows.Scan(&temp.AgentId, &temp.GameId, &temp.State); err != nil {
			return err
		}

		AgentGameCache.Add(&temp)
	}

	return nil
}

func InitGlobalAgentGameRoom(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	AgentGameRoomCache = NewGlobalAgentGameRoomCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "game_room_id", "state").
		From("agent_game_room").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.AgentGameRoom

		if err := rows.Scan(&temp.AgentId, &temp.GameRoomId, &temp.State); err != nil {
			return err
		}

		AgentGameRoomCache.Add(&temp)
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

	if !UpdateRedisData {
		return nil
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
		// 初始化定點追殺用戶
		if kds == definition.GAMEUSERS_STATUS_KILLDIVE_CONFIGKILL {
			kdMap := make(map[string]float64)
			kdMap["need_update"] = float64(0)
			kdMap["kill_dive_value"] = kdv
			GameUsersKillDiveValueList.Store(tmp.GameUserId, kdMap)
		}
	}

	AgentDataOfGameUserCache.Adds(tmps)

	return nil
}

func InitGlobalExchangeData(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if ExchangeDataCache == nil {
		ExchangeDataCache = NewGlobalExchangeDataCache(rdb, REDIS_IDX_AGENT_DATA, REDIS_HASH_EXCHANGE_DATA)
	}

	if !UpdateRedisData {
		return nil
	}

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "currency", "to_coin").
		From("exchange_data").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp ExchangeData

		if err := rows.Scan(&temp.Id, &temp.Currency, &temp.ToCoin); err != nil {
			return err
		}

		ExchangeDataCache.Add(&temp)
	}

	return nil
}

func InitGlobalPermissionList(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("feature_code", "name", "api_path", "is_enabled", "is_required", "remark", "create_time", "update_time", "action_type").
		From("permission_list").
		OrderBy("feature_code desc").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	// clear map
	PermissionList.Range(func(key interface{}, value interface{}) bool {
		PermissionList.Delete(key)
		return true
	})

	for rows.Next() {
		temp := table_model.NewEmptyPermissionList()

		if err := rows.Scan(&temp.FeatureCode, &temp.Name, &temp.ApiPath, &temp.IsEnabled, &temp.IsRequired,
			&temp.Remark, &temp.CreateTime, &temp.UpdateTime, &temp.ActionType); err != nil {
			return err
		}

		tmpMap, err := utils.StructToMap(temp)
		if err != nil {
			log.Printf("InitGlobalPermissionList() StructToMap has error: %v", err)
		} else {
			PermissionList.Store(temp.ApiPath, tmpMap)
		}
	}

	return nil
}

func InitGlobalAgentPermission(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	AgentPermissionCache = NewGlobalAgentPermissionCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("id", "agent_id", "account_type", "name", "info", "permission").
		From("agent_permission").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.AgentPermission

		if err := rows.Scan(&temp.Id, &temp.AgentId, &temp.AccountType,
			&temp.Name, &temp.Info, &temp.PermissionBytes); err != nil {
			return err
		}

		utils.ToStruct(temp.PermissionBytes, &temp.Permission)

		AgentPermissionCache.Add(&temp)
	}

	return nil
}

func InitGlobalJobScheduler(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	/*
		1. read setting from db
		2. init jbs event setting
	*/

	if JobSchedulerSwitch {

		query, args, _ := sq.StatementBuilder.
			PlaceholderFormat(sq.Dollar).
			Select("id", "spec", "info", "trigger_func", "is_enabled", "exec_limit", "last_sync_date").
			From("job_scheduler").
			ToSql()

		rows, err := db.Query(query, args...)
		if err != nil {
			return err
		}

		defer rows.Close()

		// clear map
		JobSchedulerSetting.Range(func(key interface{}, value interface{}) bool {
			PermissionList.Delete(key)
			return true
		})

		for rows.Next() {
			temp := table_model.NewEmptyJobScheduler()

			if err := rows.Scan(&temp.Id, &temp.Spec, &temp.Info, &temp.TriggerFunc, &temp.IsEnabled,
				&temp.ExecLimit, &temp.LastSyncDate); err != nil {
				return err
			}

			if err != nil {
				log.Printf("InitGlobalJobScheduler() StructToMap has error: %v", err)
			} else {
				if temp.IsEnabled {
					JobSchedulerSetting.Store(temp.Id, temp)
				}
			}
		}
	}

	return nil
}

/*
* 定義 function mapping list
 */
func InitGlobalJobfunc(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	JobSchedulerFuncCache = cache.NewLocalDataCache()

	/*
		key 對應資料庫內的 job_scheduler -> trigger_func
	*/
	// define customized job func after here
	// JobSchedulerFuncMap["TestFunc"] = jobTest

	if JobSchedulerFuncCache != nil {
		{ // 測試用func
			// JobSchedulerFuncCache.Add("TestFunc", jobTest)
		}
		{ // 業績報表任務定義 // report (視角: 代理 -> 用戶投注統計)
			JobSchedulerFuncCache.Add("job_rp_agent_stat_15min", jobRPAgentStat15Minute)
		}
		{
			// 遊戲輸贏排行榜 // realtime // mv table (視角: 代理、遊戲用戶 -> 用戶投注統計)
			// JobSchedulerFuncCache.Add("job_rt_game_stat_hour", jobRTGameStatHour)
		}
	}

	return nil
}

/*
* 實際執行 cron job
	job setting depends on InitGlobalJobScheduler()
	function mapping define depends on InitGlobalJobfunc()
*/
func StartGlobalJobScheduler(ctx context.Context, logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect, cronJob *job.JobScheduler) error {
	JobSchedulerSetting.Range(func(key interface{}, value interface{}) bool {
		val := value.(*table_model.JobScheduler)
		fn_interface, ok := JobSchedulerFuncCache.Get(val.TriggerFunc)
		if ok {
			fn, ok := fn_interface.(func(context.Context, *sql.DB, redis.IRedisCliect, []string))
			if ok {
				cronJob.CreateJob(ctx, db, rdb, val.Id, val.Spec, val.Info, val.LastSyncDate, val.IsEnabled, val.ExecLimit, fn)
			} else {
				log.Printf("StartGlobalJobScheduler has unknow TriggerFunc is: %v", val.TriggerFunc)
			}
		}
		return true
	})

	go cronJob.Start()

	return nil
}

func InitGlobalAdminUserCache(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	AdminUserCache = NewGlobalAdminUserCache()

	query, args, _ := sq.StatementBuilder.
		PlaceholderFormat(sq.Dollar).
		Select("agent_id", "username", "password", "nickname", "google_auth",
			"google_key", "allow_ip", "account_type", "is_readonly", "is_enabled",
			"update_time", "create_time", "is_added", "login_time", "role",
			"info").
		From("admin_user").
		ToSql()

	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var temp table_model.AdminUser

		if err := rows.Scan(&temp.AgentId, &temp.Username, &temp.Password, &temp.Nickname, &temp.GoogleAuth,
			&temp.GoogleKey, &temp.AllowIp, &temp.AccountType, &temp.IsReadonly, &temp.IsEnabled,
			&temp.UpdateTime, &temp.CreateTime, &temp.IsAdded, &temp.LoginTime, &temp.PermissionId,
			&temp.Info); err != nil {
			return err
		}

		// fmt.Println(temp.Username, temp.UpdateTime)

		AdminUserCache.Add(&temp)
	}

	return nil
}

func InitGlobalDBTableObject(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	// check storage exist
	if GlobalStorage == nil {
		GlobalStorage = NewStorageDatabaseCache(db)
	}

	if err := GlobalStorage.InitCacheFromDB(); err != nil {
		return err
	}

	// 遊戲服務全域開關
	if _, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO); !ok {
		return fmt.Errorf("storage key: %s is empty", definition.STORAGE_KEY_GAMESERVERINFO)
	}

	// 報表相關參數設定
	defaultServerInfoSetting := table_model.NewServerInfoSetting()
	defaultServerInfoSettingStr := utils.ToJSON(defaultServerInfoSetting)
	if _, ok := GlobalStorage.SelectOrInsertOne(definition.STORAGE_KEY_SERVERINFO_SETTING, defaultServerInfoSettingStr, false); !ok {
		logger.Printf("storage key: %s is empty, create new from default value: %v", definition.STORAGE_KEY_SERVERINFO_SETTING, defaultServerInfoSettingStr)
	}

	// default index is lobby
	// lobby idx: 0
	gg := GameCache.Get(SERVER_INFO_DEFAULT)

	re := regexp.MustCompile("[0-9]+")
	result := re.ReplaceAllString(gg.ServerInfoCode, "")
	DEF_PLATFORM = result

	addr, ok := ServerInfoCache.Load(gg.ServerInfoCode)
	if ok {
		addres, ok := addr.(table_model.ServerInfo)
		if ok {
			// 設定通知目標的 address
			notification.SetNotificationAddress(addres.Addresses.Notification)
		} else {
			logger.Info("InitRouter() SetNotificationAddress not found")
		}
	} else {
		logger.Info("InitRouter() ServerInfoCache not found")
	}

	// 遊戲殺放設定
	if err := checkAgentGameRatioCacheExist(logger, db, rdb); err != nil {
		return err
	}

	// 自定義標籤設定
	if err := checkAgentCustomTagInfoCacheExist(logger, db, rdb); err != nil {
		return err
	}

	// 自定義遊戲icon list
	if err := checkAgentGameIconListCacheExist(logger, db, rdb); err != nil {
		return err
	}

	// 風控設定
	if err := checkAutoRiskControlSettingCacheExist(logger, db, rdb); err != nil {
		return err
	}

	return nil
}

/*
	遊戲殺放設定
		STORAGE_KEY_GAMEKILLDIVEINFORESET: 重置flag
		STORAGE_KEY_GAMEKILLDIVEINFO: 殺放初始值(from game service)
*/
func checkAgentGameRatioCacheExist(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if AgentGameRatioCache == nil {
		defaultJson := ""
		s, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFO)
		if !ok {
			logger.Println("select definition.STORAGE_KEY_GAMEKILLDIVEINFO has nothing")
		} else {
			defaultJson = utils.ToString(s.Value, "")
		}
		AgentGameRatioCache = NewAgentGameRatioDatabaseCache(db, defaultJson)
	}

	isReset := false
	// 檢查代理遊戲殺放初始值是否需要重置
	if storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFORESET); !ok {
		// 未初始化就跑一次初始化，並把初始化 flag 設為已初始化狀態
		logger.Printf("storage key: %s is empty", definition.STORAGE_KEY_GAMEKILLDIVEINFORESET)
		val := make(map[string]interface{}, 0)
		val["flag"] = false
		val["update_time"] = time.Now().UTC().UnixMilli() // 初始化時間
		GlobalStorage.Insert(definition.STORAGE_KEY_GAMEKILLDIVEINFORESET, utils.ToJSON(val), false)
		isReset = true
	} else {
		tmp := utils.ToMap([]byte(storage.Value))
		if flag, ok := tmp["flag"].(bool); ok {
			isReset = flag
		}
	}

	killDiveDefaultJson := ""
	if isReset {
		// 重置就從 遊戲 取值
		// 並清除所有已有代理殺放設定資料
		if err := AgentGameRatioCache.Clear(); err != nil {
			return err
		}
		// 遊戲服務殺放初始值

		// 從 GAME SERVICE 取預設值
		// 每次開機都更新預設值
		val, _, err := notification.Getdefaultkilldiveinfo()
		if err != nil {
			return fmt.Errorf("notification.Getdefaultkilldiveinfo() has error: %v", err)
		}
		killDiveDefaultJson = utils.ToJSON(val)
		if killDiveDefaultJson != "" {
			if _, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFO); !ok {
				// storage 內沒有預設物件就新增
				if _, err := GlobalStorage.Insert(definition.STORAGE_KEY_GAMEKILLDIVEINFO, killDiveDefaultJson, false); err != nil {
					return err
				}
			} else {
				// storage 內有預設物件就更新
				if err := GlobalStorage.Update(definition.STORAGE_KEY_GAMEKILLDIVEINFO, killDiveDefaultJson); err != nil {
					return err
				}
			}

			if err := AgentGameRatioCache.InitCacheFromJson(killDiveDefaultJson); err != nil {
				return err
			}

			// 更新成功, 修改全預旗標
			tmp := make(map[string]interface{}, 0)
			tmp["flag"] = false
			tmp["update_time"] = time.Now().UTC().UnixMilli()
			GlobalStorage.Update(definition.STORAGE_KEY_GAMEKILLDIVEINFORESET, utils.ToJSON(tmp))
		} else {
			return fmt.Errorf("AgentGameRatioCache init failed, killDiveDefaultJson is empty")
		}
	} else {
		// 非重置就從 db 取值
		if err := AgentGameRatioCache.InitCacheFromDB(); err != nil {
			return err
		}
	}

	// check
	if AgentGameRatioCache.agentGameRatio.Count() == 0 {
		return fmt.Errorf("gameKillInfo init failed")
	}

	return nil
}

/*
	檢查自定義標籤設定是否存在
		STORAGE_KEY_GAMEAGENTCUSTOMTAGRESET: 重置flag
*/
func checkAgentCustomTagInfoCacheExist(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if AgentCustomTagInfoCache == nil {
		AgentCustomTagInfoCache = NewAgentCustomTagInfoDatabaseCache(db)
	}

	isReset := false
	// 檢查代理自定義標示初始值是否已初始化
	if storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEAGENTCUSTOMTAGRESET); !ok {
		// 未初始化就跑一次初始化，並把初始化 flag 設為已初始化狀態
		logger.Printf("storage key: %s is empty", definition.STORAGE_KEY_GAMEAGENTCUSTOMTAGRESET)
		val := make(map[string]interface{}, 0)
		val["flag"] = false
		val["update_time"] = time.Now().UTC().UnixMilli() // 初始化時間
		GlobalStorage.Insert(definition.STORAGE_KEY_GAMEAGENTCUSTOMTAGRESET, utils.ToJSON(val), false)
		isReset = true
	} else {
		tmp := utils.ToMap([]byte(storage.Value))
		if flag, ok := tmp["flag"].(bool); ok {
			isReset = flag
		}
	}

	if isReset {
		if err := AgentCustomTagInfoCache.Clear(); err != nil {
			return err
		}

		if err := AgentCustomTagInfoCache.InitDBAndCache(); err != nil {
			return err
		}

		// 更新成功, 修改全預旗標
		tmp := make(map[string]interface{}, 0)
		tmp["flag"] = false
		tmp["update_time"] = time.Now().UTC().UnixMilli()
		GlobalStorage.Update(definition.STORAGE_KEY_GAMEAGENTCUSTOMTAGRESET, utils.ToJSON(tmp))
	} else {
		if err := AgentCustomTagInfoCache.InitCacheFromDB(); err != nil {
			return err
		}
	}

	return nil
}

/*
	檢查自定義遊戲icon list 是否存在
		STORAGE_KEY_GAMEICONLISTDEFAULTRESET: 重置flag
		STORAGE_KEY_GAMEICONLISTDEFAULT: 初始值(from game service)
*/
func checkAgentGameIconListCacheExist(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	isReset := false
	isAddNew := false
	nowTimeMilli := time.Now().UTC().UnixMilli() // 初始化時間
	isNeedUpdateResetFlag := false
	// 檢查是否需要重置
	if storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEICONLISTDEFAULTRESET); !ok {
		// storage 內沒有值就跑一次初始化，並把初始化 flag 設為已初始化狀態
		logger.Printf("storage key: %s is empty", definition.STORAGE_KEY_GAMEICONLISTDEFAULTRESET)
		val := make(map[string]interface{}, 0)
		val["flag"] = false
		val["update_time"] = nowTimeMilli
		GlobalStorage.Insert(definition.STORAGE_KEY_GAMEICONLISTDEFAULTRESET, utils.ToJSON(val), false)
		isReset = true
	} else {
		tmp := utils.ToMap([]byte(storage.Value))
		if flag, ok := tmp["flag"].(bool); ok {
			isReset = flag
			isNeedUpdateResetFlag = true
		}
	}

	gameCachetmps := GameCache.GetAll()
	sort.Sort(table_model.GameSlice(gameCachetmps))
	gameIconListDefaults := make([]*GameIcon, 0)
	defaultJson := ""
	// 檢查是否有預設值 | 是否要初始化
	s, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEICONLISTDEFAULT)
	if !ok || isReset {
		logger.Println("select definition.STORAGE_KEY_GAMEICONLISTDEFAULT has nothing, create new")
		// 沒有預設值就新增
		/*
			type GameIcon struct {
				GameId int `json:"gameId"` // game id (union id)
				Rank   int `json:"rank"`   // 排行 (since from 0)
				Hot    int `json:"hot"`    // 熱門 (0:close, 1:open)
				Newest int `json:"newest"` // 最新 (0:close, 1:open)
				Push   int `json:"push"`   // 推廣大圖 (0:close, 1:第一張大圖, 2:第二章大圖.....)
			}
		*/
		index := 0
		for _, gameCachetmp := range gameCachetmps {
			if gameCachetmp.Id > definition.GAME_ID_ALL {
				giltmp := new(GameIcon)
				giltmp.GameId = gameCachetmp.Id
				giltmp.Hot = 0
				giltmp.Newest = 0
				giltmp.Rank = index
				index++
				gameIconListDefaults = append(gameIconListDefaults, giltmp)
			}
		}

		defaultJson = utils.ToJSON(gameIconListDefaults)
		if ok {
			if err := GlobalStorage.Update(definition.STORAGE_KEY_GAMEICONLISTDEFAULT, defaultJson); err != nil {
				return err
			}
		} else {
			if _, err := GlobalStorage.Insert(definition.STORAGE_KEY_GAMEICONLISTDEFAULT, defaultJson, false); err != nil {
				return err
			}
		}
	} else {
		defaultJson = utils.ToString(s.Value, "")
		gameIconListDefaults = transIconList([]byte(defaultJson))
	}

	// 檢查是否有新遊戲
	// 遊戲列表內有定義大廳所以要-1
	newGameIconList := make([]*GameIcon, 0)
	if len(gameIconListDefaults) < len(gameCachetmps)-1 {
		isAddNew = true
		count := 0
		// 找出新遊戲資料
		for _, gameCachetmp := range gameCachetmps {
			if gameCachetmp.Id > definition.GAME_ID_ALL {
				found := false
				for _, gameIconListDefault := range gameIconListDefaults {
					if gameCachetmp.Id == gameIconListDefault.GameId {
						found = true
						break
					}
				}
				if !found {
					giltmp := new(GameIcon)
					giltmp.GameId = gameCachetmp.Id
					giltmp.Hot = 0
					giltmp.Newest = 0
					giltmp.Rank = len(gameIconListDefaults) + count
					giltmp.Push = 0
					count++
					newGameIconList = append(newGameIconList, giltmp)
				}
			}
		}

		if len(newGameIconList) > 0 {
			// 更新 defaultJson
			gameIconListDefaults = append(gameIconListDefaults, newGameIconList...)
			defaultJson = utils.ToJSON(gameIconListDefaults)
		}
	}

	if AgentGameIconListCache == nil {
		AgentGameIconListCache = NewAgentGameIconListDatabaseCache(db, defaultJson)
	}

	if isReset {
		/*
			重置步驟如下：
			1. 清空本地端暫存
			2. 透過處理過的 game icon list json 字串初始化 DB 與本地端暫存
		*/
		if err := AgentGameIconListCache.Clear(); err != nil {
			return err
		}

		if err := AgentGameIconListCache.InitCacheAndDBFromDefaultJson(); err != nil {
			return err
		}
	} else if isAddNew {
		/*
			新增遊戲初始化步驟如下：
			1. 透過處理過的 new game icon list json 字串新增至 DB 與本地端暫存
			2. 更新預設初始 game icon list json 字串
		*/
		if err := AgentGameIconListCache.InitCacheAndDBAddNew(newGameIconList); err != nil {
			return err
		}
		if err := GlobalStorage.Update(definition.STORAGE_KEY_GAMEICONLISTDEFAULT, defaultJson); err != nil {
			return err
		}
	} else {
		if err := AgentGameIconListCache.InitCacheFromDB(); err != nil {
			return err
		}
	}

	if isNeedUpdateResetFlag {
		val := make(map[string]interface{}, 0)
		val["flag"] = false
		val["update_time"] = nowTimeMilli
		if err := GlobalStorage.Update(definition.STORAGE_KEY_GAMEICONLISTDEFAULTRESET, utils.ToJSON(val)); err != nil {
			return err
		}
	}

	return nil
}

/*
	檢查風控設定是否存在
*/
func checkAutoRiskControlSettingCacheExist(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if AutoRiskControlSettingCache == nil {
		AutoRiskControlSettingCache = NewGlobalAutoRiskControlSettingCache(rdb, REDIS_IDX_AUTO_RISK_CONTROL, REDIS_KEY_AUTO_RISK_CONTROL_SETTING)
	}

	if !UpdateRedisData {
		return nil
	}

	defaultAutoRiskControlSetting := table_model.NewAutoRiskControlSetting()
	defaultAutoRiskControlSettingStr := utils.ToJSON(defaultAutoRiskControlSetting)
	s, ok := GlobalStorage.SelectOrInsertOne(definition.STORAGE_KEY_AUTO_RISK_CONTROL_SETTING, defaultAutoRiskControlSettingStr, false)
	if !ok {
		logger.Printf("storage key: %s is empty, create new from default value: %v", definition.STORAGE_KEY_AUTO_RISK_CONTROL_SETTING, defaultAutoRiskControlSettingStr)
	}

	data := new(table_model.AutoRiskControlSetting)
	err := json.Unmarshal([]byte(s.Value), data)
	if err != nil {
		return err
	}

	return AutoRiskControlSettingCache.Add(data)
}

func InitAutoRiskControlStatCache(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if AutoRiskControlStatCache == nil {
		AutoRiskControlStatCache = NewGlobalAutoRiskControlStatCache(rdb, REDIS_IDX_AUTO_RISK_CONTROL, REDIS_KEY_AUTO_RISK_CONTROL_PREFIX)
	}
	return nil
}

// RTP監控 同步遊戲列表 (開機時同步)
func SyncRTPMonitorServiceGameList(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if GameCache != nil {

		conninfo := make(map[string]interface{}, 0)
		addr, ok := ServerInfoCache.Load("monitor")
		if ok {
			addres, ok := addr.(table_model.ServerInfo)
			if ok {
				conninfo = utils.ToMap(addres.AddressesBytes)
			}

			notification.SendGameListToMonitorService(conninfo, DEF_PLATFORM, GameCache.GetAll())
		}
	}

	return nil
}

func InitMonitorServiceCache(logger ginweb.ILogger, db *sql.DB, rdb redis.IRedisCliect) error {
	if MonitorServiceCache == nil {
		MonitorServiceCache = NewGlobalMonitorServiceDataCache(rdb, REDIS_IDX_MONITOR_SERVICE, REDIS_KEY_MONITOR_SERVICE_PREFIX)
	}
	return nil
}
