package store

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/model"
	"github.com/stretchr/testify/assert"
)

// newTestCache 创建一个新的用于测试的缓存实例。
func newTestCache(t *testing.T) *cache.Cache {
	// 创建一个内存缓存的配置
	cfg := &config.CacheConfig{
		Type: "memory",
		TTL:  30,
	}
	// 使用配置来创建新的缓存实例
	c, err := cache.New(cfg)
	assert.NoError(t, err)
	return c
}

// newTestStoreCache 创建一个新的用于测试的 StoreCache 实例
func newTestStoreCache(t *testing.T) *StoreCache {
	c := newTestCache(t)
	return NewStoreCache(c)
}

// TestNewStoreCache 测试 NewStoreCache 函数
func TestNewStoreCache(t *testing.T) {
	storeCache := newTestStoreCache(t)

	assert.NotNil(t, storeCache)
	assert.NotNil(t, storeCache.Cache)
}

// TestStoreCache 测试 StoreCache 的核心功能
func TestStoreCache(t *testing.T) {
	// 准备测试数据
	appConfig := model.AppConfig{}

	user := model.User{
		BaseModel: model.BaseModel{
			ID: 101,
		},
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	site := model.Site{
		BaseModel: model.BaseModel{
			ID: 114514,
		},
		Name:        "test site",
		Domain:      "example.com",
		Description: "This is a test site",
		Logo:        "https://example.com/logo.png",
	}

	comment := model.Comment{
		BaseModel: model.BaseModel{
			ID: 114514,
		},
		Content: "test comment",
		User: model.User{
			BaseModel: model.BaseModel{
				ID: 1919810,
			},
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		},
	}

	// 定义测试用例结构
	tests := []struct {
		name     string
		data     any
		saveFunc func(storeCache *StoreCache)
		keys     []string
		delFunc  func(storeCache *StoreCache)
	}{
		// AppConfig 缓存测试
		{
			name: "AppConfig",
			data: &appConfig,
			saveFunc: func(storeCache *StoreCache) {
				storeCache.AppConfigCacheSet(&appConfig)
			},
			keys: []string{
				AppConfigKey,
			},
			delFunc: func(storeCache *StoreCache) {
				storeCache.AppConfigCacheDelete()
			},
		},

		// User 缓存测试
		{
			name: "User",
			data: &user,
			saveFunc: func(storeCache *StoreCache) {
				storeCache.UserCacheSet(&user)
			},
			keys: []string{
				fmt.Sprintf(UserByIDKey, user.ID),
				fmt.Sprintf(UserByUsernameKey, user.Username),
				fmt.Sprintf(UserByEmailKey, user.Email),
			},
			delFunc: func(storeCache *StoreCache) {
				storeCache.UserCacheDelete(&user)
			},
		},

		// Site 缓存测试
		{
			name: "Site",
			data: &site,
			saveFunc: func(storeCache *StoreCache) {
				storeCache.SiteCacheSet(&site)
			},
			keys: []string{
				fmt.Sprintf(SiteByIDKey, site.ID),
				fmt.Sprintf(SiteByDomainKey, site.Domain),
			},
			delFunc: func(storeCache *StoreCache) {
				storeCache.SiteCacheDelete(&site)
			},
		},

		// Comment 缓存测试
		{
			name: "Comment",
			data: &comment,
			saveFunc: func(storeCache *StoreCache) {
				storeCache.CommentCacheSet(&comment)
			},
			keys: []string{
				fmt.Sprintf(CommentByIDKey, comment.ID),
			},
			delFunc: func(storeCache *StoreCache) {
				storeCache.CommentCacheDelete(&comment)
			},
		},
	}

	// 遍历并执行所有测试用例
	for _, tt := range tests {
		for _, key := range tt.keys {
			t.Run(fmt.Sprintf("%s/Key=%s", tt.name, key), func(t *testing.T) {
				storeCache := newTestStoreCache(t)

				createEmptyTypedData := func() any {
					p := reflect.ValueOf(tt.data).Elem()
					return reflect.New(p.Type()).Interface()
				}

				// 1. 测试保存前：缓存中应不存在该键
				t.Run("FindBeforeSave", func(t *testing.T) {
					data := createEmptyTypedData()
					// 直接使用底层缓存的 Get 方法进行检查
					err := storeCache.Cache.Get(key, data)
					assert.Error(t, err) // 应该返回错误，因为键不存在
				})

				// 2. 保存缓存
				t.Run("Save", func(t *testing.T) {
					tt.saveFunc(storeCache)
				})

				// 3. 测试保存后：应能从缓存中获取到正确的数据
				t.Run("FindAfterSave", func(t *testing.T) {
					data := createEmptyTypedData()
					err := storeCache.Cache.Get(key, data)
					if assert.NoError(t, err) {
						assert.EqualValues(t, tt.data, data) // 数据应该完全一致
					}
				})

				// 4. 删除缓存
				t.Run("DeleteAfterSave", func(t *testing.T) {
					tt.delFunc(storeCache)
				})

				// 5. 测试删除后：缓存中应再次不存在该键
				t.Run("FindAfterDelete", func(t *testing.T) {
					data := createEmptyTypedData()
					err := storeCache.Cache.Get(key, data)
					assert.Error(t, err)
				})
			})
		}
	}
}

// TestQueryWithCache 测试 QueryWithCache 函数
func TestQueryWithCache(t *testing.T) {
	storeCache := newTestStoreCache(t)

	assert.NotNil(t, storeCache)
	assert.NotNil(t, storeCache.Cache)

	called := 0
	fn := func() (int, error) {
		called++
		return 12345, nil
	}

	// 第一次调用，缓存未命中，函数执行
	res, err := QueryWithCache(storeCache, "test", fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res)
	assert.Equal(t, 1, called)

	// 第二次调用，缓存命中，函数不执行
	res2, err := QueryWithCache(storeCache, "test", fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res2)
	assert.Equal(t, 1, called)
}
