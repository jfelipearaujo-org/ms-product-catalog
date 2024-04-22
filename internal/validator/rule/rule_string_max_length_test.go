package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleStringMaxLength_Validate(t *testing.T) {
	t.Run("Should return nil when value is less than max length", func(t *testing.T) {
		// Arrange
		rule := NewStringMaxLength(3)

		// Act
		err := rule.Validate("ab")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return nil when value is equal to max length", func(t *testing.T) {
		// Arrange
		rule := NewStringMaxLength(3)

		// Act
		err := rule.Validate("abc")

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when value is greater than max length", func(t *testing.T) {
		// Arrange
		rule := NewStringMaxLength(3)

		// Act
		err := rule.Validate("abcd")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should panic when value is not a string", func(t *testing.T) {
		// Arrange
		rule := NewStringMaxLength(3)

		// Act & Assert
		assert.Panics(t, func() {
			err := rule.Validate(1)

			assert.NoError(t, err)
		})
	})
}
