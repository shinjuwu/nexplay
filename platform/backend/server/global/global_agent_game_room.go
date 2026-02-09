package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalAgentGameRoomCache() *GlobalAgentGameRoomCache {
	return &GlobalAgentGameRoomCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalAgentGameRoomCache struct {
	cache cache.ILocalDataCache
}

func (c *GlobalAgentGameRoomCache) Get(agentId, gameRoomId int) *table_model.AgentGameRoom {
	item, ok := c.cache.Get(agentId)
	if !ok {
		return nil
	}

	items := item.(*sync.Map)
	agentGameRoom, ok := items.Load(gameRoomId)
	if ok {
		return agentGameRoom.(*table_model.AgentGameRoom)
	}

	return nil
}

func (c *GlobalAgentGameRoomCache) Add(agentGameRoom *table_model.AgentGameRoom) {
	item, ok := c.cache.Get(agentGameRoom.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Store(agentGameRoom.GameRoomId, agentGameRoom)
		}
	} else {
		items := new(sync.Map)
		items.Store(agentGameRoom.GameRoomId, agentGameRoom)
		c.cache.Add(agentGameRoom.AgentId, items)
	}
}

func (c *GlobalAgentGameRoomCache) Remove(agentGameRoom *table_model.AgentGameRoom) {
	item, ok := c.cache.Get(agentGameRoom.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Delete(agentGameRoom.GameRoomId)
		}
	}
}
