package config

import (
	"os"
	"strconv"

	"github.com/creasty/defaults"
)

// 启动配置（不包括应用设置）

type ConfigSchema struct {
	Dev      bool   `default:"false"`                // 调试模式
	Host     string `default:"localhost"`            // 主机名
	Port     int    `default:"2716"`                 // 端口
	Database string `default:"sqlite://./retalk.db"` // 数据库连接字符串
}

var configInstance *ConfigSchema

func Config() *ConfigSchema {
	if configInstance == nil {
		port := 2716 // default port
		if portStr := os.Getenv("RETALK_PORT"); portStr != "" {
			if p, err := strconv.Atoi(portStr); err == nil {
				port = p
			} else {
				panic("无效的端口号:" + portStr)
			}
		}
		configInstance = &ConfigSchema{
			Dev:      os.Getenv("RETALK_DEV") == "true",
			Port:     port,
			Database: os.Getenv("RETALK_DATABASE"),
		}
		defaults.Set(configInstance)
	}
	return configInstance
}
