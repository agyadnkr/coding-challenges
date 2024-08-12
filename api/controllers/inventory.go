package controller

import (
	"app/model"
	"app/utility"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateInventory(c echo.Context) error {
	var reqData model.CreateInventoryRequest

	if err := c.Bind(&reqData); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	if err := model.CreateInventory(reqData); err != nil {
		if errors.Is(err, model.ErrDuplicatedData) {
			return utility.ReturnLog(c, http.StatusConflict, "Duplicated_data")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_creating_inventory")
	}

	return utility.ReturnLog(c, http.StatusCreated, "Inventory_created")
}

func FetchAllInventories(c echo.Context) error {
	var searchRequest model.Filter

	if err := c.Bind(&searchRequest); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	inventories, err := model.FetchInventories(searchRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ReturnLog(c, http.StatusInternalServerError, "Error_record_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_fetching_inventories")
	}

	return utility.ReturnLog(c, http.StatusOK, inventories)
}

func UpdateInventory(c echo.Context) error {
	inventoryID := c.Param("id")
	var reqData model.UpdateInventoryRequest

	if err := c.Bind(&reqData); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	if err := model.UpdateInventory(inventoryID, reqData); err != nil {
		if errors.Is(err, model.ErrItemNotFound) {
			return utility.ReturnLog(c, http.StatusNotFound, "Item_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_updating_inventory")
	}

	return utility.ReturnLog(c, http.StatusOK, "Inventory_updated")
}

func MoveStockInventory(c echo.Context) error {
	var stockMoveRequest model.StockMoveRequest

	if err := c.Bind(&stockMoveRequest); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_stock_move")
	}

	if err := model.MoveStock(stockMoveRequest); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_moving_stock")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Stock moved successfully",
	})
}
