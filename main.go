package main

import (
	"github.com/che-ict/DEV-DT-Microblog/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/login", handlers.LoginGetHandler)
	e.POST("/login")
	e.GET("/register", handlers.RegisterGetHandler)
	e.POST("/register", handlers.RegisterPostHandler)
	e.GET("/timeline")
	e.GET("/user/:user")
	e.POST("/post")
	e.DELETE("/post/:id")
}
