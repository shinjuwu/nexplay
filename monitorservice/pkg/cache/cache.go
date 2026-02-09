package cache

import "sync"

/*
local data cache interface.
*/
type ILocalDataCache interface {
	Count() int
	Add(key interface{}, item interface{})
	Get(key interface{}) (interface{}, bool)
	GetOrAdd(key interface{}, item interface{}) bool
	Remove(id interface{})
	GetAll() map[string]interface{}
	Clear()
	GetInstance() *sync.Map
}
