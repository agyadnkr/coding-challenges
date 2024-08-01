package routes

import (
	controllers "app/controllers"
	"app/middleware"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {

	e.POST("register", controllers.Signup)

	e.POST("login", controllers.Login)

	apiGroup := e.Group("/api", middleware.CheckToken)

	apiGroup.POST("/items", controllers.CreateItem)
	apiGroup.GET("/items/:id", controllers.FecthAllItems)
	apiGroup.PATCH("/items/:id", controllers.UpdateItem)
	apiGroup.DELETE("/items/:id", controllers.DeleteItem)

	apiGroup.POST("/warehouses", controllers.CreateWarehouse)
	// apiGroup.GET("/warehouses", controllers.FetchAllWarehouses)
	// apiGroup.PATCH("/warehouses/:id", controllers.UpdateWarehouse)
	// apiGroup.DELETE("/warehouses/:id", controllers.DeleteWarehouse)

}
