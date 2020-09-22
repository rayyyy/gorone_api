package route

import (
	"gorone_api/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) *echo.Echo {
	e.POST("/calc", controllers.RequestCalc)
	e.GET("/calc/:id/status", controllers.GetStatus)
	e.GET("/calc/:id/result", controllers.GetResult)
	return e
}
