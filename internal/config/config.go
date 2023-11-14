package config

import (
	"embed"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigSchema struct {
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	DB struct {
		Type string `json:"type"`
		DSN  string `json:"dsn"`
	} `json:"db"`
	Status struct {
		Enable bool   `json:"enable"`
		Title  string `json:"title"`
	} `json:"status"`
}

//go:embed default-config.yml
var defaultConfigFile embed.FS

var configInterface *ConfigSchema

func InitConfig() {
	configFile, err := os.Open("./retalk.yml")
	if err != nil {
		fmt.Println("未发现配置文件, 尝试加载默认配置")
		_, _ = defaultConfigFile.ReadFile("default-config.yml")
		defaultConfigFileBytes, _ := defaultConfigFile.ReadFile("default-config.yml")
		configFile, _ = os.Create("./retalk.yml")
		_, _ = configFile.Write(defaultConfigFileBytes)
		configFile, err = os.Open("./retalk.yml")
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	var config ConfigSchema
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}
	configInterface = &config
	defer configFile.Close()
}

func Config() *ConfigSchema {
	return configInterface
}
