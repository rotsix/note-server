package errors

import (
	"net/http"
)

// BadRequest is raised when a request is malformed
type BadRequest struct{}

func (e *BadRequest) Error() string {
	return "400 bad request"
}

// Unauthorized is raised when a guest is trying to access authenticated-only data
type Unauthorized struct{}

func (e *Unauthorized) Error() string {
	return "401 unauthorized"
}

// Forbidden is raised when an user is trying to access forbidden data
type Forbidden struct{}

func (e *Forbidden) Error() string {
	return "403 forbidden"
}

// NotFound is raised when an user tries to access unexisting data
type NotFound struct{}

func (e *NotFound) Error() string {
	return "404 not found"
}

// Internal is raised as a general catch-all error for server
type Internal struct{}

func (e *Internal) Error() string {
	return "500 internal"
}

// Manage handles errors
func Manage(rw http.ResponseWriter, err error) {
	sendErr := func(errCode int) {
		http.Error(rw, err.Error(), errCode)
	}

	switch err.(type) {
	case *BadRequest:
		sendErr(http.StatusBadRequest)
	case *Unauthorized:
		sendErr(http.StatusUnauthorized)
	case *Forbidden:
		sendErr(http.StatusForbidden)
	case *NotFound:
		sendErr(http.StatusNotFound)
	case *Internal:
		sendErr(http.StatusInternalServerError)
	default:
		panic(err)
	}
}
