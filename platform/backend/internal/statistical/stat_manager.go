package statistical

import (
	"backend/pkg/cache"
	"backend/pkg/redis"
	"backend/pkg/utils"

	"fmt"
	"sync"
	"time"
)

/*
數據總覽物件
存儲時機: 每日、每週
*/

const (
	STAT_RESULT_WIN  = "win"
	STAT_RESULT_LOSE = "lose"

	// 資料備份時間
	SAVE_PERIOD_MINUTE = "minute"
	SAVE_PERIOD_HOUR   = "hour"
	SAVE_PERIOD_DAY    = "day"
	SAVE_PERIOD_WEEK   = "week"
	SAVE_PERIOD_MONTH  = "month"

	// ${levelCode + _ + logtime}
	STATISTICAL_KEY_FORMAT = "%s_%s"
)

type StatisticalManager struct {
	rdb         redis.IRedisCliect
	rdbIdx      int
	logtime_now string
	data_of_day map[string]cache.ILocalDataCache // data
	update_list cache.ILocalDataCache
	lock        *sync.Mutex
}

func NewStatisticalManager(rdb redis.IRedisCliect, rdbIdx int) *StatisticalManager {
	return &StatisticalManager{
		rdb:         rdb,
		rdbIdx:      rdbIdx,
		logtime_now: "",
		data_of_day: make(map[string]cache.ILocalDataCache),
		update_list: cache.NewLocalDataCache(),
		lock:        new(sync.Mutex),
		// data_of_day: cache.NewLocalDataCache(),
	}
}

// func (p *StatisticalManager) LogTimeDay() string {
// 	return utils.TransUnsignedTimeUTCFormat(SAVE_PERIOD_DAY, time.Now().UTC())
// }

func (p *StatisticalManager) LogTimeHour(t time.Time) string {
	return utils.TransUnsignedTimeUTCFormat(SAVE_PERIOD_HOUR, t)
}

func (p *StatisticalManager) InitFromStr(jsonStr, logTime string) *Statistical {

	stat := NewStatisticalFromJson(jsonStr)

	p.checkLocalDataCacheExist(logTime)

	key := fmt.Sprintf("%d", stat.AgentId())
	p.data_of_day[logTime].Add(key, stat)
	// p.data_of_date.Add("backup", logTime)
	return stat
}

func (p *StatisticalManager) checkLocalDataCacheExist(key string) {
	p.lock.Lock()
	_, ok := p.data_of_day[key]
	if !ok {
		p.data_of_day[key] = cache.NewLocalDataCache()
	}
	p.lock.Unlock()
}

func (p *StatisticalManager) NewStat(agentId int, agentName, levelCode string, t time.Time) *Statistical {
	stat := NewStatistical(agentId, agentName, levelCode)
	logTime := p.LogTimeHour(t)

	p.checkLocalDataCacheExist(logTime)

	key := fmt.Sprintf("%d", agentId)
	p.data_of_day[logTime].Add(key, stat)
	return stat
}

func (p *StatisticalManager) GetStat(agentId int, t time.Time) *Statistical {
	logTime := p.LogTimeHour(t)

	p.checkLocalDataCacheExist(logTime)

	key := fmt.Sprintf("%d", agentId)
	val, ok := p.data_of_day[logTime].Get(key)
	if ok {
		tmp := val.(*Statistical)
		return tmp
	}

	return nil
}

func (p *StatisticalManager) SetUpdateList(logtime string, update bool) {
	if update {
		p.update_list.GetOrAdd(logtime, update)
	} else {
		p.update_list.Remove(logtime)
	}
}

func (p *StatisticalManager) GetUpdateList() map[string]interface{} {
	return p.update_list.GetAll()
}

func (p *StatisticalManager) ClearUpdateList() {
	p.update_list.Clear()
}

