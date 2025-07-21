package store

import (
	"testing"

	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	launchConfig := &config.LaunchConfigSchema{
		Database: db.GetTestDBPath(),
		Cache: config.CacheConfig{
			Type: config.CacheTypeMemory,
		},
	}
	dbInstance := db.GetTestDB()
	defer db.ClearTestDB()

	err := Init(dbInstance, launchConfig)
	assert.NoError(t, err)

	assert.NotNil(t, AppConfig)
}
