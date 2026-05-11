package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) getJSON(url string, out any) error {
	if value, ok := c.cache.Get(url); ok {
		return json.Unmarshal(value, out)
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, out); err != nil {
		return err
	}

	c.cache.Add(url, data)
	return nil
}
