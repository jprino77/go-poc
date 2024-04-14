package handler

import (
	"encoding/json"
	"github.com/jprino77/go-poc/internal/domain"
	"log"
	"net/http"
	"strconv"
)

type PokemonHandler struct {
	pokemonSrv domain.PokemonSrv
}

func NewPokemonHandler(pokemonSrv domain.PokemonSrv) *PokemonHandler {
	return &PokemonHandler{pokemonSrv: pokemonSrv}
}

func (p PokemonHandler) GetPokemonById(w http.ResponseWriter, r *http.Request) {

	r.Context()
	id, _ := strconv.Atoi(r.PathValue("id"))

	pokemon, err := p.pokemonSrv.GetPokemonBy(r.Context(), id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Pokemon found ", *pokemon)

	err = json.NewEncoder(w).Encode(fromDomain(*pokemon))

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
