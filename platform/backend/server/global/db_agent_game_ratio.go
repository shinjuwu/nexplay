package global

import (
	"backend/pkg/cache"
	"backend/pkg/utils"
	"database/sql"
	"definition"
	"fmt"
	"time"
)

/* 代理殺率設定
* tip: 不知道資料的 key 是甚麼的時候，直接使用物件內的 getKey() 產生對應資料的 key 值即可
* key 值產生的規則依照遊戲類型不同而不同
* 百人場(type:1) 沒有 agentId 的設定
 */
type AgentGameRatio struct {
	Id           string    `json:"id"`             // agentId, gameId, gameType, roomType 組合字串用_分隔
	AgentId      int       `json:"agent_id"`       // 代理id
	GameId       int       `json:"game_id"`        // 遊戲id (no search default: 0)
	GameType     int       `json:"game_type"`      // 遊戲類型(gameId/1000)
	RoomType     int       `json:"room_type"`      // 房間類型
	KillRatio    float64   `json:"kill_ratio"`     // 基礎殺率
	NewKillRatio float64   `json:"new_kill_ratio"` // 新手殺率
	ActiveNum    int       `json:"active_num"`     // 啟動人數
	Info         string    `json:"info"`           // 備註
	UpdateTime   time.Time `json:"update_time"`    // 更新時間
}

type agentGameRatioDatabaseCache struct {
	agentGameRatio          cache.ILocalDataCache
	db                      *sql.DB
	tablename               string
	defaultKillDiveInfoJson string
}

// defaultSettingJson: from storage of definition.STORAGE_KEY_GAMEKILLDIVEINFO
func NewAgentGameRatioDatabaseCache(db *sql.DB, defaultKillDiveInfoJson string) *agentGameRatioDatabaseCache {
	return &agentGameRatioDatabaseCache{
		agentGameRatio:          cache.NewLocalDataCache(),
		db:                      db,
		tablename:               "agent_game_ratio",
		defaultKillDiveInfoJson: defaultKillDiveInfoJson,
	}
}

// init local memeory data from database
func (p *agentGameRatioDatabaseCache) InitCacheFromDB() error {
	query := fmt.Sprintf(`SELECT id, agent_id, game_id, game_type, room_type, kill_ratio, new_kill_ratio, active_num, info, update_time FROM %s`, p.tablename)
	rows, err := p.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := new(AgentGameRatio)
		err := rows.Scan(&tmp.Id, &tmp.AgentId, &tmp.GameId, &tmp.GameType, &tmp.RoomType, &tmp.KillRatio,
			&tmp.NewKillRatio, &tmp.ActiveNum, &tmp.Info, &tmp.UpdateTime)
		if err != nil {
			return err
		}

		p.agentGameRatio.Add(tmp.Id, tmp)
	}

	return nil
}

