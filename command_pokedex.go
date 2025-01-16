package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("command takes no arguments")
	}

	if len(cfg.caughtPokemon) == 0 {
		return errors.New("pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for _, item := range cfg.caughtPokemon {
		fmt.Printf("  - %v\n", item.Name)
	}
	return nil
}
