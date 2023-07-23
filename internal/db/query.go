package db

import (
	"github.com/retalkgo/retalk/internal/query"

	"gorm.io/gorm"
)

func InitQuery(db *gorm.DB) {
	query.SetDefault(db)
}
