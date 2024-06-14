package request

import (
	"log"
	"net/http"
	"slices"

	"github.com/StepanAnanin/weaver/config"
	"github.com/StepanAnanin/weaver/http/cors"
	"github.com/StepanAnanin/weaver/http/response"
)

// it's not nil, just does nothing (cuz if don't give it a value app will panic due to nil pointer)
// Returns given handler if everything OK, otherwise returns empty http.HandlerFunc,
//
// # Also applies CORS headers
//
// # For request with OPTIONS method always will be returned empty http.HandlerFunc
//
// # For request with not allowed method sends error (status 405 and body with error message and list of allowed methods)
func Preprocessing(handler http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		cors.Headers.AllowMethods.Set(methods)

		cors.Headers.Apply(w)

		if config.Settings.LogIncomingRequests {
			log.Printf("[ %s ] %s %s", req.RemoteAddr, req.Method, req.RequestURI)
		}

		if req.Method == "OPTIONS" {
			return
		}

		if !slices.Contains(methods, req.Method) {
			response.New(w).Message("Method Not Allowed. Allowed methods: "+cors.Headers.AllowMethods.String(), http.StatusMethodNotAllowed)

			return
		}

		handler(w, req)
	}
}
