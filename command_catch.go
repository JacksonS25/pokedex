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

	maxBaseCatchRate := 100.0

	catchChance := maxBaseCatchRate / float64(resp.BaseExperience)

	if catchChance > 1.0 {
		catchChance = 1.0
	} else if catchChance < 0.05 {
		catchChance = 0.05
	}

	r := rand.Float64()

	if r > catchChance {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.pokedex[name] = resp
	return nil
}
