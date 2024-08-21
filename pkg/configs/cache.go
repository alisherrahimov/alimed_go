package configs

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func InitCache() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func SetItem(key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}

func GetItem(key string) (interface{}, bool) {
	return c.Get(key)
}
