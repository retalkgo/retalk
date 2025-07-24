package store

import (
	"testing"

	"github.com/retalkgo/retalk/internal/config"
	"github.com/retalkgo/retalk/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	cacheConfig := config.CacheConfig{
		Type: config.CacheTypeMemory,
	}
	dbInstance := db.GetTestDB()
	defer db.ClearTestDB()

	err := Init(dbInstance, &cacheConfig)
	assert.NoError(t, err)

	assert.NotNil(t, AppConfig)
	assert.NotNil(t, Users)
	assert.NotNil(t, Comments)
}
