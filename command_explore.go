package main

import(
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 2 {
    	return fmt.Errorf("usage: explore <area_name>")
	}
	location := args[1]
	url := "https://pokeapi.co/api/v2/location-area/"+location+"/"
	res, err := getResponseExplore(url)
	if err != nil{
		return err
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters{
		fmt.Println("-", encounter.Pokemon.Name)
	}
	return nil
}