package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateComment(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	var req model.Comment
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	req.UID = claims.UID
	err = model.CreateComment(&req)
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
	if req.Page <= 0 || req.PageSize <= 0 {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	comments, err := model.GetCommentsByPID(req.Id, req.Page, req.PageSize)
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
	if req.Page <= 0 || req.PageSize <= 0 {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	comments, err := model.GetCommentsByUID(req.Id, req.Page, req.PageSize)
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
		return echo.ErrNotFound
	}
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	if comment.UID != claims.UID && claims.Permission == 0 {
		return echo.ErrForbidden
	}
	err = model.DeleteComment(cid, comment.PID)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}
