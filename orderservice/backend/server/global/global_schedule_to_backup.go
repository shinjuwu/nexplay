package global

import (
	"backend/pkg/cache"
	table_model "backend/server/table/model"
	"sync"

	"go.uber.org/atomic"
)

func NewGlobalScheduleToBackupCache() *GlobalScheduleToBackupCache {
	return &GlobalScheduleToBackupCache{
		cache: &cache.LocalDataCache{
			DataCache: new(sync.Map),
			DataCount: atomic.NewInt32(0),
		},
	}
}

type GlobalScheduleToBackupCache struct {
	cache cache.ILocalDataCache
}

func (ScheduleToBackupCache *GlobalScheduleToBackupCache) Get(id string) *table_model.ScheduleToBackup {
	if item, ok := ScheduleToBackupCache.cache.Get(id); ok {
		return item.(*table_model.ScheduleToBackup)
	}
	return nil
}

func (ScheduleToBackupCache *GlobalScheduleToBackupCache) GetAll() []*table_model.ScheduleToBackup {
	items := make([]*table_model.ScheduleToBackup, 0)

	kvPairs := ScheduleToBackupCache.cache.GetAll()
	for _, value := range kvPairs {
		items = append(items, value.(*table_model.ScheduleToBackup))
	}

	return items
}

func (ScheduleToBackupCache *GlobalScheduleToBackupCache) Add(item *table_model.ScheduleToBackup) {
	ScheduleToBackupCache.cache.Add(item.Id, item)
}

func (ScheduleToBackupCache *GlobalScheduleToBackupCache) Remove(item *table_model.ScheduleToBackup) {
	ScheduleToBackupCache.cache.Remove(item.Id)
}
