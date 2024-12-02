package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateLike(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	var req model.Like
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	req.UID = claims.UID
	err = model.CreateLike(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}

func GetLikesByPID(c echo.Context) error {
	id := c.Param("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	likes, err := model.GetLikesByPID(pid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, likes)
}

func GetLikesByUID(c echo.Context) error {
	id := c.Param("uid")
	uid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	likes, err := model.GetLikesByUID(uid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, likes)
}

func DeleteLike(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	uid := claims.UID
	pid, err := strconv.Atoi(c.QueryParam("pid"))
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	err = model.DeleteLike(uid, pid)
	if err == model.ErrLikeNotFound {
		return param.ErrNotFound(c, "user didn't like this post")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}
