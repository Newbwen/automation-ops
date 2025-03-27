package auth

import (
	"github.com/Newbwen/automation-ops/backend/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// 生成令牌
func GenerateToken(user models.Users) (string, error) {
	claims := Claims{
		//UserID: user.ID,
		//Role:   user.Role.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Issuer:    "ops-platform",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
