package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pokemon-service/routes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonList(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/pokemon", nil)
	rec := httptest.NewRecorder()

	router := routes.RegisterRoutes()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var pokemons []map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &pokemons)
	assert.NoError(t, err)
	assert.Greater(t, len(pokemons), 0)
	assert.Contains(t, pokemons[0], "name")
	assert.Contains(t, pokemons[0], "image")
}
