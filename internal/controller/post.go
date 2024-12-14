package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/dto"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"slices"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePost(c echo.Context) error {
	var req dto.Post
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	req.UID = utils.GetUID(c)
	if req.UID == -1 {
		return param.ErrUnauthorized(c, "")
	}
	id, err := model.CreatePost(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, id)
}

func GetPostByPID(c echo.Context) error {
	id := c.QueryParam("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	post, err := model.GetPostByPID(pid)
	if err == gorm.ErrRecordNotFound {
		return param.ErrNotFound(c, "Post not found")
	} else if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, post)
}

func GetPosts(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if !req.Validate() {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	posts, err := model.GetPosts(req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}

func GetPostsByUID(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if _, err := model.GetUserByID(req.Id); err != nil {
		return param.ErrNotFound(c, "User not found")
	}
	if !req.Validate() {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	posts, err := model.GetPostsByUID(req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}

func GetPostsByNID(c echo.Context) error {
	var req param.Paging
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if _, err := model.GetNodeByNID(req.Id); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if !req.Validate() {
		return param.ErrBadRequest(c, "page or pageSize must be greater than 0")
	}
	posts, err := model.GetPostsByNID(req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, posts)
}

func DeletePost(c echo.Context) error {
	uid := utils.GetUID(c)
	if uid == -1 {
		return param.ErrUnauthorized(c, "")
	}
	id := c.QueryParam("pid")
	pid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	post, err := model.GetPostByPID(pid)
	if err != nil {
		return param.ErrNotFound(c, "Post not found")
	}
	node, err := model.GetNodeByNID(post.NID)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	if post.UID != uid && utils.CheckPermission(c, 0) && !slices.Contains(node.Moderators, uid) {
		return param.ErrForbidden(c, "")
	}
	if err := model.DeletePost(pid); err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
