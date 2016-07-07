package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("version", func() {

	DefaultMedia(ALMVersion)
	BasePath("/version")

	Action("show", func() {
		Security("jwt", func() {
			Scope("system")
		})
		Routing(
			GET(""),
		)
		Description("Show current running version")
		Response(OK)
	})
})

var _ = Resource("login", func() {

	BasePath("/login")

	Action("authorize", func() {
		Routing(
			GET("authorize"),
		)
		Params(func() {
			Param("code", String)
			Param("state", String)
		})
		Description("Authorize with the ALM")
		Response(OK, func() {
			Media(AuthToken)
		})
		Response(Unauthorized)
		Response(TemporaryRedirect)
	})

	Action("generate", func() {
		Routing(
			GET("generate"),
		)
		Description("Generates a set of Tokens for different Auth levels. NOT FOR PRODUCTION. Only available if server is running in dev mode")
		Response(OK, func() {
			Media(CollectionOf(AuthToken))
		})
		Response(Unauthorized)
	})
})
