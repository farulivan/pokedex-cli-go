package pokeapi

func (c *Client) ListLocations(pageURL *string) (ListLocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var listLocations ListLocationArea
	if err := c.getJSON(url, &listLocations); err != nil {
		return ListLocationArea{}, err
	}
	return listLocations, nil
}

func (c *Client) GetLocationArea(name string) (LocationAreaDetail, error) {
	url := baseURL + "/location-area/" + name

	var detail LocationAreaDetail
	if err := c.getJSON(url, &detail); err != nil {
		return LocationAreaDetail{}, err
	}

	return detail, nil
}
