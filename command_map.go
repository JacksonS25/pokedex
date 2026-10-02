package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(cfg *config) error {
	for i := cfg.previous; i < cfg.next; i++ {
		fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", i)
		res, err := http.Get(fullURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		var locationArea map[string]interface{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
		}

		fmt.Println(locationArea["name"])
	}

	cfg.previous = cfg.next
	cfg.next += 20
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous <= 21 {
		fmt.Println("No previous location areas visited.")
		return nil
	}

	cfg.next -= 40
	cfg.previous -= 40

	for i := cfg.previous; i < cfg.next; i++ {
		fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", i)
		res, err := http.Get(fullURL)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		var locationArea map[string]interface{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&locationArea); err != nil {
			return err
		}

		fmt.Println(locationArea["name"])
	}

	cfg.previous = cfg.next
	cfg.next += 20
	return nil
}
