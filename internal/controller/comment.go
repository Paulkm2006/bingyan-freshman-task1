package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateComment(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, nil)
	}
	var req model.Comment
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	req.UID = uid
	err := model.CreateComment(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}

func GetCommentsByPID(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	if req.Validate() {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	comments, err := model.GetCommentsByPID(req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, comments)
}

func GetCommentsByUID(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	if req.Validate() {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	comments, err := model.GetCommentsByUID(req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, comments)
}

func DeleteComment(c echo.Context) error {
	id := c.QueryParam("cid")
	cid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	comment, err := model.GetCommentByCID(cid)
	if err != nil {
		return param.ErrNotFound(c, nil)
	}
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, nil)
	}
	if comment.UID != uid && utils.CheckPermission(c, 0) {
		return param.ErrForbidden(c, nil)
	}
	err = model.DeleteComment(cid, comment.PID)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}
