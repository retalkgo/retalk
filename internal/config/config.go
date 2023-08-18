package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

var configInterface *ConfigSchema

type ConfigSchema struct {
	Lang  string            `yaml:"lang"`
	Email EmailConfigSchema `yaml:"email"`
}

type EmailConfigSchema struct {
	Enable   bool   `yaml:"enable"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

func InitConfig() {
	file, _ := os.Open("retalk.yml")
	defer file.Close()

	// 读取文件内容
	data, _ := io.ReadAll(file)

	// 解析 YAML 数据
	var config ConfigSchema
	yaml.Unmarshal(data, &config)
	configInterface = &config
}

func Config() *ConfigSchema {
	return configInterface
}
