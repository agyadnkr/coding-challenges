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
