package domain

import "context"

type Pokemon struct {
	id    int
	name  string
	order int
}

func NewPokemon(id int, name string, order int) *Pokemon {
	return &Pokemon{
		id:    id,
		name:  name,
		order: order,
	}
}

func (p Pokemon) Id() int {
	return p.id
}

func (p Pokemon) Name() string {
	return p.name
}

func (p Pokemon) Order() int {
	return p.order
}

type PokemonRepository interface {
	GetPokemonById(ctx context.Context, id int) (*Pokemon, error)
}

type PokemonSrv interface {
	GetPokemonBy(ctx context.Context, id int) (*Pokemon, error)
}
