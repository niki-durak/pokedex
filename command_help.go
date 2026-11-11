package main

import(
	"fmt"
)

func commandHelp(_ *config) error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:
`)
commands := getCommands()
	for _, command := range commands{
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}