package get_categories

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return categories", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		id := uuid.NewString()

		repository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), []entity.Category{
				{
					UUID:        id,
					Title:       "title",
					Description: "description",
				},
			}, nil).Once()

		service := NewService(repository)

		expected := struct {
			count      int64
			categories []entity.Category
		}{
			count: 1,
			categories: []entity.Category{
				{
					UUID:        id,
					Title:       "title",
					Description: "description",
				},
			},
		}

		// Act
		count, resp, err := service.Handle(context.Background(), common.Pagination{}, GetCategoriesDto{})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, resp, expected.categories)
		repository.AssertExpectations(t)
	})

	t.Run("Should return error when something got wrong", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		repository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), []entity.Category{}, errors.New("error")).
			Once()

		service := NewService(repository)

		expected := struct {
			count      int64
			categories []entity.Category
		}{
			count:      0,
			categories: []entity.Category{},
		}

		// Act
		count, resp, err := service.Handle(context.Background(), common.Pagination{}, GetCategoriesDto{})

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, resp, expected.categories)
		repository.AssertExpectations(t)
	})

	t.Run("Should return error when request is invalid", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockCategoryRepository(t)

		service := NewService(repository)

		expected := struct {
			count      int64
			categories []entity.Category
		}{
			count:      0,
			categories: []entity.Category{},
		}

		req := GetCategoriesDto{
			Title: randomString(101),
		}

		// Act
		count, resp, err := service.Handle(context.Background(), common.Pagination{}, req)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, resp, expected.categories)
		repository.AssertExpectations(t)
	})
}
