package main

import (
	"time"

	"github.com/akshelstad/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := config{
		pokeapiClient: pokeClient,
	}

	startRepl(&cfg)

}
