package create_product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return nil when dto is valid", func(t *testing.T) {
		// Arrange
		dto := CreateProductDto{
			Title:         "Title",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when title is empty", func(t *testing.T) {
		// Arrange
		dto := CreateProductDto{
			Title:         "",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})
}
