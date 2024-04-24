package get_product

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("Should return a product by ID", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		id := uuid.NewString()

		repository.On("GetByID", context.Background(), "123").
			Return(entity.Product{
				UUID:        id,
				Title:       "Title",
				Description: "Description",
			}, nil).
			Once()

		service := NewService(repository)

		req := GetProductDto{
			UUID: "123",
		}

		expected := entity.Product{
			UUID:        id,
			Title:       "Title",
			Description: "Description",
		}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp)
		repository.AssertExpectations(t)
	})

	t.Run("Should return a product by Title", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		id := uuid.NewString()

		repository.On("GetByTitle", context.Background(), "Title").
			Return(entity.Product{
				UUID:        id,
				Title:       "Title",
				Description: "Description",
			}, nil).
			Once()

		service := NewService(repository)

		req := GetProductDto{
			Title: "Title",
		}

		expected := entity.Product{
			UUID:        id,
			Title:       "Title",
			Description: "Description",
		}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when request is invalid", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		service := NewService(repository)

		req := GetProductDto{}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidRequest)
		assert.Empty(t, resp)
		repository.AssertExpectations(t)
	})
}
