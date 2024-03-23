package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type PostRequest struct {
	Name string `json:"name"`
}

func getHello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "get hello world"})
}
func postHello(c echo.Context) error {
	r := new(PostRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	fmt.Println(r)
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("%s hello world", r.Name)})
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
	e.Logger.Fatal(e.Start(":8090"))
}
