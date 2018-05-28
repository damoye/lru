package lru

import (
	"container/list"
	"sync"
)

// Cache is an LRU Cache. It is safe for concurrent access.
// Refer to https://github.com/golang/groupcache/tree/master/lru
type Cache struct {
	maxEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
	mutex      sync.Mutex
}

type entry struct {
	key   interface{}
	value interface{}
}

// newCache creates a new cache.
func newCache(maxEntries int) *Cache {
	return &Cache{
		maxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key interface{}, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.ll.Len() > c.maxEntries {
		ele = c.ll.Back()
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
	}
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return nil, false
}
