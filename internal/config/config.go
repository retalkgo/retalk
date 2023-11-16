package config

import (
	"embed"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigSchema struct {
	Lang   string `json:"lang"`
	Server struct {
		Port int `json:"port"`
		Secret string `json:"secret"`
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
	if configInterface == nil {
		InitConfig()
	}
	return configInterface
}
