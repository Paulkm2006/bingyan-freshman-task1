package controller

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func UserInfo(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	_, err := utils.ParseToken(token)
	if err != nil {
		return echo.ErrUnauthorized
	}
	var req model.User
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	var user *model.User
	if req.ID != 0 {
		user, err = model.GetUserByID(req.ID)
	} else if req.Username != "" {
		user, err = model.GetUserByUsername(req.Username)
	} else {
		return echo.ErrBadRequest
	}
	if err == model.ErrUserNotFound {
		return echo.ErrNotFound
	} else if err != nil {
		return echo.ErrInternalServerError
	}
	user.Password = ""
	return c.JSON(200, &param.Resp{
		Success: true,
		Data:    user,
	})
}

func UserLogin(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	result, err := model.GetUserByUsername(user.Username)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if result.Password != fmt.Sprintf("%x", md5.Sum([]byte(user.Password))) {
		return echo.ErrUnauthorized
	}
	token, err := utils.GenerateToken(result.ID, result.Permissiom)
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

func UserRegister(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	status, err := utils.ValidateCode(user.Email, c.QueryParam("code"))
	if err != nil {
		return echo.ErrInternalServerError
	} else if !status {
		return echo.ErrUnauthorized
	}
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	err = model.AddUser(&user)
	if err == model.ErrUserAlreadyExist {
		return echo.ErrConflict
	} else if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(201, &param.Resp{
		Success: true,
	})
}

func UserDelete(c echo.Context) error {
	var user model.User
	token := c.Get("user").(*jwt.Token)
	claims, err := utils.ParseToken(token.Raw)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if claims.Permissiom == 0 {
		return echo.ErrForbidden
	}
	if err = c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	err = model.DeleteUser(user.ID)
	if err == model.ErrUserNotFound {
		return echo.ErrNotFound
	} else if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, &param.Resp{
		Success: true,
	})
}
