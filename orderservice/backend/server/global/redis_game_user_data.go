package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	"encoding/json"
	"strconv"
)

type AgentDataOfGameUser struct {
	GameUserId int    `json:"-"`
	AgentId    int    `json:"agent_id"`
	LevelCode  string `json:"level_code"`
}

func NewGlobalGameUserCache(rdb redis.IRedisCliect, idx int, hashName string) *GlobalAgentDataOfGameUserCache {
	return &GlobalAgentDataOfGameUserCache{
		rdb:           rdb,
		rdb_idx:       idx,
		rdb_hash_name: hashName,
	}
}

type GlobalAgentDataOfGameUserCache struct {
	rdb_idx       int
	rdb_hash_name string
	rdb           redis.IRedisCliect
}

// param game_user_id
func (p *GlobalAgentDataOfGameUserCache) Get(id int) *AgentDataOfGameUser {
	ret := new(AgentDataOfGameUser)
	jsonStr, err := p.rdb.LoadHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(id))
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), ret)
	}

	if ret != nil && ret.AgentId > 0 && ret.LevelCode != "" && err == nil {
		return ret
	}

	return nil
}

// func (p *GlobalAgentDataOfGameUserCache) GetAll() []*AgentDataOfGameUser {
// 	return nil
// }

// func (p *GlobalAgentDataOfGameUserCache) GetChildAgents(agentId int) []*AgentDataOfGameUser {
// 	return nil
// }

func (p *GlobalAgentDataOfGameUserCache) Add(agent *AgentDataOfGameUser) {
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(agent.GameUserId), utils.ToJSON(agent))
}

func (p *GlobalAgentDataOfGameUserCache) Adds(agent []*AgentDataOfGameUser) {

	vals := make([]string, 0)
	for _, v := range agent {
		vals = append(vals, strconv.Itoa(v.GameUserId))
		vals = append(vals, utils.ToJSON(v))
	}
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, vals...)
}

func (p *GlobalAgentDataOfGameUserCache) Remove(id int) {
	p.rdb.DeleteHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(id))
}

func (p *GlobalAgentDataOfGameUserCache) RemoveAll() {
	p.rdb.Del(p.rdb_idx, p.rdb_hash_name)
}
