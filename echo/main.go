package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "get Hello World")
}
func postHello(c echo.Context) error {
	return c.String(http.StatusOK, "post Hello World")
}

func deleteHello(c echo.Context) error {
	return c.String(http.StatusOK, "delete Hello World")
}
func main() {
	e := echo.New()
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8090"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/getHello", getHello)
	e.POST("/postHello", postHello)
	e.DELETE("/deleteHello", deleteHello)
	e.Start(":8090")
}
