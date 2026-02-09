package data

import (
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/pkg/cache"
	"sync"

	"go.uber.org/atomic"
)

func NewServiceInfoCache() *ServiceInfoCache {
	return &ServiceInfoCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type ServiceInfoCache struct {
	cache cache.ILocalDataCache
}

func (c *ServiceInfoCache) GetAll() map[string]*model.ServiceInfo {
	tmps := c.cache.GetAll() //map[string]interface{}

	res := make(map[string]*model.ServiceInfo)
	for key, val := range tmps {
		s, ok := val.(*model.ServiceInfo)
		if ok {
			res[key] = s
		}
	}

	return res
}

func (c *ServiceInfoCache) DataCount() int {
	return c.cache.Count()
}

func (c *ServiceInfoCache) Get(id string) *model.ServiceInfo {
	item, ok := c.cache.Get(id)
	if !ok {
		return nil
	}

	return item.(*model.ServiceInfo)
}

func (c *ServiceInfoCache) Add(data *model.ServiceInfo) {
	c.cache.Add(data.ID, data)
}

func (c *ServiceInfoCache) Remove(data *model.ServiceInfo) {
	_, ok := c.cache.Get(data.ID)
	if ok {
		c.cache.Remove(data.ID)
	}
}
