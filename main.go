package main



func main() {
	cfg := config {commands: getCommands(), Pokedex: make(map[string]PokemonResponse),}
	startRepl(&cfg)
}