package weaver

import (
	"fmt"
	"log"
	"net/http"
)

// Works like "log.Print()", but also attaches some request info to output.
// (IP of device, that sent request; HTTP Method; Requested URI)
//
// Returns log result. If passed status is empty, then it will be set to "INFO".
//
// Example with `status` = "ERROR:" and `message` = "Authentication successful, user id: ...:
//
// "2020/01/01 12:00:00 [ INFO ] 127.0.0.1:50000 GET /verification | Access token expired"

func LogRequestWithStatus(message string, req *http.Request, status string) string {
	if status == "" {
		status = "INFO"
	}

	l := fmt.Sprintf("[ "+status+" ] %s %s %s | %s", req.RemoteAddr, req.Method, req.RequestURI, message)

	log.Print(l)

	return l
}

// Same as LogRequestWithStatus(message, req, "REQUEST")
func LogIncomingRequest(message string, req *http.Request) string {
	return LogRequestWithStatus(message, req, "REQUEST")
}

// Same as LogRequestWithStatus(message, req, "INFO")
func LogRequest(message string, req *http.Request) string {
	return LogRequestWithStatus(message, req, "INFO")
}

// Same as LogRequestWithStatus(message, req, "ERROR:")
func LogRequestError(message string, req *http.Request) string {
	return LogRequestWithStatus(message, req, "ERROR:")
}

// Same as LogRequestWithStatus(message, req, "WARNING:")
func LogRequestWarning(message string, req *http.Request) string {
	return LogRequestWithStatus(message, req, "WARNING:")
}
