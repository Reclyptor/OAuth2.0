package main

import (
	"github.com/labstack/echo/v4"
	"oauth/controller/api/v1"
)

func main() {
	e := echo.New()
	g := e.Group("/api/v1")
	g.POST("/applications", v1.PostApplication)
	e.Logger.Fatal(e.Start(":8080"))
}
