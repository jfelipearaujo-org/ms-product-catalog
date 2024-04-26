package get_category_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/get_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service get_category.GetCategoryService
}

func NewHandler(service get_category.GetCategoryService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := get_category.GetCategoryDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	category, err := h.service.Handle(context, req)
	if err != nil {
		if err == repository.ErrCategoryNotFound {
			return errors.NewHttpAppError(http.StatusNotFound, "error to find the category", err)
		}
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.JSON(http.StatusOK, category)
}