func (p *StatisticalManager) GetCache(logtime string) cache.ILocalDataCache {
	val, ok := p.data_of_day[logtime]
	if ok {
		return val
	}

	p.data_of_day[logtime] = cache.NewLocalDataCache()
	return p.data_of_day[logtime]
}

// 輸出基本資訊(整合)
func (p *StatisticalManager) JsonOutput(levelCode string, logtime string) string {

	obj, ok := p.data_of_day[logtime]
	if ok {
		tmps := obj.GetAll()

		var activePlayer, numberBettors, numberRegistrants, oddNumber int
		var totalBetting, gameTax, platformTotalWinScore, platformTotalLoseScore, platformTotalScore float64

		for _, val := range tmps {
			data := val.(*Statistical)

			if p.checkTargetLevelCodeIsPassing(levelCode, data.levelCode) {
				activePlayer += data.ActivePlayer()
				numberBettors += data.NumberBettors()
				numberRegistrants += data.NumberRegistrants()
				oddNumber += data.OddNumber()
				totalBetting = utils.DecimalAdd(totalBetting, data.TotalBetting())
				gameTax = utils.DecimalAdd(gameTax, data.GameTax())
				platformTotalWinScore = utils.DecimalAdd(platformTotalWinScore, data.PlatformWinScore())
				platformTotalLoseScore = utils.DecimalAdd(platformTotalLoseScore, data.PlatformLoseScore())
				p := utils.DecimalSub(data.PlatformWinScore(), data.PlatformLoseScore())
				platformTotalScore = utils.DecimalAdd(platformTotalScore, p)
			}
		}

		tmpMap := make(map[string]interface{})

		tmpMap["active_player"] = activePlayer
		tmpMap["number_bettors"] = numberBettors
		tmpMap["number_registrants"] = numberRegistrants
		tmpMap["odd_number"] = oddNumber
		tmpMap["total_betting"] = totalBetting
		tmpMap["game_tax"] = gameTax
		tmpMap["platform_total_win_score"] = platformTotalWinScore
		tmpMap["platform_total_lose_score"] = platformTotalLoseScore
		tmpMap["platform_total_score"] = platformTotalScore

		return utils.ToJSON(tmpMap)
	}
	return ""
}

// 輸出基本資訊(整合)
// agent = 0, 全部統計
func (p *StatisticalManager) JsonOutputInternal(levelCode string, formatType string, startTime, endTime time.Time) string {

	logtimes := utils.GetTimeIntervalList(formatType, startTime, endTime)

	var activePlayer, numberBettors, numberRegistrants, oddNumber int
	var totalBetting, gameTax, platformTotalWinScore, platformTotalLoseScore, platformTotalScore float64

	tmpMap := make(map[string]interface{})

	tmpMap["active_player"] = 0
	tmpMap["number_bettors"] = 0
	tmpMap["number_registrants"] = 0
	tmpMap["odd_number"] = 0
	tmpMap["total_betting"] = 0
	tmpMap["game_tax"] = 0
	tmpMap["platform_total_win_score"] = 0
	tmpMap["platform_total_lose_score"] = 0
	tmpMap["platform_total_score"] = 0

	// 統計查詢時段內，不重複人數
	activePlayerList := make(map[string]int, 0)
	numberBettorsList := make(map[string]int, 0)
	numberRegistrantsList := make(map[string]int, 0)

	for _, logtime := range logtimes {

		// 從redis 內取出資料
		statData, err := p.rdb.LoadHAllValue(p.rdbIdx, logtime)
		if err != nil {
			return utils.ToJSON(tmpMap)
		}

		for _, val := range statData {
			// josn string 轉格式
			data := NewStatisticalFromJson(val)

			// 檢查權限
			if p.checkTargetLevelCodeIsPassing(levelCode, data.levelCode) {
				tmpActivePlayerList := data.GetActivePlayerList()
				for userId := range tmpActivePlayerList {
					activePlayerList[userId] = 1
				}
				tmpNumberBettorsList := data.GetNumberBettorsList()
				for userId := range tmpNumberBettorsList {
					activePlayerList[userId] = 1
					numberBettorsList[userId] = 1
				}
				tmpNumberRegistrantsList := data.GetNumberRegistrantsList()
				for userId := range tmpNumberRegistrantsList {
					numberRegistrantsList[userId] = 1
				}
				oddNumber += data.OddNumber()
				totalBetting = utils.DecimalAdd(totalBetting, data.TotalBetting())
				gameTax = utils.DecimalAdd(gameTax, data.GameTax())
				platformTotalWinScore = utils.DecimalAdd(platformTotalWinScore, data.PlatformWinScore())
				platformTotalLoseScore = utils.DecimalAdd(platformTotalLoseScore, data.PlatformLoseScore())
				p := utils.DecimalSub(data.PlatformWinScore(), data.PlatformLoseScore())
				platformTotalScore = utils.DecimalAdd(platformTotalScore, p)
			}
		}
	}

	activePlayer = len(activePlayerList)
	numberBettors = len(numberBettorsList)
	numberRegistrants = len(numberRegistrantsList)

	tmpMap["active_player"] = activePlayer
	tmpMap["number_bettors"] = numberBettors
	tmpMap["number_registrants"] = numberRegistrants
	tmpMap["odd_number"] = oddNumber
	tmpMap["total_betting"] = totalBetting
	tmpMap["game_tax"] = gameTax
	tmpMap["platform_total_win_score"] = platformTotalWinScore
	tmpMap["platform_total_lose_score"] = platformTotalLoseScore
	tmpMap["platform_total_score"] = platformTotalScore

	return utils.ToJSON(tmpMap)
}

