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

	e.Static("/static", "assets")
	e.GET("/", func(c echo.Context) error {
		component := views.Html(views.Home(), c.Request())
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/pokemons", func(c echo.Context) error {
		var pokemons []models.Pokemon
		if err := gorm.Find(&pokemons).Error; err != nil {
			return c.String(http.StatusOK, "Something went wrong while quering the database")
		}
		component := views.Html(views.PokemonPage(pokemons), c.Request())
		return component.Render(context.Background(), c.Response().Writer)
	})

	pokemonController := controllers.CreatePokemonController(gorm)

	// json rest routes
	e.GET("/api/v1/pokemons", pokemonController.GetPokemons)
	e.GET("/api/v1/pokemons/:id", pokemonController.GetPokemon)
	e.POST("/api/v1/pokemons", pokemonController.AddPokemon)
	e.PATCH("/api/v1/pokemons", pokemonController.EditLevel)
	e.DELETE("/api/v1/pokemons", pokemonController.DeletePokemon)

	// HTML rest routes

	htmxPokemonController := controllers.CreateHTMXPokemonController(gorm)

	e.POST("/api/v2/pokemons", htmxPokemonController.AddPokemon)
	e.GET("/api/v2/pokemons/:id", htmxPokemonController.GetPokemon)
	e.DELETE("/api/v2/pokemons", htmxPokemonController.DeletePokemon)

	e.Logger.Fatal(e.Start(":1323"))
}
