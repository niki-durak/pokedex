package main

import(
	"net/http"
	"encoding/json"
	"io"
	"fmt"
	"pokedex/internal"
	"time"
)

type LocationResponse struct {
	Count int
	Next *string
	Previous *string
	Results []Location
}

type Location struct {
	Name string
	Url string
}


type ExploreResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func getResponseLocation(url string) (*LocationResponse, error) {
	var lr LocationResponse
	data, ok := cache.Get(url)
	if ok{
		err := json.Unmarshal(data, &lr)
		if err != nil {
			return nil, err
		}
		return &lr, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	err = json.Unmarshal(data, &lr)
	if err != nil {
		return nil, err
	}
	cache.Add(url, data)
	return &lr, nil
}

func getResponseExplore(url string) (*ExploreResponse, error) {
	var er ExploreResponse
	data, ok := cache.Get(url)
	if ok{
		err := json.Unmarshal(data, &er)
		if err != nil {
			return nil, err
		}
		return &er, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	data, err = io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	
	err = json.Unmarshal(data, &er)
	if err != nil {
		return nil, err
	}
	cache.Add(url, data)
	return &er, nil
}



var cache = pokecache.NewCache(5 * time.Minute)