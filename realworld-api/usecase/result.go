package usecase

type FailType int

const (
	FailTypeInternalError FailType = iota + 1
	FailTypeBadRequest
	FailTypeNotFound
	FailTypeUnauthorized
)

type FailResult struct {
	Type        FailType
	Message     string
	Description []string
}

func NewFailResult(typ FailType, message string, description ...string) *FailResult {
	return &FailResult{
		Type:        typ,
		Message:     message,
		Description: description,
	}
}

type Result[R any] struct {
	isFail        bool
	successResult R
	failResult    *FailResult
}

func Success[R any](result R) *Result[R] {
	return &Result[R]{
		isFail:        false,
		successResult: result,
	}
}

func Fail[R any](result *FailResult) *Result[R] {
	return &Result[R]{
		isFail:     true,
		failResult: result,
	}
}

func (r *Result[R]) IsFailed() bool {
	return r.isFail
}

func (r *Result[R]) Success() R {
	return r.successResult
}

func (r *Result[R]) Fail() *FailResult {
	return r.failResult
}
