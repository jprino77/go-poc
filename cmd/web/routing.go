package web

import (
	"github.com/jprino77/go-poc/cmd/web/middleware"
	"net/http"
)

func InitRoutes(dependencies *Dependencies) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/pokemons/{id}", dependencies.PokemonHandler.GetPokemonById)

	return addMiddlewares(mux)
}

func addMiddlewares(handler http.Handler) http.Handler {
	contentType := middleware.NewResponseHeader(handler, "Content-Type", "application/json")
	customType := middleware.NewResponseHeader(contentType, "X-Custom-Header", "custom")
	logger := middleware.NewLogger(customType)

	return logger
}
