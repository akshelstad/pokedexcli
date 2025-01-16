package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	_ "github.com/akshelstad/pokedexcli/internal/pokeapi"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s", pokemon.Name)

	count := 0
	maxCount := 5

	for {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		count++
		if count == maxCount {
			fmt.Println()
			break
		}
	}

	randPercent := rand.Float32()
	baseExp := float32(pokemon.BaseExperience)
	expCeiling := float32(400.0)
	difficultyChance := 1.0 - (baseExp / expCeiling)

	var state int
	if randPercent < difficultyChance {
		state = 1
	} else {
		state = 0
	}

	if state == 1 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s got away...\n", pokemon.Name)
	}

	return nil
}
