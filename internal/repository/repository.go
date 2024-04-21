package repository

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
)

type Pagination struct {
	Page int64
	Size int64
}

type CategoryRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id string) (*entity.Category, error)
	GetByTitle(ctx context.Context, title string) (*entity.Category, error)
	GetAll(ctx context.Context, filter Pagination) (int64, []*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
}

type ProductRepository interface {
	Create(ctx context.Context, category *entity.Category) error
	GetByID(ctx context.Context, id string) (*entity.Category, error)
	GetByTitle(ctx context.Context, title string) (*entity.Category, error)
	GetByCategoryID(ctx context.Context, categoryId string, filter Pagination) (int64, []entity.Product, error)
	GetByCategoryTitle(ctx context.Context, categoryTitle string, filter Pagination) (int64, []entity.Product, error)
	GetAll(ctx context.Context, filter Pagination) (int64, []*entity.Category, error)
	Update(ctx context.Context, category *entity.Category) error
	Delete(ctx context.Context, id string) error
}
