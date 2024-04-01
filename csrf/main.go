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

func postHello(c echo.Context) error {
	fmt.Println("hoge")
	r := new(PostRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("hello %s", r.Name)})
}
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))
	// e.Use(middleware.CSRF())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	e.POST("/hello", postHello)
	e.Logger.Fatal(e.Start(":8080"))
}
