package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("app", func() {
	Title("Pokemon Quiz Service")
	Description("Service for viewing Pokemon Quiz.")
    Server("app", func() {
        Host("0.0.0.0", func() {
            URI("http://0.0.0.0:8000")
            URI("grpc://0.0.0.0:8080")
        })
    })
})

var _ = Service("app", func() {
	Description("The app service provides some pokemon api.")

	Method("select", func() {
		Payload(func() {
			Field(1, "type", String, "Selected Pokemon Types")
			Required("type")
		})

		Result(String)

		HTTP(func() {
			GET("/select/{type}")
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})