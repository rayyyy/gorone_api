package controllers

import (
	"fmt"
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
	req, err := models.NewRequest("calc", params, 5)
	if err != nil {
		return err
	}
	err = calc.Publish(*req)
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

	// https://gorm.io/docs/advanced_query.html#Count
	var count int64
	db.Model(&models.CalcResult{}).Where("key_name IN ?", req.DecodedBody()).Count(&count)
	fmt.Println(count)
	ret := responses.Status{
		TotalCount: len(req.DecodedBody()),
		Completed:  int(count),
		RequestID:  int(req.ID),
	}
	if len(req.DecodedBody()) == int(count) {
		ret.Status = "OK"
	} else {
		ret.Status = "processing"
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
	if err := db.Where("key_name IN ?", req.DecodedBody()).Find(&results).Error; err != nil {
		// https://gorm.io/ja_JP/docs/error_handling.html
		return err
	}

	return c.JSON(http.StatusOK, results)
}
