package controller

import (
	"app/model"
	"app/utility"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func FecthAllItems(c echo.Context) error {

	var searchRequest model.Filter

	if err := c.Bind(&searchRequest); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	item, err := model.FetchItem(searchRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utility.ReturnLog(c, http.StatusInternalServerError, "Error_record_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_searching_products")
	}

	return utility.ReturnLog(c, http.StatusOK, item)
}

func UpdateItem(c echo.Context) error {

	return nil
}

func DeleteItem(c echo.Context) error {

	return nil
}
