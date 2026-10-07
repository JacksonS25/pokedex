package main

import (
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	pokemon, exists := cfg.pokedex[name]
	if !exists {
		return fmt.Errorf("Pokemon not found: %s", name)
	}

	// Print the Pokemon details
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
