package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaults(t *testing.T) {
	t.Run("Should do nothing when page and size are valid", func(t *testing.T) {
		// Arrange
		p := Pagination{
			Page: 0,
			Size: 10,
		}

		expected := struct {
			Page int64
			Size int64
		}{
			Page: 0,
			Size: 10,
		}

		// Act
		p.SetDefaults()

		// Assert
		assert.Equal(t, expected.Page, p.Page)
		assert.Equal(t, expected.Size, p.Size)
	})

	t.Run("Should set default values when page is invalid", func(t *testing.T) {
		// Arrange
		p := Pagination{
			Page: -1,
			Size: 10,
		}

		expected := struct {
			Page int64
			Size int64
		}{
			Page: 0,
			Size: 10,
		}

		// Act
		p.SetDefaults()

		// Assert
		assert.Equal(t, expected.Page, p.Page)
		assert.Equal(t, expected.Size, p.Size)
	})

	t.Run("Should set default values when size is invalid", func(t *testing.T) {
		// Arrange
		sizes := []struct {
			size     int64
			expected int64
		}{
			{size: 0, expected: 10},
			{size: 101, expected: 100},
		}

		for _, tt := range sizes {
			p := Pagination{
				Page: 0,
				Size: tt.size,
			}

			expected := struct {
				Page int64
				Size int64
			}{
				Page: 0,
				Size: tt.expected,
			}

			// Act
			p.SetDefaults()

			// Assert
			assert.Equal(t, expected.Page, p.Page)
			assert.Equal(t, expected.Size, p.Size)
		}
	})
}
