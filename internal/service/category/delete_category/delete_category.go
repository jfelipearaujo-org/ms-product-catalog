package delete_category

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type DeleteCategoryService interface {
	Handle(ctx context.Context, request DeleteCategoryDto) error
}

type Service struct {
	repository repository.CategoryRepository
}

func NewService(
	repository repository.CategoryRepository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Handle(ctx context.Context, request DeleteCategoryDto) error {
	if err := request.Validate(); err != nil {
		return err
	}

	if err := s.repository.Delete(ctx, request.UUID); err != nil {
		return err
	}

	return nil
}
