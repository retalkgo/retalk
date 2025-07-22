package store

import (
	"fmt"

	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/model"
)

type StoreCache struct {
	*cache.Cache
}

const (
	// AppConfig
	AppConfigKey = "app_config"

	// User
	UserByIDKey       = "user#id=%d"
	UserByUsernameKey = "user#username=%s"
	UserByEmailKey    = "user#email=%s"
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

func (s *StoreCache) AppConfigCacheDelete() {
	s.Delete(AppConfigKey)
}

func (s *StoreCache) UserCacheSet(user *model.User) {
	s.Set(user,
		fmt.Sprintf(UserByIDKey, user.ID),
		fmt.Sprintf(UserByUsernameKey, user.Username),
		fmt.Sprintf(UserByEmailKey, user.Email),
	)
}

func (s *StoreCache) UserCacheDelete(user *model.User) {
	s.Delete(
		fmt.Sprintf(UserByIDKey, user.ID),
		fmt.Sprintf(UserByUsernameKey, user.Username),
		fmt.Sprintf(UserByEmailKey, user.Email),
	)
}
