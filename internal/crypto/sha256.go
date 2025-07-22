package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(s string) string {
	hashed := sha256.Sum256([]byte(s))

	hexStr := hex.EncodeToString(hashed[:])

	return hexStr
}
