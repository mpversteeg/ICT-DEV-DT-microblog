package handlers

import (
	"github.com/che-ict/DEV-DT-Microblog/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RegisterGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func RegisterPostHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	displayName := c.FormValue("display-name")

	err := repositories.CreateUser(username, password, displayName)

	if err != nil {
		return c.Redirect(http.StatusFound, "/register")
	}
	c.SetCookie(&http.Cookie{
		Name:    "user",
		Value:   username,
		Expires: time.Now().Add(time.Hour * 12),
	})
	return c.Redirect(http.StatusFound, "/feed")

}

func LoginGetHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}

func LoginPostHandler(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result := repositories.LoginUser(username, password)

	if !result {
		return c.Redirect(300, "/login")
	}
	c.SetCookie(&http.Cookie{
		Name:    "user",
		Value:   username,
		Expires: time.Now().Add(time.Hour * 12),
	})
	return c.Redirect(300, "/feed")
}
