package create_product_handler

import (
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/product/create_product"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/product/get_products"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	createProductService create_product.CreateProductService
	getProductsService   get_products.GetProductsService
}

func NewHandler(
	createProductService create_product.CreateProductService,
	getProductsService get_products.GetProductsService,
) *Handler {
	return &Handler{
		createProductService: createProductService,
		getProductsService:   getProductsService,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	req := create_product.CreateProductDto{}

	if err := ctx.Bind(&req); err != nil {
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	context := ctx.Request().Context()

	count, _, err := h.getProductsService.Handle(context,
		common.Pagination{},
		get_products.GetProductsDto{
			Title: req.Title,
		})
	if err != nil {
		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	if count > 0 {
		return errors.NewHttpAppError(http.StatusConflict, "product cannot be created", errors.ErrProductAlreadyExists)
	}

	product, err := h.createProductService.Handle(context, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}
		if err == repository.ErrCategoryNotFound {
			return errors.NewHttpAppError(http.StatusNotFound, "error to find the category", err)
		}

		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.JSON(http.StatusCreated, product)
}
