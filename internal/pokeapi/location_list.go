package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cached, ok := c.Cache.Get(url); ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cached, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.Cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetLocation(locationArea *string) (RespDeepLocation, error) {
	url := baseURL + "/location-area/" + *locationArea

	if cached, ok := c.Cache.Get(url); ok {
		locationResp := RespDeepLocation{}
		err := json.Unmarshal(cached, &locationResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	c.Cache.Add(url, dat)

	locationResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	return locationResp, nil
}
