package get_product

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type GetProductService interface {
	Handle(ctx context.Context, request GetProductDto) (entity.Product, error)
}

type Service struct {
	repository repository.ProductRepository
}

func NewService(repository repository.ProductRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Handle(ctx context.Context, request GetProductDto) (entity.Product, error) {
	if err := request.Validate(); err != nil {
		return entity.Product{}, err
	}

	return s.repository.GetByID(ctx, request.UUID)
}
