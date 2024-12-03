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

func UserInfo(c echo.Context) error {
	var req model.User
	var err error
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	var user *model.User
	if req.ID != 0 {
		user, err = model.GetUserByID(req.ID)
	} else if req.Username != "" {
		user, err = model.GetUserByUsername(req.Username)
	} else {
		return param.ErrBadRequest(c, nil)
	}
	if err == model.ErrUserNotFound {
		return param.ErrNotFound(c, "user not found")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	user.Password = ""
	return param.Success(c, user)
}

func UserLogin(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	result, err := model.GetUserByUsername(user.Username)
	if err != nil {
		return param.ErrUnauthorized(c, "Username not found")
	}
	if result.Password != fmt.Sprintf("%x", md5.Sum([]byte(user.Password))) {
		return param.ErrUnauthorized(c, "Password incorrect")
	}
	token, err := utils.GenerateToken(result.ID, result.Permission)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
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
		return param.ErrBadRequest(c, nil)
	}
	status, err := utils.ValidateCode(user.Email, c.QueryParam("code"))
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	} else if !status {
		return param.ErrUnauthorized(c, "Invalid code")
	}
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	err = model.AddUser(&user)
	if err == model.ErrUserAlreadyExist {
		return param.ErrConflict(c, "Username already exists")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}

func UserDelete(c echo.Context) error {
	var user model.User
	if !utils.CheckPermission(c, 1) {
		return param.ErrForbidden(c, nil)
	}
	if err := c.Bind(&user); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	err := model.DeleteUser(user.ID)
	if err == model.ErrUserNotFound {
		return param.ErrNotFound(c, "User not found")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}
