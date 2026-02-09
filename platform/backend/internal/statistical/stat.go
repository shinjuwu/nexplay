package statistical

import (
	"backend/pkg/cache"
	"backend/pkg/utils"
	"sync"
)

// 數據總覽
type Statistical struct {
	agentId           int                   // 代理id
	agentName         string                // 代理名稱
	levelCode         string                // 層級碼
	activePlayer      cache.ILocalDataCache // 活躍玩家(不重複) [userId(string), trigger count(int)]
	numberRegistrants cache.ILocalDataCache // 註冊人數(不重複) [userId(string), trigger count(int)]
	numberBettors     cache.ILocalDataCache // 投注人數(不重複) [userId(string), trigger count(int)]
	oddNumber         int                   // 注單數
	totalBetting      float64               // 總投注
	gameTax           float64               // 遊戲抽水
	platformWinScore  float64               // 平台總贏分數
	platformLoseScore float64               // 平台總輸分數
	isNeedUpdate      bool                  // 是否需要更新
	lock              *sync.Mutex
}

func NewStatistical(anentId int, agentName, levelCode string) *Statistical {

	return &Statistical{
		agentId:           anentId,
		agentName:         agentName,
		levelCode:         levelCode,
		activePlayer:      cache.NewLocalDataCache(),
		numberRegistrants: cache.NewLocalDataCache(),
		numberBettors:     cache.NewLocalDataCache(),
		oddNumber:         0,
		totalBetting:      float64(0),
		gameTax:           float64(0),
		platformWinScore:  float64(0),
		platformLoseScore: float64(0),
		lock:              new(sync.Mutex),
		isNeedUpdate:      true,
	}
}

func NewStatisticalFromJson(json string) *Statistical {

	tmpMap := utils.ToMap([]byte(json))

	tmp := &Statistical{
		agentId:           0,
		agentName:         "",
		levelCode:         "",
		activePlayer:      cache.NewLocalDataCache(),
		numberRegistrants: cache.NewLocalDataCache(),
		numberBettors:     cache.NewLocalDataCache(),
		oddNumber:         0,
		totalBetting:      float64(0),
		gameTax:           float64(0),
		platformWinScore:  float64(0),
		platformLoseScore: float64(0),
		lock:              new(sync.Mutex),
		isNeedUpdate:      false,
	}

	if agentId, ok := tmpMap["agent_id"]; ok {
		tmp.agentId = int(agentId.(float64))
	}
	if agentName, ok := tmpMap["agent_name"]; ok {
		tmp.agentName = agentName.(string)
	}
	if levelCode, ok := tmpMap["level_code"]; ok {
		tmp.levelCode = levelCode.(string)
	}

	for key, val := range tmpMap["active_player_list"].(map[string]interface{}) {
		tmp.activePlayer.Add(key, val)
	}
	for key, val := range tmpMap["number_bettors_list"].(map[string]interface{}) {
		tmp.numberBettors.Add(key, val)
	}
	for key, val := range tmpMap["number_registrants_list"].(map[string]interface{}) {
		tmp.numberRegistrants.Add(key, val)
	}

	if oddNumber, ok := tmpMap["odd_number"]; ok {
		tmp.oddNumber = int(oddNumber.(float64))
	}
	if totalBetting, ok := tmpMap["total_betting"]; ok {
		tmp.totalBetting = totalBetting.(float64)
	}
	if gameTax, ok := tmpMap["game_tax"]; ok {
		tmp.gameTax = gameTax.(float64)
	}
	if platformWinScore, ok := tmpMap["platform_win_score"]; ok {
		tmp.platformWinScore = platformWinScore.(float64)
	}
	if platformLoseScore, ok := tmpMap["platform_lose_score"]; ok {
		tmp.platformLoseScore = platformLoseScore.(float64)
	}

	return tmp
}

func (p *Statistical) AgentId() int {
	return p.agentId
}

func (p *Statistical) AgentName() string {
	return p.agentName
}

func (p *Statistical) LevelCode() string {
	return p.levelCode
}

func (p *Statistical) ActivePlayer() int {
	return p.activePlayer.Count()
}

func (p *Statistical) NumberRegistrants() int {
	return p.numberRegistrants.Count()
}

func (p *Statistical) NumberBettors() int {
	return p.numberBettors.Count()
}

func (p *Statistical) OddNumber() int {
	return p.oddNumber
}

func (p *Statistical) TotalBetting() float64 {
	return p.totalBetting
}

func (p *Statistical) GameTax() float64 {
	return p.gameTax
}

func (p *Statistical) PlatformWinScore() float64 {
	return p.platformWinScore
}

func (p *Statistical) PlatformLoseScore() float64 {
	return p.platformLoseScore
}

func (p *Statistical) IsNeedUpdate() bool {
	return p.isNeedUpdate
}

func (p *Statistical) SetNeedUpdate(isNeed bool) {
	p.lock.Lock()
	p.isNeedUpdate = isNeed
	p.lock.Unlock()
}

