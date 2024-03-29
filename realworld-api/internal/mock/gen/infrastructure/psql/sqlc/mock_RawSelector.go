// Code generated by mockery v2.32.0. DO NOT EDIT.

package sqlc

import (
	context "context"

	squirrel "github.com/Masterminds/squirrel"
	sqlc "github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
	mock "github.com/stretchr/testify/mock"
)

// MockRawSelector is an autogenerated mock type for the RawSelector type
type MockRawSelector[T interface{}] struct {
	mock.Mock
}

type MockRawSelector_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockRawSelector[T]) EXPECT() *MockRawSelector_Expecter[T] {
	return &MockRawSelector_Expecter[T]{mock: &_m.Mock}
}

// Select provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockRawSelector[T]) Select(_a0 context.Context, _a1 sqlc.ContextExecutor, _a2 squirrel.SelectBuilder) ([]T, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []T
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.ContextExecutor, squirrel.SelectBuilder) ([]T, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sqlc.ContextExecutor, squirrel.SelectBuilder) []T); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]T)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, sqlc.ContextExecutor, squirrel.SelectBuilder) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRawSelector_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type MockRawSelector_Select_Call[T interface{}] struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 sqlc.ContextExecutor
//   - _a2 squirrel.SelectBuilder
func (_e *MockRawSelector_Expecter[T]) Select(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockRawSelector_Select_Call[T] {
	return &MockRawSelector_Select_Call[T]{Call: _e.mock.On("Select", _a0, _a1, _a2)}
}

func (_c *MockRawSelector_Select_Call[T]) Run(run func(_a0 context.Context, _a1 sqlc.ContextExecutor, _a2 squirrel.SelectBuilder)) *MockRawSelector_Select_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(sqlc.ContextExecutor), args[2].(squirrel.SelectBuilder))
	})
	return _c
}

func (_c *MockRawSelector_Select_Call[T]) Return(_a0 []T, _a1 error) *MockRawSelector_Select_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRawSelector_Select_Call[T]) RunAndReturn(run func(context.Context, sqlc.ContextExecutor, squirrel.SelectBuilder) ([]T, error)) *MockRawSelector_Select_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewMockRawSelector creates a new instance of MockRawSelector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRawSelector[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRawSelector[T] {
	mock := &MockRawSelector[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
