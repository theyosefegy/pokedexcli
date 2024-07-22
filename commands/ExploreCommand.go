package commands

import (
	"errors"
	"fmt"
)

func Explore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.PokeClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemons in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}