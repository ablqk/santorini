package errors

import (
	"fmt"
)

type sError int

const (
	// BadRequestE Representation of a bad request
	BadRequestE sError = 400
	// NotFoundE Representation of a not found request
	NotFoundE sError = 404
	// InternalError Representation of an internal error
	InternalError sError = 500
)

func NewFromType(errorType int, message string) *HTTPError {
	return &HTTPError{httpCode: errorType, message: message}
}

// New Creation of a new error
// First parameter is mandatory and represents the main category of the error
func New(errorType sError, message ...interface{}) *HTTPError {
	base := HTTPError{httpCode: int(errorType)}
	return Wrap(&base, message...)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err *HTTPError, items ...interface{}) *HTTPError {
	space := ""
	what := ""
	for _, item := range items {
		what += fmt.Sprint(space, item)
		space = " "
	}
	return &HTTPError{httpCode: err.httpCode, message: what}
}

// NewFromError returns a service error constructed from
// a generic error passed in parameter
// Adds what to it's starting message
func NewFromError(err error, items ...interface{}) func(*HTTPError) {
	return func(e *HTTPError) {
		space := ""
		what := ""
		for _, item := range items {
			what += fmt.Sprint(space, item)
			space = " "
		}
		e.message = what + ", " + err.Error()
	}
}

// HTTPError The error type
type HTTPError struct {
	httpCode int
	message  string
}

// GetErrorValue returns the error category
func (e HTTPError) GetErrorValue() int {
	return e.httpCode
}

// Error is the interface implementation
func (e HTTPError) Error() string {
	return e.message
}
