package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalAgentGameCache() *GlobalAgentGameCache {
	return &GlobalAgentGameCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalAgentGameCache struct {
	cache cache.ILocalDataCache
}

func (c *GlobalAgentGameCache) Get(agentId int, gameId int) *table_model.AgentGame {
	item, ok := c.cache.Get(agentId)
	if !ok {
		return nil
	}

	items := item.(*sync.Map)
	agentGame, ok := items.Load(gameId)
	if ok {
		return agentGame.(*table_model.AgentGame)
	}

	return nil
}

func (c *GlobalAgentGameCache) Add(agentGame *table_model.AgentGame) {
	item, ok := c.cache.Get(agentGame.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Store(agentGame.GameId, agentGame)
		}
	} else {
		items := new(sync.Map)
		items.Store(agentGame.GameId, agentGame)
		c.cache.Add(agentGame.AgentId, items)
	}
}

func (c *GlobalAgentGameCache) Remove(agentGame *table_model.AgentGame) {
	item, ok := c.cache.Get(agentGame.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Delete(agentGame.GameId)
		}
	}
}
