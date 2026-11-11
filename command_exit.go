package main
import (
	"fmt"
	"os"
)
func commandExit(_ *config) error {
	fmt.Printf(`Closing the Pokedex... Goodbye!
`)
	os.Exit(0)
	return nil
}