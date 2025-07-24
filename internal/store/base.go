package store

import (
	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/config"
	"gorm.io/gorm"
)

var (
	storeCache *StoreCache
	AppConfig  *AppConfigStore
	Users      *UserStore
	Sites      *SiteStore
	Comments   *CommentStore
)

func Init(db *gorm.DB, cacheConfig *config.CacheConfig) error {
	cache, err := cache.New(cacheConfig)
	if err != nil {
		return err
	}

	storeCache = NewStoreCache(cache)

	AppConfig = NewAppConfigStore(db, storeCache)
	Users = NewUsersStore(db, storeCache)
	Sites = NewSitesStore(db, storeCache)
	Comments = NewCommentsStore(db, storeCache)

	return nil
}
