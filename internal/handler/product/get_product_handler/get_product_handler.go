package get_product_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/product/get_product"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	getProductService get_product.GetProductService
}

func NewHandler(
	getProductService get_product.GetProductService,
) *Handler {
	return &Handler{
		getProductService: getProductService,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := get_product.GetProductDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	product, err := h.getProductService.Handle(context, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}
		if err == repository.ErrProductNotFound {
			return errors.NewHttpAppError(http.StatusNotFound, "error to find the product", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.JSON(http.StatusOK, product)
}
