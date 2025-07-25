// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/LexusEgorov/items-accounting/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// CategoryRepository is an autogenerated mock type for the CategoryRepository type
type CategoryRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, name
func (_m *CategoryRepository) Add(ctx context.Context, name string) (int, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *CategoryRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *CategoryRepository) Get(ctx context.Context, id int) (models.Category, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 models.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (models.Category, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) models.Category); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, id, name
func (_m *CategoryRepository) Set(ctx context.Context, id int, name string) error {
	ret := _m.Called(ctx, id, name)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCategoryRepository creates a new instance of CategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryRepository {
	mock := &CategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
