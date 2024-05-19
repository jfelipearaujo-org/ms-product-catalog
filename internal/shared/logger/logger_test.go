package logger

import (
	"log/slog"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/stretchr/testify/assert"
)

func TestSetupLog(t *testing.T) {
	t.Run("Should setup log when is development", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				EnvName: "development",
			},
		}

		// Act
		SetupLog(config)

		// Assert
		assert.IsType(t, &slog.TextHandler{}, slog.Default().Handler())
	})

	t.Run("Should setup log when is not development", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				EnvName: "production",
			},
		}

		// Act
		SetupLog(config)

		// Assert
		assert.IsType(t, &slog.JSONHandler{}, slog.Default().Handler())
	})
}
