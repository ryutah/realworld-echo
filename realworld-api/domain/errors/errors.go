package errors

import (
	"errors"

	cerrors "github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
)

type errorAndMessage struct {
	Err     error
	Message string
}

var Errors = struct {
	Validation    errorAndMessage
	NotFound      errorAndMessage
	NotAuthorized errorAndMessage
	Internal      errorAndMessage
}{
	Validation: errorAndMessage{
		Err:     errors.New("validation_error"),
		Message: "failed to validate fields",
	},
	NotFound: errorAndMessage{
		Err:     errors.New("not_found"),
		Message: "specified entity is not found",
	},
	NotAuthorized: errorAndMessage{
		Err:     errors.New("not_authorized"),
		Message: "not authorized to perform this action",
	},
	Internal: errorAndMessage{
		Err:     errors.New("internal_error"),
		Message: "internal error",
	},
}

func NewValidationError(depth int, err error) error {
	newErr := cerrors.WrapWithDepth(depth+1, Errors.Validation.Err, Errors.Validation.Message)

	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return cerrors.WithDetail(newErr, err.Error())
	}
	for _, ve := range verrs {
		newErr = cerrors.WithDetail(newErr, ve.Error())
	}
	return newErr
}

func NewNotFoundError(depth int, err error, msg string) error {
	return cerrors.WithMessage(
		cerrors.WithDetailf(
			cerrors.WrapWithDepth(depth+1, Errors.NotFound.Err, Errors.NotFound.Message),
			"%+v", err,
		),
		msg,
	)
}

func NewInternalError(depth int, err error, msg string) error {
	return cerrors.WithMessage(
		cerrors.WithDetailf(
			cerrors.WrapWithDepth(depth+1, Errors.Internal.Err, Errors.Internal.Message),
			"%v: %+v", msg, err,
		),
		msg,
	)
}
