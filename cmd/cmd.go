package main

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/router"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	config.InitConfig()
	router.InitRouter(e)
	model.InitDB()
	utils.InitRedis()
	//model.AddDefaultAdmin()
	e.Logger.Fatal(e.Start(":" + config.Config.Server.Port))
}