func (p *agentGameRatioDatabaseCache) InitCacheFromJson(defaultKillDiveInfoJson string) error {
	tmpMap := utils.ToArrayMap([]byte(defaultKillDiveInfoJson))
	nowTime := time.Now().UTC()
	agentIds := make([]int, 0)
	query := `select id from agent`
	agentRows, err := p.db.Query(query)
	if err != nil {
		return err
	}

	defer agentRows.Close()
	for agentRows.Next() {
		var agentId int
		if err := agentRows.Scan(&agentId); err != nil {
			return err
		}

		agentIds = append(agentIds, agentId)
	}

	sqlStr := fmt.Sprintf(`INSERT INTO "public"."%s" ("id", "agent_id", "game_id", "game_type", "room_type", "kill_ratio", "new_kill_ratio", "active_num") VALUES `, p.tablename)
	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),`

	idx := 1
	vals := []interface{}{}

	for _, agentId := range agentIds {
		for _, val := range tmpMap {
			gameIdF, ok := val["GameId"].(float64)
			if !ok || gameIdF == 0 {
				continue
			}
			gameId := int(gameIdF)
			game := GameCache.Get(gameId)
			if game == nil {
				continue
			}
			roomIdF, ok := val["RoomId"].(float64)
			if !ok {
				continue
			}
			roomId := int(roomIdF)
			gameType := gameId / 1000
			// 百人場、好友房不處理
			if gameType == definition.GAME_TYPE_BAIREN ||
				gameType == definition.GAME_TYPE_FRIENDSROOM {
				continue
			}
			roomType := roomId % gameId
			killrate, ok := val["Killrate"].(float64)
			if !ok {
				continue
			}
			newKillRate, ok := val["Newkillrate"].(float64)
			if !ok {
				continue
			}
			activeNumF, ok := val["Activenum"].(float64)
			if !ok {
				continue
			}
			activeNum := int(activeNumF)

			id := p.GetKey(agentId, gameId, gameType, roomType)

			vals = append(vals, id, agentId, gameId, gameType, roomType, killrate, newKillRate, activeNum)
			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7)
			idx += 8

			tmp := &AgentGameRatio{
				Id:           id,
				AgentId:      agentId,
				GameId:       gameId,
				GameType:     gameType,
				RoomType:     roomType,
				KillRatio:    killrate,
				NewKillRatio: newKillRate,
				ActiveNum:    activeNum,
				Info:         "",
				UpdateTime:   nowTime,
			}

			p.agentGameRatio.Add(tmp.Id, tmp)
		}
	}

	agentId := definition.AGENT_ID_UNKNOW
	for _, val := range tmpMap {
		gameIdF, ok := val["GameId"].(float64)
		if !ok || gameIdF == 0 {
			continue
		}
		gameId := int(gameIdF)
		game := GameCache.Get(gameId)
		if game == nil {
			continue
		}
		roomIdF, ok := val["RoomId"].(float64)
		if !ok {
			continue
		}
		roomId := int(roomIdF)
		gameType := gameId / 1000
		// 非百人場不處理
		if gameType != definition.GAME_TYPE_BAIREN {
			continue
		}
		roomType := roomId % gameId
		killrate, ok := val["Killrate"].(float64)
		if !ok {
			continue
		}
		newKillRate, ok := val["Newkillrate"].(float64)
		if !ok {
			continue
		}
		activeNumF, ok := val["Activenum"].(float64)
		if !ok {
			continue
		}
		activeNum := int(activeNumF)

		id := p.GetKey(agentId, gameId, gameType, roomType)

		vals = append(vals, id, agentId, gameId, gameType, roomType, killrate, newKillRate, activeNum)
		sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7)
		idx += 8

		tmp := &AgentGameRatio{
			Id:           id,
			AgentId:      agentId,
			GameId:       gameId,
			GameType:     gameType,
			RoomType:     roomType,
			KillRatio:    killrate,
			NewKillRatio: newKillRate,
			ActiveNum:    activeNum,
			Info:         "",
			UpdateTime:   nowTime,
		}

		p.agentGameRatio.Add(tmp.Id, tmp)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := p.db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return err
		}
	}

	return nil
}

// 新增代理時使用，創建代理殺放資料
func (p *agentGameRatioDatabaseCache) CreateNewAgentToCacheFromJson(agentId int) error {

	tmpMap := utils.ToArrayMap([]byte(p.defaultKillDiveInfoJson))

	sqlStr := fmt.Sprintf(`INSERT INTO "public"."%s" ("id", "agent_id", "game_id", "game_type", "room_type", "kill_ratio", "new_kill_ratio", "active_num") VALUES `, p.tablename)
	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),`

	idx := 1
	vals := []interface{}{}
	// for _, agentId := range agentIds {
	for _, val := range tmpMap {
		gameIdF, ok := val["GameId"].(float64)
		if !ok || gameIdF == 0 {
			continue
		}
		gameId := int(gameIdF)
		game := GameCache.Get(gameId)
		if game == nil {
			continue
		}
		roomIdF, ok := val["RoomId"].(float64)
		if !ok {
			continue
		}
		roomId := int(roomIdF)
		gameType := gameId / 1000
		// 百人場、好友房不處理
		if gameType == definition.GAME_TYPE_BAIREN ||
			gameType == definition.GAME_TYPE_FRIENDSROOM {
			continue
		}
		roomType := roomId % gameId
		killrate, ok := val["Killrate"].(float64)
		if !ok {
			continue
		}
		newKillRate, ok := val["Newkillrate"].(float64)
		if !ok {
			continue
		}
		activeNumF, ok := val["Activenum"].(float64)
		if !ok {
			continue
		}
		activeNum := int(activeNumF)

		id := p.GetKey(agentId, gameId, gameType, roomType)

		vals = append(vals, id, agentId, gameId, gameType, roomType, killrate, newKillRate, activeNum)
		sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4, idx+5, idx+6, idx+7)
		idx += 8

		tmp := &AgentGameRatio{
			Id:         id,
			AgentId:    agentId,
			GameId:     gameId,
			GameType:   gameType,
			RoomType:   roomType,
			KillRatio:  killrate,
			Info:       "",
			UpdateTime: time.Now().UTC(),
		}

		p.agentGameRatio.Add(tmp.Id, tmp)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	//prepare the statement
	stmt, err := p.db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if stmt != nil {
		//format all vals at once
		if _, err := stmt.Exec(vals...); err != nil {
			return err
		}
	}

	return nil
}

// 新增遊戲時使用，創建所有代理新遊戲殺放資料
func (p *agentGameRatioDatabaseCache) CreateNewGameToCacheFromJson(_gameId int, newDefaultKillDiveInfoJson string) error {
	p.defaultKillDiveInfoJson = newDefaultKillDiveInfoJson

	tmpMap := utils.ToArrayMap([]byte(p.defaultKillDiveInfoJson))
	fliterMap := make([]map[string]interface{}, 0)
	for _, tmp := range tmpMap {
		gameIdF, ok := tmp["GameId"].(float64)
		if !ok || gameIdF == 0 {
			continue
		}
		gameId := int(gameIdF)
		game := GameCache.Get(gameId)
		if game == nil {
			continue
		}

		if gameId == _gameId {
			fliterMap = append(fliterMap, tmp)
		}
	}

	if len(fliterMap) == 0 {
		return fmt.Errorf("newDefaultKillDiveInfoJson has no new game info,data: %s", newDefaultKillDiveInfoJson)
	}

	fliterJson := utils.ToJSON(fliterMap)

	return p.InitCacheFromJson(fliterJson)
}

