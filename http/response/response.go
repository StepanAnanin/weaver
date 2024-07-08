package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/StepanAnanin/weaver/logger"
)

type response struct {
	writer http.ResponseWriter
	isSent bool
	logged bool
	// Request corresponding to this response (nil by default)
	req *http.Request
}

func New(w http.ResponseWriter) *response {
	return &response{
		writer: w,
		isSent: false,
		logged: false,
		req:    nil,
	}
}

// Writes given `body` in response and sends it.
func (res *response) send(body []byte) error {
	if res.isSent {
		return errors.New("response was already sent")
	}

	if _, writeError := res.writer.Write(body); writeError != nil {
		return writeError
	}

	return nil
}

// Sends response with status 200 and given `body`.
func (res *response) SendBody(body []byte) error {
	res.writer.WriteHeader(http.StatusOK)

	err := res.send(body)

	res.isSent = true

	return err
}

// Sends response with passed status and JSON {"message": <message>}
func (res *response) Message(message string, status int) error {
	body, err := json.Marshal(MessageResponseBody{Message: message})

	if err != nil {
		return err
	}

	res.writer.WriteHeader(status)

	if err := res.send(body); err != nil {
		return err
	}

	res.isSent = true

	if res.logged && res.req != nil {
		logger.Print(message, res.req)
	}

	return nil
}

// If this method was called, then on sending response also will be done log in terminal (similar to "logger.Print")
func (res *response) Logged(req *http.Request) *response {
	res.logged = true
	res.req = req

	return res
}

// Sends response with status 200 and JSON {"message": "OK"}
func (res *response) OK() error {
	return res.Message("OK", http.StatusOK)
}

// Sends response with status 400 and JSON {"message": <message>}
func (res *response) BadRequest(message string) error {
	return res.Message(message, http.StatusBadRequest)
}

// Sends response with status 401 and JSON {"message": <message>}
func (res *response) Unauthorized(message string) error {
	return res.Message(message, http.StatusUnauthorized)
}

// Sends response with status 403 and JSON {"message": <message>}
func (res *response) Forbidden(message string) error {
	return res.Message(message, http.StatusForbidden)
}

// Sends response with status 404 and JSON {"message": <message>}
func (res *response) NotFound(message string) error {
	return res.Message(message, http.StatusNotFound)
}

// Sends response with status 408 and JSON {"message": <message>}
func (res *response) RequestTimeout(message string) error {
	return res.Message(message, http.StatusRequestTimeout)
}

// Sends response with status 409 and JSON {"message": <message>}
func (res *response) Conflict(message string) error {
	return res.Message(message, http.StatusConflict)
}

// Sends response with status 500 and JSON {"message": "Internal Server Error"}
func (res *response) InternalServerError() error {
	return res.Message("Internal Server Error", http.StatusInternalServerError)
}
