package utils

import (
	"bingyan-freshman-task0/internal/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTClaims struct {
	UID        int `json:"uid"`
	Permissiom int `json:"permission"` // 0: user, 1: admin
	jwt.RegisteredClaims
}

func jwtSkipper(c echo.Context) bool {
	for _, path := range config.Config.Jwt.SkippedPaths {
		i := strings.Split(path, ":")
		if i[1] == c.Path() && i[0] == c.Request().Method {
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

func GenerateToken(uid int, permission int) (string, error) {
	claims := &JWTClaims{
		uid,
		permission,
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
