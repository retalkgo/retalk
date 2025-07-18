package config

import (
	"os"

	"github.com/creasty/defaults"
	"sigs.k8s.io/yaml"
)

// 启动配置（不包括应用设置）

type LaunchConfigSchema struct {
	Dev      bool   `yaml:"dev" default:"false"`                     // 调试模式
	Host     string `yaml:"host" default:"localhost"`                // 主机名
	Port     int    `yaml:"port" default:"2716"`                     // 端口
	Database string `yaml:"database" default:"sqlite://./retalk.db"` // 数据库连接字符串
}

var launchConfigInstance *LaunchConfigSchema

func readLaunchConfig() (*LaunchConfigSchema, error) {
	file, err := os.ReadFile("retalk.yaml")
	if err != nil {
		return nil, err
	}

	var launchConfig *LaunchConfigSchema
	err = yaml.Unmarshal(file, &launchConfig)
	if err != nil {
		return nil, err
	}

	return launchConfig, nil
}

func LaunchConfig() *LaunchConfigSchema {
	if launchConfigInstance == nil {
		var err error

		launchConfigInstance, err = readLaunchConfig()
		if err != nil {
			panic(err)
		}

		err = defaults.Set(launchConfigInstance)
		if err != nil {
			panic(err)
		}
	}
	return launchConfigInstance
}
