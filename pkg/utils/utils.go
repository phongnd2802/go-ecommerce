package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomSecretKeyBase64(size int) string {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(key)
}