package weaver

import (
	"log"
	"net/http"
	"slices"
)

// Accepts endpoint `handler` and allowed `methods`.
// Returns handle function wich does request preprocessing, if preprocessing successful then calls given `handler`.
//
// # Also applies cors headers to the response.
//
// # OPTIONS requests won't pass preprocessing (can be disabled by setting `weaver.settings.PassOptionsRequestsOnPreprocessing` to true)
//
// # If methods array is empty, then will be used default value for Access-Control-Allow-Methods header (default value is "GET").
//
// # If method isn't suppored sends response with error message and list of allowed methods for this endpoint (status 405).
//
// # If request logs enabled, then also print some request info in terminal (requester ip, method and requested URI)
func Preprocessing(handler http.HandlerFunc, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if Settings.LogIncomingRequests {
			log.Printf("[ %s ] %s %s", req.RemoteAddr, req.Method, req.RequestURI)
		}

		if Settings.DisableCORS {
			if req.Method == "OPTIONS" && !Settings.PassOptionsRequestsOnPreprocessing {
				return
			}

			handler(w, req)

			return
		}

		corsHeaders := InitCORS()

		if len(methods) > 0 {
			corsHeaders.AllowMethods.Set(methods)
		}

		corsHeaders.Apply(w)

		if req.Method == "OPTIONS" && !Settings.PassOptionsRequestsOnPreprocessing {
			return
		}

		if !slices.Contains(methods, req.Method) {
			NewResponse(w).Message("Method Not Allowed. Allowed methods: "+corsHeaders.AllowMethods.String(), http.StatusMethodNotAllowed)

			return
		}

		handler(w, req)
	}
}
