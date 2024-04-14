package handler

import "github.com/jprino77/go-poc/internal/domain"

type pokemonResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func fromDomain(pokemon domain.Pokemon) *pokemonResponse {
	return &pokemonResponse{
		Id:    pokemon.Id(),
		Name:  pokemon.Name(),
		Order: pokemon.Order(),
	}
}
