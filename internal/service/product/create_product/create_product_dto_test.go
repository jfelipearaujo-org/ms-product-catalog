package create_product

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

	t.Run("Should return error when description is empty", func(t *testing.T) {
		// Arrange
		dto := CreateProductDto{
			Title:         "Title",
			Description:   "",
			Price:         10.0,
			CategoryTitle: "Category",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return error when category is empty", func(t *testing.T) {
		// Arrange
		dto := CreateProductDto{
			Title:         "Title",
			Description:   "Description",
			Price:         10.0,
			CategoryTitle: "",
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should return error when price is invalid", func(t *testing.T) {
		prices := []float64{0.4, 1000.1}

		for _, price := range prices {
			// Arrange
			dto := CreateProductDto{
				Title:         "Title",
				Description:   "Description",
				Price:         price,
				CategoryTitle: "Category",
			}

			// Act
			err := dto.Validate()

			// Assert
			assert.Error(t, err)
		}
	})

	t.Run("Should return error when title length is invalid", func(t *testing.T) {
		lengths := []int{2, 101}

		for _, length := range lengths {
			// Arrange
			dto := CreateProductDto{
				Title:         randomString(length),
				Description:   "Description",
				Price:         10.0,
				CategoryTitle: "Category",
			}

			// Act
			err := dto.Validate()

			// Assert
			assert.Error(t, err)
		}
	})

	t.Run("Should return error when description length is invalid", func(t *testing.T) {
		lengths := []int{9, 201}

		for _, length := range lengths {
			// Arrange
			dto := CreateProductDto{
				Title:         "title",
				Description:   randomString(length),
				Price:         10.0,
				CategoryTitle: "Category",
			}

			// Act
			err := dto.Validate()

			// Assert
			assert.Error(t, err)
		}
	})

	t.Run("Should return error when category length is invalid", func(t *testing.T) {
		lengths := []int{2, 101}

		for _, length := range lengths {
			// Arrange
			dto := CreateProductDto{
				Title:         "title",
				Description:   "description",
				Price:         10.0,
				CategoryTitle: randomString(length),
			}

			// Act
			err := dto.Validate()

			// Assert
			assert.Error(t, err)
		}
	})
}
