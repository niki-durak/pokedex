package main

import(
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error{
	if len(args) < 2 {
    	return fmt.Errorf("usage: catch <pokemon_name>")
	}
	pokemon := args[1]
	url := "https://pokeapi.co/api/v2/pokemon/"+pokemon+"/"
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	res, err := getPokemonResponse(url)
	if err != nil {
		return err
	}

	chanceToCatch := 100 - res.BaseExperience/5
	if chanceToCatch > rand.Intn(100){
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.Pokedex[pokemon] = *res
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}