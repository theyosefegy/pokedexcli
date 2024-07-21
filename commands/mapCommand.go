package commands

import (
	"fmt"
)

// This 'Map' command displays the locations of all pokemons like in the pokedex :)
func Map(cfg *Config) {

	locationAreas, err := cfg.PokeClient.ListLocationArea(cfg.NextLocationAreaURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Location Areas:\n")
	for i, area := range locationAreas.Results {
		fmt.Printf(" %.2d- Location: '%s'\n", i+1, area.Name)
	}
	cfg.NextLocationAreaURL = locationAreas.Next
	cfg.PreLocationAreaURL = locationAreas.Previous
} 