package routes

import (
	controllers "app/controllers"
	"app/middleware"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {

	e.POST("register", controllers.Signup)

	e.POST("login", controllers.Login)

	itemGroup := e.Group(e.POST("/items", middleware.CheckToken, controllers.CreateItem))

	itemGroup.POST("/:id", controllers.FecthAllItems)

}
