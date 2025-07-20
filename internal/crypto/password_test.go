package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "super_secret_password"

	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Positive test: correct password matches hash
	isValid := CheckPasswordHash(password, hash)
	assert.True(t, isValid)

	// Negative test: wrong password fails
	isValid = CheckPasswordHash("wrong_password", hash)
	assert.False(t, isValid)
}

func TestHashPasswordUniqueHashes(t *testing.T) {
	password := "same_password"

	hash1, err1 := HashPassword(password)
	assert.NoError(t, err1)

	hash2, err2 := HashPassword(password)
	assert.NoError(t, err2)

	// Bcrypt adds salt, hashes should differ
	assert.NotEqual(t, hash1, hash2)
}
