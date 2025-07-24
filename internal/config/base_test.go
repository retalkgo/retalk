package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleConfig = `dev: true
secret: retalk
server:
  host: example.com
  port: 8080
database:
  type: postgres
  host: localhost
  port: 5432
  username: user
  password: 11111111
  dbname: dbname
cache:
  type: redis
  addr: "redis:6379"
  username: user
  password: pass
  db: 2
  ttl: 15`

var sampleConfigStruct = LaunchConfigSchema{
	Dev:    true,
	Secret: "retalk",
	Server: ServerConfig{
		Host: "example.com",
		Port: 8080,
	},
	Database: DatabaseConfig{
		Type:     "postgres",
		Host:     "localhost",
		Port:     5432,
		Username: "user",
		Password: "11111111",
		DBName:   "dbname",
	},
	Cache: CacheConfig{
		Type:     "redis",
		Addr:     "redis:6379",
		Username: "user",
		Password: "pass",
		DB:       2,
		TTL:      15,
	},
}

var defualtConfigStruct = LaunchConfigSchema{
	Dev:    false,
	Secret: "retalk",
	Server: ServerConfig{
		Host: "localhost",
		Port: 2716,
	},
	Database: DatabaseConfig{
		Type:     "postgres",
		Host:     "localhost",
		Port:     5432,
		Username: "root",
		Password: "root",
		DBName:   "retalk",
	},
	Cache: CacheConfig{
		Type:     "memory",
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		DB:       0,
		TTL:      30,
	},
}

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

	assert.Equal(t, sampleConfigStruct, *cfg)
}

func TestLaunchConfig_DefaultsApplied(t *testing.T) {
	cfgFile := prepareTestConfig(t, "retalk.test.yaml", []byte{})
	defer cleanupTestConfig(t, cfgFile)

	configReset()

	cfg := LaunchConfig()
	assert.NotNil(t, cfg)

	assert.Equal(t, defualtConfigStruct, *cfg)
}
