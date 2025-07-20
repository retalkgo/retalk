package store

import (
	"encoding/json"

	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type AppConfigStore struct {
	db *gorm.DB
}

func NewAppConfigStore(db *gorm.DB) *AppConfigStore {
	return &AppConfigStore{
		db: db,
	}
}

func (s *AppConfigStore) Get(key string) (*model.AppConfigKV, error) {
	var config model.AppConfigKV
	if err := s.db.First(&config, "`key` = ?", key).Error; err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *AppConfigStore) Set(key string, value any) error {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	config := model.AppConfigKV{
		Key:   key,
		Value: string(jsonBytes),
	}

	if err := s.db.Save(&config).Error; err != nil {
		return err
	}

	return nil
}