// 活躍玩家(不重複)
func (p *Statistical) ActivePlayerAdd(userId int) {
	p.lock.Lock()
	repeat := 0
	count, ok := p.activePlayer.Get(userId)
	if ok {
		repeat = count.(int) + 1
		p.activePlayer.Update(userId, repeat)
	} else {
		repeat = 1
		p.activePlayer.Add(userId, repeat)
	}
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 取得活躍玩家不重複列表
func (p *Statistical) GetActivePlayerList() map[string]interface{} {
	p.lock.Lock()
	tmp := p.activePlayer.GetAll()
	p.lock.Unlock()

	return tmp
}

// 註冊人數(不重複)
func (p *Statistical) NumberRegistrantsAdd(userId int) {
	p.lock.Lock()
	repeat := 0
	count, ok := p.numberRegistrants.Get(userId)
	if ok {
		repeat = count.(int) + 1
		p.numberRegistrants.Update(userId, repeat)
	} else {
		repeat = 1
		p.numberRegistrants.Add(userId, repeat)
	}
	// p.numberRegistrants.Update(userId, repeat)
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 取得註冊人數不重複列表
func (p *Statistical) GetNumberRegistrantsList() map[string]interface{} {
	p.lock.Lock()
	tmp := p.numberRegistrants.GetAll()
	p.lock.Unlock()

	return tmp
}

// 投注人數(不重複)
func (p *Statistical) NumberBettorsAdd(userId int) {
	p.lock.Lock()
	repeat := 0
	count, ok := p.numberBettors.Get(userId)
	if ok {
		repeat = count.(int) + 1
		p.numberBettors.Update(userId, repeat)
	} else {
		repeat = 1
		p.numberBettors.Add(userId, repeat)
	}
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 取得投注人數不重複列表
func (p *Statistical) GetNumberBettorsList() map[string]interface{} {
	p.lock.Lock()
	tmp := p.numberBettors.GetAll()
	p.lock.Unlock()

	return tmp
}

// 注單數
func (p *Statistical) OddNumberAdd(order int) {
	p.lock.Lock()
	p.oddNumber += order
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 總投注
func (p *Statistical) TotalBettingAdd(bet float64) {
	p.lock.Lock()
	p.totalBetting = utils.DecimalAdd(p.totalBetting, bet)
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 遊戲抽水
func (p *Statistical) GameTaxAdd(tax float64) {
	p.lock.Lock()
	p.gameTax = utils.DecimalAdd(p.gameTax, tax)
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 平台總贏/輸分數
func (p *Statistical) PlatformTotalScoreAdd(result string, score float64) {
	p.lock.Lock()
	switch result {
	case STAT_RESULT_WIN:
		p.platformWinScore = utils.DecimalAdd(p.platformWinScore, score)
	case STAT_RESULT_LOSE:
		p.platformLoseScore = utils.DecimalAdd(p.platformLoseScore, score)
	}
	p.isNeedUpdate = true
	p.lock.Unlock()
}

// 輸出基本資訊
func (p *Statistical) JsonOutput() string {
	p.lock.Lock()
	defer p.lock.Unlock()
	tmpMap := make(map[string]interface{})

	tmpMap["agent_id"] = p.agentId
	tmpMap["agent_name"] = p.agentName
	tmpMap["level_code"] = p.levelCode
	tmpMap["active_player"] = p.activePlayer.Count()
	tmpMap["number_bettors"] = p.numberBettors.Count()
	tmpMap["number_registrants"] = p.numberRegistrants.Count()
	tmpMap["odd_number"] = p.oddNumber
	tmpMap["total_betting"] = p.totalBetting
	tmpMap["game_tax"] = p.gameTax
	tmpMap["platform_total_score"] = utils.DecimalSub(p.platformWinScore, p.platformLoseScore)

	return utils.ToJSON(tmpMap)
}

/* 輸出包含用戶列表

{
  "active_player_list": {
    "1": 2,
    "2": 1
  },
  "game_tax": 0,
  "number_bettors_list": {},
  "number_registrants_list": {},
  "odd_number": 0,
  "platform_total_score": 0,
  "total_betting": 0,
  "active_player": 2,
  "number_bettors": 0,
  "number_registrants": 0,
  "platform_win_score": 0,
  "platform_lose_score": 0
}
*/
func (p *Statistical) JsonOutputDetail() string {
	p.lock.Lock()
	defer p.lock.Unlock()
	tmpMap := make(map[string]interface{})

	tmpMap["agent_id"] = p.agentId
	tmpMap["agent_name"] = p.agentName
	tmpMap["level_code"] = p.levelCode
	tmpMap["active_player"] = p.activePlayer.Count()
	tmpMap["number_bettors"] = p.numberBettors.Count()
	tmpMap["number_registrants"] = p.numberRegistrants.Count()
	tmpMap["active_player_list"] = p.activePlayer.GetAll()
	tmpMap["number_bettors_list"] = p.numberBettors.GetAll()
	tmpMap["number_registrants_list"] = p.numberRegistrants.GetAll()
	tmpMap["odd_number"] = p.oddNumber
	tmpMap["total_betting"] = p.totalBetting
	tmpMap["game_tax"] = p.gameTax
	tmpMap["platform_win_score"] = p.platformWinScore
	tmpMap["platform_lose_score"] = p.platformLoseScore
	tmpMap["platform_total_score"] = utils.DecimalSub(p.platformWinScore, p.platformLoseScore)

	return utils.ToJSON(tmpMap)
}

func (p *Statistical) Reset() {
	p.activePlayer.Clear()
	p.numberRegistrants.Clear()
	p.numberBettors.Clear()
	p.oddNumber = int(0)
	p.totalBetting = float64(0)
	p.gameTax = float64(0)
	p.platformWinScore = float64(0)
	p.platformLoseScore = float64(0)
	p.isNeedUpdate = true
}