func (p *agentGameRatioDatabaseCache) DataLength() int {
	return p.agentGameRatio.Count()
}

func (p *agentGameRatioDatabaseCache) GetKey(agentId, gameId, gameType, roomType int) string {
	if gameType == definition.GAME_TYPE_BAIREN {
		agentId = definition.AGENT_ID_UNKNOW
	}

	if agentId == definition.AGENT_ID_UNKNOW {
		return fmt.Sprintf("%d_%d_%d", gameId, gameType, roomType)
	} else {
		return fmt.Sprintf("%d_%d_%d_%d", agentId, gameId, gameType, roomType)
	}
}

// insert data to local memory and database
func (p *agentGameRatioDatabaseCache) Insert(agentId, gameId, gameType, roomType, activeNum int, killRatio, newKillRatio float64, info string) (*AgentGameRatio, error) {
	id := p.GetKey(agentId, gameId, gameType, roomType)
	nowTime := time.Now().UTC()
	query := fmt.Sprintf(`INSERT INTO %s(id, agent_id, game_id, game_type, room_type, kill_ratio, new_kill_ratio, active_num info, update_time) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`, p.tablename)
	result, err := p.db.Exec(query, id, agentId, gameId, gameType, roomType, killRatio, info, nowTime)
	if err != nil {
		return nil, err
	}

	if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
		return nil, err
	}

	tmp := &AgentGameRatio{
		Id:           id,
		AgentId:      agentId,
		GameId:       gameId,
		GameType:     gameType,
		RoomType:     roomType,
		KillRatio:    killRatio,
		NewKillRatio: newKillRatio,
		ActiveNum:    activeNum,
		Info:         info,
		UpdateTime:   nowTime, // 不用很精確,所以不使用 sql returning
	}

	p.agentGameRatio.Add(tmp.Id, tmp)
	return tmp, nil
}

// get all data from local memory
func (p *agentGameRatioDatabaseCache) Select() ([]*AgentGameRatio, error) {
	tmps := make([]*AgentGameRatio, 0)
	mapTmp := p.agentGameRatio.GetAll()
	for _, val := range mapTmp {
		if tmp, ok := val.(*AgentGameRatio); ok {
			tmps = append(tmps, tmp)
		}
	}
	return tmps, nil
}

// get data from local memory
func (p *agentGameRatioDatabaseCache) SelectOne(key string) (*AgentGameRatio, bool) {
	if val, ok := p.agentGameRatio.Get(key); ok {
		if tmp, ok := val.(*AgentGameRatio); ok {
			return tmp, (tmp != nil)
		}
	}
	return nil, false
}

// get data from local memory
// return (*AgentGameRatio, bool), [select or insert object, is exist objct]
func (p *agentGameRatioDatabaseCache) SelectOrInsertOne(agentId, gameId, gameType, roomType, activeNum int, killRatio, newKillRatio float64, info string) (*AgentGameRatio, bool) {
	key := p.GetKey(agentId, gameId, gameType, roomType)
	if val, ok := p.agentGameRatio.Get(key); ok {
		if tmp, ok := val.(*AgentGameRatio); ok {
			return tmp, ok
		}
	} else {
		s, err := p.Insert(agentId, gameId, gameType, roomType, activeNum, killRatio, newKillRatio, info)
		if err == nil {
			return s, false
		}
	}
	return nil, false
}

// update data to local memory and database
func (p *agentGameRatioDatabaseCache) Update(key string, killRatio, newKillRatio float64, activeNum int, info string) error {
	nowTime := time.Now().UTC()
	tmp, ok := p.SelectOne(key)
	if ok {
		query := fmt.Sprintf(`UPDATE %s 
		SET kill_ratio=$1, new_kill_ratio=$2, active_num=$3, info=$4, update_time=$5
		WHERE id=$6`, p.tablename)
		_, err := p.db.Exec(query, killRatio, newKillRatio, activeNum, info, nowTime, key)
		if err != nil {
			return err
		}
		tmp.UpdateTime = nowTime
		tmp.KillRatio = killRatio
		tmp.NewKillRatio = newKillRatio
		tmp.ActiveNum = activeNum
		tmp.Info = info
		p.agentGameRatio.Add(tmp.Id, tmp)
		return nil
	} else {
		return fmt.Errorf("Update() has error, no data")
	}
}

// delete data to local memory and database
func (p *agentGameRatioDatabaseCache) Delete(key int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE agent_id=$1`, p.tablename)
	_, err := p.db.Exec(query, key)
	if err != nil {
		return err
	}

	p.agentGameRatio.Remove(key)
	return nil
}

func (p *agentGameRatioDatabaseCache) Clear() error {
	query := fmt.Sprintf(`DELETE FROM %s`, p.tablename)
	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	p.agentGameRatio.Clear()
	return nil
}
