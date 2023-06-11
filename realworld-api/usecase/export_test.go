package usecase

// export for error_handler.go

type ExportErrorHandlerConfig = errorHandlerConfig

func (e *ExportErrorHandlerConfig) AddRendrer(target error, f renderFunc) {
	e.rendrers = append(e.rendrers, errorRenderer{
		target:   target,
		renderer: f,
	})
}
