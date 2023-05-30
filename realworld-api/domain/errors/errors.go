package errors

import (
	"errors"
)

type errorAndMessage struct {
	Err     error
	Message string
}

var Errors = struct {
	ErrValidation errorAndMessage
}{
	ErrValidation: errorAndMessage{
		Err:     errors.New("validation_error"),
		Message: "failed to validate fields",
	},
}
