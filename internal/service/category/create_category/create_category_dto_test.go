package create_category

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return nil when dto is valid", func(t *testing.T) {
		// Arrange
		category := &CreateCategoryDto{
			Title:       "title",
			Description: "description",
		}

		// Act
		err := category.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when title is empty", func(t *testing.T) {
		// Arrange
		category := &CreateCategoryDto{
			Title:       "",
			Description: "description",
		}

		// Act
		err := category.Validate()

		// Assert
		assert.Error(t, err)
	})
}
