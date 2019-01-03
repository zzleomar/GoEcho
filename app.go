package main

import (
	"strconv"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// typeUser 

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
	Id string `json:"id" xml:"id"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Service Static
	e.Static("/", "views")

	// Routes
	e.GET("/aux", redirect)
	e.GET("/ruta", hello)
	e.GET("/usuario", usuario)
	e.GET("/usuario/:id", findusuario)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func usuario(c echo.Context) error {
	return c.File("views/usuarios.html")
}

func findusuario(c echo.Context) error {
	var name = "initial"
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if i>20 {
		name = "Juan"
    } else {
		name = "Carlos"
	}

	u := &User{
		Name: name,
		Email: "jon@labstack.com",
		Id:  c.Param("id"),
	}
	// return c.JSON(http.StatusOK, u)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)
}

func redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/ruta")
}
