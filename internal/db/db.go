package db

import (
	"retalk/internal/entity"
	"retalk/internal/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInterface *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("retalk.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		logger.Panic(err)
	}
	db.AutoMigrate(&entity.Comment{}, &entity.Author{}, &entity.Reply{}) // 同步数据库
	dbInterface = db

}

func DB() *gorm.DB {
	return dbInterface
}
