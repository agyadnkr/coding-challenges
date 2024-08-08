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

func CreateInventory(c echo.Context) error {
	var inventory model.Inventory

	if err := c.Bind(&inventory); err != nil {
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_bind_inventory")
	}

	if inventory.Wid == "" || inventory.Itmid == "" || inventory.Quantity == 0 {
		return helpers.ReturnLog(c, http.StatusBadRequest, "Error_empty_fields")
	}

	if err := model.CreateInventory(inventory); err != nil {
		if errors.Is(err, model.ErrDuplicateData) {
			return helpers.ReturnLog(c, http.StatusConflict, "duplicated_data")
		}
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_create_user")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":         inventory.Invid,
		"created_at": inventory.CreatedAt,
	})
}

func FetchAllInventories(c echo.Context) error {
	var searchRequest model.Filter

	if err := c.Bind(&searchRequest); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	inventory, err := model.GetAllInventories(searchRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ReturnLog(c, http.StatusInternalServerError, "Error_record_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_searching_products")
	}

	filteredInventory := make([]map[string]interface{}, len(inventory))
	for i, inventory := range inventory {
		filteredInventory[i] = map[string]interface{}{
			"id":           inventory.Invid,
			"warehouse_id": inventory.Wid,
			"item_id":      inventory.Itmid,
			"created_at":   inventory.CreatedAt,
			"last_updated": inventory.UpdatedAt,
		}
	}

	return utility.ReturnLog(c, http.StatusOK, filteredInventory)
}

func UpdateInventory(c echo.Context) error {
	inventoryID := c.Param("id")

	var updatedInventory model.Inventory
	if err := c.Bind(&updatedInventory); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_inventory")
	}

	if err := model.UpdateInventory(inventoryID, updatedInventory); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_updating_inventory")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Inventory updated successfully",
		"data":    updatedInventory,
	})
}

func DeleteInventory(c echo.Context) error {
	inventoryID := c.Param("id")

	if err := model.DeleteInventory(inventoryID); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_deleting_inventory")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Inventory deleted successfully",
	})
}

// func MoveStockInventory(c echo.Context) error {

// 	var stockMoveRequest model.StockMoveRequest

// 	if err := c.Bind(&stockMoveRequest); err != nil {
// 		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_stock_move")
// 	}

// 	// if err := model.MoveStock(stockMoveRequest); err != nil {
// 	// 	return utility.ReturnLog(c, http.StatusInternalServerError, "Error_moving_stock")
// 	// }

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Stock moved successfully",
// 	})
// }
