package get_products_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/product/get_products"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service get_products.GetProductsService
}

func NewHandler(service get_products.GetProductsService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := get_products.GetProductsDto{}

	req.SetDefaults()

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

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
