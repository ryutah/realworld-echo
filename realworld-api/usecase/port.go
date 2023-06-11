package usecase

import "context"

type OutputPort[T any] interface {
	Success(context.Context, T) error
}
