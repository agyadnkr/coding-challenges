package controller

import (
	"app/model"
	"app/utility"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateItem(c echo.Context) error {

	var ReqData []model.Item

	if err := c.Bind(&ReqData); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_items")
	}

	return nil
}
