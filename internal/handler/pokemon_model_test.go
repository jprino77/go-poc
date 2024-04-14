package handler

import (
	"github.com/jprino77/go-poc/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fromDomain(t *testing.T) {
	type args struct {
		pokemon domain.Pokemon
	}
	tests := []struct {
		name string
		args args
		want *pokemonResponse
	}{
		{
			name: "Should convert a domain pokemon to a response pokemon",
			args: args{
				pokemon: *domain.NewPokemon(1, "bulbasaur", 1),
			},
			want: &pokemonResponse{
				Id:    1,
				Name:  "bulbasaur",
				Order: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fromDomain(tt.args.pokemon))
		})
	}
}
