package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/retalkgo/retalk/internal/config"
)

type Claims struct {
	Username string `json:"usr"`
	Password string `json:"pwd"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)

	jwtSecret := config.Config().Server.Secret
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}

func Verify(tokenStr string) (bool, error) {
	jwtSecret := config.Config().Server.Secret
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, err
}
