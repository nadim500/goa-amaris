// Code generated by goa v3.7.10, DO NOT EDIT.
//
// amaris client
//
// Command:
// $ goa gen testamaris/design

package amaris

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "amaris" service client.
type Client struct {
	SortNamesEndpoint      goa.Endpoint
	PokemonNameEndpoint    goa.Endpoint
	FriendlyChainsEndpoint goa.Endpoint
}

// NewClient initializes a "amaris" service client given the endpoints.
func NewClient(sortNames, pokemonName, friendlyChains goa.Endpoint) *Client {
	return &Client{
		SortNamesEndpoint:      sortNames,
		PokemonNameEndpoint:    pokemonName,
		FriendlyChainsEndpoint: friendlyChains,
	}
}

// SortNames calls the "SortNames" endpoint of the "amaris" service.
func (c *Client) SortNames(ctx context.Context, p *SortNamesPayload) (res *Sortnamesresponse, err error) {
	var ires interface{}
	ires, err = c.SortNamesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Sortnamesresponse), nil
}

// PokemonName calls the "PokemonName" endpoint of the "amaris" service.
func (c *Client) PokemonName(ctx context.Context, p *PokemonNamePayload) (res *Pokemonnameresponse, err error) {
	var ires interface{}
	ires, err = c.PokemonNameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Pokemonnameresponse), nil
}

// FriendlyChains calls the "FriendlyChains" endpoint of the "amaris" service.
func (c *Client) FriendlyChains(ctx context.Context, p *FriendlyChainsPayload) (res *Friendlychainsresponse, err error) {
	var ires interface{}
	ires, err = c.FriendlyChainsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Friendlychainsresponse), nil
}
