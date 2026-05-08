package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		var locationResp Location
		if err := json.Unmarshal(val, &locationResp); err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return Location{}, fmt.Errorf("location does not exist: %s", locationName)
	}

	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("bad statusCode: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var locationResp Location
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, body)
	return locationResp, nil
}
