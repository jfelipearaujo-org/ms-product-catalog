package get_category

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
)

type GetCategoryDto struct {
	Title string `query:"title" validate:"required,min=3,max=100"`

	common.Pagination
}

func (dto *GetCategoryDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
