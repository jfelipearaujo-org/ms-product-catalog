package create_category

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rnd.Intn(len(letters))]
	}
	return string(b)
}

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

	t.Run("Should return error when description is empty", func(t *testing.T) {
		// Arrange
		category := &CreateCategoryDto{
			Title:       "title",
			Description: "",
		}

		// Act
		err := category.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return error when title length is invalid", func(t *testing.T) {
		// Arrange
		lengths := []int{2, 101}

		for _, length := range lengths {
			category := &CreateCategoryDto{
				Title:       randomString(length),
				Description: "description",
			}

			// Act
			err := category.Validate()

			// Assert
			assert.Error(t, err)
		}
	})

	t.Run("Should return error when description length is invalid", func(t *testing.T) {
		// Arrange
		lengths := []int{2, 201}

		for _, length := range lengths {
			category := &CreateCategoryDto{
				Title:       "title",
				Description: randomString(length),
			}

			// Act
			err := category.Validate()

			// Assert
			assert.Error(t, err)
		}
	})
}
