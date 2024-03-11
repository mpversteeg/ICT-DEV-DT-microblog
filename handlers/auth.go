package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func RegisterPostHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
}

func LoginGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}
