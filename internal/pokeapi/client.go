package pokeapi

import (
	"net/http"
	"time"

	"github.com/farulivan/pokedex-cli-go/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, cacheInternval time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInternval),
	}
}
