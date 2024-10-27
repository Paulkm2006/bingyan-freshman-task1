package controller

import (
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

func SendValidation(c echo.Context) error {
	mail := c.QueryParam("mail")
	status, t, err := utils.CheckEmailExist(mail)
	if err != nil {
		return echo.ErrInternalServerError
	} else if status {
		return c.JSON(429, &param.Resp{
			Success: false,
			Msg:     "Interval too short",
			Data:    t,
		})
	}
	code := utils.GenerateValidationCode()
	err = utils.SendValidation(mail, code)
	if err != nil {
		return echo.ErrInternalServerError
	}
	err = utils.WriteValidationCode(mail, code)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, &param.Resp{
		Success: true,
	})
}
