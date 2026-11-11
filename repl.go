package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	commands map[string]cliCommand
	Next *string
	Previous *string
}


func cleanInput(text string) []string{
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scaner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Printf("Pokedex > ")
		scaner.Scan()
		input := scaner.Text()
		cleanInput := cleanInput(input)
		command, ok := cfg.commands[cleanInput[0]]
		if !ok{
			fmt.Printf("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil{
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map":{
			name: "map",
			description: "Shows next 20 locations",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Shows previous 20 locations",
			callback: commandMapb,
		},
	}
}