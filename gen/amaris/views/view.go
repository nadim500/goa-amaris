// Code generated by goa v3.7.10, DO NOT EDIT.
//
// amaris views
//
// Command:
// $ goa gen testamaris/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Sortnamesresponse is the viewed result type that is projected based on a
// view.
type Sortnamesresponse struct {
	// Type to project
	Projected *SortnamesresponseView
	// View to render
	View string
}

// Pokemonnameresponse is the viewed result type that is projected based on a
// view.
type Pokemonnameresponse struct {
	// Type to project
	Projected *PokemonnameresponseView
	// View to render
	View string
}

// Friendlychainsresponse is the viewed result type that is projected based on
// a view.
type Friendlychainsresponse struct {
	// Type to project
	Projected *FriendlychainsresponseView
	// View to render
	View string
}

// SortnamesresponseView is a type that runs validations on a projected type.
type SortnamesresponseView struct {
	// Amount in Array
	Quantity *int
	// Names sort by first letter
	Data []string
}

// PokemonnameresponseView is a type that runs validations on a projected type.
type PokemonnameresponseView struct {
	// Pokemon Name
	Name *string
}

// FriendlychainsresponseView is a type that runs validations on a projected
// type.
type FriendlychainsresponseView struct {
	// Friendly Chains
	Result *string
}

var (
	// SortnamesresponseMap is a map indexing the attribute names of
	// Sortnamesresponse by view name.
	SortnamesresponseMap = map[string][]string{
		"default": {
			"quantity",
			"data",
		},
	}
	// PokemonnameresponseMap is a map indexing the attribute names of
	// Pokemonnameresponse by view name.
	PokemonnameresponseMap = map[string][]string{
		"default": {
			"name",
		},
	}
	// FriendlychainsresponseMap is a map indexing the attribute names of
	// Friendlychainsresponse by view name.
	FriendlychainsresponseMap = map[string][]string{
		"default": {
			"result",
		},
	}
)

// ValidateSortnamesresponse runs the validations defined on the viewed result
// type Sortnamesresponse.
func ValidateSortnamesresponse(result *Sortnamesresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateSortnamesresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidatePokemonnameresponse runs the validations defined on the viewed
// result type Pokemonnameresponse.
func ValidatePokemonnameresponse(result *Pokemonnameresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePokemonnameresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateFriendlychainsresponse runs the validations defined on the viewed
// result type Friendlychainsresponse.
func ValidateFriendlychainsresponse(result *Friendlychainsresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateFriendlychainsresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateSortnamesresponseView runs the validations defined on
// SortnamesresponseView using the "default" view.
func ValidateSortnamesresponseView(result *SortnamesresponseView) (err error) {

	return
}

// ValidatePokemonnameresponseView runs the validations defined on
// PokemonnameresponseView using the "default" view.
func ValidatePokemonnameresponseView(result *PokemonnameresponseView) (err error) {

	return
}

// ValidateFriendlychainsresponseView runs the validations defined on
// FriendlychainsresponseView using the "default" view.
func ValidateFriendlychainsresponseView(result *FriendlychainsresponseView) (err error) {

	return
}
