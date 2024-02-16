package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheValues map[string]cacheEntry
	mut         sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.cacheValues[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()

	entry, ok := c.cacheValues[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		t := time.Now().Add(-interval)
		for key, val := range c.cacheValues {
			if val.createdAt.Before(t) {
				delete(c.cacheValues, key)
			}
		}
	}
}

func NewCache(reapInterval time.Duration) Cache {
	cache := Cache{
		cacheValues: make(map[string]cacheEntry),
		mut:         sync.RWMutex{},
	}
	go cache.ReapLoop(reapInterval)

	return cache
}
