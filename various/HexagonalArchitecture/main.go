package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	http.Handle("/", e)
	g := e.Group("/api/v1")
	g.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"name": "ダミーユーザー",
			"mail": "dummy@example.com",
		})
	})
	e.Logger.Fatal(e.Start(":8085"))
}
