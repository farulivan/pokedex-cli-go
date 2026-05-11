package main

import (
	"time"

	"github.com/farulivan/pokedex-cli-go/internal/pokeapi"
	"github.com/farulivan/pokedex-cli-go/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex.NewPokedex(),
	}
	startRepl(cfg)
}
