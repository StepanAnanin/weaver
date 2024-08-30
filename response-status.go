package weaver

import "net/http"

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
