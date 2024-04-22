package create_category

import (
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator/rule"
)

const (
	TITLE_MIN_LENGTH = 3
	TITLE_MAX_LENGTH = 100

	DESCRIPTION_MIN_LENGTH = 10
	DESCRIPTION_MAX_LENGTH = 200
)

type CreateCategoryDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (dto *CreateCategoryDto) Validate() error {
	rules := rule.NewBuilder().
		StringNotEmpty(dto.Title).
		StringNotEmpty(dto.Description).
		StringMinMaxLength(dto.Title, TITLE_MIN_LENGTH, TITLE_MAX_LENGTH).
		StringMinMaxLength(dto.Description, DESCRIPTION_MIN_LENGTH, DESCRIPTION_MAX_LENGTH).
		Build()

	validator := validator.NewValidator(rules)

	if err := validator.Validate(); err != nil {
		return err
	}

	return nil
}
