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

	for _, user := range ReqData {
		if err := model.CreateItem(user); err != nil {
			return utility.ReturnLog(c, http.StatusInternalServerError, "item_with_the_same_name_already_exist")
		}
	}

	return nil
}
