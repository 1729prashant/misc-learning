package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Define commands and their actions
	commands := map[string]func(){
		"help": func() {
			fmt.Print(`
Welcome to the Pokedex!
Usage:
  help: Displays a help message
  exit: Exit the Pokedex

`)
		},

		// Exit the program
		"exit": func() {
			fmt.Println("Exiting the Pokedex. Goodbye!")
			os.Exit(0) // Exit the program
		},
		
		//The map command displays the names of 20 location areas in the Pokemon world. 
		//Each subsequent call to map should display the next 20 locations, and so on. 
		//The idea is that the map command will let us explore the world of Pokemon.
		"map": func() {
			//GET https://pokeapi.co/api/v2/location-area/{id or name}/

		},

		//Similar to the map command, however, instead of displaying the next 20 
		//locations, it displays the previous 20 locations. It's a way to go back.
		//If you're on the first "page" of results, this command should just print an 
		//error message.
		"mapb": func() {

		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')
		command := strings.TrimSpace(text)

		// Execute the command if it exists in the map
		if action, exists := commands[command]; exists {
			action()
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
