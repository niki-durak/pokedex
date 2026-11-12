package pokecache

import(
	"time"
	"sync"
)

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
	
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.mu.Lock()
		for tick, val := range c.data {
			if time.Since(val.createdAt) > interval {
				delete(c.data, tick)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.data[key] = cacheEntry{
	createdAt: time.Now(),
	val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	val, ok := c.data[key]
	if ok{
		defer c.mu.Unlock()
		return val.val, true
	}
	defer c.mu.Unlock()
	return nil, false
}