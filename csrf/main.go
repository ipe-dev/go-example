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
	if err := c.Ech; err != nil {
		return echo.ErrForbidden
	}
	r := new(PostRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("hello %s", r.Name)})
}
func csrfToken(c echo.Context) error {
	token := c.Get("csrf")
	fmt.Println(token)
	return c.JSON(http.StatusOK, map[string]string{"token": token.(string)})
}
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	e.GET("/csrf_token", csrfToken)
	e.POST("/hello", postHello)
	e.Logger.Fatal(e.Start(":8080"))

}
