package main

import (
	"fmt"
)

func commandExplore(cfg *config, arg string) error {
	resp, err := cfg.pokeapiClient.GetLocation(&arg)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", arg)
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
