package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5 - not for production use
func GetMD5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// Implement BCrypt
