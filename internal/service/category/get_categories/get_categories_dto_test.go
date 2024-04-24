package get_categories

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
		dto := GetCategoriesDto{
			Title: "title",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when title is empty", func(t *testing.T) {
		// Arrange
		dto := GetCategoriesDto{
			Title: "",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return error when title is length is invalid", func(t *testing.T) {
		// Arrange
		lengths := []int{1, 101}

		for _, length := range lengths {
			dto := GetCategoriesDto{
				Title: randomString(length),
			}

			// Act
			err := dto.Validate()

			// Assert
			assert.Error(t, err)
		}
	})
}
