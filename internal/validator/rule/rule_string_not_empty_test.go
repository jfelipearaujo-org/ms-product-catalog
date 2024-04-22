package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleStringNotEmpty_Validate(t *testing.T) {
	t.Run("Should return nil when value is not empty", func(t *testing.T) {
		// Arrange
		rule := NewStringNotEmpty()

		// Act
		err := rule.Validate("abc")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when value is empty", func(t *testing.T) {
		// Arrange
		rule := NewStringNotEmpty()

		// Act
		err := rule.Validate("")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should panic when value is not a string", func(t *testing.T) {
		// Arrange
		rule := NewStringNotEmpty()

		// Act & Assert
		assert.Panics(t, func() {
			err := rule.Validate(1)

			assert.NoError(t, err)
		})
	})
}
