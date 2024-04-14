package rest

import "github.com/jprino77/go-poc/internal/domain"

type PokemonResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (p PokemonResponse) ToDomain() *domain.Pokemon {
	return domain.NewPokemon(p.Id, p.Name, p.Order)
}
