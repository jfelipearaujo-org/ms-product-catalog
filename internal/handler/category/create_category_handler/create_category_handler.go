package create_category_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/create_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service create_category.CreateCategoryService
}

func NewHandler(service create_category.CreateCategoryService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := create_category.CreateCategoryDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	category, err := h.service.Handle(context, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.JSON(http.StatusCreated, category)
}
