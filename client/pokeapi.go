package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/brunopma/gokedex-cli/model"
	"github.com/patrickmn/go-cache"
)

const (
	//BaseURL of pokeapi.co
	BaseURL string = "https://pokeapi.co/api/v2/"
	//PokemonURI of pokeapi.co
	PokemonURI string = "pokemon/"
	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 30 * time.Second
)

type pokemonID int
type pokemonName string

var c *cache.Cache

func init() {
	c = cache.New(defaultCacheSettings.MinExpire, defaultCacheSettings.MaxExpire)
}

// PokeapiClient is the client for pokeapi
type PokeapiClient struct {
	client  *http.Client
	baseURL string
}

// NewPokeapiClient creates a new PokeapiClient
func NewPokeapiClient() *PokeapiClient {
	return &PokeapiClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

// SetTimeout overrides the default ClientTimeout
func (hc *PokeapiClient) SetTimeout(d time.Duration) {
	hc.client.Timeout = d
}

// FetchPokemon retrieves the Resource as per provided Pokemon parameter
func (hc *PokeapiClient) FetchPokemon(p interface{}) (model.Pokemon, error) {
	finalURL := hc.buildURL(p)

	// returning cached response if exists
	cached, found := c.Get(finalURL)
	if found && CacheSettings.UseCache {
		var cachedResource model.PokemonResponse
		json.Unmarshal(cached.([]byte), &cachedResource)
		return cachedResource.Pokemon(), nil
	}

	resp, err := hc.client.Get(finalURL)
	if err != nil {
		return model.Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return model.Pokemon{}, errors.New("Resource not found, maybe name mispelling pokemon name or inexistent search iD?")
	}

	var PokemonResp model.PokemonResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &PokemonResp); err != nil {
		return model.Pokemon{}, err
	}
	setCache(finalURL, body)
	return PokemonResp.Pokemon(), nil
}

func (hc *PokeapiClient) buildURL(parameter interface{}) string {
	switch parameter.(type) {
	case int:
		return fmt.Sprintf("%s%s%s", hc.baseURL, PokemonURI, strconv.Itoa(parameter.(int)))
	case string:
		return fmt.Sprintf("%s%s%s", hc.baseURL, PokemonURI, parameter)
	}
	return ""
}
