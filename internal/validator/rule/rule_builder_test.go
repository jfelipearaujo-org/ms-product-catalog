package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRule(t *testing.T) {
	t.Run("Should add rule to builder and build correctly", func(t *testing.T) {
		// Arrange
		rb := NewBuilder()

		expected := 10

		// Act
		rules := rb.
			addRule(NewStringNotEmpty(), "test").
			StringNotEmpty("test").
			StringMinLength("test", 1).
			StringMaxLength("test", 1).
			StringMinMaxLength("test", 1, 1).
			CurrencyMinAmount(1.0, 1.0).
			CurrencyMaxAmount(1.0, 1.0).
			CurrencyMinMaxAmount(1.0, 1.0, 1.0).
			Build()

		// Assert
		assert.Equal(t, expected, len(rules))
	})
}
