package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache {
		cacheEntries: make(map[string]cacheEntry),
		mu: sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add (key string, value []byte) {
	c.mu.Lock()
    defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get (key string) ([]byte, bool) {
	c.mu.Lock()
    defer c.mu.Unlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop (interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.cacheEntries {
		if now.Sub(entry.createdAt) > interval {
			delete(c.cacheEntries, key)
		}
	}
}