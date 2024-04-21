package health_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should return a new handler", func(t *testing.T) {
		// Arrange

		// Act
		handler := NewHandler()

		// Assert
		assert.NotNil(t, handler)
	})
}

func TestHandler_Handle(t *testing.T) {
	t.Run("Should return a map with database status", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest(echo.GET, "/health", nil)
		resp := httptest.NewRecorder()

		echo := echo.New()
		ctx := echo.NewContext(req, resp)

		handler := NewHandler()

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"database": "healthy"}`, resp.Body.String())
	})
}
