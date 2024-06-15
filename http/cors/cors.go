// Cross-Origin Resource Sharing (CORS): https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
package cors

import (
	"net/http"

	"github.com/StepanAnanin/weaver/http/header"
)

type corsHeaders struct {
	// Access-Control-Allow-Credentials header, default value:
	// true
	AllowCreditinals *header.HttpHeader[bool]
	// Access-Control-Allow-Methods header, default value:
	// ["GET"]
	AllowMethods *header.HttpHeader[[]string]
	// Access-Control-Allow-Headers header, default value:
	// ["Accept", "Date", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"]
	AllowHeaders *header.HttpHeader[[]string]
	// Access-Control-Allow-AllowOrigin header, default value:
	// "*"
	AllowOrigin *header.HttpHeader[string]
}

func (headers *corsHeaders) Apply(writer http.ResponseWriter) {
	headers.AllowCreditinals.Apply(writer)
	headers.AllowMethods.Apply(writer)
	headers.AllowHeaders.Apply(writer)
	headers.AllowOrigin.Apply(writer)
}

var Headers *corsHeaders = &corsHeaders{
	AllowCreditinals: header.New("Access-Control-Allow-Credentials", true),
	AllowOrigin:      header.New("Access-Control-Allow-Origin", "*"),
	AllowMethods:     header.New("Access-Control-Allow-Methods", []string{"GET"}),
	AllowHeaders: header.New("Access-Control-Allow-Headers", []string{
		"Accept",
		"Date",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
	}),
}
