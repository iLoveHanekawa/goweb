package controllers

import (
	"context"
	"goweb/models"
	"goweb/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PokemonController struct {
	gorm *gorm.DB
}

func CreatePokemonController(db *gorm.DB) *PokemonController {
	return &PokemonController{gorm: db}
}

func (controller PokemonController) GetPokemons(c echo.Context) error {
	var pokemons []models.Pokemon
	if err := controller.gorm.Find(&pokemons).Error; err != nil {
		return c.String(http.StatusOK, "Something went wrong while quering the database")
	}
	component := views.Html(views.PokemonPage(pokemons), c.Request())
	return component.Render(context.Background(), c.Response().Writer)
}

func (controller PokemonController) GetPokemon(c echo.Context) error {
	var pokemon models.Pokemon
	id := c.Param("id")
	if result := controller.gorm.First(&pokemon, id); result.Error != nil {
		return c.String(http.StatusOK, "Something went wrong while quering the database")
	}
	return c.JSON(http.StatusOK, pokemon)
}
