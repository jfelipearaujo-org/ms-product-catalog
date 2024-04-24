package create_product

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	provider_mock "github.com/jfelipearaujo-org/ms-product-catalog/internal/provider/mocks"
	repository_mock "github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return an error when the request is invalid", func(t *testing.T) {
		// Arrange
		productRepository := repository_mock.NewMockProductRepository(t)
		categoryRepository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		service := NewService(
			productRepository,
			categoryRepository,
			timeProvider)

		req := CreateProductDto{
			Title:         "",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
		categoryRepository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})

	t.Run("Should return an error when the category does not exist", func(t *testing.T) {
		// Arrange
		productRepository := repository_mock.NewMockProductRepository(t)
		categoryRepository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		categoryRepository.On("GetByTitle", mock.Anything, mock.Anything).
			Return(entity.Category{}, errors.New("error")).
			Once()

		service := NewService(
			productRepository,
			categoryRepository,
			timeProvider)

		req := CreateProductDto{
			Title:         "Title",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
		categoryRepository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})

	t.Run("Should return an error when the product can't be created", func(t *testing.T) {
		// Arrange
		productRepository := repository_mock.NewMockProductRepository(t)
		categoryRepository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		categoryRepository.On("GetByTitle", mock.Anything, mock.Anything).
			Return(entity.Category{}, nil).
			Once()

		timeProvider.On("GetTime").Return(time.Now())

		productRepository.On("Create", mock.Anything, mock.Anything).
			Return(errors.New("error")).
			Once()

		service := NewService(
			productRepository,
			categoryRepository,
			timeProvider)

		req := CreateProductDto{
			Title:         "Title",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.Error(t, err)
		productRepository.AssertExpectations(t)
		categoryRepository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})

	t.Run("Should return nil when the product is created successfully", func(t *testing.T) {
		// Arrange
		productRepository := repository_mock.NewMockProductRepository(t)
		categoryRepository := repository_mock.NewMockCategoryRepository(t)
		timeProvider := provider_mock.NewMockTimeProvider(t)

		categoryRepository.On("GetByTitle", mock.Anything, mock.Anything).
			Return(entity.Category{}, nil).
			Once()

		timeProvider.On("GetTime").Return(time.Now())

		productRepository.On("Create", mock.Anything, mock.Anything).
			Return(nil).
			Once()

		service := NewService(
			productRepository,
			categoryRepository,
			timeProvider)

		req := CreateProductDto{
			Title:         "Title",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := service.Handle(context.Background(), req)

		// Assert
		assert.NoError(t, err)
		productRepository.AssertExpectations(t)
		categoryRepository.AssertExpectations(t)
		timeProvider.AssertExpectations(t)
	})
}
