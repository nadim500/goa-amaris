package design

import . "goa.design/goa/v3/dsl"

var FriendlyChainsRequest = Type("FriendlyChainsRequest", func() {
	Description("Represents a friendlychains request.")

	Attribute("a", String, "First word.", func() {
		Example("to")
	})

	Attribute("b", String, "Second word.", func() {
		Example("kyo")
	})

	Required("a", "b")
})

var FriendlyChainsPayload = Type("FriendlyChainsPayload", func() {
	Field(1, "request", FriendlyChainsRequest, func() {
		Description("The friendlychains request")
	})

	Required("request")
})
