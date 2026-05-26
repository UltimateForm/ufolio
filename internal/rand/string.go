package rand

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomStr(length int) (string, error) {
	b := make([]byte, length/2)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// unsure we need to encode to hex, could just return b as-is tbh
	return hex.EncodeToString(b), nil
}
