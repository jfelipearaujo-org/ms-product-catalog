package delete_category_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/delete_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service delete_category.DeleteCategoryService
}

func NewHandler(
	service delete_category.DeleteCategoryService,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := delete_category.DeleteCategoryDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	err := h.service.Handle(context, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}
		if err == repository.ErrCategoryNotFound {
			return errors.NewHttpAppError(http.StatusNotFound, "error to find the category", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
