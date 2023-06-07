package cmd

import (
	"retalk/internal/core"
	"retalk/internal/db"
	"retalk/internal/logger"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成查询代码",
	Run: func(cmd *cobra.Command, args []string) {
		core.InitCore()
		logger.Info("正在生成查询代码")
		db.Gen()
	},
}
