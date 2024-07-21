package commands

import (
	"errors"
	"fmt"
)

// This 'Map' command displays the locations of all pokemons like in the pokedex :)
func Map(cfg *Config) error {

	resp, err := cfg.PokeClient.ListLocationArea(cfg.NextLocationAreaURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Location Areas:\n")
	for i, area := range resp.Results {
		fmt.Printf(" %.2d- Location: '%s'\n", i+1, area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PreLocationAreaURL = resp.Previous
	return nil
} 

func Mapb(cfg *Config) error {
	resp, err := cfg.PokeClient.ListLocationArea(cfg.PreLocationAreaURL)
	
	if cfg.PreLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Location Areas:\n")
	for i, area := range resp.Results {
		fmt.Printf(" %.2d- Location: '%s'\n", i+1, area.Name)
	}
	cfg.NextLocationAreaURL = resp.Next
	cfg.PreLocationAreaURL = resp.Previous
	return nil
} 