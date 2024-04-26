package get_products

import (
	"context"
	"errors"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	t.Run("Should return the products", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		repository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(1), []entity.Product{{}}, nil).
			Once()

		service := NewService(repository)

		req := GetProductsDto{}

		expected := struct {
			count    int64
			products []entity.Product
		}{
			count:    1,
			products: []entity.Product{{}},
		}

		// Act
		count, products, err := service.Handle(context.Background(), common.Pagination{}, req)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, expected.products, products)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when the request is invalid", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		service := NewService(repository)

		req := GetProductsDto{
			Title: randomString(101),
		}

		expected := struct {
			count    int64
			products []entity.Product
		}{
			count:    0,
			products: []entity.Product{},
		}

		// Act
		count, products, err := service.Handle(context.Background(), common.Pagination{}, req)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, expected.products, products)
		repository.AssertExpectations(t)
	})

	t.Run("Should return an error when the products can't be retrieved", func(t *testing.T) {
		// Arrange
		repository := mocks.NewMockProductRepository(t)

		repository.On("GetAll", mock.Anything, mock.Anything, mock.Anything).
			Return(int64(0), []entity.Product{}, errors.New("error")).
			Once()

		service := NewService(repository)

		req := GetProductsDto{}

		expected := struct {
			count    int64
			products []entity.Product
		}{
			count:    0,
			products: []entity.Product{},
		}

		// Act
		count, products, err := service.Handle(context.Background(), common.Pagination{}, req)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expected.count, count)
		assert.Equal(t, expected.products, products)
		repository.AssertExpectations(t)
	})
}
