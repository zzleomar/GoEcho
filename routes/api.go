package routes

import (
	"app/api/controllers"

	"github.com/labstack/echo"
)

func RoutesApi(e *echo.Echo) {
	e.GET("/api/usuarios", controllers.Usuarios)
}
