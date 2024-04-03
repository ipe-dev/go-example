package main

import (
	"fmt"
	"net/http"

	  "github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type PostRequest struct {
	Name string `json:"name"`
}
type Session struct {
	csrfToken string
}
func postHello(c echo.Context) error {
	token := c.Get("session").(*Session)
	requestToken := c.Request().Header.Get("X-XSRF-TOKEN")
	fmt.Println(requestToken)
	if token.csrfToken != requestToken {
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
	sess, _ := session.Get("session", c)
	sess.Values["csrfToken"] = token
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, map[string]string{"token": token.(string)})
}
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))
	e.GET("/csrf_token", csrfToken)
	e.POST("/hello", postHello)
	e.Logger.Fatal(e.Start(":8080"))

}
