package pokedex

import "github.com/farulivan/pokedex-cli-go/internal/pokeapi"

type Pokedex struct {
	entries map[string]pokeapi.Pokemon
}
