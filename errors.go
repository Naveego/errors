package errors

import "fmt"

// Error is the main error type that will be used by components of the
// pipeline.  It adds a machine readable code that can be used to deteremine
// the type of error.
type Error struct {
	Code    int    // The machine readable error code
	Message string // The human readable error message
}

// Error implements the native golang error interface
func (e Error) Error() string {
	return e.Message
}

// HTTPStatus returns the HTTP status code for the error
func (e Error) HTTPStatus() int {
	if e.Code >= 4000000 {
		return e.Code / 10000
	}

	return 500
}

// New implements the built-in errors.New method
func New(message string) error {
	return Error{Code: 5000000, Message: message}
}

// NewWithCode returns a error with the given code
func NewWithCode(code int, message string) error {
	return Error{Code: code, Message: message}
}

// NewWithCodef returns an error with formatting
func NewWithCodef(code int, message string, args ...interface{}) error {
	errMsg := fmt.Sprintf(message, args...)
	return NewWithCode(code, errMsg)
}
