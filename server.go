package main

import (
	"goweb/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
	_, err := db.Connect()
	if err != nil {
		e.Logger.Errorf("Error connecting to database: %v", err)
	}
	e.Logger.Fatal(e.Start(":1323"))
}
