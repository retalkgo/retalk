package core

import (
	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"
	"github.com/retalkgo/retalk/internal/i18n"
)

// 挂载核心功能
func InitCore() {
	config.InitConfig()
	db.InitDB()
	db.InitQuery(db.DB())
	lang := config.Config().Lang
	i18n.InitI18n(lang)
}
