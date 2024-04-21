package health_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx echo.Context) error {
	data := map[string]string{
		"database": "healthy",
	}

	return ctx.JSON(http.StatusOK, data)
}
