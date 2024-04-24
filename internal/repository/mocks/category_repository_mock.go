// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

// MockCategoryRepository is an autogenerated mock type for the CategoryRepository type
type MockCategoryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, category
func (_m *MockCategoryRepository) Create(ctx context.Context, category *entity.Category) error {
	ret := _m.Called(ctx, category)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Category) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockCategoryRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, filter
func (_m *MockCategoryRepository) GetAll(ctx context.Context, filter repository.Pagination) (int64, []entity.Category, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 int64
	var r1 []entity.Category
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Pagination) (int64, []entity.Category, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, repository.Pagination) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, repository.Pagination) []entity.Category); ok {
		r1 = rf(ctx, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]entity.Category)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, repository.Pagination) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *MockCategoryRepository) GetByID(ctx context.Context, id string) (entity.Category, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 entity.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Category, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Category); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *MockCategoryRepository) GetByTitle(ctx context.Context, title string) (entity.Category, error) {
	ret := _m.Called(ctx, title)

	if len(ret) == 0 {
		panic("no return value specified for GetByTitle")
	}

	var r0 entity.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Category, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Category); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(entity.Category)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, category
func (_m *MockCategoryRepository) Update(ctx context.Context, category *entity.Category) error {
	ret := _m.Called(ctx, category)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Category) error); ok {
		r0 = rf(ctx, category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockCategoryRepository creates a new instance of MockCategoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCategoryRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCategoryRepository {
	mock := &MockCategoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}