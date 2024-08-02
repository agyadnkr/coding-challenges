package controller

import (
	"app/model"
	helpers "app/utility"
	"net/http"

	"github.com/labstack/echo/v4"
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
		return helpers.ReturnLog(c, http.StatusInternalServerError, "Error_create_inventory")
	}

	return c.JSON(http.StatusCreated, inventory)
}
