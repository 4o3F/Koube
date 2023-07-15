package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(str string) string {
	hashInstance := sha256.New()
	hashInstance.Write([]byte(str))
	result := hashInstance.Sum(nil)
	return hex.EncodeToString(result)
}
