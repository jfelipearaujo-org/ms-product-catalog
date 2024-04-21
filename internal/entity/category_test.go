package entity

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	// Arrange
	now := time.Now()

	expected := Category{
		Title:       "Category 1",
		Description: "Description 1",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// Act
	category := NewCategory("Category 1", "Description 1", now)

	// Assert
	assert.NotNil(t, category)
	assert.NoError(t, uuid.Validate(category.UUID))
	assert.Equal(t, expected.Title, category.Title)
	assert.Equal(t, expected.Description, category.Description)
	assert.Equal(t, expected.CreatedAt, category.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, category.UpdatedAt)
}

func TestUpdateCategory(t *testing.T) {
	// Arrange
	old := time.Now().Add(-time.Hour)
	now := time.Now()

	expected := Category{
		Title:       "Category 2",
		Description: "Description 2",
		CreatedAt:   old,
		UpdatedAt:   now,
	}

	category := NewCategory("Category 1", "Description 1", old)

	// Act
	category.Update("Category 2", "Description 2", now)

	// Assert
	assert.NotNil(t, category)
	assert.Equal(t, expected.Title, category.Title)
	assert.Equal(t, expected.Description, category.Description)
	assert.Equal(t, expected.CreatedAt, category.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, category.UpdatedAt)
}
