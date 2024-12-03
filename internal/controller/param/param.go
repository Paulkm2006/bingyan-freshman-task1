package param

import (
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(200, &Resp{
		Success: true,
		Data:    data,
	})
}

type TokenResponse struct {
	Token   string `json:"token"`
	Expires int    `json:"expires_in"`
}

type Paging struct {
	Id       int `query:"id"`
	Page     int `query:"page"`
	PageSize int `query:"pageSize"`
	Sort     int `query:"sort"` // 0: by time desc, 1: by comment num desc, 2: by like num desc
}

func (p *Paging) Validate() bool {
	return p.Page > 0 && p.PageSize > 0
}

func (p *Paging) SortingStatement() string {
	switch p.Sort {
	case 0:
		return "created desc"
	case 1:
		return "comments desc"
	case 2:
		return "likes desc"
	default:
		return "created desc"
	}
}

type ErrResp struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func ErrBadRequest(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Bad Request"
	}
	return c.JSON(400, &ErrResp{
		Code: 400,
		Msg:  msg,
	})
}

func ErrUnauthorized(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Unauthorized"
	}
	return c.JSON(401, &ErrResp{
		Code: 401,
		Msg:  msg,
	})
}

func ErrNotFound(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Not Found"
	}
	return c.JSON(404, &ErrResp{
		Code: 404,
		Msg:  msg,
	})
}

func ErrInternalServerError(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Internal Server Error"
	}
	utils.Logger.Error(msg.(string))
	return c.JSON(500, &ErrResp{
		Code: 500,
		Msg:  msg,
	})
}

func ErrConflict(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Conflict"
	}
	return c.JSON(409, &ErrResp{
		Code: 409,
		Msg:  msg,
	})
}

func ErrForbidden(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Forbidden"
	}
	return c.JSON(403, &ErrResp{
		Code: 403,
		Msg:  msg,
	})
}

func ErrIntervalTooShort(c echo.Context, msg interface{}) error {
	if msg == nil {
		msg = "Interval too short"
	}
	return c.JSON(429, &ErrResp{
		Code: 429,
		Msg:  msg,
	})
}
