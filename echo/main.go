package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "get hello world"})
}
func postHello(c echo.Context) error {
	name := c.FormValue("name")
	fmt.Println(name)
	return c.JSON(http.StatusOK, map[string]string{"response": fmt.Sprintf("%s hello world", name)})
}

func deleteHello(c echo.Context) error {
	return c.String(http.StatusOK, "delete Hello World")
}
func main() {
	e := echo.New()
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/getHello", getHello)
	e.POST("/postHello", postHello)
	e.DELETE("/deleteHello", deleteHello)
	e.Start(":8090")
}
