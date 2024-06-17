package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/benskia/PokeDex/internal/cache"
)

const (
	LocationAreaEndpoint string = "https://pokeapi.co/api/v2/location-area/"
	PokemonEndpoint      string = "https://pokeapi.co/api/v2/pokemon/"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type apiResponse interface {
	FetchJSON()
}

func getByteData(key string, c *cache.Cache) ([]byte, error) {
	entry, ok := c.Get(key)
	if ok {
		return entry, nil
	}
	client := NewClient()
	response, err := client.httpClient.Get(key)
	if err != nil {
		fmt.Println("Error making GET request.")
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body.")
		return nil, err
	}
	c.Add(key, body)
	return body, nil
}
