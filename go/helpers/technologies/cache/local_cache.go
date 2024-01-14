package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type LocalCache struct {
	cache *cache.Cache
}

func NewLocalCache(expire time.Duration) *LocalCache {
	return &LocalCache{
		cache: cache.New(expire, 2*expire),
	}
}

func (c *LocalCache) Get(key string) (interface{}, bool) {
	value, found := c.cache.Get(key)
	return value, found
}

func (c *LocalCache) Set(key string, value interface{}) {
	c.cache.Set(key, value, cache.DefaultExpiration)
}

func (c *LocalCache) Delete(key string) {
	c.cache.Delete(key)
}

func (c *LocalCache) Clear() {
	c.cache.Flush()
}
