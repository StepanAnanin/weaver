// Cross-Origin Resource Sharing (CORS): https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
package cors

import (
	"net/http"

	"github.com/StepanAnanin/weaver"
	"github.com/StepanAnanin/weaver/http/header"
)

type headers struct {
	// Access-Control-Allow-Credentials header, default value:
	// true
	AllowCreditinals *header.HttpHeader[bool]
	// Access-Control-Allow-Methods header, default value:
	// ["GET"]
	AllowMethods *header.HttpHeader[[]string]
	// Access-Control-Allow-Headers header, default value:
	// ["Accept", "Date", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"]
	AllowHeaders *header.HttpHeader[[]string]
	// Access-Control-Allow-Origin header, default value:
	// "*"
	AllowOrigin *header.HttpHeader[string]
}

func New() *headers {
	return &headers{
		AllowCreditinals: header.New("Access-Control-Allow-Credentials", true),
		AllowOrigin:      header.New("Access-Control-Allow-Origin", weaver.Settings.AccessControlAllowOrigin),
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
}

func (h *headers) Apply(writer http.ResponseWriter) {
	h.AllowCreditinals.Apply(writer)
	h.AllowMethods.Apply(writer)
	h.AllowHeaders.Apply(writer)
	h.AllowOrigin.Apply(writer)
}
