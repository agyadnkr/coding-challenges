package controller

import (
	"app/model"
	"app/utility"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateWarehouse(c echo.Context) error {
	var warehouse model.Warehouse

	if err := c.Bind(&warehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_bind_warehouse")
	}

	warehouse.Wid = uuid.New().String()
	if err := model.CreateWarehouse(warehouse); err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "warehouse_with_the_same_name_already_exist")
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Warehouse created successfully",
	})
}

func FetchAllWarehouses(c echo.Context) error {
	warehouses, err := model.GetAllWarehouses()
	if err != nil {
		return utility.ReturnLog(c, http.StatusInternalServerError, "Error_fetching_warehouses")
	}

	return c.JSON(http.StatusOK, warehouses)
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
