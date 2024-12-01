package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/1729prashant/misc-learning/go_learn/bootdev-pokedex/internal/pokecache"
)

var cache = pokecache.NewCache(5 * time.Minute)

/*
------------------------------------------------------------------------------------
							START fetch location area list
------------------------------------------------------------------------------------
*/

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
	// Check cache first
	if cachedData, found := cache.Get(url); found {
		var response LocationAreaResponse
		err := json.Unmarshal(cachedData, &response)
		if err == nil {
			return &response, nil
		}
		fmt.Println("Warning: failed to unmarshal cached data, refetching...")
	}

	// If not in cache, fetch from API
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

	// Add response to cache
	cache.Add(url, body)

	return &response, nil
}

// Display location area names
func displayLocationAreas(response *LocationAreaResponse) {
	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
}

/*
------------------------------------------------------------------------------------
							  END fetch location area list
------------------------------------------------------------------------------------
*/

/*
------------------------------------------------------------------------------------
							START fetch location area list
------------------------------------------------------------------------------------
*/

type LocationAreaPokemonsResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// Fetch location areas from the given URL
func fetchLocationAreaPokemons(url string) (*LocationAreaPokemonsResponse, error) {
	// Check cache first
	if cachedData, found := cache.Get(url); found {
		var response LocationAreaPokemonsResponse
		err := json.Unmarshal(cachedData, &response)
		if err == nil {
			return &response, nil
		}
		fmt.Println("Warning: failed to unmarshal cached data, refetching...")
	}

	// If not in cache, fetch from API
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

	var response LocationAreaPokemonsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Add response to cache
	cache.Add(url, body)

	return &response, nil
}

