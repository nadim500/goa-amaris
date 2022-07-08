package design

import . "goa.design/goa/v3/dsl"

var SortNamesRequest = Type("SortNamesRequest", func() {
	Description("Represents a sortnames request.")

	Attribute("text", String, "The text string.", func() {
		Example("Luis,Camilo,Andres,Laura")
	})

	Required("text")
})

var SortNamesPayload = Type("SortNamesPayload", func() {
	Field(1, "request", SortNamesRequest, func() {
		Description("The sortnames request")
	})

	Required("request")
})
