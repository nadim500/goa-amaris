// Code generated by goa v3.7.10, DO NOT EDIT.
//
// HTTP request path constructors for the amaris service.
//
// Command:
// $ goa gen testamaris/design

package server

import (
	"fmt"
)

// SortNamesAmarisPath returns the URL path to the amaris service SortNames HTTP endpoint.
func SortNamesAmarisPath() string {
	return "/sortnames"
}

// PokemonNameAmarisPath returns the URL path to the amaris service PokemonName HTTP endpoint.
func PokemonNameAmarisPath(id int) string {
	return fmt.Sprintf("/pokemonname/%v", id)
}

// FriendlyChainsAmarisPath returns the URL path to the amaris service FriendlyChains HTTP endpoint.
func FriendlyChainsAmarisPath() string {
	return "/friendlychains"
}