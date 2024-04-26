package repository

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
	ErrProductNotFound  = errors.New("product not found")
)

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id string) (entity.Category, error)
	GetByTitle(ctx context.Context, title string) (entity.Category, error)
	GetAll(ctx context.Context, pagination common.Pagination, filter GetAllCategoriesFilter) (int64, []entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
}

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	GetByID(ctx context.Context, id string) (entity.Product, error)
	GetByTitle(ctx context.Context, title string) (entity.Product, error)
	GetByCategoryID(ctx context.Context, categoryId string, pagination common.Pagination) (int64, []entity.Product, error)
	GetByCategoryTitle(ctx context.Context, categoryTitle string, pagination common.Pagination) (int64, []entity.Product, error)
	GetAll(ctx context.Context, pagination common.Pagination, filter GetAllProductsFilter) (int64, []entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id string) error
}
