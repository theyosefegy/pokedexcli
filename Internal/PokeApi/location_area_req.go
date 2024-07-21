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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreaResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreaResp, nil
}
