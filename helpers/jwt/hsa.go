package jwt

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"

	"github.com/dgrijalva/jwt-go"
	jwtGo "github.com/dgrijalva/jwt-go"
)

func HS256GenerateHashKey(length int) []byte {
	key := []byte{}
	for i := 0; i < 256; i++ {
		key = append(key, byte(rand.Intn(256)))
	}
	return key
}

func HS256Sign(key []byte, claim jwtGo.MapClaims) (string, error) {
	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func HS256Verify(key []byte, tokenString string) error {
	parts := strings.Split(tokenString, ".")

	method := jwt.SigningMethodHS256
	err := method.Verify(strings.Join(parts[0:2], "."), parts[2], key)
	return err
}

// HS256GetClaim retreives claim but not verify the signature
func HS256GetClaim(tokenString string) (string, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) < 3 {
		return "", errors.New("invalid signed token format")
	}

	var decodedCliam []byte
	decodedCliam, err := base64.RawURLEncoding.DecodeString(parts[1])
	return string(decodedCliam), err
}
