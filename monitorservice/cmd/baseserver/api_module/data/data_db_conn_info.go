package data

import (
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/pkg/cache"
	"sync"

	"go.uber.org/atomic"
)

func NewDBConnInfoCache() *DBConnInfoCache {
	return &DBConnInfoCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type DBConnInfoCache struct {
	cache cache.ILocalDataCache
}

func (c *DBConnInfoCache) GetAll() map[string]*model.DBConnInfo {
	tmps := c.cache.GetAll() //map[string]interface{}

	res := make(map[string]*model.DBConnInfo)
	for key, val := range tmps {
		s, ok := val.(*model.DBConnInfo)
		if ok {
			res[key] = s
		}
	}

	return res
}

func (c *DBConnInfoCache) DataCount() int {
	return c.cache.Count()
}

func (c *DBConnInfoCache) Get(id string) *model.DBConnInfo {
	item, ok := c.cache.Get(id)
	if !ok {
		return nil
	}

	return item.(*model.DBConnInfo)
}

func (c *DBConnInfoCache) Add(data *model.DBConnInfo) {
	c.cache.Add(data.ID, data)
}

func (c *DBConnInfoCache) Remove(data *model.DBConnInfo) {
	_, ok := c.cache.Get(data.ID)
	if ok {
		c.cache.Remove(data.ID)
	}
}
