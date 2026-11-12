package main

import(
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 2 {
    	return fmt.Errorf("usage: inspect <pokemon_name>")
	}
	name := args[1]
	pokemon, ok := cfg.Pokedex[name]
	if !ok{
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
    	fmt.Printf("- %s\n", pokemonType.Type.Name)
	}
	return nil
}