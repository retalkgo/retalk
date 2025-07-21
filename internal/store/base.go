package store

import (
	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/config"
	"gorm.io/gorm"
)

var (
	storeCache *StoreCache
	AppConfig  *AppConfigStore
)

func Init(db *gorm.DB, launchConfig *config.LaunchConfigSchema) error {
	cache, err := cache.New(&launchConfig.Cache)
	if err != nil {
		return err
	}

	storeCache = NewStoreCache(cache)

	AppConfig = NewAppConfigStore(db, storeCache)

	return nil
}
