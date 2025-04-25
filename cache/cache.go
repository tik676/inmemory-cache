package cache

import (
	"sync"
	"time"
)

type T interface{}

type item struct {
	value   T
	expires time.Time
}

type Cache struct {
	mu    sync.RWMutex
	items map[string]item
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]item),
		ttl:   ttl,
	}
	go cache.regularClearCache()
	return cache
}

func (c *Cache) Set(key string, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
}

func (c *Cache) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.items[key]
	if ok {
		return v.value, true
	}

	return nil, false
}

func (c *Cache) regularClearCache() {
	for {
		time.Sleep(1 * time.Second)

		c.mu.Lock()
		for k, v := range c.items {
			if time.Now().After(v.expires) {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]item)
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}
