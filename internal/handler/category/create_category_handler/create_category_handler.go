package create_category_handler

import (
	"log/slog"
	"net/http"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/create_category"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/service/category/get_categories"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	createCategoryService create_category.CreateCategoryService
	getCategoriesService  get_categories.GetCategoriesService
}

func NewHandler(
	createCategoryService create_category.CreateCategoryService,
	getCategoriesService get_categories.GetCategoriesService,
) *Handler {
	return &Handler{
		createCategoryService: createCategoryService,
		getCategoriesService:  getCategoriesService,
	}
}

func (h *Handler) Handle(ctx echo.Context) error {
	context := ctx.Request().Context()

	req := create_category.CreateCategoryDto{}

	if err := ctx.Bind(&req); err != nil {
		slog.ErrorContext(context, "error binding request", slog.String("error", err.Error()))
		return errors.NewHttpAppError(http.StatusBadRequest, "invalid request", err)
	}

	count, _, err := h.getCategoriesService.Handle(context,
		common.Pagination{},
		get_categories.GetCategoriesDto{
			Title: req.Title,
		})
	if err != nil {
		slog.ErrorContext(context, "error getting categories", slog.String("error", err.Error()))
		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	if count > 0 {
		return errors.NewHttpAppError(http.StatusConflict, "category cannot be created", errors.ErrCategoryAlreadyExists)
	}

	category, err := h.createCategoryService.Handle(context, req)
	if err != nil {
		if err == errors.ErrRequestNotValid {
			return errors.NewHttpAppError(http.StatusUnprocessableEntity, "validation error", err)
		}

		slog.ErrorContext(context, "error creating category", slog.String("error", err.Error()))
		return errors.NewHttpAppError(http.StatusInternalServerError, "internal server error", err)
	}

	return ctx.JSON(http.StatusCreated, category)
}
