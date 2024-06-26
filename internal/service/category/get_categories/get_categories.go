package get_categories

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type GetCategoriesService interface {
	Handle(ctx context.Context, pagination common.Pagination, request GetCategoriesDto) (int64, []entity.Category, error)
}

type Service struct {
	repository repository.CategoryRepository
}

func NewService(repository repository.CategoryRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Handle(
	ctx context.Context,
	pagination common.Pagination,
	request GetCategoriesDto,
) (int64, []entity.Category, error) {
	pagination.SetDefaults()

	if err := request.Validate(); err != nil {
		return 0, []entity.Category{}, err
	}

	count, categories, err := s.repository.GetAll(ctx, pagination, repository.GetAllCategoriesFilter{
		Title: request.Title,
	})
	if err != nil {
		return 0, categories, err
	}

	return count, categories, nil
}
