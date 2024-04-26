package delete_product

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should delete the product", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		repository.On("Delete", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		service := NewService(repository)

		req := DeleteProductDto{
			UUID: uuid.NewString(),
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when the request is invalid", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		service := NewService(repository)

		req := DeleteProductDto{
			UUID: "abc",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when the product can't be deleted", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		repository.On("Delete", mock.Anything, mock.Anything).
			Return(errors.New("error")).
			Once()

		service := NewService(repository)

		req := DeleteProductDto{
			UUID: uuid.NewString(),
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
	})
}
