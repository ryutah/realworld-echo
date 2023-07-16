package xvalidator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	_validatorOnce sync.Once
	_validator     *validator.Validate
)

func Validator() *validator.Validate {
	_validatorOnce.Do(func() {
		_validator = validator.New()
	})
	return _validator
}
