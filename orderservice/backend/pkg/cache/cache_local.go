package cache

import (
	"fmt"
	"sync"

	"go.uber.org/atomic"
)

type LocalDataCache struct {
	DataCache *sync.Map
	DataCount *atomic.Int32
}

func NewLocalDataCache() ILocalDataCache {
	return &LocalDataCache{
		DataCache: new(sync.Map),
		DataCount: atomic.NewInt32(0),
	}
}

func (s *LocalDataCache) Count() int {
	return int(s.DataCount.Load())
}

func (s *LocalDataCache) Add(id interface{}, a interface{}) {
	_, ok := s.DataCache.LoadOrStore(id, a)
	if !ok {
		s.DataCount.Inc()
	}
}

func (s *LocalDataCache) Get(id interface{}) (interface{}, bool) {
	return s.DataCache.Load(id)
}

/* only update value, no count

 * something updated, return true

 * no value updated, return false
 */
func (s *LocalDataCache) Update(id interface{}, a interface{}) bool {
	_, ok := s.DataCache.Load(id)
	if ok {
		s.DataCache.Store(id, a)
	}
	return ok
}

func (s *LocalDataCache) GetOrAdd(id interface{}, a interface{}) bool {
	_, ok := s.DataCache.LoadOrStore(id, a)
	if !ok {
		s.DataCount.Inc()
	}
	return ok
}

func (s *LocalDataCache) Remove(id interface{}) {
	s.DataCache.Delete(id)
	s.DataCount.Dec()
}

func (s *LocalDataCache) Clear() {
	s.DataCache.Range(func(key interface{}, value interface{}) bool {
		s.DataCache.Delete(key)
		return true
	})
	s.DataCount.Store(0)
}

func (s *LocalDataCache) GetAll() map[string]interface{} {
	temp := make(map[string]interface{}, 0)
	s.DataCache.Range(func(key interface{}, value interface{}) bool {
		// temp = append(temp, value)
		str := fmt.Sprintf("%v", key)
		temp[str] = value
		return true
	})

	return temp
}

func (s *LocalDataCache) GetInstance() *sync.Map {
	return s.DataCache
}