// Display location area names
func displayLocationAreaPokemons(response *LocationAreaPokemonsResponse) {

	if len(response.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return
	}

	fmt.Println("Found Pokémon:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}

/*
------------------------------------------------------------------------------------
							  END pokemon list in location area
------------------------------------------------------------------------------------
*/

/*
------------------------------------------------------------------------------------
						START catch pokemon
------------------------------------------------------------------------------------
*/

type CaughtPokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func catchPokemon(url string) (*CaughtPokemon, error) {
	//url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	// Check cache first
	if cachedData, found := cache.Get(url); found {
		var pokemon CaughtPokemon
		err := json.Unmarshal(cachedData, &pokemon)
		if err == nil {
			return &pokemon, nil
		}
		fmt.Println("Warning: failed to unmarshal cached data, refetching...")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pokemon: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var pokemon CaughtPokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Add response to cache
	cache.Add(url, body)

	return &pokemon, nil
}

func calculateCatchProbability(baseExperience int) float64 {
	// Base catch rate decreases as experience increases
	// Maximum base experience is around 600
	baseRate := 1.0 - (float64(baseExperience) / 600.0)
	// Ensure minimum catch rate of 10%
	return 0.1 + (baseRate * 0.6)
}

/*
------------------------------------------------------------------------------------
							  END catch pokemon
------------------------------------------------------------------------------------
*/

/*
------------------------------------------------------------------------------------
							START pokedex
------------------------------------------------------------------------------------
*/

type Pokedex struct {
	CaughtPokemon map[string]CaughtPokemon
}

var pokedex Pokedex

const pokedexFile = "./data/pokedex.json"

// Add these functions
func initPokedex() {
	pokedex.CaughtPokemon = make(map[string]CaughtPokemon)
	loadPokedex()
}

func loadPokedex() {
	data, err := os.ReadFile(pokedexFile)
	if err != nil {
		if os.IsNotExist(err) {
			return // File doesn't exist yet, that's okay
		}
		fmt.Printf("Error loading Pokedex: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &pokedex.CaughtPokemon)
	if err != nil {
		fmt.Printf("Error parsing Pokedex data: %v\n", err)
	}
}

func savePokedex() {
	data, err := json.Marshal(pokedex.CaughtPokemon)
	if err != nil {
		fmt.Printf("Error serializing Pokedex: %v\n", err)
		return
	}

	err = os.MkdirAll(filepath.Dir(pokedexFile), 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.WriteFile(pokedexFile, data, 0644)
	if err != nil {
		fmt.Printf("Error saving Pokedex: %v\n", err)
	}
}

/*
------------------------------------------------------------------------------------
							  END pokedex
------------------------------------------------------------------------------------
*/

/*
------------------------------------------------------------------------------------
							START inspect pokemon
------------------------------------------------------------------------------------
*/

/*
------------------------------------------------------------------------------------
							  END inspect pokemon
------------------------------------------------------------------------------------
*/

func main() {
	//rand.Seed(time.Now().UnixNano())

	locationAreaURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	locationAreaPokemonsURL := "https://pokeapi.co/api/v2/location-area/"
	catchPokemonURL := "https://pokeapi.co/api/v2/pokemon/"

	initPokedex()

	config := &Config{
		Next:     &locationAreaURL,
		Previous: nil,
	}

	commands := map[string]func(*Config, ...string){
		"help": func(cfg *Config, args ...string) {
			fmt.Print(`
Welcome to the Pokedex!
Usage:
  help:               Displays a help message
  exit:               Exit the Pokedex
  map:                Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map displays the next 20 locations.
  mapb:               Displays the previous 20 locations. If you're on the first page, it will print an error message.
  explore <area>:     Fetch list of all the Pokémon in a given location area.
  catch <pokemon>:    Attempt to catch pokemon !
  inspect <pokemon>:  List details of caught pokemon.
  pokedex:            List all caught pokemon.
`)
		},

		"exit": func(cfg *Config, args ...string) {
			fmt.Println("Exiting the Pokedex. Goodbye!")
			os.Exit(0)
		},

		"map": func(cfg *Config, args ...string) {
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

		"mapb": func(cfg *Config, args ...string) {
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

		"explore": func(cfg *Config, args ...string) {
			if len(args) < 1 {
				fmt.Println("Usage: explore <area>, where <area> is the location area.")
				return
			}

			pokemonListURL := locationAreaPokemonsURL + args[0]
			response, err := fetchLocationAreaPokemons(pokemonListURL)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			displayLocationAreaPokemons(response)
		},

		"catch": func(cfg *Config, args ...string) {
			if len(args) < 1 {
				fmt.Println("Usage: catch <pokemon>.")
				return
			}
			pokemonName := strings.ToLower(args[0])
			fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
			catchingPokemon := catchPokemonURL + pokemonName
			pokemon, err := catchPokemon(catchingPokemon)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			catchProbability := calculateCatchProbability(pokemon.BaseExperience)
			if rand.Float64() > catchProbability {
				fmt.Printf("%s escaped!\n", pokemonName)
				return
			}

			pokedex.CaughtPokemon[pokemonName] = *pokemon
			savePokedex()
			fmt.Printf("%s was caught!\n", pokemonName)
			fmt.Println("You may now inspect it with the inspect command.")
		},

		"inspect": func(cfg *Config, args ...string) {
			if len(args) < 1 {
				fmt.Println("Usage: inspect <pokemon>")
				return
			}

			pokemonName := strings.ToLower(args[0])
			pokemon, exists := pokedex.CaughtPokemon[pokemonName]
			if !exists {
				fmt.Println("You have not caught that pokemon")
				return
			}

			fmt.Printf("Name: %s\n", pokemon.Name)
			fmt.Printf("Height: %d\n", pokemon.Height)
			fmt.Printf("Weight: %d\n", pokemon.Weight)

			fmt.Println("Stats:")
			for _, stat := range pokemon.Stats {
				fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
			}

			fmt.Println("Types:")
			for _, typeInfo := range pokemon.Types {
				fmt.Printf(" - %s\n", typeInfo.Type.Name)
			}
		},

		"pokedex": func(cfg *Config, args ...string) {
			if len(pokedex.CaughtPokemon) == 0 {
				fmt.Println("Your Pokedex is empty")
				return
			}

			fmt.Println("Your Pokedex:")
			for name := range pokedex.CaughtPokemon {
				fmt.Printf(" - %s\n", name)
			}
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if input == "" {
			continue
		}

		parts := strings.Split(input, " ")
		cmd, args := parts[0], parts[1:]

		if command, found := commands[cmd]; found {
			command(config, args...)
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
