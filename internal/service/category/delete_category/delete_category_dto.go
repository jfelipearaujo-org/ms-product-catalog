package delete_category

import (
	"github.com/go-playground/validator/v10"
)

type DeleteCategoryDto struct {
	UUID string `param:"id" validate:"required,uuid4"`
}

func (dto *DeleteCategoryDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
