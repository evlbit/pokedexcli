package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/evlbit/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		words := cleanInput(scanner.Text())

		if len(words) > 0 {
			command, ok := getCommands()[words[0]]
			if ok {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}

		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading stdin: %v", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
