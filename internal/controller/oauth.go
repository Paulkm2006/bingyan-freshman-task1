package controller

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

func OauthCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return param.ErrBadRequest(c, "code is required")
	}

	accessToken, err := utils.GetAccessToken(code)
	if err != nil {
		return param.ErrForbidden(c, err.Error())
	}

	user, err := utils.GetOrCreateUser(accessToken)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}

	token, err := utils.GenerateToken(user.ID, user.Permission)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}

	return param.Success(c, &param.TokenResponse{
		Token:   token,
		Expires: int(config.Config.Jwt.Expire),
	})
}
