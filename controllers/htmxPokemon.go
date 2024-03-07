package controllers

import (
	"gorm.io/gorm"
)

type HTMXPokemonController struct {
	gorm *gorm.DB
}

func CreateHTMXPokemonController(db *gorm.DB) *HTMXPokemonController {
	return &HTMXPokemonController{gorm: db}
}

// func (controller HTMXPokemonController) GetPokemons(c echo.Context) error {
// 	var pokemons []models.Pokemon
// 	if err := controller.gorm.Find(&pokemons).Error; err != nil {
// 		return c.String(http.StatusOK, "Failed to fetch pokemons from database");
// 	}

// }
