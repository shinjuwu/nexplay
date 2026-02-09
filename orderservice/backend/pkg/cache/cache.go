package cache

import "sync"

/*
local data cache interface.
*/
type ILocalDataCache interface {
	Count() int
	Add(key interface{}, item interface{})
	// get value by key
	Get(key interface{}) (interface{}, bool)
	/* only update value, no count
	 * something updated, return true
	 * no value updated, return false
	 */
	Update(key interface{}, item interface{}) bool
	GetOrAdd(key interface{}, item interface{}) bool
	Remove(id interface{})
	GetAll() map[string]interface{}
	Clear()
	GetInstance() *sync.Map
}
