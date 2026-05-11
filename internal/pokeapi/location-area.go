package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (ListLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var listLocations ListLocationArea

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &listLocations); err != nil {
			return ListLocationArea{}, err
		}
		return listLocations, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return ListLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ListLocationArea{}, err
	}

	if err := json.Unmarshal(data, &listLocations); err != nil {
		return ListLocationArea{}, err
	}

	c.cache.Add(url, data)
	return listLocations, nil
}
