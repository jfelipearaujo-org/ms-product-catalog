package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaginationResponse(t *testing.T) {
	t.Run("Should return a new pagination response", func(t *testing.T) {
		// Arrange
		data := []string{"a", "b", "c"}

		expected := &PaginationResponse[string]{
			Page:       1,
			TotalPages: 2,
			TotalItems: 3,
			Data:       data,
		}

		// Act
		resp := NewPaginationResponse(1, 2, 3, data)

		// Assert
		assert.Equal(t, expected, resp)
	})

	t.Run("Should return a new pagination response with empty data", func(t *testing.T) {
		// Arrange
		expected := &PaginationResponse[string]{
			Page:       1,
			TotalPages: 2,
			TotalItems: 3,
			Data:       []string{},
		}

		// Act
		resp := NewPaginationResponse[string](1, 2, 3, nil)

		// Assert
		assert.Equal(t, expected, resp)
	})
}
