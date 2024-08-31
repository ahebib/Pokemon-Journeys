package pokemon

type Pokedex struct {
	Count   int
	Results []PokedexEntry
}

type PokedexEntry struct {
	Url  string      `json:"url"`
	Name interface{} `json:"name"`
	Weight int `json: weight`
	Height int `json: height`
	Abilities []Abilities `json:"abilities"`
	Sprites Sprite `json:"sprites"`
}

type Sprite struct {
	FrontDefault string `json:"front_default"`
}

type Abilities struct {
	Ability AbilityValues
	Hidden bool `json:"is_hidden"`
}

type AbilityValues struct {
	Name string
	Url string
}