package cmd

import "github.com/retalkgo/retalk/internal/logger"

func Excute() {
	Init()
	if err := rootCmd.Execute(); err != nil {
		logger.Panic(err)
	}
}
