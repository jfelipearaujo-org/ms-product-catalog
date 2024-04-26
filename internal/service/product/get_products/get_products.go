package get_products

import (
	"context"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

type GetProductsService interface {
	Handle(ctx context.Context, pagination common.Pagination, request GetProductsDto) (int64, []entity.Product, error)
}

type Service struct {
	repository repository.ProductRepository
}

func NewService(repository repository.ProductRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) Handle(
	ctx context.Context,
	pagination common.Pagination,
	request GetProductsDto,
) (int64, []entity.Product, error) {
	pagination.SetDefaults()

	count, products, err := s.repository.GetAll(ctx, pagination, repository.GetAllProductsFilter{
		Title:         request.Title,
		CategoryTitle: request.CategoryTitle,
	})
	if err != nil {
		return 0, products, err
	}

	return count, products, nil
}
