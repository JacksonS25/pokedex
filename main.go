package main

import (
	"time"

	"github.com/JacksonS25/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokedex:	   make(map[string]pokeapi.RespPokemon),
	}

	startRepl(cfg)
}
