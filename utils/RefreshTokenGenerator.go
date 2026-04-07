package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateRefreshTokens() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("There is an error in rand.Read", err)
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil

}
