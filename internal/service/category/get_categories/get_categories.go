package get_categories

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type GetCategories interface {
	Handle(ctx context.Context, pagination common.Pagination) (int64, []entity.Category, error)
}

type Service struct {
	repository repository.CategoryRepository
}

func NewService(repository repository.CategoryRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Handle(ctx context.Context, pagination common.Pagination) (int64, []entity.Category, error) {
	pagination.SetDefaults()

	count, categories, err := s.repository.GetAll(ctx, pagination)
	if err != nil {
		return 0, categories, err
	}

	return count, categories, nil
}
