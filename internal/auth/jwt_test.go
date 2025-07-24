package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJwtManager(t *testing.T) {
	jwtManager := NewJwtManager("test-secret")

	assert.NotNil(t, jwtManager)
}

func TestJwt(t *testing.T) {
	jwtManager := NewJwtManager("test-secret")

	token, err := jwtManager.GenerateToken("test-username", true)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := jwtManager.VerifyJwtToken(token)

	assert.NoError(t, err)
	assert.Equal(t, "test-username", claims.Username)
	assert.True(t, claims.IsAdmin)

	claims, err = jwtManager.VerifyJwtToken("invalid-token")

	assert.Error(t, err)
	assert.Nil(t, claims)
}
