package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"` // Nullable string
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Fetch location areas from the given URL
func fetchLocationAreas(url string) (*LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch location areas: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response LocationAreaResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

// Display location area names
func displayLocationAreas(response *LocationAreaResponse) {
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
}

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	config := &Config{
		Next:     &baseURL,
		Previous: nil,
	}

	commands := map[string]func(*Config){
		"help": func(cfg *Config) {
			fmt.Print(`
Welcome to the Pokedex!
Usage:
  help : Displays a help message
  exit : Exit the Pokedex
  map  : Displays the names of 20 location areas in the Pokemon world.
		 Each subsequent call to map displays the next 20 locations.
  mapb : Displays the previous 20 locations.
		 If you're on the first page, it will print an error message.
`)
		},

		"exit": func(cfg *Config) {
			fmt.Println("Exiting the Pokedex. Goodbye!")
			os.Exit(0)
		},

		"map": func(cfg *Config) {
			if cfg.Next == nil {
				fmt.Println("No more locations to display.")
				return
			}

			response, err := fetchLocationAreas(*cfg.Next)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			displayLocationAreas(response)
			cfg.Next = response.Next
			cfg.Previous = response.Previous
		},

		"mapb": func(cfg *Config) {
			if cfg.Previous == nil {
				fmt.Println("You're already at the first page of locations.")
				return
			}

			response, err := fetchLocationAreas(*cfg.Previous)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			displayLocationAreas(response)
			cfg.Next = response.Next
			cfg.Previous = response.Previous
		},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		text, _ := reader.ReadString('\n')
		command := strings.TrimSpace(text)

		if action, exists := commands[command]; exists {
			action(config)
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.\n")
		}
	}
}
