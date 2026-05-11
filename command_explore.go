package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide one location name")
	}

	areaName := args[0]
	detail, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, enc := range detail.PokemonEncounters {
		fmt.Printf("  - %s\n", enc.Pokemon.Name)
	}
	return nil
}
