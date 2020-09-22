package controllers

import (
	"gorone_api/db"
	calc "gorone_api/lib/queue"
	"gorone_api/lib/util"
	"gorone_api/models"
	"gorone_api/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequestCalc(c echo.Context) error {
	params := new(postParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	err := c.Validate(params)
	if err != nil {
		return err
	}
	params.Values = util.UniqString(params.Values)
	req, err := models.NewRequest(params)
	if err != nil {
		return err
	}
	err = calc.Publish(params.Values)
	if err != nil {
		return err
	}

	ret := responses.Request{
		Message:   "ok",
		RequestID: req.ID,
	}

	return c.JSON(http.StatusOK, ret)
}

type postParams struct {
	Values []string `json:"values" validate:"required"`
}

func GetStatus(c echo.Context) error {
	db := db.DbManager()
	var req models.Request
	if err := db.First(&req, c.Param("id")).Error; err != nil {
		// https://gorm.io/ja_JP/docs/error_handling.html
		return err
	}
	ret := new(responses.Status)
	ret.Status = "ok"

	return c.JSON(http.StatusOK, ret)
}

func GetResult(c echo.Context) error {
	ret := new(responses.Request)
	ret.Message = "ok"

	return c.JSON(http.StatusOK, ret)
}
