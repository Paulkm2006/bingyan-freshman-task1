package controller

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"crypto/md5"
	"fmt"

	"github.com/labstack/echo/v4"
)

func AdminLogin(c echo.Context) error {
	var adminUser model.Admin
	if err := c.Bind(&adminUser); err != nil {
		return echo.ErrBadRequest
	}
	result, err := model.GetAdmin(adminUser.Username)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if result.Password != fmt.Sprintf("%x", md5.Sum([]byte(adminUser.Password))) {
		return echo.ErrUnauthorized
	}
	token, err := utils.GenerateToken(result.ID, true)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, &param.Resp{
		Success: true,
		Data: &param.TokenResponse{
			Token:   token,
			Expires: int(config.Config.Jwt.Expire),
		},
	})
}
