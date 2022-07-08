package design

import . "goa.design/goa/v3/dsl"

var FriendlyChainsResponse = ResultType("FriendlyChainsResponse", func() {

	Description("Represent the response from FriendlyChains")

	Attributes(func() {
		Attribute("result", String, "Friendly Chains", func() {
			Example("Las cadenas son amigas")
		})
	})

})
