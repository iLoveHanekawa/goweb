package main

import (
	"goweb/controllers"
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
		component := views.Html(views.Home(), c.Request())
		return component.Render(context.Background(), c.Response().Writer)
	})

	pokemonController := controllers.CreatePokemonController(gorm)

	e.GET("/pokemons", pokemonController.GetPokemons)

	e.Static("/static", "assets")

	e.GET("/api/v1/pokemons/:id", pokemonController.GetPokemon)

	e.POST("/api/v1/pokemons", pokemonController.AddPokemon)

	e.PATCH("/api/v1/pokemons", pokemonController.EditLevel)

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
