package design

import . "goa.design/goa/v3/dsl"

var _ = API("amaris", func() {
	Title("Amaris Service")
	Description("HTTP service for amaris test")
	Server("amaris", func() {
		Host("localhost", func() { URI("http://localhost:8080") })
	})
})

var _ = Service("amaris", func() {
	Description("The amaris service provides endpoints to responde the amaris' test.")

	Method("SortNames", func() {
		Description("Receives text string and returns two parameters: an array string and the mount of array")

		Payload(SortNamesPayload)

		Result(SortNamesResponse)

		HTTP(func() {

			POST("/sortnames")

			Body("request")

			Response(StatusOK)
		})
	})

	Method("PokemonName", func() {
		Description("Receives pokemon id and returns its name")

		Payload(func() {
			Field(1, "id", Int, "Pokemon id")

			Required("id")
		})

		Result(PokemonNameResponse)

		HTTP(func() {

			GET("/pokemonname/{id}")

			Response(StatusOK)
		})
	})

	Method("FriendlyChains", func() {
		Description("Receives two words and returns if these are friendly")

		Payload(FriendlyChainsPayload)

		Result(FriendlyChainsResponse)

		HTTP(func() {

			POST("/friendlychains")

			Body("request")

			Response(StatusOK)
		})
	})
})

var _ = Service("openapi", func() {
	// Serve the file gen/http/openapi3.json for requests sent to
	// /openapi.json. The HTTP file system is created below.
	Files("/openapi.json", "openapi3.json")
})
