package get_category_handler

import (
	baseErr "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	repository "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/mocks"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return the category", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetCategoryService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Category{}, nil).
			Once()

		req := httptest.NewRequest(echo.GET, "/categories", nil)
		resp := httptest.NewRecorder()

		echo := echo.New()
		ctx := echo.NewContext(req, resp)

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"id":"", "title":"", "description":"", "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z"}`, resp.Body.String())
		service.AssertExpectations(t)
	})

	t.Run("Should return validation error", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetCategoryService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Category{}, errors.ErrRequestNotValid).
			Once()

		req := httptest.NewRequest(echo.GET, "/categories", nil)
		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusUnprocessableEntity, he.Code)
		assert.Equal(t, errors.AppError{
			Code:    http.StatusUnprocessableEntity,
			Message: "validation error",
			Details: "request not valid, please check the fields",
		}, he.Message)
		service.AssertExpectations(t)
	})

	t.Run("Should return error when the category is not found", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetCategoryService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Category{}, repository.ErrCategoryNotFound).
			Once()

		req := httptest.NewRequest(echo.GET, "/categories", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusNotFound, he.Code)
		assert.Equal(t, errors.AppError{
			Code:    http.StatusNotFound,
			Message: "error to find the category",
			Details: "category not found",
		}, he.Message)
		service.AssertExpectations(t)
	})

	t.Run("Should return internal server error", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetCategoryService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(entity.Category{}, baseErr.New("error")).
			Once()

		req := httptest.NewRequest(echo.GET, "/categories", nil)
		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusInternalServerError, he.Code)
		assert.Equal(t, errors.AppError{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
			Details: "error",
		}, he.Message)
		service.AssertExpectations(t)
	})
}
