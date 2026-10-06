package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	resp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	r := rand.Intn(resp.BaseExperience)

	if r < resp.BaseExperience/2 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.pokedex[name] = resp
	return nil
}
