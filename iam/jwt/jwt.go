package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("dGVzdHNlY3JldEtleQo=") //TODO 这里只是做测试

// GenerateToken 生成 JWT Token
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // 过期时间为 24 小时
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析并验证 JWT Token
func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Failed to parse token claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("Invalid user ID in token")
	}

	return uint(userID), nil
}
