package utils

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	userId primitive.ObjectID `json:"_id"`
	Email  string
	jwt.RegisteredClaims
}

var privateKey *rsa.PrivateKey

func GenerateJWTToken(userId primitive.ObjectID, email string) string {

	claims := JWTPayload{
		Email:  email,
		userId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.Hex(),
			Issuer:    "api-service",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedKey, err := tokens.SignedString(privateKey)
	if err != nil {
		log.Fatal("There was an error signing jwt strings")
		return ""
	}
	return signedKey

}
