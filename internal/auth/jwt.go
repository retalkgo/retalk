package auth

import (
	"crypto/ed25519"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Jwt *JwtManager

type JwtClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

var jwtSigningMethod = &jwt.SigningMethodEd25519{}

type JwtManager struct {
	signingMethod jwt.SigningMethod
	privateKey    ed25519.PrivateKey
	publicKey     ed25519.PublicKey
}

func NewJwtManager(secret string) *JwtManager {
	hashedSecret := sha256.Sum256([]byte(secret))

	privateKey := ed25519.NewKeyFromSeed(hashedSecret[:])
	publicKey := privateKey.Public().(ed25519.PublicKey)

	return &JwtManager{
		signingMethod: jwtSigningMethod,
		privateKey:    privateKey,
		publicKey:     publicKey,
	}
}

func (j *JwtManager) GenerateToken(username string, isAdmin bool) (string, error) {
	claims := JwtClaims{
		Username: username,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "retalk",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	token := jwt.NewWithClaims(j.signingMethod, claims)
	tokenString, err := token.SignedString(j.privateKey)

	return tokenString, err
}

func (j *JwtManager) VerifyJwtToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (any, error) {
		return j.publicKey, nil
	})

	if err != nil && errors.Is(err, jwt.ErrTokenExpired) {
		return nil, fmt.Errorf("token 已过期")
	}

	if err != nil {
		return nil, fmt.Errorf("无效的 token")
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("无效的 token")
}
