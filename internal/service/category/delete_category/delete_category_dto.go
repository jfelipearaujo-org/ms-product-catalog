package delete_category

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
)

type DeleteCategoryDto struct {
	UUID string `param:"id" validate:"required,uuid4"`
}

func (dto *DeleteCategoryDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return errors.ErrRequestNotValid
	}

	return nil
}
