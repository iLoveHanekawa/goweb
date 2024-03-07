package main

import (
	"goweb/controllers"
	"goweb/db"

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

	// json rest routes

	e.Static("/static", "assets")
	e.GET("/pokemons", pokemonController.GetPokemons)
	e.GET("/api/v1/pokemons/:id", pokemonController.GetPokemon)
	e.POST("/api/v1/pokemons", pokemonController.AddPokemon)
	e.PATCH("/api/v1/pokemons", pokemonController.EditLevel)
	e.DELETE("/api/v1/pokemons", pokemonController.DeletePokemon)

	e.Logger.Fatal(e.Start(":1323"))
}
