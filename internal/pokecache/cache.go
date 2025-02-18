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
        var keysToDelete []string
        c.mutex.Lock()
        for key, value := range c.cache {
            elapsedTime := time.Since(value.createdAt)
            if elapsedTime > interval {
                keysToDelete = append(keysToDelete, key)
            }
        }
        c.mutex.Unlock()
        for _, key := range keysToDelete {
            c.mutex.Lock()
            delete(c.cache, key)
            c.mutex.Unlock()
        }
    }
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        cache: make(map[string]cacheEntry),
    }

    go c.reapLoop(interval)
    return c
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