// 只輸出不重複人數資訊
func (p *StatisticalManager) OutputInternalPartInfo(levelCode string, formatType string, startTime, endTime time.Time) map[string]interface{} {

	logtimes := utils.GetTimeIntervalList(formatType, startTime, endTime)

	var activePlayer, numberBettors int

	tmpMap := make(map[string]interface{})

	tmpMap["active_player"] = 0
	tmpMap["number_bettors"] = 0

	// 統計查詢時段內，不重複人數
	activePlayerList := make(map[string]int, 0)
	numberBettorsList := make(map[string]int, 0)

	for _, logtime := range logtimes {

		// 從redis 內取出資料
		statData, err := p.rdb.LoadHAllValue(p.rdbIdx, logtime)
		if err != nil {
			return tmpMap
		}

		for _, val := range statData {
			// josn string 轉格式
			data := NewStatisticalFromJson(val)

			// 檢查權限
			if p.checkTargetLevelCodeIsPassing(levelCode, data.levelCode) {
				tmpActivePlayerList := data.GetActivePlayerList()
				for userId := range tmpActivePlayerList {
					activePlayerList[userId] = 1
				}
				tmpNumberBettorsList := data.GetNumberBettorsList()
				for userId := range tmpNumberBettorsList {
					activePlayerList[userId] = 1
					numberBettorsList[userId] = 1
				}
			}
		}
	}

	activePlayer = len(activePlayerList)
	numberBettors = len(numberBettorsList)

	tmpMap["active_player"] = activePlayer
	tmpMap["number_bettors"] = numberBettors

	return tmpMap
}

func (p *StatisticalManager) checkTargetLevelCodeIsPassing(my string, target string) (isOk bool) {

	isOk = false
	myLen := len(my)
	targetLen := len(target)
	// level code 長度越短，權限越大
	// 被查詢目標權限大於查詢者
	// level code 格式不對不給查
	if myLen > targetLen || myLen%4 != 0 || targetLen%4 != 0 {
		return
	}
	// 開發者帳號 4碼
	// 如果檢查的是自己
	if my == target {
		isOk = true
		return
	}

	// 查詢目標如果同是開發者，就不給查
	// 平級只能查自己，其他不給查
	if targetLen > 4 {
		if my == target[:myLen] {
			isOk = true
			return
		}
	}

	return
}
