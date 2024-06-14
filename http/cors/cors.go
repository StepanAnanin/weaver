// Cross-Origin Resource Sharing (CORS): https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
package cors

import (
	"net/http"
	"strings"
)

type header[T string | bool | int | float64 | float32 | []string] struct {
	name  string
	value T
}

func (h *header[T]) Name() string {
	return h.name
}

func (h *header[T]) Get() T {
	return h.value
}

func (h *header[T]) Set(newValue T) {
	h.value = newValue
}

func (h *header[T]) String() string {
	out := any(h.value)

	switch out.(type) {
	case []string:
		return strings.Join(out.([]string), ", ")
	case bool, int, string, float32, float64:
		return out.(string)
	default:
		panic("Invalid CORS header value")
	}
}

func (h *header[T]) Apply(writer http.ResponseWriter) {
	writer.Header().Set(h.name, h.String())
}

type headers struct {
	// Access-Control-Allow-Credentials header, default value:
	// true
	AllowCreditinals *header[bool]
	// Access-Control-Allow-AllowMethods header, default value:
	// ["GET"]
	AllowMethods *header[[]string]
	// Access-Control-Allow-Headers header, default value:
	// ["Accept", "Date", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"]
	AllowHeaders *header[[]string]
	// Access-Control-Allow-AllowOrigin header, default value:
	// "*"
	AllowOrigin *header[string]
}

func (headers *headers) Apply(writer http.ResponseWriter) {
	headers.AllowCreditinals.Apply(writer)
	headers.AllowMethods.Apply(writer)
	headers.AllowHeaders.Apply(writer)
	headers.AllowOrigin.Apply(writer)
}

var Headers *headers = &headers{
	AllowCreditinals: &header[bool]{
		name:  "Access-Control-Allow-Credentials",
		value: true,
	},
	AllowMethods: &header[[]string]{
		name:  "Access-Control-Allow-Methods",
		value: []string{"GET"},
	},
	AllowHeaders: &header[[]string]{
		name: "Access-Control-Allow-Headers",
		value: []string{
			"Accept",
			"Date",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	},
	AllowOrigin: &header[string]{
		name:  "Access-Control-Allow-Origin",
		value: "*",
	},
}
