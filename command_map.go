package main

import(
	"fmt"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	lp, err := getResponse(url)
	if err != nil{
		return err
	}
	for _, location := range lp.Results{
		fmt.Println(location.Name)
	}
	cfg.Next = lp.Next
	cfg.Previous = lp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Previous != nil {
		url = *cfg.Previous
	}
	if cfg.Previous == nil{
		fmt.Println("You're on the first page")
	}
	lp, err := getResponse(url)
	if err != nil{
		return err
	}
	for _, location := range lp.Results{
		fmt.Println(location.Name)
	}
	cfg.Next = lp.Next
	cfg.Previous = lp.Previous
	return nil
}