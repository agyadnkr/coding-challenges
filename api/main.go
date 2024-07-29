package main

import (
	"app/interactor"
	"app/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	data, err := interactor.InitDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to Database")
	}

	interactor.DB = data

	routes.User(e)

	e.Logger.Fatal(e.Start(":8080"))

}
