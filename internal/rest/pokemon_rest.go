package rest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/jprino77/go-poc/internal/domain"
	"strconv"
	"time"
)

const url = "https://pokeapi.co/api/v2"

type PokemonRest struct {
	client *resty.Client
}

func NewPokemonRest() *PokemonRest {
	client := resty.New().
		SetBaseURL(url).
		SetTimeout(15 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(15 * time.Second)

	return &PokemonRest{client: client}
}

func (p PokemonRest) GetPokemonById(_ context.Context, id int) (*domain.Pokemon, error) {

	response, err := p.client.
		R().
		SetPathParam("id", strconv.Itoa(id)).
		Get("/pokemon/{id}")

	if err != nil {
		return nil, errors.New("an error")
	}

	var pokemonResponse = &PokemonResponse{}

	if err := json.Unmarshal(response.Body(), pokemonResponse); err != nil {
		return nil, errors.New("an error when parse response")
	}

	return pokemonResponse.ToDomain(), nil
}
