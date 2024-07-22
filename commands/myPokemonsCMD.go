package commands

import (
	"fmt"
)

func MyPokemons(cfg *Config, args ...string) error {
	if len(cfg.CaughtPokemons) == 0 {
		return fmt.Errorf("you did not caught any pokemon yet")
	}

	fmt.Println("Your caught pokemons list:")
	for _, mypokemon := range cfg.CaughtPokemons {
		fmt.Printf("-- %s : %v\n", mypokemon.Name, mypokemon.ID)
	}

	return nil
}
