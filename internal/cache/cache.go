package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	return
}

func (c *Cache) Get(key string, val []byte, found bool) cacheEntry {
	return cacheEntry{}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			timeSinceEntryCreated := time.Now().Sub(entry.createdAt)
			if timeSinceEntryCreated > c.interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}