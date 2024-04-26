package delete_product

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return validation error", func(t *testing.T) {
		// Arrange
		dto := DeleteProductDto{}

		// Act
		err := dto.Validate()

		// Assert
		assert.Equal(t, errors.ErrRequestNotValid, err)
	})

	t.Run("Should return nil", func(t *testing.T) {
		// Arrange
		dto := DeleteProductDto{
			UUID: uuid.NewString(),
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Nil(t, err)
	})
}
