package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// The 'ListLocationArea' function is a method of the 'Client' struct,
func (c *Client) ListLocationArea(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint

	// if the pageURL is not 'null', it will assign it to the fullURL variable
	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		locationAreaResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreaResp)
		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreaResp, nil
	}
	fmt.Println("cache miss!")

	// Creating a new HTTP request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// The 'Do' method sends a HTTP request and returns HTTP response.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("connection lost, status code: %d", resp.StatusCode)
	}

	data, err= io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	
	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)
	
	return locationAreaResp, nil
}
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("cache hit!")
		var locationArea LocationArea
		err := json.Unmarshal(data, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("cache miss!")

	// Creating a new HTTP request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// The 'Do' method sends a HTTP request and returns HTTP response.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	// Status Code
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("connection lost, status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	
	var locationArea LocationArea
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)
	
	return locationArea, nil
}