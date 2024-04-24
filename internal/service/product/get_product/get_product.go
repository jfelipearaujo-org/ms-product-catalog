package get_product

import (
	"context"
	"errors"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/entity"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/repository"
)

var (
	ErrInvalidRequest = errors.New("invalid request, you must provide a ID or a Title")
)

type GetProduct interface {
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
	if request.UUID != "" {
		return s.repository.GetByID(ctx, request.UUID)
	}

	if request.Title != "" {
		return s.repository.GetByTitle(ctx, request.Title)
	}

	return entity.Product{}, ErrInvalidRequest
}
