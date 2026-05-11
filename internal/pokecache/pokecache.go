package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entry:    make(map[string]cacheEntry),
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entry[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for range ticker.C {
		c.mu.Lock()

		for key, entry := range c.entry {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entry, key)
			}
		}

		c.mu.Unlock()
	}
}
