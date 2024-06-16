package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	ce, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return ce.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			timeSinceEntryCreated := time.Now().Sub(entry.createdAt)
			if timeSinceEntryCreated > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
