package usecase

// export for error_handler.go

type ErrorHandlerConifg = errorHandlerConfig

type ErrorHandlerFunc = errorHandlerFunc

func (e *ErrorHandlerConifg) Handlers() []errorHandlerFunc {
	return e.handlers
}

func (e *ErrorHandlerConifg) AddHandlers(opt errorHandlerFunc) {
	e.addHandlers(opt)
}
