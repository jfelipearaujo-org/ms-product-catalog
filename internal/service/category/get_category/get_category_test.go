package get_category

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return category", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		id := uuid.NewString()

		repository.On("GetByID", mock.Anything, mock.Anything).
			Return(entity.Category{
				UUID:        id,
				Title:       "title",
				Description: "description",
			}, nil).
			Once()

		service := NewService(repository)

		req := GetCategoryDto{
			UUID: id,
		}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		repository.AssertExpectations(t)
	})

	t.Run("Should return error when request is invalid", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		service := NewService(repository)

		req := GetCategoryDto{
			UUID: "",
		}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		assert.Empty(t, resp)
		repository.AssertExpectations(t)
	})

	t.Run("Should return error when repository returns error", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		id := uuid.NewString()

		repository.On("GetByID", mock.Anything, mock.Anything).
			Return(entity.Category{}, errors.New("error")).
			Once()

		service := NewService(repository)

		req := GetCategoryDto{
			UUID: id,
		}

		// Act
		resp, err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		assert.Empty(t, resp)
		repository.AssertExpectations(t)
	})
}
