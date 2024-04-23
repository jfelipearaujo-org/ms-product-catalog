package create_category

import (
	"github.com/go-playground/validator/v10"
)

type CreateCategoryDto struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10,max=200"`
}

func (dto *CreateCategoryDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
