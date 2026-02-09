package global

import (
	"backend/pkg/cache"
	"backend/pkg/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
	"unicode/utf8"
)

/*
CREATE TABLE "public"."agent_custom_tag_info" (
  "agent_id" int4 NOT NULL,
  "level_code" varchar(128) NOT NULL,
  "custom_tag_info" jsonb NOT NULL DEFAULT '{}',
  "update_time" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("agent_id")
)
;
*/

const (
	defaultTagInfoIndex_black_lsit = -1 // 黑名單
	defaultTagInfoIndex_high_risk  = -2 // 高風險
	defaultTagInfoIndex_kill_point = -3 // 定額追殺

	customTagInfo_length_default = 8
	customTagInfo_idx_first      = 0
)

type DefaultTagInfo struct {
	Idx int `json:"idx"` // 固定順序
}

func NewDefaultTagInfo(idx int) *DefaultTagInfo {
	return &DefaultTagInfo{
		Idx: idx,
	}
}

func NewDefaultTagInfoOutput() map[int]*DefaultTagInfo {
	tmp := make(map[int]*DefaultTagInfo, 0)
	tmp[-3] = NewDefaultTagInfo(defaultTagInfoIndex_black_lsit)
	tmp[-2] = NewDefaultTagInfo(defaultTagInfoIndex_high_risk)
	tmp[-1] = NewDefaultTagInfo(defaultTagInfoIndex_kill_point)
	return tmp
}

type CustomTagInfo struct {
	Idx   int    `json:"idx"`   // 固定順序
	Color string `json:"color"` // 自定義顏色
	Name  string `json:"name"`  // 自定義名稱
	Info  string `json:"info"`  // 自定義備註
}

func NewCustomTagInfo(idx int, color, name, info string) *CustomTagInfo {
	return &CustomTagInfo{
		Idx:   idx,
		Color: color,
		Name:  name,
		Info:  info,
	}
}

type AgentCustomTagInfo struct {
	AgentId            int                    `json:"agent_id"`        // 代理id
	LevelCode          string                 `json:"level_code"`      // 層級碼
	CustomTagInfoBytes []byte                 `json:"-"`               // byte for db
	CustomTagInfo      map[int]*CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
	UpdateTime         time.Time              `json:"update_time"`     // 更新時間
}

func NewAgentCustomTagInfo(agentId int, levelCode string) *AgentCustomTagInfo {
	return &AgentCustomTagInfo{
		AgentId:   agentId,
		LevelCode: levelCode,
	}
}

type AgentTagInfoList struct {
	DefaultTagList map[int]*DefaultTagInfo
	CustomTagInfo  map[int]*CustomTagInfo `json:"custom_tag_info"` // 自定義標籤資訊(array in json)
}

type agentCustomTagInfoDatabaseCache struct {
	databaseCache cache.ILocalDataCache
	db            *sql.DB
	tablename     string
}

// defaultSettingJson: from storage of definition.STORAGE_KEY_GAMEKILLDIVEINFO
func NewAgentCustomTagInfoDatabaseCache(db *sql.DB) *agentCustomTagInfoDatabaseCache {
	return &agentCustomTagInfoDatabaseCache{
		databaseCache: cache.NewLocalDataCache(),
		db:            db,
		tablename:     "agent_custom_tag_info",
	}
}

/* 第一次初始化使用，包含DB資料的初始化
從 DB 取 agent 的資料做資料的初始化
*/
func (p *agentCustomTagInfoDatabaseCache) InitDBAndCache() error {

	type TmpData struct {
		AgentId   int
		LevelCode string
	}

	agentObjs := make([]TmpData, 0)
	query := `select id, level_code from agent where length(level_code)>=8`
	agentRows, err := p.db.Query(query)
	if err != nil {
		return err
	}

	defer agentRows.Close()
	for agentRows.Next() {
		var obj TmpData
		if err := agentRows.Scan(&obj.AgentId, &obj.LevelCode); err != nil {
			return err
		}

		agentObjs = append(agentObjs, obj)
	}

	defaultCustomTagInfo := p.GenInitCustomTagInfo(customTagInfo_length_default)
	defaultCustomTagInfoJson := utils.ToJSON(defaultCustomTagInfo)

	initTime := time.Now().UTC()

	// INSERT INTO %s(agent_id, level_code, custom_tag_info, update_time)
	sqlStr := fmt.Sprintf(`INSERT INTO %s(agent_id, level_code, custom_tag_info, update_time) VALUES `, p.tablename)
	sqlVauleStr := `($%d,$%d,$%d,$%d),`

	idx := 1
	vals := []interface{}{}

	for _, agentObj := range agentObjs {

		vals = append(vals, agentObj.AgentId, agentObj.LevelCode, defaultCustomTagInfoJson, initTime)
		sqlStr += fmt.Sprintf(sqlVauleStr, idx, idx+1, idx+2, idx+3)
		idx += 4

		tmp := &AgentCustomTagInfo{
			AgentId:            agentObj.AgentId,
			LevelCode:          agentObj.LevelCode,
			CustomTagInfo:      defaultCustomTagInfo,
			CustomTagInfoBytes: []byte(defaultCustomTagInfoJson),
			UpdateTime:         initTime,
		}

		p.databaseCache.Add(tmp.AgentId, tmp)

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

	if p.DataLength() == 0 {
		return fmt.Errorf("InitDBAndCache() fialed, data length is zero")
	}

	return nil
}

// 新增代理時使用，創建代理自定義資料
func (p *agentCustomTagInfoDatabaseCache) CreateNewAgentToCacheFromJson(agentId int, levelCode string) error {

	p.Insert(agentId, levelCode, nil)

	return nil
}

// init local memeory data from database
func (p *agentCustomTagInfoDatabaseCache) InitCacheFromDB() error {
	query := fmt.Sprintf(`SELECT agent_id, level_code, custom_tag_info, update_time FROM %s`, p.tablename)
	rows, err := p.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		tmp := new(AgentCustomTagInfo)
		err := rows.Scan(&tmp.AgentId, &tmp.LevelCode, &tmp.CustomTagInfoBytes, &tmp.UpdateTime)
		if err != nil {
			return err
		}

		err = json.Unmarshal(tmp.CustomTagInfoBytes, &tmp.CustomTagInfo)
		if err != nil {
			return err
		}
		p.databaseCache.Add(tmp.AgentId, tmp)
	}
	return nil
}

