package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/akshelstad/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")

	for {

		fmt.Print("Pokedex >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("invalid command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists 20 location areas. Repeat to list 20 more",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 location areas. Repeat to list the previous 20.",
			callback:    callbackMapB,
		},
		"explore": {
			name:        "explore",
			description: "Pass a location area as an argument to retrieve a list of pokemon in the location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw a pokeball at a pokemon and try to catch it",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists information from the pokedex about a pokemon that you have caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemon in the pokedex that you have caught",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Turns off Pokedex",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
