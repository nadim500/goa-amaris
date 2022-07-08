package design

import . "goa.design/goa/v3/dsl"

var SortNamesResponse = ResultType("SortNamesResponse", func() {

	Description("Represent the response from SortNames")

	Attributes(func() {
		Attribute("quantity", Int, "Amount in Array", func() {
			Example(5)
		})
		Attribute("data", ArrayOf(String), "Names sort by first letter", func() {
			Example([]string{"Andres", "Camilo", "Laura", "Luis"})
		})
	})

})
