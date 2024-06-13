package response

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/StepanAnanin/weaver/logger"
)

type response struct {
	writer http.ResponseWriter
	isSent bool
}

func New(w http.ResponseWriter) *response {
	return &response{
		writer: w,
		isSent: false,
	}
}

// Writes given `body` in response and sends it.
func (res *response) send(body []byte) error {
	if res.isSent {
		return errors.New("response was already sent")
	}

	if _, writeError := res.writer.Write(body); writeError != nil {
		log.Println("[ ERROR ] Failed to write in response body (status 500)")

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

// Sends response with given message and status, also log request info in terminal.
//
// Returns error if failed to send response (also does log in this case), nil otherwise.
func (res *response) SendError(message string, status int, req *http.Request) error {
	err := res.Message(message, status)

	if err != nil {
		logger.PrintError("Failed to send response", http.StatusInternalServerError, req)

		return err
	}

	logger.PrintError(message, status, req)

	return nil
}

// Sends response with passed status and JSON {"message": <message>}
func (res *response) Message(message string, status int) error {
	body, err := json.Marshal(MessageResponseBody{Message: message})

	if err != nil {
		log.Println("[ ERROR ] Failed to marshal json.")

		return err
	}

	res.writer.WriteHeader(status)

	if err := res.send(body); err != nil {
		return err
	}

	res.isSent = true

	return nil
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
	return res.Message(message, http.StatusNotFound)
}

// Sends response with status 409 and JSON {"message": <message>}
func (res *response) Conflict(message string) error {
	return res.Message(message, http.StatusNotFound)
}

// Sends response with status 500 and JSON {"message": "Internal Server Error"}
func (res *response) InternalServerError() error {
	return res.Message("Internal Server Error", http.StatusInternalServerError)
}
