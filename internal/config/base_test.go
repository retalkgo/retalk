package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleConfig = `dev: true
server:
  host: "example.com"
  port: 8080
database: "postgres://user:pass@localhost/db"
cache:
  type: "redis"
  addr: "redis:6379"
  username: "user"
  password: "pass"
  db: 2
  ttl: 15`

var originalConfigPath string

func init() {
	originalConfigPath = configFilePath
}

func writeConfig(path string, data []byte) error {
	return os.WriteFile(path, data, 0o600)
}

func prepareTestConfig(t *testing.T, filename string, content []byte) string {
	dir, err := os.Getwd()
	assert.NoError(t, err)

	fullPath := filepath.Join(dir, filename)
	err = writeConfig(fullPath, content)
	assert.NoError(t, err)

	configFilePath = fullPath
	return fullPath
}

func cleanupTestConfig(t *testing.T, path string) {
	err := os.Remove(path)
	assert.NoError(t, err)
	configFilePath = originalConfigPath
}

func configReset() {
	conf = nil
	once = sync.Once{}
}

func TestLaunchConfig_LoadAndDefaults(t *testing.T) {
	cfgFile := prepareTestConfig(t, "retalk.test.yaml", []byte(sampleConfig))
	defer cleanupTestConfig(t, cfgFile)

	configReset()

	cfg := LaunchConfig()
	assert.NotNil(t, cfg)

	assert.Equal(t, true, cfg.Dev)
	assert.Equal(t, "example.com", cfg.Server.Host)
	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "postgres://user:pass@localhost/db", cfg.Database)
	assert.Equal(t, "redis", cfg.Cache.Type)
	assert.Equal(t, "redis:6379", cfg.Cache.Addr)
	assert.Equal(t, "user", cfg.Cache.Username)
	assert.Equal(t, "pass", cfg.Cache.Password)
	assert.Equal(t, 2, cfg.Cache.DB)
	assert.Equal(t, 15, cfg.Cache.TTL)
}

func TestLaunchConfig_DefaultsApplied(t *testing.T) {
	cfgFile := prepareTestConfig(t, "retalk.test.yaml", []byte{})
	defer cleanupTestConfig(t, cfgFile)

	configReset()

	cfg := LaunchConfig()
	assert.NotNil(t, cfg)

	assert.Equal(t, false, cfg.Dev)
	assert.Equal(t, "localhost", cfg.Server.Host)
	assert.Equal(t, 2716, cfg.Server.Port)
	assert.Equal(t, "sqlite://./retalk.db", cfg.Database)
	assert.Equal(t, "memory", cfg.Cache.Type)
	assert.Equal(t, "localhost:6379", cfg.Cache.Addr)
	assert.Equal(t, "", cfg.Cache.Username)
	assert.Equal(t, "", cfg.Cache.Password)
	assert.Equal(t, 0, cfg.Cache.DB)
	assert.Equal(t, 30, cfg.Cache.TTL)
}
