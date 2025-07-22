package store

import (
	"fmt"

	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

type UsersStore struct {
	db         *gorm.DB
	storeCache *StoreCache
}

func NewUsersStore(db *gorm.DB) *UsersStore {
	return &UsersStore{
		db: db,
	}
}

func (s *UsersStore) Create(user *model.User) error {
	err := s.db.Model(&model.User{}).Create(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheSet(user)

	return nil
}

func (s *UsersStore) FindByID(id uint) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByIDKey, id), func() (*model.User, error) {
		var user model.User
		err := s.db.Model(&model.User{}).Where("id = ?", id).First(&user).Error
		if err != nil {
			return nil, err
		}

		return &user, nil
	})
}

func (s *UsersStore) FindByUsername(username string) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByUsernameKey, username), func() (*model.User, error) {
		var user model.User
		err := s.db.Model(&model.User{}).Where("username = ?", username).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
}

func (s *UsersStore) FindByEmail(email string) (*model.User, error) {
	return QueryWithCache(s.storeCache, fmt.Sprintf(UserByEmailKey, email), func() (*model.User, error) {
		var user model.User
		err := s.db.Model(&model.User{}).Where("email = ?", email).First(&user).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	})
}

func (s *UsersStore) Update(user *model.User) error {
	err := s.db.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheSet(user)

	return nil
}

func (s *UsersStore) Delete(user *model.User) error {
	err := s.db.Model(&model.User{}).Where("id = ?", user.ID).Delete(user).Error
	if err != nil {
		return err
	}

	s.storeCache.UserCacheDelete(user)

	return nil
}
