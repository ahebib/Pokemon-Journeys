package pokeapi

import (
	pokemon "dev/pokemon-journeys/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetAllPokemon() pokemon.Pokedex {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon?limit=10000&offset=0", nil)
	if err != nil {
		fmt.Println("req error", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error", err)
	}

	var data pokemon.Pokedex
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

// Get specific Pokemon with PokedexEntry
func GetPokemon(url string) pokemon.PokedexEntry {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("req error", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("resp error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error", err)
	}

	var data pokemon.PokedexEntry
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
