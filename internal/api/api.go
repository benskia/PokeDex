package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ApiResponse struct {
	Count   int        `json:"count"`
	Next    *string    `json:"next"`
	Prev    *string    `json:"previous"`
	Results []Location `json:"results"`
}

type Config struct {
	Next *string
	Prev *string
}

func FetchLocations(url string) (*ApiResponse, error) {
	response, err := http.Get(url)
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

	locations := ApiResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Struct.")
		return nil, err
	}

	return &locations, nil
}
