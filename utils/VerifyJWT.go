package utils

import (
	"crypto/rsa"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(tokenString string, publicKey *rsa.PublicKey) (*JWTPayload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTPayload{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodRS256.Alg() {
				return nil, fmt.Errorf("invalid algorithm")
			}
			return publicKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTPayload)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
