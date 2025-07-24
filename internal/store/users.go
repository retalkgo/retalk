package store

import (
	"fmt"

	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type UserStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewUsersStore(db *gorm.DB, storeCache *StoreCache) *UserStore {
	return &UserStore{
		db:         db,
		storeCache: storeCache,
	}
}

func (s *UserStore) Create(user *model.User) error {
	err := s.db.Create(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheSet(user)

	return nil
}

func (s *UserStore) FindByID(id uint) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByIDKey, id), func() (*model.User, error) {
		var user model.User
		err := s.db.Where("id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}

		return &user, nil
	})
}

func (s *UserStore) FindByUsername(username string) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByUsernameKey, username), func() (*model.User, error) {
		var user model.User
		err := s.db.Where("username = ?", username).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
}

func (s *UserStore) FindByEmail(email string) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByEmailKey, email), func() (*model.User, error) {
		var user model.User
		err := s.db.Where("email = ?", email).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
}

func (s *UserStore) Update(user *model.User) error {
	err := s.db.Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheSet(user)

	return nil
}

func (s *UserStore) Delete(user *model.User) error {
	err := s.db.Where("id = ?", user.ID).Delete(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheDelete(user)

	return nil
}
