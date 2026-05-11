package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (ListLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return ListLocationArea{}, err
	}
	defer res.Body.Close()

	var listLocations ListLocationArea

	err = json.NewDecoder(res.Body).Decode(&listLocations)
	if err != nil {
		return ListLocationArea{}, err
	}

	return listLocations, nil
}
