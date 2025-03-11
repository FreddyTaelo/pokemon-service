// controllers/pokemon.go
package controllers

import (
	"encoding/json"
	"net/http"
	"pokemon-service/services"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Get Pokémon List
// @Description Get a list of Pokémon with names and images.
// @Tags Pokémon
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Pokemon
// @Failure 500 {object} map[string]string
// @Router /api/pokemon [get]
func GetPokemonList(w http.ResponseWriter, r *http.Request) {
	pokemons, err := services.GetPokemonList()
	if err != nil {
		http.Error(w, "Error fetching Pokémon list", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(pokemons)
	if err != nil {
		return
	}
}

// @Summary Get Pokémon Details
// @Description Get detailed information about a Pokémon by ID.
// @Tags Pokémon
// @Accept  json
// @Produce  json
// @Param id path int true "Pokémon ID"
// @Success 200 {object} models.Pokemon
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/pokemon/{id} [get]
func GetPokemonDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	pokemon, err := services.GetPokemonDetails(id)
	if err != nil {
		http.Error(w, "Error fetching Pokémon details", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}
