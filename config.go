package main

import "github.com/farulivan/pokedex-cli-go/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}
