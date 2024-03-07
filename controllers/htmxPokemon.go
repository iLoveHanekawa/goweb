package controllers

import (
	"context"
	"goweb/models"
	"goweb/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HTMXPokemonController struct {
	gorm *gorm.DB
}

func CreateHTMXPokemonController(db *gorm.DB) *HTMXPokemonController {
	return &HTMXPokemonController{gorm: db}
}

func (controller HTMXPokemonController) AddPokemon(c echo.Context) error {
	pokemon := &models.Pokemon{}
	if err := c.Bind(pokemon); err != nil {
		return c.String(http.StatusOK, "Failed to bind parameters to pokemon struct")
	}
	if err := controller.gorm.Create(&pokemon).Error; err != nil {
		return c.String(http.StatusOK, "Failed to fetch pokemons from database")
	}
	component := views.Pokemon(pokemon.ID, pokemon.Name, pokemon.Type, pokemon.Level)
	return component.Render(context.Background(), c.Response().Writer)
}

func (controller HTMXPokemonController) GetPokemon(c echo.Context) error {
	var pokemon = models.Pokemon{}
	id := c.Param("id")
	if err := c.Bind(&pokemon); err != nil {
		return c.String(http.StatusOK, "Failed to bind parameters to pokemon struct")
	}
	if err := controller.gorm.First(&pokemon, id).Error; err != nil {
		return c.String(http.StatusOK, "Failed to fetch pokemons from database")
	}
	component := views.Pokemon(pokemon.ID, pokemon.Name, pokemon.Type, pokemon.Level)
	return component.Render(context.Background(), c.Response().Writer)
}

func (controller HTMXPokemonController) DeletePokemon(c echo.Context) error {
	pokemon := new(models.Pokemon)
	if err := c.Bind(&pokemon); err != nil {
		return c.String(http.StatusOK, "Failed to bind parameters to pokemon struct")
	}
	if err := controller.gorm.Delete(pokemon).Error; err != nil {
		return c.String(http.StatusOK, "Failed to delete pokemon record.")
	}
	return c.HTML(http.StatusOK, "")

}
