package cmd

import (
	"retalk/internal/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "retalk",
	Short: "Retalk " + version.Version + "-" + version.CommitHash + " 一个快速的, 便捷的自托管评论系统",
}

func Init() {
	rootCmd.AddCommand(startCmd, genCmd)
}
