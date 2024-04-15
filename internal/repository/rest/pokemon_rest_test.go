package rest

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/jprino77/go-poc/internal/domain"
	"github.com/jprino77/go-poc/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPokemonRest(t *testing.T) {
	tests := []struct {
		name string
		want *PokemonRest
	}{
		{
			name: "Create new pokemon rest",
			want: &PokemonRest{
				client: resty.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPokemonRest()
			assert.IsType(t, tt.want, got)
		})
	}
}

func TestPokemonRest_GetPokemonById(t *testing.T) {
	type args struct {
		id  int
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Pokemon
		wantErr error
		mock    func()
	}{
		{
			name: "Return a pokemon from api, when request is successful",
			args: args{
				id:  1,
				ctx: context.Background(),
			},
			want:    testdata.PokemonDomain(),
			wantErr: nil,
			mock: func() {
				httpmock.RegisterResponder(
					"GET",
					"/pokemon/1",
					httpmock.NewStringResponder(200, `{"id": 1, "name": "bulbasaur", "order": 1}`),
				)
			},
		},
		{
			name: "Returns an error when request fails",
			args: args{
				id:  1,
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: errors.New("an error"),
			mock: func() {
				httpmock.RegisterResponder(
					"GET",
					"/pokemon/1",
					httpmock.NewErrorResponder(assert.AnError),
				)
			},
		},
		{
			name: "Returns an error when parse response fails",
			args: args{
				id:  1,
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: errors.New("an error when parse response"),
			mock: func() {
				httpmock.RegisterResponder(
					"GET",
					"/pokemon/1",
					httpmock.NewStringResponder(200, `some invalid json`),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PokemonRest{
				client: resty.New(),
			}
			if tt.mock != nil {
				httpmock.ActivateNonDefault(p.client.GetClient())
				tt.mock()
			}

			got, err := p.GetPokemonById(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
