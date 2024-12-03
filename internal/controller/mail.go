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
		return param.ErrInternalServerError(c, err.Error())
	} else if status {
		return param.ErrIntervalTooShort(c, t.String())
	}
	code := utils.GenerateValidationCode()
	err = utils.SendValidation(mail, code)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	err = utils.WriteValidationCode(mail, code)
	if err != nil {
		return param.ErrInternalServerError(c, err.Error())
	}
	return param.Success(c, "")
}
