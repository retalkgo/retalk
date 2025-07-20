package store

import "gorm.io/gorm"

type UsersStore struct {
	db *gorm.DB
}

func NewUsersStore(db *gorm.DB) *UsersStore {
	return &UsersStore{
		db: db,
	}
}

type CreateUserParams struct {
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email"`
	Website    string `json:"website"`
	IsAdmin    *bool  `json:"is_admin"`
	BadgeName  string `json:"badge_name"`
	BadgeColor string `json:"badge_color"`
	Avator     string `json:"avator"`
	Password   string `json:"password"`
}

// func (s *UsersStore) Create(params *CreateUserParams) error {
// 	if params.Avator == "" {
// 	}

// 	return s.db.Create(user).Error
// }
