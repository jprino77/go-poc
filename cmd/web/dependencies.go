package web

import (
	"github.com/jprino77/go-poc/internal/app"
	"github.com/jprino77/go-poc/internal/handler"
	"github.com/jprino77/go-poc/internal/rest"
)

type Dependencies struct {
	PokemonHandler *handler.PokemonHandler
}

func InitDependencies() *Dependencies {

	pokemonRest := rest.NewPokemonRest()
	pokemonSrv := app.NewPokemonService(pokemonRest)
	pokemonHdr := handler.NewPokemonHandler(pokemonSrv)

	return &Dependencies{
		pokemonHdr,
	}
}
