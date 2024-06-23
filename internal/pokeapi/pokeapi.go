package pokeapi

import (
	"encoding/json"
	"errors"
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

func RequestLocationAreas(url *string, c *cache.Cache) (*LocationAreaResponse, error) {
	if url == nil {
		return nil, errors.New("Tried to query with a null endpoint value.")
	}
	body, err := getByteData(*url, c)
	if err != nil {
		return nil, err
	}
	locations := LocationAreaResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Struct.")
		return nil, err
	}
	return &locations, nil
}

func RequestAreaDetails(url *string, c *cache.Cache) (*LocationArea, error) {
	if url == nil {
		return nil, errors.New("Tried to query with a null endpoint value.")
	}
	body, err := getByteData(*url, c)
	if err != nil {
		return nil, err
	}
	areaDetails := LocationArea{}
	err = json.Unmarshal(body, &areaDetails)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Struct.")
		return nil, err
	}
	return &areaDetails, nil
}

func RequestPokemonDetails(url *string, c *cache.Cache) (*PokemonDetails, error) {
	if url == nil {
		return nil, errors.New("Tried to query with a null endpoint value.")
	}
	body, err := getByteData(*url, c)
	if err != nil {
		return nil, err
	}
	pokemon := PokemonDetails{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Struct.")
		return nil, err
	}
	return &pokemon, nil
}
