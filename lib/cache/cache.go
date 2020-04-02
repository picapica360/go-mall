package cache

import (
	"fmt"
	"time"
)

// Cache interface contains all behaviors for cache adapter.
type Cache interface {
	// get cached value by key.
	Get(key string) interface{}
	// get a batch cached values.
	GetMulti(keys []string) []interface{}
	// set cached value.
	Put(key string, val interface{}, timeout time.Duration) error
	// delete cached value by key.
	Delete(key string) error
	// check whether cached value exists.
	IsExist(key string) bool
	// increment counter.
	Increment(key string) error
	// decrement counter.
	Decrement(key string) error
	// clear all cache.
	ClearAll() error

	// start gc routine based on config string settings.
	StartAndGC(config string) error
}

// Instance is a function create a new cache instance.
type Instance func() Cache

var adapters = make(map[string]Instance)

// Register register a cache adapter by the adapter name.
func Register(name string, adapter Instance) {
	if adapters == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}

	adapters[name] = adapter
}

// NewCache create a new cache driver by adapter name and config string.
// config need to be correct JSON as string: {"interval":360}.
func NewCache(adapterName string, config string) (Cache, error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("cache: unknown adapter name %q", adapterName)
	}

	adapter := instanceFunc()
	if err := adapter.StartAndGC(config); err != nil {
		return nil, err
	}

	return adapter, nil
}
