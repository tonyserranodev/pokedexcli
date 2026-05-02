// Package pokeapi provides functions for interacting with the Pokemon API at https://pokeapi.co.
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tonyserranodev/pokedexcli/internal/pokecache"
)

const cacheInterval = 5 * time.Second

var cache = pokecache.NewCache(cacheInterval)

type LocationsResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(pageURL *string) (LocationsResponse, error) {

	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := cache.Get(url); ok {
		var locations LocationsResponse
		if err := json.Unmarshal(val, &locations); err != nil {
			return LocationsResponse{}, err
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, err
	}

	cache.Add(url, body)

	var locations LocationsResponse
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationsResponse{}, err
	}
	return locations, nil

}
