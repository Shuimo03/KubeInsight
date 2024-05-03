package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("dGVzdHNlY3JldEtleQo=") //TODO 这里只是做测试

// GenerateToken 生成 JWT Token
func GenerateToken(userName, Password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     userName,
		"password": Password,
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析并验证 JWT Token
func ParseToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to parse token claims")
	}

	_, ok = claims["user"].(string)
	if !ok {
		return errors.New("invalid user in token")
	}

	return nil
}
