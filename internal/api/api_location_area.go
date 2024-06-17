package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/benskia/PokeDex/internal/cache"
)

type LocationAreaResponse struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
