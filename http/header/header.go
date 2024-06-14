package header

import (
	"net/http"
	"strings"
)

type httpHeaderValue interface {
	string | bool | int | float64 | float32 | []string
}

type HttpHeader[V httpHeaderValue] struct {
	name  string
	value V
}

func New[V httpHeaderValue](name string, value V) *HttpHeader[V] {
	return &HttpHeader[V]{name, value}
}

func (h *HttpHeader[V]) Name() string {
	return h.name
}

func (h *HttpHeader[V]) Get() V {
	return h.value
}

func (h *HttpHeader[V]) Set(newValue V) {
	h.value = newValue
}

func (h *HttpHeader[V]) String() string {
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

func (h *HttpHeader[V]) Apply(writer http.ResponseWriter) {
	writer.Header().Set(h.name, h.String())
}
