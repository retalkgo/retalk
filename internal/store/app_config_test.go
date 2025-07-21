package store

import (
	"encoding/json"
	"testing"

	"github.com/retalkgo/retalk/internal/db"
	"github.com/retalkgo/retalk/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type DummyConfig struct {
	Name  string
	Count int
}

func setupTestDB() *gorm.DB {
	dbConn := db.GetTestDB()
	return dbConn
}

func TestAppConfigStore_SetAndGet(t *testing.T) {
	dbConn := setupTestDB()
	defer db.ClearTestDB()

	s := NewAppConfigStore(dbConn, nil)

	key := "dummy"
	val := DummyConfig{Name: "test", Count: 42}
	// 测试 Set
	err := s.Set(key, val)
	assert.NoError(t, err)

	// 测试 Get
	var result DummyConfig
	err = s.Get(key, &result)
	assert.NoError(t, err)
	assert.Equal(t, val, result)
}

func TestApplyDefault_NewRecord(t *testing.T) {
	dbConn := setupTestDB()
	defer db.ClearTestDB()

	s := NewAppConfigStore(dbConn, nil)

	key := "new_key"
	defaultVal := DummyConfig{Name: "def", Count: 1}

	got, err := applyDefault(s, key, defaultVal)
	assert.NoError(t, err)
	assert.Equal(t, defaultVal, got)

	// 验证数据库中已创建记录
	var kv model.AppConfigKV
	err = dbConn.First(&kv, "`key` = ?", key).Error
	assert.NoError(t, err)
	var stored DummyConfig
	err = json.Unmarshal([]byte(kv.Value), &stored)
	assert.NoError(t, err)
	assert.Equal(t, defaultVal, stored)
}

func TestApplyDefault_ExistingRecord(t *testing.T) {
	dbConn := setupTestDB()
	defer db.ClearTestDB()

	s := NewAppConfigStore(dbConn, nil)

	key := "exists"
	original := DummyConfig{Name: "orig", Count: 10}
	b, err := json.Marshal(original)
	assert.NoError(t, err)
	entry := model.AppConfigKV{Key: key, Value: string(b)}
	err = dbConn.Create(&entry).Error
	assert.NoError(t, err)

	got, err := applyDefault(s, key, DummyConfig{Name: "def", Count: 2})
	assert.NoError(t, err)
	assert.Equal(t, original, got)
}

func TestApplyGravatarDefault(t *testing.T) {
	dbConn := setupTestDB()
	defer db.ClearTestDB()

	s := NewAppConfigStore(dbConn, nil)

	cfg, err := s.ApplyGravatarDefault()
	assert.NoError(t, err)
	assert.Equal(t, model.DefaultGravatarConfig, cfg)

	// 验证记录存在
	var kv model.AppConfigKV
	err = dbConn.First(&kv, "`key` = ?", model.GravatarConfigKey).Error
	assert.NoError(t, err)
}
