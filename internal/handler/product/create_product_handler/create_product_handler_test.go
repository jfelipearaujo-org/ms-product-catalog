package create_product_handler

import (
	"bytes"
	baseErr "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/mocks"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should create the product", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createProductService.On("Handle", mock.Anything, mock.Anything).
			Return(&entity.Product{
				Title:       "product",
				Description: "description",
			}, nil).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.JSONEq(t, `{"id":"", "title":"product", "description":"description", "price": 0, "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z", "category": {"id":"", "title":"", "description":"", "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z"}}`, resp.Body.String())
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})

	t.Run("Should return error when something got wrong while try to check if product exists", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, baseErr.New("error")).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

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
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})

	t.Run("Should return error when already exists at least one product with same title", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), nil, nil).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.Error(t, err)

		he, ok := err.(*echo.HTTPError)
		assert.True(t, ok)

		assert.Equal(t, http.StatusConflict, he.Code)
		assert.Equal(t, errors.AppError{
			Code:    http.StatusConflict,
			Message: "product cannot be created",
			Details: "product already exists",
		}, he.Message)
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})

	t.Run("Should return error when request is not valid", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createProductService.On("Handle", mock.Anything, mock.Anything).
			Return(nil, errors.ErrRequestNotValid).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

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
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})

	t.Run("Should return error when category not found", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createProductService.On("Handle", mock.Anything, mock.Anything).
			Return(nil, repository.ErrCategoryNotFound).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

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
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})

	t.Run("Should return error when something got wrong while try to create the product", func(t *testing.T) {
		// Arrange
		createProductService := mocks.NewMockCreateProductService(t)
		getProductsService := mocks.NewMockGetProductsService(t)

		getProductsService.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), nil, nil).
			Once()

		createProductService.On("Handle", mock.Anything, mock.Anything).
			Return(nil, baseErr.New("error")).
			Once()

		reqBody := []byte(`{"title":"product", "description":"description"}`)

		req := httptest.NewRequest(echo.POST, "/products", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)

		handler := NewHandler(createProductService, getProductsService)

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
		createProductService.AssertExpectations(t)
		getProductsService.AssertExpectations(t)
	})
}
