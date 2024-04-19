package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	// Arrange
	now := time.Now()

	expected := Product{
		Title:       "Product 1",
		Description: "Description 1",
		Price:       10.5,
		Category: Category{
			UUID:        "uuid",
			Title:       "Category 1",
			Description: "Description 1",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Act
	product := NewProduct(
		"Product 1",
		"Description 1",
		10.5,
		Category{
			UUID:        "uuid",
			Title:       "Category 1",
			Description: "Description 1",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		now)

	// Assert
	assert.Equal(t, expected.Title, product.Title)
	assert.Equal(t, expected.Description, product.Description)
	assert.Equal(t, expected.Price, product.Price)
	assert.Equal(t, expected.Category.UUID, product.Category.UUID)
	assert.Equal(t, expected.Category.Title, product.Category.Title)
	assert.Equal(t, expected.Category.Description, product.Category.Description)
	assert.Equal(t, expected.CreatedAt, product.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, product.UpdatedAt)
}

func TestUpdateProduct(t *testing.T) {
	// Arrange
	old := time.Now()
	now := time.Now()

	product := NewProduct(
		"Product 1",
		"Description 1",
		10.5,
		Category{
			UUID:        "uuid",
			Title:       "Category 1",
			Description: "Description 1",
			CreatedAt:   old,
			UpdatedAt:   old,
		},
		old)

	expected := Product{
		UUID:        product.UUID,
		Title:       "Product 2",
		Description: "Description 2",
		Price:       22.5,
		Category: Category{
			UUID:        "uuid",
			Title:       "Category 2",
			Description: "Description 2",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		CreatedAt: old,
		UpdatedAt: now,
	}

	// Act
	product.Update(
		"Product 2",
		"Description 2",
		22.5,
		Category{
			UUID:        "uuid",
			Title:       "Category 2",
			Description: "Description 2",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		now)

	// Assert
	assert.Equal(t, &expected, product)
}
