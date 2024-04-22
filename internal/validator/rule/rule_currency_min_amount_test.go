package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleCurrencyMinAmount(t *testing.T) {
	ptBR := "BRL"

	t.Run("Should return nil when value is greater than min amount", func(t *testing.T) {
		// Arrange
		rule := NewCurrencyMinAmount(10.0, ptBR)

		// Act
		err := rule.Validate(11.0)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return nil when value is equal to min amount", func(t *testing.T) {
		// Arrange
		rule := NewCurrencyMinAmount(10.0, ptBR)

		// Act
		err := rule.Validate(10.0)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when value is less than min amount", func(t *testing.T) {
		// Arrange
		rule := NewCurrencyMinAmount(10.0, ptBR)

		// Act
		err := rule.Validate(9.0)

		// Assert
		assert.Error(t, err)
	})

	t.Run("Should panic when value is not a float64", func(t *testing.T) {
		// Arrange
		rule := NewCurrencyMinAmount(10.0, ptBR)

		// Act & Assert
		assert.Panics(t, func() {
			err := rule.Validate(1)

			assert.NoError(t, err)
		})
	})
}
