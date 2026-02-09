package global

import (
	"backend/pkg/cache"
	"backend/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type GameIcon struct {
	GameId int `json:"gameId"` // game id (union id)
	Rank   int `json:"rank"`   // 排行 (since from 0)
	Hot    int `json:"hot"`    // 熱門 (0:none, 1:熱門)
	Newest int `json:"newest"` // 最新 (0:none, 1:最新)
	Push   int `json:"push"`   // 推廣大圖 (0:none, 1:第一張大圖, 2:第二章大圖.....)
}

// json string -> array point of GameIcon
func transIconList(defaultJson []byte) []*GameIcon {
	defaultGameIconList := make([]*GameIcon, 0)
	defaultGameIconListMap := utils.ToArrayMap(defaultJson)
	if len(defaultGameIconListMap) <= 0 {
		return nil
	}

	for _, v := range defaultGameIconListMap {
		jsonData := utils.ToJSON(v)

		gi := new(GameIcon)

		err := json.Unmarshal([]byte(jsonData), gi)
		if err != nil {
			return nil
		}
		defaultGameIconList = append(defaultGameIconList, gi)
	}

	return defaultGameIconList
}

type AgentGameIconList struct {
	AgentId      int         `json:"agent_id"`       // 代理id
	LevelCode    string      `json:"level_code"`     // 層級碼
	IsAdmin      bool        `json:"is_admin"`       // 是否為總代理預設值(只能有一個)
	IsDefault    bool        `json:"is_default"`     // 是否讀取預設值
	GameIconList []*GameIcon `json:"game_icon_list"` // 遊戲 icon 設定
	UpdateTime   time.Time   `json:"update_time"`    // 更新時間
	CreateTime   time.Time   `json:"create_time"`    // 創建時間
}

type agentGameIconListDatabaseCache struct {
	localCache  cache.ILocalDataCache
	db          *sql.DB
	tablename   string
	defaultJson string
}

// defaultSettingJson: from storage of definition.STORAGE_KEY_ADMINGAMEICONLISTDEFAULT
func NewAgentGameIconListDatabaseCache(db *sql.DB, defaultJson string) *agentGameIconListDatabaseCache {
	return &agentGameIconListDatabaseCache{
		localCache:  cache.NewLocalDataCache(),
		db:          db,
		tablename:   "agent_game_icon_list",
		defaultJson: defaultJson,
	}
}

func (p *agentGameIconListDatabaseCache) TransIconList(defaultJson []byte) []*GameIcon {
	return transIconList(defaultJson)
}

// init local memeory data from database
func (p *agentGameIconListDatabaseCache) InitCacheFromDB() error {
	query := fmt.Sprintf(`SELECT agent_id, level_code, is_admin, is_default, icon_list, update_time, create_time FROM %s`, p.tablename)
	rows, err := p.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := new(AgentGameIconList)
		var gameIconList []byte
		err := rows.Scan(&tmp.AgentId, &tmp.LevelCode, &tmp.IsAdmin, &tmp.IsDefault, &gameIconList, &tmp.UpdateTime,
			&tmp.CreateTime)
		if err != nil {
			return err
		}

		tmp.GameIconList = p.TransIconList(gameIconList)

		p.localCache.Add(tmp.LevelCode, tmp)
	}

	return nil
}

