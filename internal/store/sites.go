package store

import (
	"fmt"

	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type SiteStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewSitesStore(db *gorm.DB, storeCache *StoreCache) *SiteStore {
	return &SiteStore{db: db, storeCache: storeCache}
}

func (s *SiteStore) Create(site *model.Site) error {
	err := s.db.Create(site).Error
	if err != nil {
		return err
	}

	s.storeCache.SiteCacheSet(site)

	return nil
}

func (s *SiteStore) Update(site *model.Site) error {
	err := s.db.Save(site).Error
	if err != nil {
		return err
	}

	s.storeCache.SiteCacheSet(site)

	return nil
}

func (s *SiteStore) FindByID(id uint) (*model.Site, error) {
	site, err := QueryWithCache(s.storeCache, fmt.Sprintf(SiteByIDKey, id), func() (*model.Site, error) {
		var site model.Site
		err := s.db.Where("id = ?", site.ID).First(&site).Error
		return &site, err
	})

	return site, err
}

func (s *SiteStore) FindByDomain(domain string) (*model.Site, error) {
	site, err := QueryWithCache(s.storeCache, fmt.Sprintf(SiteByDomainKey, domain), func() (*model.Site, error) {
		var site model.Site
		err := s.db.Where("domain = ?", site.Domain).First(&site).Error

		return &site, err
	})

	return site, err
}

func (s *SiteStore) Delete(site *model.Site) error {
	err := s.db.Delete(site).Error
	if err != nil {
		return err
	}

	s.storeCache.SiteCacheDelete(site)

	return nil
}
