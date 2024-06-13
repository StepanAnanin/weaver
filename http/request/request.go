package request

import (
	"log"
	"net/http"
	"strings"

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

// Return true if method supported, false otherwise
// (Except OPTIONS, for this method always will be returned false)
//
// If method isn't supported sends error response (status 405)
func (req *request) Preprocessing(methods []string) bool {
	strMethods := strings.Join(methods, ", ")

	req.writer.Header().Set("Content-Type", "application/json")
	req.writer.Header().Set("Access-Control-Allow-Credentials", "true")
	req.writer.Header().Set("Access-Control-Allow-Methods", strMethods)
	req.writer.Header().Set("Access-Control-Allow-Headers",
		"Accept, Date, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	req.writer.Header().Set("Access-Control-Allow-Origin", cors.GetOrigin())

	// TODO
	// if origin := req.Header.Get("Origin"); origin != "" {
	// 	req.writer.Header().Set("Access-Control-Allow-Origin", origin)
	// }

	log.Printf("[ %s ] %s %s", req.origin.RemoteAddr, req.origin.Method, req.origin.RequestURI)

	if req.origin.Method == "OPTIONS" {
		return false
	}

	allowed := false

	for _, method := range methods {
		if req.origin.Method == method {
			allowed = true
			break
		}
	}

	if allowed {
		response.New(req.writer).Message("Method Not Allowed. Allowed methods: "+strMethods, http.StatusMethodNotAllowed)

		return false
	}

	return true
}
