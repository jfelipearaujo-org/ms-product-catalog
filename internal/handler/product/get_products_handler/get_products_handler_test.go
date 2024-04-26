package get_products_handler

import (
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
	t.Run("Should return the products", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetProductsService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), []entity.Product{
				{
					UUID:        "123",
					Title:       "Product",
					Description: "Product description",
				},
			}, nil).
			Once()

		req := httptest.NewRequest(echo.GET, "/products", nil)
		resp := httptest.NewRecorder()

		echo := echo.New()
		ctx := echo.NewContext(req, resp)

		handler := NewHandler(service)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{ "page":1, "total_items":1, "total_pages":1, "data":[{"id":"123", "title":"Product", "description":"Product description", "price": 0, "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z", "category": {"id":"", "title":"", "description":"", "created_at":"0001-01-01T00:00:00Z", "updated_at":"0001-01-01T00:00:00Z"}}]}`, resp.Body.String())
		service.AssertExpectations(t)
	})

	t.Run("Should return validation error", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetProductsService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), []entity.Product{}, errors.ErrRequestNotValid).
			Once()

		req := httptest.NewRequest(echo.GET, "/products", nil)
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

	t.Run("Should return internal server error", func(t *testing.T) {
		// Arrange
		service := mocks.NewMockGetProductsService(t)

		service.On("Handle", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), []entity.Product{}, baseErr.New("error")).
			Once()

		req := httptest.NewRequest(echo.GET, "/products", nil)
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
