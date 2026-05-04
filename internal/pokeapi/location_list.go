package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {

	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		var locationsResponse RespShallowLocations
		if err := json.Unmarshal(val, &locationsResponse); err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var locations RespShallowLocations
	if err := json.Unmarshal(body, &locations); err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, body)
	return locations, nil

}
