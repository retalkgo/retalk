package cmd

import (
	"retalk/internal/core"
	"retalk/server"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务器",
	Run: func(cmd *cobra.Command, args []string) {
		core.InitCore()
		server.Start()
	},
}
