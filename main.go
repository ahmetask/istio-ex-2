package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/b", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "service-response2")
	})

	e.Logger.Fatal(e.Start(":8081"))
}
