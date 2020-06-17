package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	mainCache               = cache.New(5*time.Minute, 10*time.Minute)
	Duration  time.Duration = 5 * time.Minute
)

func Set(key string, value interface{}) {
	mainCache.Set(key, value, Duration)
}

func Get(key string) (value interface{}, f bool) {
	value, f = mainCache.Get(key)
	if !f {
		return nil, f
	}
	mainCache.Set(key, value, Duration)
	return value, f
}
