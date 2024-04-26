package get_product

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when the request is invalid", func(t *testing.T) {
		// Arrange
		dto := GetProductDto{
			UUID: "abc",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return nil when the request is valid", func(t *testing.T) {
		// Arrange
		dto := GetProductDto{
			UUID: uuid.NewString(),
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})
}
