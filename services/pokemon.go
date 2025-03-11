// services/pokemon.go
package services

import (
	"encoding/json"
	"fmt"
	"io"
	"pokemon-service/api"
	"pokemon-service/config"
	"pokemon-service/models"
)

var cfg = config.LoadConfig()

func GetPokemonList() ([]models.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon?limit=100", cfg.PokeAPIURL) // consider pagination if there is time
	resp, err := api.GetClient().Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var result struct {
		Results []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var pokemons []models.Pokemon
	for i, p := range result.Results {
		pokemons = append(pokemons, models.Pokemon{
			ID:    i + 1,
			Name:  p.Name,
			Image: fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/%d.png", i+1),
		})
	}
	return pokemons, nil
} /*GetPokemonList()*/

func GetPokemonDetails(id int) (*models.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%d", cfg.PokeAPIURL, id)
	resp, err := api.GetClient().Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var data struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Height  int    `json:"height"`
		Weight  int    `json:"weight"`
		Sprites struct {
			Other struct {
				OfficialArtwork struct {
					FrontDefault string `json:"front_default"`
				} `json:"official-artwork"`
			} `json:"other"`
		} `json:"sprites"`
		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
		Stats []struct {
			Stat struct {
				Name string `json:"name"`
			} `json:"stat"`
			BaseStat int `json:"base_stat"`
		} `json:"stats"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	pokemon := &models.Pokemon{
		ID:     data.ID,
		Name:   data.Name,
		Image:  data.Sprites.Other.OfficialArtwork.FrontDefault,
		Height: data.Height,
		Weight: data.Weight,
		Types:  []string{},
		Stats:  []models.Stat{},
	}

	for _, t := range data.Types {
		pokemon.Types = append(pokemon.Types, t.Type.Name)
	}
	for _, s := range data.Stats {
		pokemon.Stats = append(pokemon.Stats, models.Stat{Name: s.Stat.Name, Value: s.BaseStat})
	}

	return pokemon, nil
} /*GetPokemonDetails()*/
