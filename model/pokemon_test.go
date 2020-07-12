package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const defaultError = "expected %v, but result was %v."

func TestPokemon(t *testing.T) {
	//importing the response fixture from pokeapi.co/api/v2/pokemon/1
	pokemonJSON, err := os.Open("fixtures/pokemonResponse.json")
	defer pokemonJSON.Close()
	if err != nil {
		t.Fatal("Missing file fixtures/pokemonResponse.json")
	}
	byteValue, _ := ioutil.ReadAll(pokemonJSON)
	var pokemonResponse PokemonResponse
	json.Unmarshal(byteValue, &pokemonResponse)

	pokemon := pokemonResponse.Pokemon()

	// Using reflection to get info about the structs
	// like size, field values etc
	valuePokemonResponse := reflect.ValueOf(pokemonResponse)
	valuePokemon := reflect.ValueOf(pokemon)

	// Check if number of fields in struct pokemon are correct
	if valuePokemon.NumField() != 6 {
		t.Errorf(defaultError, 6, valuePokemon.NumField())
	}

	// Checking if each field in Pokemon struct has the same value of
	// PokemonResponse
	for i := 0; i < valuePokemon.NumField(); i++ {
		pokemonFieldName := valuePokemon.Type().Field(i).Name

		// transforming Values into interfaces to be able to compare the fields using the
		// cmp package
		pokemonField := valuePokemon.FieldByName(pokemonFieldName).Interface()
		pokemonResponseField := valuePokemonResponse.FieldByName(pokemonFieldName).Interface()

		if !cmp.Equal(pokemonField, pokemonResponseField) {
			t.Errorf(defaultError, pokemonField, pokemonResponseField)
		}
	}
}

func TestPrettyString(t *testing.T) {
	//importing the pokemon.json already extracted from pokemon api response
	pokemonJSON, err := os.Open("fixtures/pokemon.json")
	defer pokemonJSON.Close()
	if err != nil {
		t.Fatal("Missing file pokemon.json")
	}
	byteValue, _ := ioutil.ReadAll(pokemonJSON)
	var pokemon Pokemon
	json.Unmarshal(byteValue, &pokemon)

	pokemonString := pokemon.PrettyString()
	expectedString := "#1\nName: Bulbasaur\nHeight: 0.7m\nWeight: 6.9Kg\nTypes: grass / poison\nAbilities: overgrow / chlorophyll"
	if pokemonString != expectedString {
		t.Errorf(defaultError, pokemonString, expectedString)
	}
}

func TestPokemonJSON(t *testing.T) {
	//importing the pokemon.json already extracted from pokemon api response
	pokemonJSON, err := os.Open("fixtures/pokemon.json")
	defer pokemonJSON.Close()
	if err != nil {
		t.Fatal("Missing file pokemon.json")
	}
	byteValue, _ := ioutil.ReadAll(pokemonJSON)
	var pokemon Pokemon
	json.Unmarshal(byteValue, &pokemon)

	marshaledPokemon := pokemon.JSON()

	if cmp.Equal(marshaledPokemon, string(byteValue)) {
		t.Errorf(defaultError, string(byteValue), marshaledPokemon)
	}
}
