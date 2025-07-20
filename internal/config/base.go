package config

import (
	"os"
	"sync"

	"github.com/creasty/defaults"
	"sigs.k8s.io/yaml"
)

type ServerConfig struct {
	Host string `yaml:"host" default:"localhost"` // 主机名
	Port int    `yaml:"port" default:"2716"`      // 端口
}

type CacheConfig struct {
	Type     string `yaml:"type" default:"memory"`
	Addr     string `yaml:"addr" default:"localhost:6379"`
	Username string `yaml:"username" default:""`
	Password string `yaml:"password" default:""`
	DB       int    `yaml:"db" default:"0"`
	TTL      int    `yaml:"ttl" default:"30"` // 缓存过期时间（分钟）
}

type LaunchConfigSchema struct {
	Dev      bool         `yaml:"dev" default:"false"` // 调试模式
	Server   ServerConfig `yaml:"server"`
	Database string       `yaml:"database" default:"sqlite://./retalk.db"` // 数据库连接字符串
	Cache    CacheConfig  `yaml:"cache"`
}

var (
	conf           *LaunchConfigSchema
	once           sync.Once
	configFilePath = "retalk.yaml"
)

func readLaunchConfig() (*LaunchConfigSchema, error) {
	launchConfig := &LaunchConfigSchema{}

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 返回空结构体以避免默认值应用失败
			return launchConfig, nil
		}
		return nil, err
	}

	err = yaml.Unmarshal(file, launchConfig)
	if err != nil {
		return nil, err
	}

	return launchConfig, nil
}

func LaunchConfig() *LaunchConfigSchema {
	once.Do(func() {
		var err error
		conf, err = readLaunchConfig()
		if err != nil {
			panic(err)
		}

		err = defaults.Set(conf)
		if err != nil {
			panic(err)
		}
	})
	return conf
}
