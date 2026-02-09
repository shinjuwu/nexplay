package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalGameCache() *GlobalGameCache {
	return &GlobalGameCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalGameCache struct {
	cache cache.ILocalDataCache
}

func (gameCache *GlobalGameCache) Get(id int) *table_model.Game {
	if game, ok := gameCache.cache.Get(id); ok {
		return game.(*table_model.Game)
	}
	return nil
}

func (gameCache *GlobalGameCache) GetAll() []*table_model.Game {
	games := make([]*table_model.Game, 0)

	kvPairs := gameCache.cache.GetAll()
	for _, value := range kvPairs {
		games = append(games, value.(*table_model.Game))
	}

	return games
}

func (gameCache *GlobalGameCache) Add(game *table_model.Game) {
	gameCache.cache.Add(game.Id, game)
}

func (gameCache *GlobalGameCache) Remove(game *table_model.Game) {
	gameCache.cache.Remove(game.Id)
}
