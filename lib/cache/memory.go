package cache

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

var (
	// DefaultEvery the clock time of recycling the expired cache items in memory.
	DefaultEvery = 60 // 1 minute
)

func init() {
	Register("memery", NewMemoryCache)
}

// MemoryItem store memory cache item.
type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

// check is expired.
func (m *MemoryItem) isExpire() bool {
	if m.lifespan == 0 {
		return false
	}
	return time.Now().Sub(m.createdTime) > m.lifespan
}

// MemoryCache is memory cache adapter.
type MemoryCache struct {
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
	Every int // run an expiration check Every clock time
}

// NewMemoryCache create a new memorycache
func NewMemoryCache() Cache {
	return &MemoryCache{items: make(map[string]*MemoryItem)}
}

// Get get cache from memory.
func (c *MemoryCache) Get(name string) interface{} {
	c.RLock()
	defer c.RUnlock()
	if item, ok := c.items[name]; ok {
		if item.isExpire() {
			return nil
		}
		return item.val
	}
	return nil
}

// GetMulti get caches from memory.
func (c *MemoryCache) GetMulti(names []string) []interface{} {
	var values []interface{}
	for _, name := range names {
		values = append(values, c.Get(name))
	}
	return values
}

// Put set a cache to memory.
func (c *MemoryCache) Put(name string, value interface{}, lifespan time.Duration) error {
	c.Lock()
	defer c.Unlock()
	c.items[name] = &MemoryItem{
		val:         value,
		createdTime: time.Now(),
		lifespan:    lifespan,
	}

	return nil
}

// Delete delete the cache in memory.
func (c *MemoryCache) Delete(name string) error {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.items[name]; !ok {
		return errors.New("key not exist")
	}
	delete(c.items, name)
	if _, ok := c.items[name]; ok {
		return errors.New("delete key error")
	}
	return nil
}

// IsExist check cache whether exists in memory.
func (c *MemoryCache) IsExist(name string) bool {
	c.RLock()
	defer c.RUnlock()
	if v, ok := c.items[name]; ok {
		return !v.isExpire()
	}

	return false
}

// Increment increment cache counter in memory.
func (c *MemoryCache) Increment(name string) error {
	c.Lock()
	defer c.Unlock()

	item, ok := c.items[name]
	if !ok {
		return errors.New("key not exist")
	}

	switch val := item.val.(type) {
	case int:
		item.val = val + 1
	case int32:
		item.val = val + 1
	case int64:
		item.val = val + 1
	case uint:
		item.val = val + 1
	case uint32:
		item.val = val + 1
	case uint64:
		item.val = val + 1
	default:
		return errors.New("item val is not (u)int (u)int32 (u)int64")
	}
	return nil
}

// Decrement decrement cache counter in memory.
func (c *MemoryCache) Decrement(name string) error {
	c.Lock()
	defer c.Unlock()

	item, ok := c.items[name]
	if !ok {
		return errors.New("key not exist")
	}

	switch val := item.val.(type) {
	case int:
		item.val = val - 1
	case int32:
		item.val = val - 1
	case int64:
		item.val = val - 1
	case uint:
		if val > 0 {
			item.val = val - 1
		} else {
			return errors.New("item val is less than 0")
		}
	case uint32:
		if val > 0 {
			item.val = val - 1
		} else {
			return errors.New("item val is less than 0")
		}
	case uint64:
		if val > 0 {
			item.val = val - 1
		} else {
			return errors.New("item val is less than 0")
		}
	default:
		return errors.New("item val is not (u)int (u)int32 (u)int64")
	}

	return nil
}

// ClearAll delete all cache in memory.
func (c *MemoryCache) ClearAll() error {
	c.Lock()
	defer c.Unlock()
	c.items = make(map[string]*MemoryItem)

	return nil
}

// StartAndGC start gc routine based on config string settings.
func (c *MemoryCache) StartAndGC(config string) error {
	var conf map[string]int
	json.Unmarshal([]byte(config), &conf)
	if _, ok := conf["interval"]; !ok {
		conf = make(map[string]int)
		conf["interval"] = DefaultEvery
	}

	dur := time.Duration(conf["interval"]) * time.Second
	c.Every = conf["interval"]
	c.dur = dur
	go c.vacuum()

	return nil
}

func (c *MemoryCache) vacuum() {
	c.RLock()
	every := c.Every
	c.RUnlock()

	if every < 1 {
		return
	}

	for {
		<-time.After(c.dur) // block, not sleep
		if c.items == nil {
			return
		}
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

// return keys which are expired.
func (c *MemoryCache) expiredKeys() []string {
	c.RLock()
	defer c.RUnlock()

	var keys []string
	for key, item := range c.items {
		if item.isExpire() {
			keys = append(keys, key)
		}
	}

	return keys
}

func (c *MemoryCache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()

	for _, key := range keys {
		delete(c.items, key)
	}
}
