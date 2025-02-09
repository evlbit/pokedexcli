package main

import (
	"time"

	"github.com/evlbit/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
		pokeapiClient: pokeapiClient,
	}

	startRepl(cfg)
}
