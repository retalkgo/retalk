package store

import (
	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type SitesStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewSitesStore(db *gorm.DB, storeCache *StoreCache) *SitesStore {
	return &SitesStore{db: db, storeCache: storeCache}
}

func (s *SitesStore) Create(site *model.Site) error {
	err := s.db.Create(site).Error
	if err != nil {
		return err
	}

	s.storeCache.SiteCacheSet(site)
	return nil
}