func (p *agentCustomTagInfoDatabaseCache) DataLength() int {
	return p.databaseCache.Count()
}

// 產生初始化 CustomTagInfo 資料
func (p *agentCustomTagInfoDatabaseCache) GenInitCustomTagInfo(length int) map[int]*CustomTagInfo {
	tmps := make(map[int]*CustomTagInfo, 0)

	for i := customTagInfo_idx_first; i < length; i++ {
		tmp := &CustomTagInfo{
			Idx:   i,
			Color: "",
			Name:  "",
			Info:  "",
		}

		tmps[i] = tmp
	}

	return tmps
}

// 檢查 CustomTagInfo 資料數量與順序
func (p *agentCustomTagInfoDatabaseCache) CheckCustomTagInfoFormat(value map[int]*CustomTagInfo) bool {
	if len(value) != customTagInfo_length_default {
		return false
	}

	for i := customTagInfo_idx_first; i < customTagInfo_length_default; i++ {
		if value[i] == nil {
			return false
		} else {
			// 名字4個中文字,備註30個中文字
			// http://c.biancheng.net/view/36.html
			// Go 语言的字符串都以 UTF-8 格式保存，每个中文占用 3 个字节
			// ASCII 字符串长度使用 len() 函数。
			// Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
			if utf8.RuneCountInString(value[i].Name) > 4 || utf8.RuneCountInString(value[i].Info) > 30 || value[i].Idx != i {
				return false
			}
		}
	}

	return true
}

// insert data to local memory and database
func (p *agentCustomTagInfoDatabaseCache) Insert(agentId int, levelCode string, value map[int]*CustomTagInfo) (*AgentCustomTagInfo, error) {

	if len(value) == 0 {
		value = p.GenInitCustomTagInfo(customTagInfo_length_default)
	}

	if !p.CheckCustomTagInfoFormat(value) {
		return nil, fmt.Errorf("CheckCustomTagInfoFormat() is failed")
	}

	tmp := &AgentCustomTagInfo{
		AgentId:            agentId,
		LevelCode:          levelCode,
		CustomTagInfo:      value,
		CustomTagInfoBytes: []byte(utils.ToJSON(value)),
		UpdateTime:         time.Now().UTC(),
	}

	query := fmt.Sprintf(`INSERT INTO %s(agent_id, level_code, custom_tag_info, update_time) 
		VALUES($1, $2, $3, $4);`, p.tablename)
	result, err := p.db.Exec(query, agentId, levelCode, tmp.CustomTagInfoBytes, tmp.UpdateTime)
	if err != nil {
		return nil, err
	}

	if rowsAffectedCount, _ := result.RowsAffected(); rowsAffectedCount != 1 {
		return nil, err
	}

	p.databaseCache.Add(tmp.AgentId, tmp)
	return tmp, nil
}

// get all data from local memory
func (p *agentCustomTagInfoDatabaseCache) Select() ([]*AgentCustomTagInfo, error) {
	tmps := make([]*AgentCustomTagInfo, 0)
	mapTmp := p.databaseCache.GetAll()
	for _, val := range mapTmp {
		if tmp, ok := val.(*AgentCustomTagInfo); ok {
			tmps = append(tmps, tmp)
		}
	}
	return tmps, nil
}

// get data from local memory
func (p *agentCustomTagInfoDatabaseCache) SelectOne(key int) (*AgentCustomTagInfo, bool) {
	if val, ok := p.databaseCache.Get(key); ok {
		if tmp, ok := val.(*AgentCustomTagInfo); ok {
			return tmp, (tmp != nil)
		}
	}
	return nil, false
}

// update data to local memory and database
func (p *agentCustomTagInfoDatabaseCache) Update(agentId int, value map[int]*CustomTagInfo) error {
	tmp, ok := p.SelectOne(agentId)
	if !ok {
		return fmt.Errorf("can't update date, key: %d", tmp.AgentId)
	}

	if !p.CheckCustomTagInfoFormat(value) {
		return fmt.Errorf("CheckCustomTagInfoFormat() is failed")
	}

	tmp.CustomTagInfoBytes = []byte(utils.ToJSON(value))
	tmp.CustomTagInfo = value

	query := fmt.Sprintf(`UPDATE %s 
		SET custom_tag_info=$1, update_time=now() 
		WHERE agent_id=$2`, p.tablename)
	_, err := p.db.Exec(query, tmp.CustomTagInfoBytes, agentId)
	if err != nil {
		return err
	}

	p.databaseCache.Add(tmp.AgentId, tmp)
	return nil

}

// delete data to local memory and database
func (p *agentCustomTagInfoDatabaseCache) Delete(agentId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE agent_id=$1`, p.tablename)
	_, err := p.db.Exec(query, agentId)
	if err != nil {
		return err
	}

	p.databaseCache.Remove(agentId)
	return nil
}

func (p *agentCustomTagInfoDatabaseCache) Clear() error {
	query := fmt.Sprintf(`DELETE FROM %s`, p.tablename)
	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}

	p.databaseCache.Clear()
	return nil
}
