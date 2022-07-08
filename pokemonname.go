package amarisapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	amaris "testamaris/gen/amaris"
	"time"
)

type PokemonData struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type PokemonSprites struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShiniFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

type PokemonTypeType struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type PokemonType struct {
	Slot *int             `json:"slot"`
	Type *PokemonTypeType `json:"type"`
}

type PokemonVersionGroup struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type Pokemon struct {
	FormName     *string              `json:"form_name"`
	FormNames    *[]string            `json:"form_names"`
	FormOrder    *int                 `json:"form_order"`
	ID           int                  `json:"id"`
	IsBattleOnly *bool                `json:"is_battle_only"`
	IsDefault    *bool                `json:"is_default"`
	IsMega       *bool                `json:"is_mega"`
	Name         *string              `json:"name"`
	Names        *[]string            `json:"names"`
	Order        *int                 `json:"order"`
	Pokemon      *PokemonData         `json:"pokemon"`
	Sprites      *PokemonSprites      `pokemon:"sprites"`
	Types        *[]PokemonType       `pokemon:"types"`
	VersionGroup *PokemonVersionGroup `pokemon:"version_group"`
}

// Receives pokemon id and returns its name
func (s *amarissrvc) PokemonName(ctx context.Context, p *amaris.PokemonNamePayload) (res *amaris.Pokemonnameresponse, err error) {
	res = &amaris.Pokemonnameresponse{}

	res.Name, _ = retrieveNamePokemon(p.ID)

	return
}

func retrieveNamePokemon(id int) (name *string, err error) {
	client := http.Client{
		Timeout: time.Duration(6 * time.Minute),
	}

	oldEndpoint := "https://pokeapi.co/api/v2/pokemon-form/{{id}}"
	endpoint := strings.Replace(oldEndpoint, "{{id}}", fmt.Sprint(id), -1)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var pokemon Pokemon
	err = json.Unmarshal(res, &pokemon)
	if err != nil {
		return
	}

	name = pokemon.Name

	return
}
