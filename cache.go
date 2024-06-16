package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mut     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Time) Cache {
	return Cache{}
}

func (c *Cache) Add(key string, val []byte) {
	return
}

func (c *Cache) Get(key string, val []byte, found bool) cacheEntry {
	return cacheEntry{}
}
