package controllers

import (
	"app/api/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Usuario(c echo.Context) error {
	return c.File("views/usuarios.html")
}

func FindUsuario(c echo.Context) error {
	var name = "initial"
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if i > 20 {
		name = "Juan"
	} else {
		name = "Carlos"
	}

	u := &models.User{
		Name:   name,
		Email:  "jon@labstack.com",
		IdUser: c.Param("id"),
	}
	// return c.JSON(http.StatusOK, u)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)
}

func Redirect(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/ruta")
}

func Dividir(c echo.Context) error {
	data := c.QueryParam("d")
	dataInt, _ := strconv.Atoi(data)
	result := 2500 / dataInt
	return c.String(http.StatusOK, strconv.Itoa(result))

}

func Usuarios(c echo.Context) error {
	data := []struct {
		Name string `json:"name"`
	}{
		{"Leomar"},
		{"Jose"},
		{"Carmen"},
		{"Ramon"},
	}

	return c.JSON(http.StatusOK, data)
}
