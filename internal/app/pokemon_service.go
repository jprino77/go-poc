package app

import (
	"context"
	"github.com/jprino77/go-poc/internal/domain"
)

type PokemonService struct {
	pokemonRepository domain.PokemonRepository
}

func NewPokemonService(pokemonRepository domain.PokemonRepository) *PokemonService {
	return &PokemonService{
		pokemonRepository,
	}
}

func (ps PokemonService) GetPokemonBy(ctx context.Context, id int) (*domain.Pokemon, error) {
	return ps.pokemonRepository.GetPokemonById(ctx, id)
}
