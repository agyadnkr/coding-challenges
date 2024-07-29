package main

import (
	"app/model"
	"app/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	data, err := model.InitDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to Database")
	}

	model.DB = data

	routes.User(e)

	e.Logger.Fatal(e.Start(":8080"))

}
