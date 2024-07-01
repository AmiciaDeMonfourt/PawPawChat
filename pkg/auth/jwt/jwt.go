package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

// REFACTOR
var (
	signedKey = []byte("key")
)

func GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userId})

	tokenStr, err := token.SignedString(signedKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func CheckToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, keyFunc)

	if err != nil || !token.Valid {
		return err
	}

	return nil
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return signedKey, nil
}
