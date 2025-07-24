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

func New(dbConf config.DatabaseConfig) (*gorm.DB, error) {
	dbType := dbConf.Type

	var db *gorm.DB
	var err error

	if dbType == "sqlite" {
		// 排除 dsn 中的 sqlite:// 前缀
		path := dbConf.DBName

		if path == "" {
			panic("SQLite 数据库连接字符串不能为空")
		}

		db, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	}

	if dbType == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConf.Username,
			dbConf.Password,
			dbConf.Host,
			dbConf.Port,
			dbConf.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if dbType == "postgres" {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			dbConf.Host,
			dbConf.Username,
			dbConf.Password,
			dbConf.DBName,
			dbConf.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	dbConf := config.DatabaseConfig{
		Type:   "sqlite",
		DBName: testDBPath,
	}

	db, err := New(dbConf)
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
