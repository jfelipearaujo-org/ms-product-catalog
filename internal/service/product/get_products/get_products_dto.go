package get_products

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
)

type GetProductsDto struct {
	Title         string `query:"title" validate:"required,min=3,max=100"`
	CategoryTitle string `query:"category" validate:"required,min=3,max=100"`

	common.Pagination
}

func (dto *GetProductsDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
