package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpAppError(t *testing.T) {
	t.Run("Should return a new HTTP error", func(t *testing.T) {
		// Arrange
		expected := AppError{
			Code:    404,
			Message: "content not found",
			Details: "content not found",
		}

		// Act
		err := NewHttpAppError(404, "content not found", ErrNotFound)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, expected.Code, err.Code)
		assert.Equal(t, expected, err.Message)
	})
}
