package gravatar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGravatarURL(t *testing.T) {
	baseURL := "https://www.gravatar.com/avatar/"
	email := "test@example.com"

	url := GetGravatarURL(baseURL, email)
	expectedURL := "https://www.gravatar.com/avatar/test@example.com"

	assert.Equal(t, url, expectedURL)
}
