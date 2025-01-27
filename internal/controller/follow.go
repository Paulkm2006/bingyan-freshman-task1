package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/dto"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Follow(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	followee, err := strconv.Atoi(c.QueryParam("followee"))
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	var req dto.Follow
	req.UID = uid
	req.Followee = followee
	err = model.CreateFollow(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}

func GetFollows(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	follows, err := model.GetFollowsByUID(uid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, follows)
}

func GetFollowers(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	followers, err := model.GetFollowersByUID(uid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, followers)
}

func Unfollow(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	followee, err := strconv.Atoi(c.QueryParam("followee"))
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	err = model.DeleteFollow(uid, followee)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
