package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrateModels(t *testing.T) {
	db := GetTestDB()
	defer ClearTestDB()

	err := MigrateModels(db)
	assert.NoError(t, err)
}
