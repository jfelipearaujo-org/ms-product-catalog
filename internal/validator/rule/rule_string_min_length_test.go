package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleStringMinLength_Validate(t *testing.T) {
	t.Run("Should return nil when value is more than min length", func(t *testing.T) {
		// Arrange
		rule := NewStringMinLength(3)

		// Act
		err := rule.Validate("abcd")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return nil when value is equal to min length", func(t *testing.T) {
		// Arrange
		rule := NewStringMinLength(3)

		// Act
		err := rule.Validate("abc")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when value is less than min length", func(t *testing.T) {
		// Arrange
		rule := NewStringMinLength(3)

		// Act
		err := rule.Validate("ab")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should panic when value is not a string", func(t *testing.T) {
		// Arrange
		rule := NewStringMinLength(3)

		// Act & Assert
		assert.Panics(t, func() {
			err := rule.Validate(1)

			assert.NoError(t, err)
		})
	})
}
