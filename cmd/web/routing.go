package web

import (
	"github.com/jprino77/go-poc/cmd/web/middlewere"
	"net/http"
)

func InitRoutes(dependencies *Dependencies) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/pokemons/{id}", dependencies.PokemonHandler.GetPokemonById)

	return addMiddlewares(mux)
}

func addMiddlewares(handler http.Handler) http.Handler {
	contentType := middlewere.NewResponseHeader(handler, "Content-Type", "application/json")
	customType := middlewere.NewResponseHeader(contentType, "X-Custom-Header", "custom")
	logger := middlewere.NewLogger(customType)

	return logger
}
