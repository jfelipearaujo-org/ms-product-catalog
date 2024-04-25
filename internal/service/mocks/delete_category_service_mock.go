// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	delete_category "github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/delete_category"
	mock "github.com/stretchr/testify/mock"
)

// MockDeleteCategoryService is an autogenerated mock type for the DeleteCategoryService type
type MockDeleteCategoryService struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, request
func (_m *MockDeleteCategoryService) Handle(ctx context.Context, request delete_category.DeleteCategoryDto) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, delete_category.DeleteCategoryDto) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockDeleteCategoryService creates a new instance of MockDeleteCategoryService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteCategoryService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteCategoryService {
	mock := &MockDeleteCategoryService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
