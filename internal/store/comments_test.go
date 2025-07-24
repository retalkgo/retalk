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

// setupCommentsStoreTest 用于设置评论存储的测试环境
func setupCommentsStoreTest(t *testing.T) (*CommentStore, *StoreCache) {
	testDB := db.GetTestDB()

	cacheCfg := config.CacheConfig{Type: "memory"}
	c, err := cache.New(&cacheCfg)
	require.NoError(t, err)
	storeCache := NewStoreCache(c)

	Init(testDB, &cacheCfg)

	commentsStore := NewCommentsStore(testDB, storeCache)

	return commentsStore, storeCache
}

func TestCommentsStore(t *testing.T) {
	db.GetTestDB()
	defer db.ClearTestDB()

	commentsStore, storeCache := setupCommentsStoreTest(t)

	// 预先创建一个用户用于评论关联
	user := &model.User{
		Username:    "commenter",
		Email:       "commenter@example.com",
		Password:    "password",
		HashedEmail: "hash_for_commenter",
		Nickname:    "评论者昵称",
		Website:     "https://commenter.example.com",
	}
	err := Users.Create(user)
	require.NoError(t, err)

	site := &model.Site{
		Name:   "测试站点",
		Domain: "example.com",
	}

	err = Sites.Create(site)
	require.NoError(t, err)

	comment := &model.Comment{
		UserID:  &user.ID,
		User:    *user,
		SiteID:  site.ID,
		Path:    "/test-path",
		Content: "这是一条测试评论。",
	}

	t.Run("Create Comment", func(t *testing.T) {
		err := commentsStore.Create(comment)
		require.NoError(t, err)
		assert.NotZero(t, comment.ID)

		var cachedComment model.Comment
		key := fmt.Sprintf(CommentByIDKey, comment.ID)
		err = storeCache.Cache.Get(key, &cachedComment)

		assert.NoError(t, err)
		assert.Equal(t, comment.Content, cachedComment.Content) // 断言缓存中的内容与原内容相同
	})

	t.Run("Find Comment By ID", func(t *testing.T) {
		found, err := commentsStore.FindByID(comment.ID)
		require.NoError(t, err)
		assert.Equal(t, comment.Content, found.Content)

		_, err = commentsStore.FindByID(99999) // 查找一个不存在的ID
		assert.Error(t, err)
	})

	t.Run("Find By Domain And Path", func(t *testing.T) {
		// 创建更多评论用于分页测试
		for i := range 5 {
			err := commentsStore.Create(&model.Comment{
				UserID:  &user.ID,
				User:    *user,
				SiteID:  1,
				Path:    "/test-path",
				Content: fmt.Sprintf("另一条评论 %d", i),
			})
			require.NoError(t, err)
		}

		comments, count, err := commentsStore.FindByDomainAndPath("example.com", "/test-path", 1, 10)
		require.NoError(t, err)
		assert.Len(t, comments, 6) // 断言找到6条评论 (包括之前创建的那一条)
		assert.Equal(t, int64(6), count)
	})

	t.Run("Update Comment", func(t *testing.T) {
		newContent := "这是一条更新后的评论。"
		comment.Content = newContent

		err := commentsStore.Update(comment)
		require.NoError(t, err)

		// 从数据库中查找并验证
		found, err := commentsStore.FindByID(comment.ID)
		require.NoError(t, err)
		assert.Equal(t, newContent, found.Content)

		// 从缓存中查找并验证
		var cachedComment model.Comment
		key := fmt.Sprintf(CommentByIDKey, comment.ID)
		err = storeCache.Cache.Get(key, &cachedComment)
		assert.NoError(t, err)
		assert.Equal(t, newContent, cachedComment.Content)
	})

	t.Run("Delete Comment", func(t *testing.T) {
		err := commentsStore.Delete(comment)
		require.NoError(t, err)

		// 验证数据库中已删除
		_, err = commentsStore.FindByID(comment.ID)
		assert.Error(t, err)

		// 验证缓存中也已删除
		key := fmt.Sprintf(CommentByIDKey, comment.ID)
		err = storeCache.Cache.Get(key, &model.Comment{})
		assert.Error(t, err)
	})

	t.Run("Cook Comment", func(t *testing.T) {
		t.Run("For Registered User", func(t *testing.T) {
			cookedComment, err := commentsStore.Cook(comment)
			require.NoError(t, err)

			assert.Equal(t, comment.UserID, cookedComment.UserID)
			assert.Equal(t, user.Nickname, cookedComment.NickName)
			assert.Equal(t, user.HashedEmail, cookedComment.HashedEmail)
			assert.Equal(t, user.Website, cookedComment.Website)
			assert.NotEmpty(t, cookedComment.Avatar) // 断言头像不为空
		})

		t.Run("For Guest", func(t *testing.T) {
			guestComment := &model.Comment{
				GuestName:        "游客用户",
				GuestHashedEmail: "hash_for_guest",
				GuestWebsite:     "https://guest.example.com",
				Content:          "一条来自游客的评论。",
			}

			cookedComment, err := commentsStore.Cook(guestComment)
			require.NoError(t, err)

			assert.Nil(t, cookedComment.UserID) // 游客的 UserID 应为 nil
			assert.Equal(t, guestComment.GuestName, cookedComment.NickName)
			assert.Equal(t, guestComment.GuestHashedEmail, cookedComment.HashedEmail)
			assert.Equal(t, guestComment.GuestWebsite, cookedComment.Website)
			assert.NotEmpty(t, cookedComment.Avatar)
		})
	})
}
