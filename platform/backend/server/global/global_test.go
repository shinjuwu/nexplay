package global_test

import (
	"backend/internal/notification"
	"backend/internal/statistical"
	"backend/pkg/utils"
	"backend/server/global"
	"database/sql"
	"definition"
	"log"
	"strings"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {

	var money int64 = 10000

	log.Println(float64(int(money)))

	utcTimeNow := time.Now().UTC()

	log.Println(utcTimeNow)
	log.Println(global.RELOADDATA_COOLTIME_LAST_TIME)
	log.Println(utcTimeNow.After(global.RELOADDATA_COOLTIME_LAST_TIME.Add(global.RELOADDATA_COOLTIME_SEC)))

	if utcTimeNow.After(global.RELOADDATA_COOLTIME_LAST_TIME.Add(global.RELOADDATA_COOLTIME_SEC)) {
		global.RELOADDATA_COOLTIME_LAST_TIME = time.Now().UTC().Add(global.RELOADDATA_COOLTIME_SEC)
	}
	// log.Println(utcTimeNow)
	// log.Println(global.RELOADDATA_COOLTIME_LAST_TIME)
	// log.Println(utcTimeNow.After(global.RELOADDATA_COOLTIME_LAST_TIME.Add(global.RELOADDATA_COOLTIME_SEC)))
	log.Println(global.RELOADDATA_COOLTIME_LAST_TIME)

	// uri := `https://www.cnblogs.com/wanghui-garcia/p/10424463.html?s=0&account=kenny&money=1000&orderid=120220907151305551kenny&ip=&kind=0`
	// uri := `https://www.cnblogs.com/wanghui-garcia/p/10424463.html`
	uri := `?`
	idx := strings.Index(uri, "?")
	log.Println(idx)
	if idx >= 0 {
		log.Println(uri[:idx])
	}
}

func Test123456(t *testing.T) {
	nowTime := time.Now().UTC()
	startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day()-29, 0, 0, 0, 0, nowTime.UTC().Location())
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.UTC().Location())
	dateList := utils.GetTimeIntervalList(statistical.SAVE_PERIOD_DAY, startTime, endTime)

	t.Log(dateList)

}

func TestStorage(t *testing.T) {
	// GlobalStorage := new(global.StorageDatabaseCache)
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}
	GlobalStorage := global.NewStorageDatabaseCache(db)
	GlobalStorage.InitCacheFromDB()
	killDiveInfo := ""
	s, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFO)
	if ok {
		killDiveInfo = utils.ToString(s.Value, "")
	}
	AgentGameRatioCache := global.NewAgentGameRatioDatabaseCache(db, killDiveInfo)

	AgentGameRatioCache.InitCacheFromJson(killDiveInfo)

	err = AgentGameRatioCache.CreateNewGameToCacheFromJson(1001, killDiveInfo)
	if err != nil {
		t.Logf("CreateNewGameToCacheFromJson() has error: %v", err)
	}

	// 遊戲服務全域開關
	if _, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMESERVERINFO); ok {
		// return fmt.Errorf("storage key: %s is empty", definition.STORAGE_KEY_GAMEKILLDIVEINFO)
		// t.Errorf("storage key: %s is empty", definition.STORAGE_KEY_GAMESERVERINFO)
		val := make(map[string]interface{}, 0)
		val["state"] = 1
		GlobalStorage.Insert(definition.STORAGE_KEY_GAMESERVERINFO, utils.ToJSON(val), false)
	}

	isReset := false
	// 檢查代理遊戲殺放初始值是否已初始化
	if storage, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFORESET); ok {
		// 未初始化就跑一次初始化，並把初始化 flag 設為已初始化狀態
		// t.Logf("storage key: %s is empty", definition.STORAGE_KEY_GAMEKILLDIVEINFORESET)
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

	if isReset {
		// 遊戲服務殺放初始值
		if _, ok := GlobalStorage.SelectOne(definition.STORAGE_KEY_GAMEKILLDIVEINFO); ok {
			// 從 GAME SERVICE 取預設值
			val, _, err := notification.Getdefaultkilldiveinfo()
			if err != nil {
				t.Errorf("notification.Getdefaultkilldiveinfo() has error: %v", err)
				// return fmt.Errorf("notification.Getdefaultkilldiveinfo() has error: %v", err)
			} else {
				killDiveDefaultJson := utils.ToJSON(val)
				GlobalStorage.Insert(definition.STORAGE_KEY_GAMEKILLDIVEINFO, killDiveDefaultJson, false)
				AgentGameRatioCache.InitCacheFromJson(killDiveDefaultJson)
			}
		}
		// AgentGameRatioCache.InitCacheFromJson()
	} else {
		// 從 db 取值
		AgentGameRatioCache.InitCacheFromDB()
	}

}

func TestAgentCustomTagInfoDatabaseCache(t *testing.T) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}

	acti := global.NewAgentCustomTagInfoDatabaseCache(db)
	acti.InitDBAndCache()
}

func TestAgentGameIconListDatabaseCache(t *testing.T) {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dcc_game password=123456 sslmode=disable")
	if err != nil {
		t.Errorf("%v", err)
	}
	j := `[{"gameId":1,"rank":1,"hot":1},{"gameId":1,"rank":1,"hot":1}]`

	// hhh := utils.ToArrayMap([]byte(j))

	// t.Log(hhh)
	// ts := make([]*global.GameIcon, 0)
	// for _, v := range hhh {
	// 	jsonData := utils.ToJSON(v)

	// 	too := new(global.GameIcon)

	// 	err = json.Unmarshal([]byte(jsonData), too)
	// 	if err != nil {
	// 		t.Log(err)
	// 	}
	// 	ts = append(ts, too)
	// }

	// t.Log(ts)

	acti := global.NewAgentGameIconListDatabaseCache(db, j)
	acti.InitCacheAndDBFromDefaultJson()
}
