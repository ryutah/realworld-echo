// Code generated by mockery v2.32.0. DO NOT EDIT.

package repository

import (
	context "context"

	model "github.com/ryutah/realworld-echo/realworld-api/domain/article/model"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/ryutah/realworld-echo/realworld-api/domain/article/repository"
)

// MockArticle is an autogenerated mock type for the Article type
type MockArticle struct {
	mock.Mock
}

type MockArticle_Expecter struct {
	mock *mock.Mock
}

func (_m *MockArticle) EXPECT() *MockArticle_Expecter {
	return &MockArticle_Expecter{mock: &_m.Mock}
}

// GenerateID provides a mock function with given fields: _a0
func (_m *MockArticle) GenerateID(_a0 context.Context) (model.Slug, error) {
	ret := _m.Called(_a0)

	var r0 model.Slug
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (model.Slug, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) model.Slug); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(model.Slug)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticle_GenerateID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateID'
type MockArticle_GenerateID_Call struct {
	*mock.Call
}

// GenerateID is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockArticle_Expecter) GenerateID(_a0 interface{}) *MockArticle_GenerateID_Call {
	return &MockArticle_GenerateID_Call{Call: _e.mock.On("GenerateID", _a0)}
}

func (_c *MockArticle_GenerateID_Call) Run(run func(_a0 context.Context)) *MockArticle_GenerateID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockArticle_GenerateID_Call) Return(_a0 model.Slug, _a1 error) *MockArticle_GenerateID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticle_GenerateID_Call) RunAndReturn(run func(context.Context) (model.Slug, error)) *MockArticle_GenerateID_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockArticle) Get(_a0 context.Context, _a1 model.Slug) (*model.Article, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Slug) (*model.Article, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Slug) *model.Article); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Slug) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticle_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockArticle_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.Slug
func (_e *MockArticle_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockArticle_Get_Call {
	return &MockArticle_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockArticle_Get_Call) Run(run func(_a0 context.Context, _a1 model.Slug)) *MockArticle_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Slug))
	})
	return _c
}

func (_c *MockArticle_Get_Call) Return(_a0 *model.Article, _a1 error) *MockArticle_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticle_Get_Call) RunAndReturn(run func(context.Context, model.Slug) (*model.Article, error)) *MockArticle_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: _a0, _a1
func (_m *MockArticle) Save(_a0 context.Context, _a1 model.Article) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Article) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockArticle_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockArticle_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 model.Article
func (_e *MockArticle_Expecter) Save(_a0 interface{}, _a1 interface{}) *MockArticle_Save_Call {
	return &MockArticle_Save_Call{Call: _e.mock.On("Save", _a0, _a1)}
}

func (_c *MockArticle_Save_Call) Run(run func(_a0 context.Context, _a1 model.Article)) *MockArticle_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Article))
	})
	return _c
}

func (_c *MockArticle_Save_Call) Return(_a0 error) *MockArticle_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockArticle_Save_Call) RunAndReturn(run func(context.Context, model.Article) error) *MockArticle_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockArticle) Search(_a0 context.Context, _a1 repository.ArticleSearchParam) ([]model.Article, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []model.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.ArticleSearchParam) ([]model.Article, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.ArticleSearchParam) []model.Article); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.ArticleSearchParam) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockArticle_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockArticle_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 repository.ArticleSearchParam
func (_e *MockArticle_Expecter) Search(_a0 interface{}, _a1 interface{}) *MockArticle_Search_Call {
	return &MockArticle_Search_Call{Call: _e.mock.On("Search", _a0, _a1)}
}

func (_c *MockArticle_Search_Call) Run(run func(_a0 context.Context, _a1 repository.ArticleSearchParam)) *MockArticle_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(repository.ArticleSearchParam))
	})
	return _c
}

func (_c *MockArticle_Search_Call) Return(_a0 []model.Article, _a1 error) *MockArticle_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockArticle_Search_Call) RunAndReturn(run func(context.Context, repository.ArticleSearchParam) ([]model.Article, error)) *MockArticle_Search_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockArticle creates a new instance of MockArticle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockArticle(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockArticle {
	mock := &MockArticle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}