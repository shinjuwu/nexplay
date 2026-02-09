package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalGameRoomCache() *GlobalGameRoomCache {
	return &GlobalGameRoomCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalGameRoomCache struct {
	cache cache.ILocalDataCache
}

func (gameRoomCache *GlobalGameRoomCache) Get(id int) *table_model.GameRoom {
	if room, ok := gameRoomCache.cache.Get(id); ok {
		return room.(*table_model.GameRoom)
	}
	return nil
}

func (gameRoomCache *GlobalGameRoomCache) GetAll() []*table_model.GameRoom {
	gameRooms := make([]*table_model.GameRoom, 0)

	kvPairs := gameRoomCache.cache.GetAll()
	for _, value := range kvPairs {
		gameRooms = append(gameRooms, value.(*table_model.GameRoom))
	}

	return gameRooms
}

func (gameRoomCache *GlobalGameRoomCache) GetGameRooms(gameId int) []*table_model.GameRoom {
	gameRooms := make([]*table_model.GameRoom, 0)
	for _, gameRoom := range gameRoomCache.GetAll() {
		if gameRoom.GameId != gameId {
			continue
		}

		gameRooms = append(gameRooms, gameRoom)
	}
	return gameRooms
}

func (gameRoomCache *GlobalGameRoomCache) GetGameRoomMaps(gameId int) map[int]*table_model.GameRoom {
	gameRoomMaps := make(map[int]*table_model.GameRoom)
	for _, gameRoom := range gameRoomCache.GetGameRooms(gameId) {
		gameRoomMaps[gameRoom.Id] = gameRoom
	}
	return gameRoomMaps
}

func (gameRoomCache *GlobalGameRoomCache) Add(gameRoom *table_model.GameRoom) {
	gameRoomCache.cache.Add(gameRoom.Id, gameRoom)
}

func (gameRoomCache *GlobalGameRoomCache) Remove(gameRoom *table_model.GameRoom) {
	gameRoomCache.cache.Remove(gameRoom.Id)
}
