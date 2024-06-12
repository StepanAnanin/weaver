package logger

import (
	"log"
	"net/http"
)

// Works like "log.Print()", but also attaches some request data to output.
// (IP of device, that sent request; HTTP Method; Requested URI)
//
// Example with `message` = "Authentication successful, user id: ...:
//
// "2020/01/01 12:00:00 [ 127.0.0.1:5000 ] POST /login | Authentication successful, user id: ..."
func Print(message string, req *http.Request) {
	log.Printf("[ %s ] %s %s | %s", req.RemoteAddr, req.Method, req.RequestURI, message)
}

// Works like "log.Print()", but also attaches some request data to output.
// (IP of device, that sent request; HTTP Method; Requested URI)
//
// Example with `message` = "Access token expired":
//
// "2020/01/01 12:00:00 [ 127.0.0.1:50000 ] Error: GET /verification | Access token expired"
func PrintError(message string, status int, req *http.Request) {
	log.Printf("[ %s ] Error: %s %s | %s", req.RemoteAddr, req.Method, req.RequestURI, message)
}
