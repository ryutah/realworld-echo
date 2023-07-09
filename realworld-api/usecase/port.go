package usecase

import "context"

type OutputPort[T any] interface {
	Success(context.Context, T) error
}

type NewOutputPort[Arg, Ret any] interface {
	GenericsErrorOutputPort[Ret]
	Success(context.Context, Arg) (Ret, error)
}
