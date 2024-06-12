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

// Writes given `body` in response. Returns nil if success, error otherwise.
//
// Also in error case sends error response with status 500. (Using `InternalServerError` method)
func (res *response) writeBody(body []byte) error {
	if res.isSent {
		return errors.New("response was already sent")
	}

	if _, writeError := res.writer.Write(body); writeError != nil {
		log.Println("[ ERROR ] Failed to write in response body (status 500)")

		err := res.InternalServerError()

		if err != nil {
			return err
		}

		return writeError
	}

	return nil
}

// Sends response with status 200 and given `body`.
// If error was return, that mean response has already been sent.
func (res *response) Send(body []byte) error {
	res.writer.WriteHeader(http.StatusOK)

	err := res.writeBody(body)

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

// Sends response with status 500 and JSON {"message": "Internal Server Error"}
func (res *response) InternalServerError() error {
	body, err := json.Marshal(MessageResponseBody{Message: "Internal Server Error"})

	if err != nil {
		log.Println("[ ERROR ] Failed to marshal json (status 500)")

		return err
	}

	res.writer.WriteHeader(http.StatusInternalServerError)

	// IMPORTANT
	// Don't use `writeBody` here, it may cause infinite recursion.
	// (cuz this method used inside `writeBody`)
	if _, writeError := res.writer.Write(body); writeError != nil {
		log.Println("[ ERROR ] Failed to write in response body (status 500)")

		return writeError
	}

	res.isSent = true

	return nil
}

// Sends response with passed status and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) Message(message string, status int) error {
	body, err := json.Marshal(MessageResponseBody{Message: message})

	// TODO duplicates, move to a new function
	if err != nil {
		log.Println("[ ERROR ] Failed to marshal json.")

		if e := res.InternalServerError(); e != nil {
			panic(e)
		}

		return err
	}

	res.writer.WriteHeader(status)

	if err := res.writeBody(body); err != nil {
		return err
	}

	res.isSent = true

	return nil
}

// Sends response with status 200 and JSON {"message": "OK"}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) OK() error {
	return res.Message("OK", http.StatusOK)
}

// Sends response with status 400 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) BadRequest(message string) error {
	return res.Message(message, http.StatusBadRequest)
}

// Sends response with status 401 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) Unauthorized(message string) error {
	return res.Message(message, http.StatusUnauthorized)
}

// Sends response with status 403 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) Forbidden(message string) error {
	return res.Message(message, http.StatusForbidden)
}

// Sends response with status 404 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) NotFound(message string) error {
	return res.Message(message, http.StatusNotFound)
}

// Sends response with status 408 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) RequestTimeout(message string) error {
	return res.Message(message, http.StatusNotFound)
}

// Sends response with status 409 and JSON {"message": <message>}
//
// Also handles all possible errors, if one was return, that mean response has already been sent.
func (res *response) Conflict(message string) error {
	return res.Message(message, http.StatusNotFound)
}
