package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	data := "3.1415926535"
	target := "330548c742a7c77a612f6d5c2ba2b2917c1533c0cb71a163feef53efe3cbee09"

	assert.Equal(t, target, Sha256(data))
}
