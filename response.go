package weaver

import (
	"encoding/json"
	"errors"
	"net/http"
)

type response struct {
	writer http.ResponseWriter
	isSent bool
	logged bool
	// Request corresponding to this response (nil by default)
	req *http.Request
}

func NewResponse(w http.ResponseWriter) *response {
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

type MessageResponseBody struct {
	Message string `json:"message"`
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
		LogRequest(message, res.req)
	}

	return nil
}

// If this method was called, then on sending response also will be done log in terminal (similar to "logger.Print")
func (res *response) Logged(req *http.Request) *response {
	res.logged = true
	res.req = req

	return res
}
