package create_product

import (
	"github.com/go-playground/validator/v10"
)

type CreateProductDto struct {
	Title         string  `json:"title" validate:"required,min=3,max=100"`
	Description   string  `json:"description" validate:"required,min=10,max=200"`
	Price         float64 `json:"price" validate:"required,min=0.5,max=1000.0"`
	CategoryTitle string  `json:"category_title" validate:"required,min=3,max=100"`
}

func (dto *CreateProductDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
