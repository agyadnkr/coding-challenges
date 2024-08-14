package controller

import (
	"app/model"
	"app/utility"
	helpers "app/utility"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateItem(c echo.Context) error {
	var item model.Item

	if err := c.Bind(&item); err != nil {
		fmt.Println(err)
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_items")
	}

	var createdItem map[string]interface{}
	if item.ItemName == "" {
		return helpers.ReturnLog(c, http.StatusBadRequest, "Error_empty_fields")
	}

	if err := model.CreateItem(&item); err != nil {
		return helpers.ReturnLog(c, 409, "")
	}

	createdItem = map[string]interface{}{
		"id":         item.Itmid,
		"item_name":  item.ItemName,
		"Price":      item.ItemPrice,
		"created_at": item.CreatedAt,
	}

	return c.JSON(http.StatusCreated, createdItem)
}

func CreateMultipleItems(c echo.Context) error {
	var items []model.Item

	if err := c.Bind(&items); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_items")
	}

	tx := model.DB.Begin()
	if tx.Error != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_starting_transaction")
	}

	var createdItems []map[string]interface{}
	for _, item := range items {
		if item.ItemName == "" {
			tx.Rollback()
			return helpers.ReturnLog(c, http.StatusBadRequest, "Error_empty_fields")
		}

		if err := model.CreateItems(tx, &item); err != nil {
			tx.Rollback()
			return helpers.ReturnLog(c, http.StatusInternalServerError, err.Error())
		}

		createdItems = append(createdItems, map[string]interface{}{
			"id":         item.Itmid,
			"item_name":  item.ItemName,
			"Price":      item.ItemPrice,
			"created_at": item.CreatedAt,
		})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_committing_transaction")
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

func FetchSingleItem(c echo.Context) error {
	itemID := c.Param("id")

	item, err := model.FetchItemByID(itemID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helpers.ReturnLog(c, http.StatusNotFound, "Item_not_found")
		}
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_fetching_item")
	}

	return c.JSON(http.StatusOK, item)
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
