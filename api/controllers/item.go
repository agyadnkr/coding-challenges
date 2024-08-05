package controller

import (
	"app/model"
	"app/utility"
	helpers "app/utility"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateItem(c echo.Context) error {
	var items []model.Item

	if err := c.Bind(&items); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_items")
	}

	var createdItems []map[string]interface{}
	for _, item := range items {
		if item.ItemName == "" || item.ItemPrice == 0 {
			return helpers.ReturnLog(c, http.StatusBadRequest, "Error_empty_fields")
		}

		if err := model.CreateItem(&item); err != nil {
			return helpers.ReturnLog(c, http.StatusInternalServerError, "item_with_the_same_name_already_exist")
		}

		createdItems = append(createdItems, map[string]interface{}{
			"id":         item.Itmid,
			"created_at": item.CreatedAt,
		})
	}

	return c.JSON(http.StatusCreated, createdItems)
}

func FetchAllItems(c echo.Context) error {
	var searchRequest model.Filter

	if err := c.Bind(&searchRequest); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	items, err := model.FetchItem(searchRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ReturnLog(c, http.StatusInternalServerError, "Error_record_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_searching_products")
	}

	filteredItems := make([]map[string]interface{}, len(items))
	for i, item := range items {
		filteredItems[i] = map[string]interface{}{
			"id":         item.Itmid,
			"created_at": item.CreatedAt,
		}
	}

	return utility.ReturnLog(c, http.StatusOK, filteredItems)
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         id,
		"updated_at": itemData.UpdatedAt,
	})
}

func DeleteItem(c echo.Context) error {
	itemID := c.Param("id")

	if err := model.DeleteItem(itemID); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_deleting_warehouse")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item deleted successfully",
	})
}
