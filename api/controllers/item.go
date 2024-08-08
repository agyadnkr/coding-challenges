package controller

import (
	"app/model"
	"app/utility"
	helpers "app/utility"
	"errors"
	"net/http"
	"time"

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
			"item_name":  item.ItemName,
			"Price":      item.ItemPrice,
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
			"id":           item.Itmid,
			"item_name":    item.ItemName,
			"price":        item.ItemPrice,
			"created_at":   item.CreatedAt,
			"last_updated": item.UpdatedAt,
		}
	}

	return utility.ReturnLog(c, http.StatusOK, filteredItems)
}

func UpdateItem(c echo.Context) error {
	itemID := c.Param("id")

	var updatedItem model.Item
	if err := c.Bind(&updatedItem); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_item")
	}

	updatedItem.UpdatedAt = time.Now()

	if err := model.UpdateItem(itemID, updatedItem); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_updating_item")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         itemID,
		"Response":   "Ok",
		"updated_at": updatedItem.UpdatedAt,
	})
}

func DeleteItem(c echo.Context) error {
	itemID := c.Param("id")

	if err := model.DeleteItem(itemID); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_deleting_item")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "item deleted successfully",
	})
}
