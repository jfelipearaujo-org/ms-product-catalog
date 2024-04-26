package get_products

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
	t.Run("Should return nil when the request is valid", func(t *testing.T) {
		// Arrange
		dto := GetProductsDto{
			Title:         randomString(10),
			CategoryTitle: randomString(10),
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return an error when the title length is more than max", func(t *testing.T) {
		// Arrange
		dto := GetProductsDto{
			Title: randomString(101),
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return an error when the category title length is more than max", func(t *testing.T) {
		// Arrange
		dto := GetProductsDto{
			CategoryTitle: randomString(101),
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})
}
