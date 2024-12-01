package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	/*EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`*/
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		/*VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`*/
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

type CatchPokemon struct {
	BaseExperience int `json:"base_experience"`
}

/*TODO
add logic to catch pokemons based on base_experience value and random value, higher the base experiece larger the range of random values. pokemon is caught
add logic to store caught pokemons, storage should persist even after
*/

/*
------------------------------------------------------------------------------------
							  END catch pokemon
------------------------------------------------------------------------------------
*/

func main() {
	locationAreaURL := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	locationAreaPokemonsURL := "https://pokeapi.co/api/v2/location-area/"
	//catchPokemonURL := "https://pokeapi.co/api/v2/pokemon/"

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
  catch <pokemon>:    
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

			/* TODO add logic to catch pokemons absed on experience   */
			/* coughtPokemonURL := catchPokemonURL + args[0] */
			/*response, err := caughtPokemon(coughtPokemonURL)*/
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
