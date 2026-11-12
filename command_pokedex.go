package main

import(
	"fmt"
)

func commandPokedex(cfg *config, _ []string) error {
	if len(cfg.Pokedex) == 0{
		fmt.Println("Your pokedex is empty")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex{
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}