package create_category

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/provider"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type CreateCategoryService interface {
	Handle(ctx context.Context, request CreateCategoryDto) (*entity.Category, error)
}

type Service struct {
	repository   repository.CategoryRepository
	timeProvider provider.TimeProvider
}

func NewService(
	repository repository.CategoryRepository,
	timeProvider provider.TimeProvider,
) *Service {
	return &Service{
		repository:   repository,
		timeProvider: timeProvider,
	}
}

func (s *Service) Handle(ctx context.Context, request CreateCategoryDto) (*entity.Category, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	category := entity.NewCategory(
		request.Title,
		request.Description,
		s.timeProvider.GetTime())

	if err := s.repository.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}
