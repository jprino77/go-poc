package testdata

import "github.com/jprino77/go-poc/internal/domain"

func PokemonDomain() *domain.Pokemon {
	return domain.NewPokemon(1, "bulbasaur", 1)
}
