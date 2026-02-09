package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalAdminUserCache() *GlobalAdminUserCache {
	return &GlobalAdminUserCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalAdminUserCache struct {
	cache cache.ILocalDataCache
}

func (c *GlobalAdminUserCache) Get(agentId int, userName string) *table_model.AdminUser {
	item, ok := c.cache.Get(agentId)
	if !ok {
		return nil
	}

	items := item.(*sync.Map)
	user, ok := items.Load(userName)
	if ok {
		return user.(*table_model.AdminUser)
	}

	return nil
}

func (c *GlobalAdminUserCache) GetAgentAdminUsers(agentId int) []*table_model.AdminUser {
	adminUsers := make([]*table_model.AdminUser, 0)

	item, ok := c.cache.Get(agentId)
	if !ok {
		return adminUsers
	}

	items := item.(*sync.Map)
	items.Range(func(key interface{}, value interface{}) bool {
		adminUsers = append(adminUsers, value.(*table_model.AdminUser))
		return true
	})

	return adminUsers
}

func (c *GlobalAdminUserCache) Add(adminUser *table_model.AdminUser) {
	item, ok := c.cache.Get(adminUser.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Store(adminUser.Username, adminUser)
		}
	} else {
		items := new(sync.Map)
		items.Store(adminUser.Username, adminUser)
		c.cache.Add(adminUser.AgentId, items)
	}
}

func (c *GlobalAdminUserCache) Remove(adminUser *table_model.AdminUser) {
	item, ok := c.cache.Get(adminUser.AgentId)
	if ok {
		items, ok := item.(*sync.Map)
		if ok {
			items.Delete(adminUser.Username)
		}
	}
}
