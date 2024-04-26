package delete_product

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type DeleteProductService interface {
	Handle(ctx context.Context, request DeleteProductDto) error
}

type Service struct {
	repository repository.ProductRepository
}

func NewService(
	repository repository.ProductRepository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Handle(ctx context.Context, request DeleteProductDto) error {
	if err := request.Validate(); err != nil {
		return err
	}

	if err := s.repository.Delete(ctx, request.UUID); err != nil {
		return err
	}

	return nil
}
