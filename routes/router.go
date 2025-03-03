// routes/router.go
package routes

import (
	"net/http"
	"pokemon-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/pokemon", controllers.GetPokemonList).Methods("GET")
	router.HandleFunc("/api/pokemon/{id:[0-9]+}", controllers.GetPokemonDetails).Methods("GET")

	// 🟢 Health Check Endpoint
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	}).Methods("GET")

	return router
} /*RegisterRoutes()*/
