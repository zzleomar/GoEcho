package main

import (
	"app/api/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Service Static
	e.Static("/", "views")
	routes.RoutesTesting(e)
	routes.RoutesApi(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
