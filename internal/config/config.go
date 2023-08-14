package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

var configInterface *ConfigSchema

type ConfigSchema struct {
	Lang string `yaml:"lang"`
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
