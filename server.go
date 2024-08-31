package main

import (
	pokeapi "dev/pokemon-journeys/services"
	"html/template"
	"net/http"
)

func main() {
	// Views
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		results := pokeapi.GetAllPokemon()

		indexTmpl := template.Must(template.ParseFiles("./views/go/index.html"))
		indexTmpl.Execute(w, results)
	})

	http.HandleFunc("GET /pokemon", func(w http.ResponseWriter, r *http.Request) {
		var apiURL string = r.URL.Query().Get("url")
		selectedPokemon := pokeapi.GetPokemon(apiURL)
		pokemonTmpl := template.Must(template.ParseFiles("./views/go/pokemon.html"))
		pokemonTmpl.Execute(w, selectedPokemon)
	})

	// Images, CSS, and JS will be stored in the assets folder
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3080", nil)
}
