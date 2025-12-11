package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateCode() (string, error) {
	b := make([]byte, 8)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
