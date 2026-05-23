package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(data []byte) string {
	sum := sha1.Sum(data)
	return hex.EncodeToString(sum[:])
}
