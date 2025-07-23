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

func setupUsersStoreTest(t *testing.T) (*UsersStore, *StoreCache) {
	testDB := db.GetTestDB()

	cacheCfg := &config.CacheConfig{Type: "memory"}
	c, err := cache.New(cacheCfg)
	require.NoError(t, err)
	storeCache := NewStoreCache(c)

	usersStore := NewUsersStore(testDB, storeCache)

	return usersStore, storeCache
}

func TestUsersStore(t *testing.T) {
	db.GetTestDB()
	defer db.ClearTestDB()

	usersStore, storeCache := setupUsersStoreTest(t)

	user := &model.User{
		Username:    "tester",
		Email:       "tester@example.com",
		Password:    "password",
		HashedEmail: "hash_for_tester",
	}

	t.Run("Create User", func(t *testing.T) {
		err := usersStore.Create(user)
		require.NoError(t, err)
		assert.NotZero(t, user.ID)

		var cachedUser model.User
		key := fmt.Sprintf(UserByIDKey, user.ID)
		err = storeCache.Cache.Get(key, &cachedUser)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, cachedUser.Username)
	})

	t.Run("Find User", func(t *testing.T) {
		found, err := usersStore.FindByID(user.ID)
		require.NoError(t, err)
		assert.Equal(t, user.Username, found.Username)

		found, err = usersStore.FindByUsername(user.Username)
		require.NoError(t, err)
		assert.Equal(t, user.ID, found.ID)

		found, err = usersStore.FindByEmail(user.Email)
		require.NoError(t, err)
		assert.Equal(t, user.ID, found.ID)

		_, err = usersStore.FindByID(99999)
		assert.Error(t, err)
	})

	t.Run("Find User with Cache Miss", func(t *testing.T) {
		missUser := &model.User{
			Username:    "miss_user",
			Email:       "miss@example.com",
			HashedEmail: "hash_for_miss_user",
		}
		err := usersStore.db.Create(missUser).Error
		require.NoError(t, err)

		missKey := fmt.Sprintf(UserByIDKey, missUser.ID)
		err = storeCache.Cache.Get(missKey, &model.User{})
		assert.Error(t, err)

		found, err := usersStore.FindByID(missUser.ID)
		require.NoError(t, err)
		assert.Equal(t, missUser.Username, found.Username)

		var cachedUser model.User
		err = storeCache.Cache.Get(missKey, &cachedUser)
		assert.NoError(t, err)
		assert.Equal(t, missUser.Username, cachedUser.Username)
	})

	t.Run("Update User", func(t *testing.T) {
		newUsername := "tester_updated"
		user.Username = newUsername

		err := usersStore.Update(user)
		require.NoError(t, err)

		found, err := usersStore.FindByID(user.ID)
		require.NoError(t, err)
		assert.Equal(t, newUsername, found.Username)

		var cachedUser model.User
		key := fmt.Sprintf(UserByIDKey, user.ID)
		err = storeCache.Cache.Get(key, &cachedUser)
		assert.NoError(t, err)
		assert.Equal(t, newUsername, cachedUser.Username)
	})

	t.Run("Delete User", func(t *testing.T) {
		err := usersStore.Delete(user)
		require.NoError(t, err)

		_, err = usersStore.FindByID(user.ID)
		assert.Error(t, err)

		key := fmt.Sprintf(UserByIDKey, user.ID)
		err = storeCache.Cache.Get(key, &model.User{})
		assert.Error(t, err)
	})
}
