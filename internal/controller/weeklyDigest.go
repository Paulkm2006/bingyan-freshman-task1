package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/service"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

func SendWeeklyDigest(c echo.Context) error {

	if !utils.CheckPermission(c, 1) {
		return param.ErrUnauthorized(c, "")
	}
	posts, err := model.GetWeeklyPosts()
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	users, err := model.GetUsers()
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	err = service.SendWeeklyDigest(users, posts)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
