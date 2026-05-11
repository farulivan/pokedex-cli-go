package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const catchThreshold = 50

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one pokemon name to catch")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if rand.IntN(pokemon.BaseExperience) < catchThreshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex.Add(pokemon)
	} else {
		fmt.Printf("%s was escaped!\n", pokemon.Name)
	}

	return nil
}
