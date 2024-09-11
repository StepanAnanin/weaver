package weaver

import (
	"net/http"
)

type corsHeaders struct {
	// Access-Control-Allow-Credentials header, default value:
	// true
	AllowCreditinals *HttpHeader[bool]
	// Access-Control-Allow-Methods header, default value:
	// ["GET"]
	AllowMethods *HttpHeader[[]string]
	// Access-Control-Allow-Headers header, default value:
	// ["Accept", "Date", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"]
	AllowHeaders *HttpHeader[[]string]
	// Access-Control-Allow-Origin header, default value:
	// "*" (can be changed: Settings.DefaultOrigin)
	AllowOrigin *HttpHeader[string]
}

// Generates CORS headers which can be changed.
// Don't forget to apply them to the response (.Apply method of `corsHeaders`).
//
// CORS headers are: Access-Control-Allow-Credentials, Access-Control-Allow-Methods, Access-Control-Allow-Headers, Access-Control-Allow-Origin.
//
// You can see their default values in comments inside of `corsHeaders` struct.
func InitCORS() *corsHeaders {
	return &corsHeaders{
		AllowCreditinals: NewHttpHeader("Access-Control-Allow-Credentials", true),
		AllowOrigin:      NewHttpHeader("Access-Control-Allow-Origin", Settings.DefaultOrigin),
		AllowMethods:     NewHttpHeader("Access-Control-Allow-Methods", []string{"GET"}),
		AllowHeaders: NewHttpHeader("Access-Control-Allow-Headers", []string{
			"Accept",
			"Date",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Methods",
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
		}),
	}
}

func (h *corsHeaders) Apply(writer http.ResponseWriter) {
	h.AllowCreditinals.Apply(writer)
	h.AllowMethods.Apply(writer)
	h.AllowHeaders.Apply(writer)
	h.AllowOrigin.Apply(writer)
}
