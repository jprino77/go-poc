package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/jprino77/go-poc/internal/domain"
	mockdomain "github.com/jprino77/go-poc/mocks"
	"github.com/jprino77/go-poc/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handlerDependencies struct {
	pokemonSrv *mockdomain.MockPokemonSrv
}

func makeHandlerDependencies(t *testing.T) *handlerDependencies {
	return &handlerDependencies{
		pokemonSrv: mockdomain.NewMockPokemonSrv(gomock.NewController(t)),
	}
}

func TestNewPokemonHandler(t *testing.T) {
	srv := makeHandlerDependencies(t).pokemonSrv
	type args struct {
		pokemonSrv domain.PokemonSrv
	}
	tests := []struct {
		name string
		args args
		want *PokemonHandler
	}{
		{
			name: "Should return a new PokemonHandler instance",
			args: args{pokemonSrv: srv},
			want: &PokemonHandler{pokemonSrv: srv},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPokemonHandler(tt.args.pokemonSrv)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPokemonHandler_GetPokemonById(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name     string
		args     args
		mock     func(*handlerDependencies)
		wantBody string
		wantCode int
	}{
		{
			name: "Given a valid pokemon id, it should return the pokemon",
			args: args{id: "1"},
			mock: func(d *handlerDependencies) {
				d.pokemonSrv.EXPECT().
					GetPokemonBy(context.Background(), 1).
					Return(testdata.PokemonDomain(), nil)
			},
			wantBody: `{"id":1,"name":"bulbasaur","order":1}`,
			wantCode: 200,
		},
		{
			name: "Given an invalid pokemon id, it should return an error",
			args: args{id: "invalid"},
			mock: func(d *handlerDependencies) {
				d.pokemonSrv.EXPECT().
					GetPokemonBy(context.Background(), 0).
					Return(nil, fmt.Errorf("invalid id"))
			},
			wantBody: "",
			wantCode: 500,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dependencies := makeHandlerDependencies(t)
			if test.mock != nil {
				test.mock(dependencies)
			}

			handler := NewPokemonHandler(dependencies.pokemonSrv)

			router := http.NewServeMux()
			router.HandleFunc("/v1/pokemons/{id}", handler.GetPokemonById)

			s := &http.Server{
				Handler: router,
			}

			req := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/v1/pokemons/%s", test.args.id),
				bytes.NewReader([]byte(test.wantBody)),
			)
			rec := httptest.NewRecorder()
			s.Handler.ServeHTTP(rec, req)

			assert.Equal(t, test.wantCode, rec.Code)

			if test.wantBody != "" {
				buff := new(bytes.Buffer)
				err := json.Compact(buff, []byte(test.wantBody))
				assert.NoError(t, err)
				require.JSONEq(t, buff.String(), rec.Body.String())
			}
		})
	}
}
