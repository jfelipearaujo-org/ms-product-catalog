package server

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("Should return a new server", func(t *testing.T) {
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 8080,
			},
		}

		server := NewServer(config)

		assert.NotNil(t, server)
		assert.Equal(t, ":8080", server.Addr)
	})
}
