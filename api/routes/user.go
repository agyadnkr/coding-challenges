package routes

import (
	"app/middleware"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {

	e.POST("Signup", controllers.Signup)

	e.POST("Login", controllers.Login)

	userGroup := e.Group("/User", middleware.CheckToken)

	userGroup.POST("/Create", controllers.CreateUser)

}
