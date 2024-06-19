package cache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
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
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	ce, ok := c.entries[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}
	return ce.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
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
