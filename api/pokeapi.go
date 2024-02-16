package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LocationArea struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

type HttpClient struct {
	httpClient http.Client
}

func CreateHttpClient() HttpClient {
	return HttpClient{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocConfig struct {
	nextLocationUrl *string
	prevLocationUrl *string
}

const pokeUrl = "https://pokeapi.co/api/v2"

func (c *HttpClient) LocationAreaRequest(config *LocConfig, forward bool) (LocationArea, error) {
	endpoint := "/location/?offset=0&limit=20"
	url := pokeUrl + endpoint
	if forward {
		if config.nextLocationUrl != nil {
			url = *config.nextLocationUrl
		}
	} else {
		if config.prevLocationUrl != nil {
			url = *config.prevLocationUrl
		} else {
			return LocationArea{}, fmt.Errorf("error. you are already on the first page")
		}
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return LocationArea{}, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("error. status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll((response.Body))
	if err != nil {
		return LocationArea{}, err
	}

	locationResponse := LocationArea{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationArea{}, err
	}

	config.nextLocationUrl = locationResponse.Next
	config.prevLocationUrl = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}

	return locationResponse, nil
}
