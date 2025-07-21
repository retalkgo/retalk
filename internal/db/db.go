package db

import (
	"fmt"
	"os"
	"path/filepath"

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

func New(dsn string) (*gorm.DB, error) {
	dbType := getDBType(dsn)

	var db *gorm.DB
	var err error

	if dbType == "sqlite" {
		// 排除 dsn 中的 sqlite:// 前缀
		path := dsn[9:]

		if path == "" {
			panic("SQLite 数据库连接字符串不能为空")
		}

		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	}

	if dbType == "mysql" {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if dbType == "postgres" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if dbType == "unknown" {
		return nil, fmt.Errorf("无效的数据库连接字符串: %s", dsn)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func DB() *gorm.DB {
	if dbInstance == nil {
		var err error
		dbInstance, err = New(config.LaunchConfig().Database)
		MigrateModels(dbInstance)
		if err != nil {
			panic("[DB] 数据库连接失败: " + err.Error())
		}
	}

	return dbInstance
}

func GetTestDBPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dbPath := filepath.Join(wd, "reblog.test.db")
	return dbPath
}

func GetTestDB() *gorm.DB {
	testDBPath := GetTestDBPath()

	dsn := "sqlite://" + testDBPath

	db, err := New(dsn)
	if err != nil {
		panic(err)
	}

	err = MigrateModels(db)
	if err != nil {
		panic(err)
	}

	return db
}

func ClearTestDB() {
	testDBPath := GetTestDBPath()

	err := os.Remove(testDBPath)
	if err != nil {
		panic(err)
	}
}
