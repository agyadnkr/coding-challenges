package controller

import (
	"app/model"
	"app/utility"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateItem(c echo.Context) error {
	var reqData []model.Item

	if err := c.Bind(&reqData); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_items")
	}

	for _, item := range reqData {
		item.Itmid = uuid.New().String()
		if err := model.CreateItem(item); err != nil {
			return utility.ReturnLog(c, http.StatusInternalServerError, "item_with_the_same_name_already_exist")
		}
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Items created successfully",
	})
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
	id := c.Param("id")
	var itemData model.Item

	if err := c.Bind(&itemData); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	if err := model.UpdateItem(id, itemData); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_updating_warehouse")
	}

	return utility.ReturnLog(c, http.StatusOK, "Item_updated_successfully")
}

func DeleteItem(c echo.Context) error {
	itemID := c.Param("id")

	if err := model.DeleteItem(itemID); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_deleting_warehouse")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Warehouse deleted successfully",
	})
}
