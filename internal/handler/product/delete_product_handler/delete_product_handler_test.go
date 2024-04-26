package delete_product_handler

import (
	baseErr "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/mocks"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler(t *testing.T) {
	t.Run("Should delete the product", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockDeleteProductService(t)

		service.On("Handle", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		req := httptest.NewRequest(echo.DELETE, "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.Code)
		assert.Empty(t, resp.Body.String())
		service.AssertExpectations(t)
	})

	t.Run("Should return error when request is invalid", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockDeleteProductService(t)

		service.On("Handle", mock.Anything, mock.Anything).
			Return(errors.ErrRequestNotValid).
			Once()

		req := httptest.NewRequest(echo.DELETE, "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

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

	t.Run("Should return error when product does not exist", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockDeleteProductService(t)

		service.On("Handle", mock.Anything, mock.Anything).
			Return(repository.ErrProductNotFound).
			Once()

		req := httptest.NewRequest(echo.DELETE, "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

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
			Message: "error to find the product",
			Details: "product not found",
		}, he.Message)
		service.AssertExpectations(t)
	})

	t.Run("Should return error when product cannot be deleted", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockDeleteProductService(t)

		service.On("Handle", mock.Anything, mock.Anything).
			Return(baseErr.New("error")).
			Once()

		req := httptest.NewRequest(echo.DELETE, "/products", nil)
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		e := echo.New()
		ctx := e.NewContext(req, resp)
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")

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
