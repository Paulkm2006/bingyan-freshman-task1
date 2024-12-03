package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateNode(c echo.Context) error {
	var req model.Node
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	node, err := model.CreateNode(&req)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, node.NID)
}

func GetNodes(c echo.Context) error {
	nodes, err := model.GetNodes()
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, nodes)
}

func GetNodeByNID(c echo.Context) error {
	id := c.QueryParam("nid")
	nid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	node, err := model.GetNodeByNID(nid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, node)
}

func AddModerator(c echo.Context) error {
	type Request struct {
		UID int `json:"uid"`
		NID int `json:"nid"`
	}
	var req Request
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if !utils.CheckPermission(c, 1) {
		return param.ErrForbidden(c, "")
	}
	err := model.AddModerator(req.UID, req.NID)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}

func DeleteModerator(c echo.Context) error {
	type Request struct {
		UID int `json:"uid"`
		NID int `json:"nid"`
	}
	var req Request
	if err := c.Bind(&req); err != nil {
		return param.ErrBadRequest(c, "")
	}
	if !utils.CheckPermission(c, 1) {
		return param.ErrForbidden(c, "")
	}
	err := model.DeleteModerator(req.UID, req.NID)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}

func DeleteNode(c echo.Context) error {
	id := c.QueryParam("nid")
	nid, err := strconv.Atoi(id)
	if err != nil {
		return param.ErrBadRequest(c, "")
	}
	if !utils.CheckPermission(c, 1) {
		return param.ErrForbidden(c, "")
	}
	err = model.DeleteNode(nid)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
