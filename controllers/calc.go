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

	var results []models.CalcResult
	res := db.Find(&results, req.GetValues())
	ret := responses.Status{Status: "progress",
		TotalCount: len(req.GetValues()),
		Completed:  int(res.RowsAffected),
	}
	return c.JSON(http.StatusOK, ret)
}

func GetResult(c echo.Context) error {
	db := db.DbManager()
	var req models.Request
	if err := db.First(&req, c.Param("id")).Error; err != nil {
		// https://gorm.io/ja_JP/docs/error_handling.html
		return err
	}
	var results []models.CalcResult
	if err := db.Find(&results, req.GetValues()).Error; err != nil {
		// https://gorm.io/ja_JP/docs/error_handling.html
		return err
	}

	return c.JSON(http.StatusOK, results)
}
