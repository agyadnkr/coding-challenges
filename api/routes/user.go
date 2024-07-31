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

}
