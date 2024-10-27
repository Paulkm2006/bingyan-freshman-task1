package utils

import (
	"bingyan-freshman-task0/internal/config"
	"log"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UID   int   `json:"uid"`
	Admin bool  `json:"admin"`
	Exp   int64 `json:"exp"`
}

func (claim JWTClaims) Valid() error {
	if jwt.TimeFunc().Unix() > claim.Exp {
		log.Println(jwt.TimeFunc().Unix())
		log.Println(claim.Exp)
		return jwt.NewValidationError("token expired", jwt.ValidationErrorExpired)
	}
	return nil
}

func GenerateToken(claims JWTClaims) (string, error) {
	var jwtSecret = []byte(config.Config.Jwt.Secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*JWTClaims, error) {
	var jwtSecret = []byte(config.Config.Jwt.Secret)
	if tokenString[:7] != "Bearer " {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}
	tokenString = tokenString[7:]
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	err = token.Claims.Valid()
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}
	return claims, nil
}
