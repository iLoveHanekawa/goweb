package main

import (
	"goweb/db"
	"goweb/models"
	"net/http"

	"context"

	"goweb/views"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	gorm, err := db.Connect()
	if err != nil {
		e.Logger.Errorf("Error connecting to database: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		component := views.Html()
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/static", "assets")

	e.GET("/api/v1/pokemons", func(c echo.Context) error {
		var pokemons []models.Pokemon
		if err := gorm.Find(&pokemons).Error; err != nil {
			return c.String(http.StatusOK, "Something went wrong while quering the database")
		}
		return c.JSON(http.StatusOK, pokemons)
	})

	e.GET("/api/v1/pokemons/:id", func(c echo.Context) error {
		var pokemon models.Pokemon
		id := c.Param("id")
		if result := gorm.First(&pokemon, id); result.Error != nil {
			return c.String(http.StatusOK, "Something went wrong while quering the database")
		}
		return c.JSON(http.StatusOK, pokemon)
	})

	e.POST("/api/v1/pokemons", func(c echo.Context) error {
		pokemon := new(models.Pokemon)
		if err := c.Bind(pokemon); err != nil {
			return c.JSON(http.StatusOK, err)
		}
		if err := gorm.Create(&pokemon).Error; err != nil {
			return c.String(http.StatusOK, "Something went wrong while writing to the database.")
		}
		return c.JSON(http.StatusOK, pokemon)
	})

	e.PATCH("/api/v1/pokemons", func(c echo.Context) error {
		pokemon := new(models.Pokemon)
		if err := c.Bind(pokemon); err != nil {
			return c.JSON(http.StatusOK, err)
		}
		if err = gorm.Model(pokemon).Updates(pokemon).First(pokemon).Error; err != nil {
			return c.String(http.StatusOK, "Something went wrong while writing to the database.")
		}
		return c.JSON(http.StatusOK, pokemon)
	})

	e.DELETE("/api/v1/pokemons", func(c echo.Context) error {
		pokemon := new(models.Pokemon)
		if err := c.Bind(pokemon); err != nil {
			return c.JSON(http.StatusOK, err)
		}
		if err = gorm.Delete(pokemon).Error; err != nil {
			return c.JSON(http.StatusOK, err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"message": "Delete operation successful",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
