package db

import (
	"retalk/internal/query"

	"gorm.io/gorm"
)

func InitQuery(db *gorm.DB) {
	query.SetDefault(db)
}
