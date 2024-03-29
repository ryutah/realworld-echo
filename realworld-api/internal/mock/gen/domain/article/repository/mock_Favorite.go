// Code generated by mockery v2.32.0. DO NOT EDIT.

package repository

import (
	context "context"

	authmodel "github.com/ryutah/realworld-echo/realworld-api/domain/auth/model"

	mock "github.com/stretchr/testify/mock"

	model "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
)

// MockFavorite is an autogenerated mock type for the Favorite type
type MockFavorite struct {
	mock.Mock
}

type MockFavorite_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFavorite) EXPECT() *MockFavorite_Expecter {
	return &MockFavorite_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *MockFavorite) Count(_a0 context.Context, _a1 model.Slug) (int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Slug) (int, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Slug) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Slug) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFavorite_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockFavorite_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.Slug
func (_e *MockFavorite_Expecter) Count(_a0 interface{}, _a1 interface{}) *MockFavorite_Count_Call {
	return &MockFavorite_Count_Call{Call: _e.mock.On("Count", _a0, _a1)}
}

func (_c *MockFavorite_Count_Call) Run(run func(_a0 context.Context, _a1 model.Slug)) *MockFavorite_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Slug))
	})
	return _c
}

func (_c *MockFavorite_Count_Call) Return(_a0 int, _a1 error) *MockFavorite_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFavorite_Count_Call) RunAndReturn(run func(context.Context, model.Slug) (int, error)) *MockFavorite_Count_Call {
	_c.Call.Return(run)
	return _c
}

// CountList provides a mock function with given fields: _a0, _a1
func (_m *MockFavorite) CountList(_a0 context.Context, _a1 ...model.Slug) (model.FavoriteCountMap, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 model.FavoriteCountMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...model.Slug) (model.FavoriteCountMap, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...model.Slug) model.FavoriteCountMap); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.FavoriteCountMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...model.Slug) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFavorite_CountList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountList'
type MockFavorite_CountList_Call struct {
	*mock.Call
}

// CountList is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 ...model.Slug
func (_e *MockFavorite_Expecter) CountList(_a0 interface{}, _a1 ...interface{}) *MockFavorite_CountList_Call {
	return &MockFavorite_CountList_Call{Call: _e.mock.On("CountList",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *MockFavorite_CountList_Call) Run(run func(_a0 context.Context, _a1 ...model.Slug)) *MockFavorite_CountList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]model.Slug, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(model.Slug)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockFavorite_CountList_Call) Return(_a0 model.FavoriteCountMap, _a1 error) *MockFavorite_CountList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFavorite_CountList_Call) RunAndReturn(run func(context.Context, ...model.Slug) (model.FavoriteCountMap, error)) *MockFavorite_CountList_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockFavorite) Exists(_a0 context.Context, _a1 authmodel.UserID, _a2 model.Slug) (bool, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authmodel.UserID, model.Slug) (bool, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authmodel.UserID, model.Slug) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, authmodel.UserID, model.Slug) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFavorite_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockFavorite_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 authmodel.UserID
//   - _a2 model.Slug
func (_e *MockFavorite_Expecter) Exists(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockFavorite_Exists_Call {
	return &MockFavorite_Exists_Call{Call: _e.mock.On("Exists", _a0, _a1, _a2)}
}

func (_c *MockFavorite_Exists_Call) Run(run func(_a0 context.Context, _a1 authmodel.UserID, _a2 model.Slug)) *MockFavorite_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(authmodel.UserID), args[2].(model.Slug))
	})
	return _c
}

func (_c *MockFavorite_Exists_Call) Return(_a0 bool, _a1 error) *MockFavorite_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFavorite_Exists_Call) RunAndReturn(run func(context.Context, authmodel.UserID, model.Slug) (bool, error)) *MockFavorite_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// ExistsList provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockFavorite) ExistsList(_a0 context.Context, _a1 authmodel.UserID, _a2 ...model.Slug) (model.FavoriteExistsMap, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 model.FavoriteExistsMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, authmodel.UserID, ...model.Slug) (model.FavoriteExistsMap, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, authmodel.UserID, ...model.Slug) model.FavoriteExistsMap); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.FavoriteExistsMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, authmodel.UserID, ...model.Slug) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFavorite_ExistsList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExistsList'
type MockFavorite_ExistsList_Call struct {
	*mock.Call
}

// ExistsList is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 authmodel.UserID
//   - _a2 ...model.Slug
func (_e *MockFavorite_Expecter) ExistsList(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *MockFavorite_ExistsList_Call {
	return &MockFavorite_ExistsList_Call{Call: _e.mock.On("ExistsList",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *MockFavorite_ExistsList_Call) Run(run func(_a0 context.Context, _a1 authmodel.UserID, _a2 ...model.Slug)) *MockFavorite_ExistsList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]model.Slug, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(model.Slug)
			}
		}
		run(args[0].(context.Context), args[1].(authmodel.UserID), variadicArgs...)
	})
	return _c
}

func (_c *MockFavorite_ExistsList_Call) Return(_a0 model.FavoriteExistsMap, _a1 error) *MockFavorite_ExistsList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFavorite_ExistsList_Call) RunAndReturn(run func(context.Context, authmodel.UserID, ...model.Slug) (model.FavoriteExistsMap, error)) *MockFavorite_ExistsList_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFavorite creates a new instance of MockFavorite. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFavorite(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFavorite {
	mock := &MockFavorite{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
