package main

import (
	"github.com/farulivan/pokedex-cli-go/internal/pokeapi"
	"github.com/farulivan/pokedex-cli-go/internal/pokedex"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          *pokedex.Pokedex
}
