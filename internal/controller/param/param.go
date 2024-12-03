package param

import (
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	Code int         `json:"success"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c echo.Context, data interface{}) error {
	return c.JSON(200, &Resp{
		Code: 200,
		Msg:  "Success",
		Data: data,
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

func ErrBadRequest(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Bad Request"
	}
	return c.JSON(400, &Resp{
		Code: 400,
		Msg:  msg,
	})
}

func ErrUnauthorized(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Unauthorized"
	}
	return c.JSON(401, &Resp{
		Code: 401,
		Msg:  msg,
	})
}

func ErrNotFound(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Not Found"
	}
	return c.JSON(404, &Resp{
		Code: 404,
		Msg:  msg,
	})
}

func ErrInternalServerError(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Internal Server Error"
	}
	utils.Logger.Error(msg)
	return c.JSON(500, &Resp{
		Code: 500,
		Msg:  msg,
	})
}

func ErrConflict(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Conflict"
	}
	return c.JSON(409, &Resp{
		Code: 409,
		Msg:  msg,
	})
}

func ErrForbidden(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Forbidden"
	}
	return c.JSON(403, &Resp{
		Code: 403,
		Msg:  msg,
	})
}

func ErrIntervalTooShort(c echo.Context, msg string) error {
	if msg == "" {
		msg = "Interval too short"
	}
	return c.JSON(429, &Resp{
		Code: 429,
		Msg:  msg,
	})
}
