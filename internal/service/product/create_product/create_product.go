package create_product

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/provider"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type CreateProductService interface {
	Handle(ctx context.Context, request CreateProductDto) (*entity.Product, error)
}

type Service struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	timeProvider       provider.TimeProvider
}

func NewService(
	productRepository repository.ProductRepository,
	categoryRepository repository.CategoryRepository,
	timeProvider provider.TimeProvider,
) *Service {
	return &Service{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		timeProvider:       timeProvider,
	}
}

func (s *Service) Handle(ctx context.Context, request CreateProductDto) (*entity.Product, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	category, err := s.categoryRepository.GetByTitle(ctx, request.CategoryTitle)
	if err != nil {
		return nil, err
	}

	product := entity.NewProduct(request.Title,
		request.Description,
		request.Price,
		category,
		s.timeProvider.GetTime())

	if err := s.productRepository.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}
