package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBType(t *testing.T) {
	tests := []struct {
		dsn      string
		expected string
	}{
		{"", "sqlite"},
		{"sqlite:///tmp/test.db", "sqlite"},
		{"postgres://user:pass@localhost/db", "postgres"},
		{"mysql://user:pass@localhost/db", "mysql"},
		{"unknown://foo", "unknown"},
	}

	for _, tt := range tests {
		result := getDBType(tt.dsn)
		assert.Equal(t, tt.expected, result, "dsn: %s", tt.dsn)
	}
}