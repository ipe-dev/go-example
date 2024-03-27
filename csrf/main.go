package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
type PostRequest struct {
	Name string `json:"name"`
}
func postHello(c echo.Context) {
	request := c.Bind(i interface{})
	return c.JSON(http.StatusOK, )
}
func main() {

}