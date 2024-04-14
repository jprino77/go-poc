package app

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jprino77/go-poc/internal/domain"
	mockdomain "github.com/jprino77/go-poc/mocks"
	"github.com/jprino77/go-poc/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type pokemonServiceDependencies struct {
	pokemonRepository *mockdomain.MockPokemonRepository
}

func makePokemonServiceDependencies(t *testing.T) pokemonServiceDependencies {
	return pokemonServiceDependencies{
		pokemonRepository: mockdomain.NewMockPokemonRepository(gomock.NewController(t)),
	}
}

func TestNewPokemonService(t *testing.T) {
	repository := makePokemonServiceDependencies(t).pokemonRepository
	type args struct {
		pokemonRepository domain.PokemonRepository
	}
	tests := []struct {
		name string
		args args
		want *PokemonService
	}{
		{
			name: "Should return a new PokemonService instance",
			args: args{
				pokemonRepository: repository,
			},
			want: &PokemonService{
				pokemonRepository: repository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPokemonService(tt.args.pokemonRepository)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPokemonService_GetPokemonBy(t *testing.T) {
	pokemon := testdata.PokemonDomain()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Pokemon
		wantErr error
		mock    func(pokemonServiceDependencies)
	}{
		{
			name: "Given a valid pokemon id, it should return the pokemon",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want:    pokemon,
			wantErr: nil,
			mock: func(dependencies pokemonServiceDependencies) {
				dependencies.pokemonRepository.EXPECT().
					GetPokemonById(context.Background(), 1).
					Return(pokemon, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dependencies := makePokemonServiceDependencies(t)

			if tt.mock != nil {
				tt.mock(dependencies)
			}

			ps := NewPokemonService(dependencies.pokemonRepository)

			got, err := ps.GetPokemonBy(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
