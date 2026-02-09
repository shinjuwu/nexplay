package global

import (
	"backend/pkg/redis"
	"backend/pkg/utils"
	table_model "backend/server/table/model"
	"encoding/json"
	"strconv"
	"strings"
)

func NewGlobalAgentCache(rdb redis.IRedisCliect, idx int, hashName string) *GlobalAgentCache {
	return &GlobalAgentCache{
		rdb:           rdb,
		rdb_idx:       idx,
		rdb_hash_name: hashName,
	}
}

type GlobalAgentCache struct {
	rdb_idx       int
	rdb_hash_name string
	rdb           redis.IRedisCliect
}

// param agent_id
func (p *GlobalAgentCache) Get(id int) *table_model.Agent {
	ret := new(table_model.Agent)
	jsonStr, err := p.rdb.LoadHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(id))
	if err == nil {
		err = json.Unmarshal([]byte(jsonStr), ret)
	}

	if ret != nil && ret.Id > 0 && err == nil {
		return ret
	}

	return nil
}

func (p *GlobalAgentCache) GetAll() []*table_model.Agent {
	agents := make([]*table_model.Agent, 0)
	kvPairs, err := p.rdb.LoadHAllValue(p.rdb_idx, p.rdb_hash_name)
	if err == nil {
		for _, jsonStr := range kvPairs {
			ret := new(table_model.Agent)
			err = json.Unmarshal([]byte(jsonStr), ret)
			if err != nil {
				continue
			}
			agents = append(agents, ret)
		}
		if len(agents) > 0 {
			return agents
		}
	}
	return nil
}

func (p *GlobalAgentCache) GetChildAgents(agentId int) []*table_model.Agent {

	agent := p.Get(agentId)
	if agent != nil {
		childAgents := make([]*table_model.Agent, 0)
		kvPairs, err := p.rdb.LoadHAllValue(p.rdb_idx, p.rdb_hash_name)
		if err == nil {
			for _, jsonStr := range kvPairs {
				targetAgent := new(table_model.Agent)
				err = json.Unmarshal([]byte(jsonStr), targetAgent)
				if err != nil {
					continue
				}
				if targetAgent.Id == agent.Id || !strings.HasPrefix(targetAgent.LevelCode, agent.LevelCode) {
					continue
				}
				childAgents = append(childAgents, targetAgent)
			}
			if len(childAgents) > 0 {
				return childAgents
			}
		}
	}
	return nil
}

func (p *GlobalAgentCache) Add(agent *table_model.Agent) {
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(agent.Id), utils.ToJSON(agent))
}

func (p *GlobalAgentCache) Adds(agent []*table_model.Agent) {

	vals := make([]string, 0)
	for _, v := range agent {
		vals = append(vals, strconv.Itoa(v.Id))
		vals = append(vals, utils.ToJSON(v))
	}
	p.rdb.StoreHValue(p.rdb_idx, p.rdb_hash_name, vals...)
}

func (p *GlobalAgentCache) Remove(agent *table_model.Agent) {
	p.rdb.DeleteHValue(p.rdb_idx, p.rdb_hash_name, strconv.Itoa(agent.Id))
}

func (p *GlobalAgentCache) RemoveAll() {
	p.rdb.Del(p.rdb_idx, p.rdb_hash_name)
}
