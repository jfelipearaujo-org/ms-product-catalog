package server

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("Should return a new server", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 8080,
			},
			DbConfig: &environment.DatabaseConfig{
				DbName:   "db",
				Host:     "localhost",
				Port:     1234,
				User:     "user",
				Password: "pass",
			},
		}

		// Act
		server := NewServer(config)

		// Assert
		assert.NotNil(t, server)
		assert.Equal(t, ":8080", server.Addr)
	})
}
