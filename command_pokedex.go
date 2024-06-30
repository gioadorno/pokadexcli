package main

import (
	"fmt"
)

func callbackPokdex(cfg *config, args ...string) error {
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
