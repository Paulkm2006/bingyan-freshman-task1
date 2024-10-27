package router

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	var ver = config.Config.Server.Ver
	e.POST(fmt.Sprintf("/%s/user/token", ver), controller.UserLogin)   //Login and retrieve token
	e.POST(fmt.Sprintf("/%s/user", ver), controller.UserRegister)      // Register
	e.DELETE(fmt.Sprintf("/%s/user", ver), controller.UserDelete)      // Delete user, requires admin token
	e.GET(fmt.Sprintf("/%s/user", ver), controller.UserInfo)           // Get user information, requires token
	e.POST(fmt.Sprintf("/%s/admin/token", ver), controller.AdminLogin) // Login and retrieve admin token
	e.GET(fmt.Sprintf("/%s/verify", ver), controller.SendValidation)   // Get user information, requires admin token
}
