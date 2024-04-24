package get_category

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
)

type GetCategoryService interface {
	Handle(ctx context.Context, pagination common.Pagination, request GetCategoryDto) (entity.Category, error)
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
	request GetCategoryDto,
) (entity.Category, error) {
	pagination.SetDefaults()

	if err := request.Validate(); err != nil {
		return entity.Category{}, errors.ErrRequestNotValid
	}

	category, err := s.repository.GetByTitle(ctx, request.Title)
	if err != nil {
		return category, err
	}

	return category, nil
}
