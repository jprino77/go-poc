package rest

import (
	"github.com/jprino77/go-poc/internal/domain"
	"testing"
)

func TestPokemonResponse_ToDomain(t *testing.T) {

	tests := []struct {
		name            string
		pokemonResponse PokemonResponse
		want            *domain.Pokemon
	}{
		{
			name: "Convert a response pokemon to a domain pokemon",
			pokemonResponse: PokemonResponse{
				Id:    1,
				Name:  "bulbasaur",
				Order: 1,
			},
			want: domain.NewPokemon(1, "bulbasaur", 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//test without assert library
			if got := *tt.pokemonResponse.ToDomain(); got != *tt.want {
				t.Errorf("PokemonResponse.ToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
