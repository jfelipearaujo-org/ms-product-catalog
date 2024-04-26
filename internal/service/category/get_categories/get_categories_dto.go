package get_categories

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/common"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
)

type GetCategoriesDto struct {
	Title string `query:"title" validate:"max=100"`

	common.Pagination
}

func (dto *GetCategoriesDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return errors.ErrRequestNotValid
	}

	return nil
}
