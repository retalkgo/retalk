package store

import (
	"fmt"
	"testing"

	"github.com/retalkgo/retalk/internal/cache"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"
	"github.com/retalkgo/retalk/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupSitesStoreTest 用于设置站点存储的测试环境
func setupSitesStoreTest(t *testing.T) (*SitesStore, *StoreCache) {
	testDB := db.GetTestDB()

	cacheCfg := config.CacheConfig{Type: "memory"}
	c, err := cache.New(&cacheCfg)
	require.NoError(t, err)
	storeCache := NewStoreCache(c)

	Init(testDB, &cacheCfg)

	sitesStore := NewSitesStore(testDB, storeCache)

	return sitesStore, storeCache
}

func TestSitesStore(t *testing.T) {
	db.GetTestDB()
	defer db.ClearTestDB()

	sitesStore, storeCache := setupSitesStoreTest(t)

	site := &model.Site{
		Name:   "测试站点",
		Domain: "example.com",
	}

	t.Run("Create Site", func(t *testing.T) {
		err := sitesStore.Create(site)
		require.NoError(t, err)
		assert.NotZero(t, site.ID)

		// 验证缓存
		var cachedSite model.Site
		key := fmt.Sprintf(SiteByIDKey, site.ID)
		err = storeCache.Cache.Get(key, &cachedSite)

		assert.NoError(t, err)
		assert.Equal(t, site.Name, cachedSite.Name)
	})

	t.Run("Find Site By ID", func(t *testing.T) {
		// 第一次查找，应从数据库加载并写入缓存
		found, err := sitesStore.FindByID(site.ID)
		require.NoError(t, err)
		assert.Equal(t, site.Name, found.Name)

		// 第二次查找，应直接从缓存命中
		foundFromCache, err := sitesStore.FindByID(site.ID)
		require.NoError(t, err)
		assert.Equal(t, site.Name, foundFromCache.Name)

		// 查找一个不存在的ID
		_, err = sitesStore.FindByID(99999)
		assert.Error(t, err) // 断言应该会发生错误
	})

	t.Run("Find Site By Domain", func(t *testing.T) {
		// 第一次查找，应从数据库加载并写入缓存
		found, err := sitesStore.FindByDomain("example.com")
		require.NoError(t, err)
		assert.Equal(t, site.Name, found.Name)

		// 第二次查找，应直接从缓存命中
		foundFromCache, err := sitesStore.FindByDomain("example.com")
		require.NoError(t, err)
		assert.Equal(t, site.Name, foundFromCache.Name)

		// 查找一个不存在的域名
		_, err = sitesStore.FindByDomain("non-existent-domain.com")
		assert.Error(t, err)
	})

	t.Run("Update Site", func(t *testing.T) {
		newName := "更新后的站点名称"
		site.Name = newName

		err := sitesStore.Update(site)
		require.NoError(t, err)

		// 从数据库中查找并验证
		found, err := sitesStore.FindByID(site.ID)
		require.NoError(t, err)
		assert.Equal(t, newName, found.Name)

		// 从缓存中查找并验证
		var cachedSite model.Site
		key := fmt.Sprintf(SiteByIDKey, site.ID)
		err = storeCache.Cache.Get(key, &cachedSite)
		assert.NoError(t, err)
		assert.Equal(t, newName, cachedSite.Name)
	})

	t.Run("Delete Site", func(t *testing.T) {
		err := sitesStore.Delete(site)
		require.NoError(t, err)

		// 验证数据库中已删除
		_, err = sitesStore.FindByID(site.ID)
		assert.Error(t, err)

		// 验证缓存中也已删除
		key := fmt.Sprintf(SiteByIDKey, site.ID)
		err = storeCache.Cache.Get(key, &model.Site{})
		assert.Error(t, err, "从缓存中获取已删除的站点应该返回错误")
	})
}
