package controllers

import (
	"goweb/models"
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
	return c.JSON(http.StatusOK, pokemons)
}

func (controller PokemonController) GetPokemon(c echo.Context) error {
	var pokemon models.Pokemon
	id := c.Param("id")
	if result := controller.gorm.First(&pokemon, id); result.Error != nil {
		return c.String(http.StatusOK, "Something went wrong while quering the database")
	}
	return c.JSON(http.StatusOK, pokemon)
}

func (controller PokemonController) AddPokemon(c echo.Context) error {
	pokemon := new(models.Pokemon)
	if err := c.Bind(pokemon); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	if err := controller.gorm.Create(&pokemon).Error; err != nil {
		return c.String(http.StatusOK, "Something went wrong while writing to the database.")
	}
	return c.JSON(http.StatusOK, pokemon)
}

func (controller PokemonController) EditLevel(c echo.Context) error {
	pokemon := new(models.Pokemon)
	if err := c.Bind(pokemon); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	if err := controller.gorm.Model(pokemon).Updates(pokemon).First(pokemon).Error; err != nil {
		return c.String(http.StatusOK, "Something went wrong while writing to the database.")
	}
	return c.JSON(http.StatusOK, pokemon)
}

func (controller PokemonController) DeletePokemon(c echo.Context) error {
	pokemon := new(models.Pokemon)
	if err := c.Bind(pokemon); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	if err := controller.gorm.Delete(pokemon).Error; err != nil {
		return c.JSON(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Delete operation successful",
	})
}
