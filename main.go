package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brunopma/gokedex-cli/client"
	"github.com/brunopma/gokedex-cli/model"
)

func main() {
	PokemonID := flag.Int(
		"i", 0, "Pokemon ID to fetch, currently there's 807 available pokemon!",
	)
	PokemonName := flag.String(
		"n", "text", "Pokemon ID to fetch",
	)
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)
	flag.Parse()

	pokeapiClient := client.NewPokeapiClient()
	pokeapiClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	var pokemonResp model.Pokemon
	var err error
	switch {
	case *PokemonID > 0:
		pokemonResp, err = pokeapiClient.FetchPokemon(*PokemonID)
	case *PokemonName != "text":
		pokemonResp, err = pokeapiClient.FetchPokemon(*PokemonName)
	default:
		flag.Usage()
		os.Exit(0)
	}
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(pokemonResp.JSON())
	} else {
		fmt.Println(pokemonResp.PrettyString())
	}
}
