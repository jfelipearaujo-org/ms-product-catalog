package health_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/adapter/database"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/health"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	db database.DatabaseService
}

func NewHandler(db database.DatabaseService) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	dbStatus := h.db.Health()

	data := map[string]*health.HealthStatus{
		"database": dbStatus,
	}

	code := http.StatusOK

	if dbStatus.HasError() {
		code = http.StatusBadRequest
	}

	return ctx.JSON(code, data)
}
