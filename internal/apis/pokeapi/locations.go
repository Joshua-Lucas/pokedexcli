package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationsResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []Location
}

type Location struct {
	Name string
	Url  string
}

func GetLocations(url string) (*LocationsResponse, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var LocationResponse LocationsResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &LocationResponse)

	}

	response := LocationResponse
	return &response, nil
}
