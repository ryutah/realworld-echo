package usecase

// export for error_handler.go

type ExportErrorHandlerWithoutOutputPortConfig = errorHandlerConfig

type ErrorHandlerWithoutOutputPortErrorHandleOption = errorErrorHandleOption

func (e *ExportErrorHandlerWithoutOutputPortConfig) ErrorHandlerOptions() []errorErrorHandleOption {
	return e.errorHandlerOptions
}

func (e *ExportErrorHandlerWithoutOutputPortConfig) AddErrorHandlerOption(opt errorErrorHandleOption) {
	e.addErrorHandlerOption(opt)
}
