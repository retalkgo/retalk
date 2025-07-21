package store

import (
	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/model"
)

type StoreCache struct {
	*cache.Cache
}

const (
	AppConfigKey = "app_config"
)

func QueryWithCache[T any](storeCache *StoreCache, key string, fn func() (T, error)) (T, error) {
	if storeCache == nil {
		return fn()
	}

	return cache.QueryWithCache(storeCache.Cache, key, fn)
}

func NewStoreCache(cache *cache.Cache) *StoreCache {
	return &StoreCache{
		Cache: cache,
	}
}

func (s *StoreCache) AppConfigCacheSet(appConfig *model.AppConfig) {
	s.Set(appConfig, AppConfigKey)
}
