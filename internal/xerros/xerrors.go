package xerros

import (
	"errors"
	"fmt"
)

var (
	ErrValidation   = errors.New("validation_error")
	ErrNoSuchEntity = errors.New("no_such_entity")
)

func NewErrValidation(msg string, causes ...[]error) error {
	return fmt.Errorf("%w: %s\ncauses: [%v]", ErrValidation, msg, causes)
}
