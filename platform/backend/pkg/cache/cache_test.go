package cache_test

import (
	"backend/pkg/cache"
	"testing"
)

func TestCache(t *testing.T) {
	// load default config
	c := cache.NewLocalDataCache()

	// Add
	c.Add("1", 1111)
	c.Add("2", 2222)
	t.Logf("cache 1,2 add finish, count is: %v", c.Count())

	// Get
	a1, _ := c.Get("1")
	a2, _ := c.Get("2")
	t.Logf("cache content 1 is: %v", a1)
	t.Logf("cache content 2 is: %v", a2)

	// GetOrAdd duplicate
	t.Logf("cache count is: %v", c.Count())
	c.GetOrAdd("1", 1111)
	t.Logf("add duplicate value 1 cache count is: %v", c.Count())
	t.Logf("cache content object is: %v", c.GetAll())
	// GetOrAdd new value
	c.GetOrAdd("3", 3333)
	a3, _ := c.Get("3")
	t.Logf("cache 3 add finish, count is: %v", c.Count())
	t.Logf("cache content 3 is: %v", a3)
	t.Logf("cache content object is: %v", c.GetAll())

	// Remove
	t.Log("cache remove value 1")
	c.Remove("1")
	t.Logf("cache count is: %v", c.Count())
	t.Logf("cache content object is: %v", c.GetAll())
	// Clear
	c.Clear()
	t.Logf("cache count is: %v", c.Count())
	t.Logf("cache content object is: %v", c.GetAll())
	// GetAll
	// No testing required
}
