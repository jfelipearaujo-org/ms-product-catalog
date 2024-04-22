package create_category

import (
	"context"
	"errors"
	"testing"
	"time"

	provider_mock "github.com/jfelipearaujo-org/ms-product-catalog/internal/provider/mocks"
	repository_mock "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return an error when the request is invalid", func(t *testing.T) {
		// Arrange
		repository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		service := NewService(
			repository,
			timeProvider)

		req := &CreateCategoryDto{
			Title:       "",
			Description: "Description",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})

	t.Run("Should return an error when the category can't be created", func(t *testing.T) {
		// Arrange
		repository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		timeProvider.On("GetTime").Return(time.Now())

		repository.On("Create", mock.Anything, mock.Anything).
			Return(errors.New("error")).
			Once()

		service := NewService(
			repository,
			timeProvider)

		req := &CreateCategoryDto{
			Title:       "Title",
			Description: "Description",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		repository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})

	t.Run("Should return nil when the category is created successfully", func(t *testing.T) {
		// Arrange
		repository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		timeProvider.On("GetTime").Return(time.Now())

		repository.On("Create", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		service := NewService(
			repository,
			timeProvider)

		req := &CreateCategoryDto{
			Title:       "Title",
			Description: "Description",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		repository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})
}
