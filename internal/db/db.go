package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dbInterface *gorm.DB

func DB() *gorm.DB {
	if dbInterface == nil {
		db, err := gorm.Open(sqliteDriver("demo.sql"), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
		dbInterface = db
	}
	return dbInterface
}

func sqliteDriver(dsn string) gorm.Dialector {
	return sqlite.Open(dsn)
}

func mysqlDriver(dsn string) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN: dsn,
	})
}

func postgresDriver(dsn string) gorm.Dialector {
	return postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	})
}

func sqlserverDriver(dsn string) gorm.Dialector {
	return sqlserver.New(sqlserver.Config{
		DSN: dsn,
	})
}