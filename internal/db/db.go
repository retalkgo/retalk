package db

import (
	"github.com/retalkgo/retalk/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// 由连接字符串获取数据库类型
func getDBType(dsn string) string {
	if dsn == "" {
		return "sqlite"
	}
	if dsn[:6] == "sqlite" {
		return "sqlite"
	} else if dsn[:8] == "postgres" {
		return "postgres"
	} else if dsn[:5] == "mysql" {
		return "mysql"
	}
	return "unknown"
}

func DB() *gorm.DB {
	if dbInstance == nil {
		dbType := getDBType(config.Config().Database)

		var err error

		if dbType == "sqlite" {
			// 排除 dsn 中的 sqlite:// 前缀
			dsn := config.Config().Database[9:]

			if dsn == "" {
				panic("SQLite 数据库连接字符串不能为空")
			}

			dbInstance, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		}

		if dbType == "mysql" {
			dbInstance, err = gorm.Open(mysql.Open(config.Config().Database), &gorm.Config{})
		}

		if dbType == "postgres" {
			dbInstance, err = gorm.Open(postgres.Open(config.Config().Database), &gorm.Config{})
		}

		if dbType == "unknown" {
			panic("[DB] 无效的数据库连接字符串: " + config.Config().Database)
		}

		if err != nil {
			panic("[DB] 数据库连接失败: " + err.Error())
		}
	}

	return dbInstance
}
