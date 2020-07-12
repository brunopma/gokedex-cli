package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// PokemonResponse is a struct that stores the pokemon JSON response from API
type PokemonResponse struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string      `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        string      `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     string      `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       string      `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

// Pokemon is a struct that only have the desired information to our response
type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Height int    `json:"height"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Types  []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

// Pokemon is a function that store only the desired information
// from the API to the Pokemon struct
func (pr PokemonResponse) Pokemon() Pokemon {
	return Pokemon{
		Abilities: pr.Abilities,
		Height:    pr.Height,
		ID:        pr.ID,
		Name:      pr.Name,
		Types:     pr.Types,
		Weight:    pr.Weight,
	}
}

// PrettyString creates a pretty string of the Pokemon that we'll use as output
func (p Pokemon) PrettyString() string {
	s := fmt.Sprintf(
		"#%d\nName: %s\nHeight: %.1fm\nWeight: %.1fKg\nTypes: %s\nAbilities: %s",
		p.ID, strings.Title(p.Name), float64(p.Height)/10.0, float64(p.Weight)/10.0,
		p.stringifyPokemon("Types"), p.stringifyPokemon("Abilities"))
	return s
}

// JSON converts the Pokemon struct to JSON, we'll use the JSON string as output
func (p Pokemon) JSON() string {
	pJSON, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(pJSON)
}

// stringifyPokemon iterates over some slice field of Pokemon array and returns
// a pretty-formated string
func (p Pokemon) stringifyPokemon(field string) string {
	value := reflect.ValueOf(p)
	structField := value.FieldByName(field)

	if !structField.IsValid() {
		return ""
	}
	if structField.Type().Kind() != reflect.Slice {
		return ""
	}

	resultList := make([]string, 0)
	for i := 0; i < structField.Len(); i++ {
		f1 := structField.Index(i)
		if f1.Type().Kind() != reflect.Struct {
			continue
		}

		for j := 0; j < f1.NumField(); j++ {
			f2 := f1.Field(j)
			if f2.Type().Kind() != reflect.Struct {
				continue
			}
			name := f2.FieldByName("Name")
			if name.Kind() != reflect.String {
				continue
			}
			resultList = append(resultList, name.String())
		}
	}

	return strings.Join(resultList, " / ")
}
