package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	if len(resp.Results) == 0 {
		fmt.Println("no location area")
		return nil
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	if len(resp.Results) == 0 {
		fmt.Println("no location area")
		return nil
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
