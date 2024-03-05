package main

import (
	"goweb/db"
	"goweb/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dbInstance, err := db.Connect()
	if err != nil {
		e.Logger.Errorf("Error connecting to database: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})

	e.GET("/api/v1/pokemons", func(c echo.Context) error {
		var pokemons []models.Pokemon
		if err := dbInstance.Find(&pokemons).Error; err != nil {
			return c.String(http.StatusOK, "Something went wrong while quering the database")
		}
		return c.JSON(http.StatusOK, pokemons)
	})

	e.GET("/api/v1/pokemons/:id", func(c echo.Context) error {
		var pokemon models.Pokemon
		id := c.Param("id")
		if result := dbInstance.First(&pokemon, id); result.Error != nil {
			return c.String(http.StatusOK, "Something went wrong while quering the database")
		}
		return c.JSON(http.StatusOK, pokemon)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
