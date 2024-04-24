package get_category

import "github.com/go-playground/validator/v10"

type GetCategoryDto struct {
	Title string `json:"title" validate:"required,min=3,max=100"`
}

func (dto *GetCategoryDto) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return err
	}

	return nil
}
