package validator

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator/rule"
	"github.com/stretchr/testify/assert"
)

func TestNewValidator(t *testing.T) {
	// Arrange
	rules := rule.NewBuilder().Build()

	expected := &Validator{
		Rules: rules,
	}

	// Act
	resp := NewValidator(rules)

	// Assert
	assert.IsType(t, expected, resp)
}

func TestValidator_Validate(t *testing.T) {
	t.Run("Should return nil when all rules are empty", func(t *testing.T) {
		// Arrange
		rules := rule.NewBuilder().Build()

		validator := NewValidator(rules)

		// Act
		err := validator.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return nil when all rules are valid", func(t *testing.T) {
		// Arrange
		rules := rule.NewBuilder().
			StringNotEmpty("abc").
			StringMinLength("abc", 3).
			Build()

		validator := NewValidator(rules)

		// Act
		err := validator.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when one rule is invalid", func(t *testing.T) {
		// Arrange
		rules := rule.NewBuilder().
			StringNotEmpty("abc").
			StringMinLength("ac", 3).
			Build()

		validator := NewValidator(rules)

		// Act
		err := validator.Validate()

		// Assert
		assert.Error(t, err)
	})
}
