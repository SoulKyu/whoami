package main

import (
	"whoami/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "templates")
	e.Static("/static", "static")

	// Routes
	e.GET("/", routes.Index)
	e.GET("/tutoriels", routes.Tutoriel)
	e.GET("/whoami", routes.Whoami)

	e.Logger.Fatal(e.Start(":8080"))
}
