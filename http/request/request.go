package request

import (
	"log"
	"net/http"
	"slices"

	"github.com/StepanAnanin/weaver"
	"github.com/StepanAnanin/weaver/http/cors"
	"github.com/StepanAnanin/weaver/http/response"
)

// Accepts endpoint `handler` and array of allowed `methods`.
// Returns handle function wich does request preprocessing, if preprocessing successful then calls given `handler`.
//
// # If methods array is empty, then will be used default value for Access-Control-Allow-Methods header (default value is "GET").
//
// # Also applies cors headers to the response.
//
// # If method isn't suppored sends response with error message and list of allowed methods for this endpoint (status 405).
//
// # If request logs enabled, then also print some request info in terminal (requester ip, method and requested URI)
func Preprocessing(handler http.HandlerFunc, methods []string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		corsHeaders := cors.New()

		if len(methods) > 0 {
			corsHeaders.AllowMethods.Set(methods)
		}

		corsHeaders.Apply(w)

		if weaver.Settings.LogIncomingRequests {
			log.Printf("[ %s ] %s %s", req.RemoteAddr, req.Method, req.RequestURI)
		}

		if req.Method == "OPTIONS" {
			return
		}

		if !slices.Contains(methods, req.Method) {
			response.New(w).Message("Method Not Allowed. Allowed methods: "+corsHeaders.AllowMethods.String(), http.StatusMethodNotAllowed)

			return
		}

		handler(w, req)
	}
}
