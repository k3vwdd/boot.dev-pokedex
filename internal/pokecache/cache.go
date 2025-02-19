package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
    createdAt   time.Time
    val     []byte
}

type Cache struct {
    cache map[string]cacheEntry
    mutex sync.Mutex
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for {
        <- ticker.C
        c.reap(time.Now(), interval)
    }
}

func (c *Cache) reap(time time.Time, last time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    for k, val := range c.cache {
        if val.createdAt.Before(time.Add(-last)) {
            delete(c.cache, k)
        }
    }
}

func NewCache(interval time.Duration) *Cache {
    cache := &Cache{
        cache: make(map[string]cacheEntry),
    }

    go cache.reapLoop(interval)
    return cache
}

func (c *Cache) Add(key string, val []byte) {
    // lock mutex before accessing map
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now(),
        val: val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    value, exists := c.cache[key]
    if exists {
        return value.val, true
    } else {
        return nil, false
    }
}


