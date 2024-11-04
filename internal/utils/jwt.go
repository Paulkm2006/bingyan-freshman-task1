package utils

import (
	"bingyan-freshman-task0/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UID   int  `json:"uid"`
	Admin bool `json:"admin"`
	jwt.RegisteredClaims
}

func jwtSkipper(c echo.Context) bool {
	for _, path := range config.Config.Jwt.SkippedPaths {
		if path == c.Path() {
			return true
		}
	}
	return false
}

func InitJWT(e *echo.Echo) {
	conf := echojwt.Config{
		SigningKey:  []byte(config.Config.Jwt.Secret),
		TokenLookup: "header:Authorization:Bearer ",
		Skipper:     jwtSkipper,
	}
	e.Use(echojwt.WithConfig(conf))
}

func GenerateToken(uid int, admin bool) (string, error) {
	claims := &JWTClaims{
		uid,
		admin,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Config.Jwt.Expire * int64(time.Minute)))),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config.Jwt.Secret))
	return token, err
}

func ParseToken(token string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})
	return claims, err
}
