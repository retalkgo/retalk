package core

import (
	"retalk/internal/db"
)

// 挂载核心功能
func InitCore() {
	db.InitDB()
	db.InitQuery(db.DB())
}
