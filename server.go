package main

import (
	"gorone_api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/calc/", controllers.RequestCalc)
	e.GET("/calc/:id/status", controllers.GetStatus)
	e.GET("/calc/:id/result", controllers.GetResult)

	e.Logger.Fatal(e.Start(":5555"))
}
