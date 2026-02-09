package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalAgentPermissionCache() *GlobalAgentPermissionCache {
	return &GlobalAgentPermissionCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalAgentPermissionCache struct {
	cache cache.ILocalDataCache
}

func (c *GlobalAgentPermissionCache) Get(id string) *table_model.AgentPermission {
	if agentPermission, ok := c.cache.Get(id); ok {
		return agentPermission.(*table_model.AgentPermission)
	}
	return nil
}

func (c *GlobalAgentPermissionCache) GetTemplates(accountType int) map[int]*table_model.AgentPermission {
	templates := make(map[int]*table_model.AgentPermission)
	for _, iAgentPermission := range c.cache.GetAll() {
		agentPermission := iAgentPermission.(*table_model.AgentPermission)
		if agentPermission.AgentId != -1 {
			continue
		}

		if agentPermission.AccountType <= accountType+1 {
			templates[agentPermission.AccountType] = agentPermission
		}
	}
	return templates
}

func (c *GlobalAgentPermissionCache) GetByAgentAccountType(agentId, accountType int) []*table_model.AgentPermission {
	resp := make([]*table_model.AgentPermission, 0)
	for _, iAgentPermission := range c.cache.GetAll() {
		agentPermission := iAgentPermission.(*table_model.AgentPermission)
		if agentPermission.AgentId != agentId || agentPermission.AccountType != accountType {
			continue
		}
		resp = append(resp, agentPermission)
	}
	return resp
}

func (c *GlobalAgentPermissionCache) GetByAgentId(agentId int) []*table_model.AgentPermission {
	resp := make([]*table_model.AgentPermission, 0)
	for _, iAgentPermission := range c.cache.GetAll() {
		agentPermission := iAgentPermission.(*table_model.AgentPermission)
		if agentPermission.AgentId != agentId {
			continue
		}
		resp = append(resp, agentPermission)
	}
	return resp
}

func (c *GlobalAgentPermissionCache) Add(agentPermission *table_model.AgentPermission) {
	c.cache.Add(agentPermission.Id, agentPermission)
}

func (c *GlobalAgentPermissionCache) Remove(agentPermission *table_model.AgentPermission) {
	c.cache.Remove(agentPermission.Id)
}
