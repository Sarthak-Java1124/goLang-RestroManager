package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})

	os.WriteFile("private.pem", privPem, 0600)

	pubBytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	pubPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	os.WriteFile("public.pem", pubPem, 0644)
}
