package health_handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/adapter/database/mocks"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/health"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should return a new handler", func(t *testing.T) {
		// Arrange
		db := mocks.NewMockDatabaseService(t)

		// Act
		handler := NewHandler(db)

		// Assert
		assert.NotNil(t, handler)
	})
}

func TestHandler_Handle(t *testing.T) {
	t.Run("Should return a map with healthy database status", func(t *testing.T) {
		// Arrange
		db := mocks.NewMockDatabaseService(t)

		db.On("Health").Return(&health.HealthStatus{
			Status: "healthy",
		}, nil)

		req := httptest.NewRequest(echo.GET, "/health", nil)
		resp := httptest.NewRecorder()

		echo := echo.New()
		ctx := echo.NewContext(req, resp)

		handler := NewHandler(db)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"database": {"status":"healthy"}}`, resp.Body.String())
	})

	t.Run("Should return a map with unhealthy database status", func(t *testing.T) {
		// Arrange
		db := mocks.NewMockDatabaseService(t)

		db.On("Health").Return(&health.HealthStatus{
			Status: "unhealthy",
			Err:    "error",
		}, nil)

		req := httptest.NewRequest(echo.GET, "/health", nil)
		resp := httptest.NewRecorder()

		echo := echo.New()
		ctx := echo.NewContext(req, resp)

		handler := NewHandler(db)

		// Act
		err := handler.Handle(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.JSONEq(t, `{"database": {"status":"unhealthy", "err": "error"}}`, resp.Body.String())
	})
}
