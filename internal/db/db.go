package db

import (
	"github.com/retalkgo/retalk/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var dbInterface *gorm.DB
var dbDriver gorm.Dialector

func initDB() {
	c := config.Config().DB
	switch c.Type {
	case "sqlite":
		dbDriver = sqliteDriver(c.DSN)
	case "mysql":
		dbDriver = mysqlDriver(c.DSN)
	case "postgres":
		dbDriver = postgresDriver(c.DSN)
	case "sqlserver":
		dbDriver = sqlserverDriver(c.DSN)
	default:
		dbDriver = sqliteDriver(c.DSN)
	}
	db, err := gorm.Open(dbDriver, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	dbInterface = db
}

func DB() *gorm.DB {
	if dbInterface == nil {
		initDB()
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
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
}

func sqlserverDriver(dsn string) gorm.Dialector {
	return sqlserver.New(sqlserver.Config{
		DSN: dsn,
	})
}
