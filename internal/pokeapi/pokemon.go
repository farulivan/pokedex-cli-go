package pokeapi

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	var pokemon Pokemon
	if err := c.getJSON(url, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
