package errors

import (
	"errors"
)

type errorAndMessage struct {
	Err     error
	Message string
}

var Errors = struct {
	Validation errorAndMessage
	NotFound   errorAndMessage
}{
	Validation: errorAndMessage{
		Err:     errors.New("validation_error"),
		Message: "failed to validate fields",
	},
	NotFound: errorAndMessage{
		Err:     errors.New("not_found"),
		Message: "specified entity is not found",
	},
}
