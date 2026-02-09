package data

import (
	"monitorservice/cmd/baseserver/api_module/model"
	"monitorservice/pkg/cache"
	"sync"

	"go.uber.org/atomic"
)

func NewGameCache() *GameCache {
	return &GameCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GameCache struct {
	cache cache.ILocalDataCache
}

func (c *GameCache) GetAll() map[string]*model.Game {
	tmps := c.cache.GetAll() //map[string]interface{}

	res := make(map[string]*model.Game)
	for key, val := range tmps {
		s, ok := val.(*model.Game)
		if ok {
			res[key] = s
		}
	}

	return res
}

func (c *GameCache) DataCount() int {
	return c.cache.Count()
}

func (c *GameCache) Get(id string) *model.Game {
	item, ok := c.cache.Get(id)
	if !ok {
		return nil
	}

	return item.(*model.Game)
}

func (c *GameCache) Add(data *model.Game) {
	c.cache.Add(data.ID, data)
}

func (c *GameCache) Remove(data *model.Game) {
	_, ok := c.cache.Get(data.ID)
	if ok {
		c.cache.Remove(data.ID)
	}
}
