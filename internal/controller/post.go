package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	var req model.Post
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	req.UID = claims.UID
	id, err := model.CreatePost(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, id)
}

func GetPostByPID(c echo.Context) error {
	id := c.Param("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	post, err := model.GetPostByPID(pid)
	if err == model.ErrPostNotFound {
		return param.ErrNotFound(c, "Post not found")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, post)
}

func GetPosts(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	if req.Page <= 0 || req.PageSize <= 0 {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	posts, err := model.GetPosts(req.Page, req.PageSize)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}

func GetPostsByUID(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	if _, err := model.GetUserByID(req.Id); err != nil {
		return param.ErrBadRequest(c, nil)
	}
	if req.Page <= 0 || req.PageSize <= 0 {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	posts, err := model.GetPostsByUID(req.Id, req.Page, req.PageSize)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}

func DeletePost(c echo.Context) error {
	token := c.Get("user").(*jwt.Token).Raw
	claims, err := utils.ParseToken(token)
	if err != nil {
		return param.ErrUnauthorized(c, nil)
	}
	id := c.Param("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, nil)
	}
	post, err := model.GetPostByPID(pid)
	if err != nil {
		return param.ErrNotFound(c, "Post not found")
	}
	if post.UID != claims.UID && claims.Permission == 0 {
		return param.ErrForbidden(c, nil)
	}
	if err := model.DeletePost(pid); err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nil)
}