// init local memeory data and database from default json
func (p *agentGameIconListDatabaseCache) InitCacheAndDBFromDefaultJson() error {
	defaultGameIconList := p.TransIconList([]byte(p.defaultJson))

	agentIds := make([]int, 0)
	levelCodes := make([]string, 0)
	query := `select id, level_code from agent`
	agentRows, err := p.db.Query(query)
	if err != nil {
		return err
	}

	defer agentRows.Close()
	for agentRows.Next() {
		var agentId int
		var levelCode string
		if err := agentRows.Scan(&agentId, &levelCode); err != nil {
			return err
		}

		agentIds = append(agentIds, agentId)
		levelCodes = append(levelCodes, levelCode)
	}

	sqlStr := fmt.Sprintf(`INSERT INTO %s (agent_id, level_code, is_admin, is_default, icon_list) VALUES `, p.tablename)
	sqlVauleStr := `($%d,$%d,$%d,$%d,$%d),`

	idx := 1
	vals := []interface{}{}

	for i, agentId := range agentIds {

		isAdmin := false
		isDefault := false

		if len(levelCodes[i]) == 4 {
			isAdmin = true
		} else {
			isDefault = true
		}
		vals = append(vals, agentId, levelCodes[i], isAdmin, isDefault, p.defaultJson)
		sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4)
		idx += 5

		tmp := &AgentGameIconList{
			AgentId:    agentId,
			LevelCode:  levelCodes[i],
			IsAdmin:    isAdmin,
			IsDefault:  isDefault,
			UpdateTime: time.Now().UTC(),
			CreateTime: time.Now().UTC(),
		}

		copy(tmp.GameIconList, defaultGameIconList)

		p.localCache.Add(tmp.LevelCode, tmp)

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

// init local memeory data and database when add new game
func (p *agentGameIconListDatabaseCache) InitCacheAndDBAddNew(newGameIconList []*GameIcon) error {
	if err := p.InitCacheFromDB(); err != nil {
		return err
	}

	agentGameIconListMap := p.localCache.GetAll()
	if len(agentGameIconListMap) > 0 {

		// update all GameIconList of agent
		for _, val := range agentGameIconListMap {
			// agentId := utils.StringToInt(str, -1)
			agentGameIconList := val.(*AgentGameIconList)
			agentGameIconList.GameIconList = append(agentGameIconList.GameIconList, newGameIconList...)
		}

		// clear table
		p.Clear()

		// insert all of new
		sqlStr := fmt.Sprintf(`INSERT INTO %s (agent_id, level_code, is_admin, is_default, icon_list) VALUES `, p.tablename)
		sqlVauleStr := `($%d,$%d,$%d,$%d,$%d),`

		idx := 1
		vals := []interface{}{}

		for _, val := range agentGameIconListMap {
			agentGameIconList := val.(*AgentGameIconList)

			vals = append(vals, agentGameIconList.AgentId, agentGameIconList.LevelCode, agentGameIconList.IsAdmin, agentGameIconList.IsDefault,
				utils.ToJSON(agentGameIconList.GameIconList))
			sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3, idx+4)
			idx += 5

			p.localCache.Add(agentGameIconList.LevelCode, agentGameIconList)
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
	}

	return nil
}

// 新增代理時使用，創建代理自定義資料
func (p *agentGameIconListDatabaseCache) CreateNewAgentToCacheFromJson(parentAgentId, agentId int, levelCode string, isAdmin, isDefault bool) error {
	if v, ok := p.localCache.Get(parentAgentId); ok {
		parentAgentGameIconList := v.(*AgentGameIconList)
		p.Insert(agentId, levelCode, isAdmin, isDefault, utils.ToJSON(parentAgentGameIconList.GameIconList))
	} else {
		p.Insert(agentId, levelCode, isAdmin, isDefault, p.defaultJson)
	}

	return nil
}

func (p *agentGameIconListDatabaseCache) DataLength() int {
	return p.localCache.Count()
}

// insert data to local memory and database
func (p *agentGameIconListDatabaseCache) Insert(agentId int, levelCode string, isAdmin, isDefault bool, iconList string) (*AgentGameIconList, error) {
	nowTime := time.Now().UTC()
	query := fmt.Sprintf(`INSERT INTO %s(agent_id, level_code, is_admin, is_default, icon_list, update_time, create_time) 
	VALUES($1, $2, $3, $4, $5, $6, $7);`, p.tablename)
	result, err := p.db.Exec(query, agentId, levelCode, isAdmin, isDefault, iconList, nowTime, nowTime)
	if err != nil {
		return nil, err
	}

	if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
		return nil, err
	}

	tmp := &AgentGameIconList{
		AgentId:    agentId,
		LevelCode:  levelCode,
		IsAdmin:    isAdmin,
		IsDefault:  isDefault,
		UpdateTime: nowTime,
		CreateTime: nowTime,
	}

	tmp.GameIconList = p.TransIconList([]byte(iconList))

	p.localCache.Add(tmp.LevelCode, tmp)
	return tmp, nil
}

// get all data from local memory
func (p *agentGameIconListDatabaseCache) Select() ([]*AgentGameIconList, error) {
	tmps := make([]*AgentGameIconList, 0)
	mapTmp := p.localCache.GetAll()
	for _, val := range mapTmp {
		if tmp, ok := val.(*AgentGameIconList); ok {
			tmps = append(tmps, tmp)
		}
	}
	return tmps, nil
}

// get data from local memory
func (p *agentGameIconListDatabaseCache) SelectOne(key string) (*AgentGameIconList, bool) {
	if val, ok := p.localCache.Get(key); ok {
		if tmp, ok := val.(*AgentGameIconList); ok {
			return tmp, (tmp != nil)
		}
	}
	return nil, false
}

// update data to local memory and database
func (p *agentGameIconListDatabaseCache) Update(key string, isDefault bool, gameIconList string) error {
	nowTime := time.Now().UTC()

	tmp, ok := p.SelectOne(key)
	if ok {
		query := fmt.Sprintf(`UPDATE %s 
		SET is_default=$1, icon_list=$2, update_time=$3
		WHERE level_code=$4`, p.tablename)
		_, err := p.db.Exec(query, isDefault, gameIconList, nowTime, key)
		if err != nil {
			return err
		}

		tmp.IsDefault = isDefault
		tmp.UpdateTime = nowTime
		tmp.GameIconList = p.TransIconList([]byte(gameIconList))

		p.localCache.Add(tmp.AgentId, tmp)
		return nil
	} else {
		return fmt.Errorf("Update() has error, no data")
	}
}

// delete data to local memory and database
func (p *agentGameIconListDatabaseCache) Delete(key string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE level_code=$1`, p.tablename)
	_, err := p.db.Exec(query, key)
	if err != nil {
		return err
	}

	p.localCache.Remove(key)
	return nil
}

func (p *agentGameIconListDatabaseCache) Clear() error {
	query := fmt.Sprintf(`DELETE FROM %s`, p.tablename)
	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	p.localCache.Clear()
	return nil
}
