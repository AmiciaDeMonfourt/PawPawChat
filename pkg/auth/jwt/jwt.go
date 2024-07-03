package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// REFACTOR
var (
	signedKey = []byte("key")
)

func GenerateToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_id": userId})

	tokenStr, err := token.SignedString(signedKey)
	if err != nil {
		return "", status.Error(codes.InvalidArgument, "token didn't pass the signature: "+err.Error())
	}

	return tokenStr, nil
}

func checkToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return status.Error(codes.InvalidArgument, "failed to parse token: "+err.Error())
	}

	if !token.Valid {
		return status.Error(codes.InvalidArgument, "token is invalid")
	}

	return nil
}

func keyFunc(*jwt.Token) (interface{}, error) {
	return signedKey, nil
}

func CheckAuth(tokenStr string) error {
	return checkToken(tokenStr)
}

func ExtractUserId(tokenStr string) (uint64, error) {
	if err := checkToken(tokenStr); err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	userId, exists := claims["user_id"].(float64)
	if !exists {
		return 0, status.Error(codes.NotFound, "user id not found in token")
	}

	return uint64(userId), nil
}
