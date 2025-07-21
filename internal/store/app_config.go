package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type AppConfigStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewAppConfigStore(db *gorm.DB, storeCache *StoreCache) *AppConfigStore {
	return &AppConfigStore{
		db:         db,
		storeCache: storeCache,
	}
}

func (s *AppConfigStore) Get(key string, dest any) error {
	var config model.AppConfigKV

	if reflect.ValueOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("dest 必须为指针")
	}

	if err := s.db.First(&config, "`key` = ?", key).Error; err != nil {
		return err
	}

	err := json.Unmarshal([]byte(config.Value), dest)
	if err != nil {
		return err
	}

	return nil
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

func applyDefault[T any](s *AppConfigStore, key string, defaultVal T) (T, error) {
	var kv model.AppConfigKV
	// 查询指定 key 的记录
	err := s.db.Where("`key` = ?", key).First(&kv).Error
	if err != nil {
		// 记录不存在则写入默认值
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b, marshalErr := json.Marshal(defaultVal)
			if marshalErr != nil {
				return defaultVal, marshalErr
			}
			kv = model.AppConfigKV{
				Key:   key,
				Value: string(b),
			}
			if createErr := s.db.Create(&kv).Error; createErr != nil {
				return defaultVal, createErr
			}
			return defaultVal, nil
		}
		// 其他数据库错误直接返回
		return defaultVal, err
	}

	// 解析 JSON 为目标类型
	var cfg T
	if unmarshalErr := json.Unmarshal([]byte(kv.Value), &cfg); unmarshalErr != nil {
		return defaultVal, unmarshalErr
	}
	return cfg, nil
}

// 获取 Gravatar 配置，若不存在则写入并返回默认值
func (s *AppConfigStore) ApplyGravatarDefault() (model.GravatarConfig, error) {
	return applyDefault(s, model.GravatarConfigKey, model.DefaultGravatarConfig)
}

// 获取所有应用配置
func (s *AppConfigStore) GetAll() (*model.AppConfig, error) {
	// 批量调用 applyDefault
	appConfig, err := QueryWithCache(s.storeCache, AppConfigKey, func() (*model.AppConfig, error) {
		gravatarCfg, err := s.ApplyGravatarDefault()
		if err != nil {
			return nil, err
		}

		return &model.AppConfig{
			Gravatar: gravatarCfg,
		}, nil
	})
	return appConfig, err
}
