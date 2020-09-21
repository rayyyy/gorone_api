package controllers

import (
	"gorone_api/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequestCalc(c echo.Context) error {
	ret := new(responses.Request)
	ret.Message = "ok"

	return c.JSON(http.StatusOK, ret)
}

func GetStatus(c echo.Context) error {
	ret := new(responses.Status)
	ret.Status = "ok"

	return c.JSON(http.StatusOK, ret)
}

func GetResult(c echo.Context) error {
	ret := new(responses.Request)
	ret.Message = "ok"

	return c.JSON(http.StatusOK, ret)
}
