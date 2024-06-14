package request

import (
	"log"
	"net/http"

	"github.com/StepanAnanin/weaver/config"
	"github.com/StepanAnanin/weaver/http/cors"
	"github.com/StepanAnanin/weaver/http/response"
)

type request struct {
	origin *http.Request
	writer http.ResponseWriter
}

func New(origin *http.Request, w http.ResponseWriter) *request {
	return &request{
		origin: origin,
		writer: w,
	}
}

func (req *request) SetContentType(contentType string) {
	req.writer.Header().Set("Content-Type", contentType)
}

// Returns given handler if everything OK, otherwise returns empty http.HandlerFunc,
// it's not nil, just does nothing (cuz if don't give it a value app will panic due to nil pointer)
//
// # For request with OPTIONS method always will be returned empty http.HandlerFunc
//
// # For request with not allowed method sends error (status 405 and body with error message and list of allowed methods)
func (req *request) Preprocessing(handler http.HandlerFunc, methods []string) http.HandlerFunc {
	// Don't use empty "invalidOut" or you'll get nil pointer error
	var invalidOut http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {}

	cors.Headers.AllowMethods.Set(methods)

	cors.Headers.Apply(req.writer)

	if config.Settings.LogIncomingRequests {
		log.Printf("[ %s ] %s %s", req.origin.RemoteAddr, req.origin.Method, req.origin.RequestURI)
	}

	if req.origin.Method == "OPTIONS" {
		return invalidOut
	}

	allowed := false

	for _, method := range methods {
		if req.origin.Method == method {
			allowed = true
			break
		}
	}

	if allowed {
		response.New(req.writer).Message("Method Not Allowed. Allowed methods: "+cors.Headers.AllowMethods.String(), http.StatusMethodNotAllowed)

		return invalidOut
	}

	return handler
}
