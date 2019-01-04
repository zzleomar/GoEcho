package routes

import (
	"app/api/controllers"

	"github.com/labstack/echo"
)

func RoutesTesting(e *echo.Echo) {
	e.GET("/aux", controllers.Redirect)
	e.GET("/ruta", controllers.Hello)
	e.GET("/dividir", controllers.Dividir)
	e.GET("/usuario", controllers.Usuario)
	e.GET("/usuario/:id", controllers.FindUsuario)
}
