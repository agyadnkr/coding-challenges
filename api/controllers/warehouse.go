package controller

import (
	"app/model"
	"app/utility"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateWarehouse(c echo.Context) error {
	var warehouse model.Warehouse

	if err := c.Bind(&warehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_warehouse")
	}

	warehouse.Wid = uuid.New().String()
	if err := model.CreateWarehouse(&warehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "warehouse_with_the_same_name_already_exist")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":    "Warehouse created successfully",
		"id":         warehouse.Wid,
		"created_at": warehouse.CreatedAt,
	})
}

func FetchAllWarehouses(c echo.Context) error {
	var searchRequest model.Filter

	if err := c.Bind(&searchRequest); err != nil {
		return utility.ReturnLog(c, http.StatusBadRequest, "Invalid_request_body")
	}

	warehouse, err := model.GetAllWarehouses(searchRequest)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utility.ReturnLog(c, http.StatusInternalServerError, "Error_record_not_found")
		}
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_searching_warehouse")
	}

	filteredWarehouse := make([]map[string]interface{}, len(warehouse))
	for i, warehouse := range warehouse {
		filteredWarehouse[i] = map[string]interface{}{
			"id":             warehouse.Wid,
			"warehouse_name": warehouse.WarehouseName,
			"address":        warehouse.WarehouseAddress,
			"created_at":     warehouse.CreatedAt,
			"last_updated":   warehouse.UpdatedAt,
		}
	}

	return utility.ReturnLog(c, http.StatusOK, filteredWarehouse)
}

func UpdateWarehouse(c echo.Context) error {
	warehouseID := c.Param("id")

	var updatedWarehouse model.Warehouse
	if err := c.Bind(&updatedWarehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_warehouse")
	}

	if err := model.UpdateWarehouse(warehouseID, updatedWarehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_updating_warehouse")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Warehouse updated successfully",
	})
}

func DeleteWarehouse(c echo.Context) error {
	warehouseID := c.Param("id")

	if err := model.DeleteWarehouse(warehouseID); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_deleting_warehouse")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Warehouse deleted successfully",
	})
}
