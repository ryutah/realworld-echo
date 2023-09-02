package usecase

// export for error_handler.go

type ErrorHandlerFunc = errorHandlerFunc

func (e *ErrorHandlerConfig) Handlers() []errorHandlerFunc {
	return e.handlers
}

func (e *ErrorHandlerConfig) AddHandlers(opt errorHandlerFunc) {
	e.addHandlers(opt)
}
