package create_product

import (
	"context"
	"fmt"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/provider"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type CreateProduct interface {
	Handle(ctx context.Context, request *CreateProductDto) error
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

func (s *Service) Handle(ctx context.Context, request *CreateProductDto) error {
	if err := request.Validate(); err != nil {
		return err
	}

	category, err := s.categoryRepository.GetByTitle(ctx, request.CategoryTitle)
	if err != nil {
		return fmt.Errorf("error to search the category: %w", err)
	}

	product := entity.NewProduct(request.Title,
		request.Description,
		request.Price,
		category,
		s.timeProvider.GetTime())

	if err := s.productRepository.Create(ctx, product); err != nil {
		return fmt.Errorf("error to create the product: %w", err)
	}

	return nil
}
