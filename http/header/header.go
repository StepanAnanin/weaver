package header

import (
	"net/http"
	"strconv"
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

	switch v := out.(type) {
	case string:
		return v
	case []string:
		return strings.Join(v, ", ")
	case bool:
		return strconv.FormatBool(v)
	case int, int64:
		return strconv.FormatInt(out.(int64), 10)
	case float32, float64:
		return strconv.FormatFloat(out.(float64), 'f', -1, 64)
	default:
		panic("Invalid header value")
	}
}

func (h *HttpHeader[V]) Apply(writer http.ResponseWriter) {
	writer.Header().Set(h.name, h.String())
}
