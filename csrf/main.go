package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostRequest struct {
	Name string `json:"name"`
}

func postHello(c echo.Context) error {
	r := new(PostRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("hello %s", r.Name)})
}
func main() {
	e := echo.New()
	
	e.POST("/hello", postHello)
	e.Logger.Fatal(e.Start(":8080"))
}
