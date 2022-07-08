package design

import . "goa.design/goa/v3/dsl"

var PokemonNameResponse = ResultType("PokemonNameResponse", func() {

	Description("Represent the response from PokenName")

	Attributes(func() {
		Attribute("name", String, "Pokemon Name", func() {
			Example("Picachu")
		})
	})

})
