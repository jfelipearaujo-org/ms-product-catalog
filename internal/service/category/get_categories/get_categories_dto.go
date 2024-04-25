package get_categories

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
)

type GetCategoriesDto struct {
	Title string `query:"title" validate:"required,min=3,max=100"`

	common.Pagination
}

func (dto *GetCategoriesDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
