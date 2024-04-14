package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPokemon(t *testing.T) {
	type args struct {
		id    int
		name  string
		order int
	}
	tests := []struct {
		name string
		args args
		want *Pokemon
	}{
		{
			name: "Should create a new pokemon instance",
			args: args{
				id:    1,
				name:  "bulbasaur",
				order: 1,
			},
			want: &Pokemon{
				id:    1,
				name:  "bulbasaur",
				order: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPokemon(tt.args.id, tt.args.name, tt.args.order))

		})
	}
}

func TestPokemon_Id(t *testing.T) {
	tests := []struct {
		name   string
		domain *Pokemon
		want   int
	}{
		{
			name:   "Given a valid pokemon, should return its id",
			domain: NewPokemon(1, "bulbasaur", 1),
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.domain.Id())
		})
	}
}

func TestPokemon_Name(t *testing.T) {

	tests := []struct {
		name   string
		domain *Pokemon
		want   string
	}{
		{
			name:   "Given a valid pokemon, should return its name",
			domain: NewPokemon(1, "bulbasaur", 1),
			want:   "bulbasaur",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.domain.Name())
		})
	}
}

func TestPokemon_Order(t *testing.T) {

	tests := []struct {
		name   string
		domain *Pokemon
		want   int
	}{
		{
			name:   "Given a valid pokemon, should return its order",
			domain: NewPokemon(1, "bulbasaur", 1),
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.domain.Order())
		})
	}
}
