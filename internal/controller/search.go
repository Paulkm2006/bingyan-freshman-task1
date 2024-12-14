package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/service"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

func SearchPost(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "User must be logged in to search posts")
	}
	keyword := c.QueryParam("keyword")
	posts, err := service.SearchPost(keyword)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}
