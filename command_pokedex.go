package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	listPokemon := cfg.pokedex.ListPokemon()
	if len(listPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range listPokemon {
		fmt.Printf("  - %s\n", pokemon)
	}

	return nil
}
