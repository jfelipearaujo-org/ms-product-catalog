package create_category_handler

import (
	"bytes"
	baseErr "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/mocks"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should create the category", func(t *testing.T) {
		// Arrange
		createCategoryService := mocks.NewMockCreateCategoryService(t)
		getCategoriesService := mocks.NewMockGetCategoriesService(t)

		getCategoriesService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createCategoryService.On("Handle", mock.Anything, mock.Anything).
			Return(&entity.Category{
				Title:       "category",
				Description: "description",
			}, nil).
			Once()

		reqBody := []byte(`{"title":"category", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/categories", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createCategoryService, getCategoriesService)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.JSONEq(t, `{"id":"", "title":"category", "description":"description", "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z"}`, resp.Body.String())
		createCategoryService.AssertExpectations(t)
		getCategoriesService.AssertExpectations(t)
	})

	t.Run("Should return error when something got wrong while try to check if category exists", func(t *testing.T) {
		// Arrange
		createCategoryService := mocks.NewMockCreateCategoryService(t)
		getCategoriesService := mocks.NewMockGetCategoriesService(t)

		getCategoriesService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, baseErr.New("error")).
			Once()

		reqBody := []byte(`{"title":"category", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/categories", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createCategoryService, getCategoriesService)

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
		createCategoryService.AssertExpectations(t)
		getCategoriesService.AssertExpectations(t)
	})

	t.Run("Should return error when already exists at least one category with same title", func(t *testing.T) {
		// Arrange
		createCategoryService := mocks.NewMockCreateCategoryService(t)
		getCategoriesService := mocks.NewMockGetCategoriesService(t)

		getCategoriesService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), nil, nil).
			Once()

		reqBody := []byte(`{"title":"category", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/categories", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createCategoryService, getCategoriesService)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusConflict, he.Code)
		assert.Equal(t, errors.AppError{
			Code:    http.StatusConflict,
			Message: "category cannot be created",
			Details: "category already exists",
		}, he.Message)
		createCategoryService.AssertExpectations(t)
		getCategoriesService.AssertExpectations(t)
	})

	t.Run("Should return error when request is invalid", func(t *testing.T) {
		// Arrange
		createCategoryService := mocks.NewMockCreateCategoryService(t)
		getCategoriesService := mocks.NewMockGetCategoriesService(t)

		getCategoriesService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createCategoryService.On("Handle", mock.Anything, mock.Anything).
			Return(nil, errors.ErrRequestNotValid).
			Once()

		reqBody := []byte(`{"title":"category", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/categories", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createCategoryService, getCategoriesService)

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
		createCategoryService.AssertExpectations(t)
		getCategoriesService.AssertExpectations(t)
	})

	t.Run("Should return a generic error", func(t *testing.T) {
		// Arrange
		createCategoryService := mocks.NewMockCreateCategoryService(t)
		getCategoriesService := mocks.NewMockGetCategoriesService(t)

		getCategoriesService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createCategoryService.On("Handle", mock.Anything, mock.Anything).
			Return(nil, baseErr.New("error")).
			Once()

		reqBody := []byte(`{"title":"category", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/categories", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createCategoryService, getCategoriesService)

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
		createCategoryService.AssertExpectations(t)
		getCategoriesService.AssertExpectations(t)
	})
}
