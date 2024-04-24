package get_categories_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/get_categories"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service get_categories.GetCategoriesService
}

func NewHandler(service get_categories.GetCategoriesService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := get_categories.GetCategoriesDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	req.SetDefaults()

	context := ctx.Request().Context()

	count, categories, err := h.service.Handle(context, req.Pagination, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	resp := common.NewPaginationResponse(req.Page, (count/req.Size)+1, count, categories)

	return ctx.JSON(http.StatusOK, resp)
}
