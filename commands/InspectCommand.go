package commands

import (
	"fmt"
	"strings"
)

func Inspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("no caught pokemon was provided")
	}

	caughtPokemon := strings.ToLower(args[0])
	pokemon, isCaught := cfg.CaughtPokemons[caughtPokemon]

	if !isCaught {
		return fmt.Errorf("you didn't caught %s to inspect them", caughtPokemon)
	}

	fmt.Printf("\nName: %v\n", pokemon.Name)
	fmt.Printf(" -- Height: %v\n", pokemon.Height)
	fmt.Printf(" -- Weight: %v\n", pokemon.Weight)
	fmt.Printf(" -- Ability: %v\n", pokemon.Abilities[0].Ability.Name)

	fmt.Printf("\n%v Stats: \n", pokemon.Name)
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("\n%v Types: \n", pokemon.Name)
	for _, typ := range pokemon.Types {
		fmt.Printf(" -- %s\n", typ.Type.Name)
	}
	return nil
}
