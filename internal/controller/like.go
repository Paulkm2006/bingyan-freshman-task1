package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateLike(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	var req model.Like
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	req.UID = uid
	err := model.CreateLike(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}

func GetLikesByPID(c echo.Context) error {
	id := c.QueryParam("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	likes, err := model.GetLikesByPID(pid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, likes)
}

func GetLikesByUID(c echo.Context) error {
	id := c.QueryParam("uid")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	likes, err := model.GetLikesByUID(uid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, likes)
}

func DeleteLike(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	pid, err := strconv.Atoi(c.QueryParam("pid"))
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	err = model.DeleteLike(uid, pid)
	if err == model.ErrLikeNotFound {
		return param.ErrNotFound(c, "user didn't like this post")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
