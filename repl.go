package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JacksonS25/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.RespPokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}

		commandName := text[0]

		command, exists := cfg.commands[commandName]
		if exists {
			err := command.callback(cfg, strings.Join(text[1:], " "))
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location and list pokemon. Use: explore <location>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon. Use: catch <pokemon>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Retrieve details on Pokemon in your Pokedex. Use: inspect <pokemon>",
			callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	cleanedText := []string{}

	for _, word := range strings.Fields(text) {
		cleanedText = append(cleanedText, strings.ToLower(word))
	}

	return cleanedText
}
