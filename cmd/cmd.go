package main

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/router"
	"bingyan-freshman-task0/internal/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.InitConfig()
	router.InitRouter(e)
	utils.InitLogger(e)
	defer utils.Logger.Sync()
	model.InitDB()
	utils.InitRedis()
	utils.InitJWT(e)
	//model.AddDefaultAdmin()
	//e.Logger.Fatal(e.Start(":" + config.Config.Server.Port))
	e.Start(":" + config.Config.Server.Port)
}
