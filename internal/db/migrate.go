package db

import (
	"github.com/retalkgo/retalk/internal/model"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.AppConfigKV{},
		&model.User{},
		&model.Site{},
	)

	return err
}
