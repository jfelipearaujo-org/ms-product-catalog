package create_product

import (
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator/rule"
)

const (
	TITLE_MIN_LENGTH = 3
	TITLE_MAX_LENGTH = 100

	DESCRIPTION_MIN_LENGTH = 10
	DESCRIPTION_MAX_LENGTH = 200

	CATEGORY_TITLE_MIN_LENGTH = 3
	CATEGORY_TITLE_MAX_LENGTH = 100

	PRICE_MIN_AMOUNT = 0.5
	PRICE_MAX_AMOUNT = 1000.00
)

type CreateProductDto struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CategoryTitle string  `json:"category_title"`
}

func (dto *CreateProductDto) Validate() error {
	rules := rule.NewBuilder().
		StringNotEmpty(dto.Title).
		StringNotEmpty(dto.Description).
		StringNotEmpty(dto.CategoryTitle).
		StringMinMaxLength(dto.Title, TITLE_MIN_LENGTH, TITLE_MAX_LENGTH).
		StringMinMaxLength(dto.Description, DESCRIPTION_MIN_LENGTH, DESCRIPTION_MAX_LENGTH).
		StringMinMaxLength(dto.CategoryTitle, CATEGORY_TITLE_MIN_LENGTH, CATEGORY_TITLE_MAX_LENGTH).
		CurrencyMinMaxAmount(dto.Price, PRICE_MIN_AMOUNT, PRICE_MAX_AMOUNT).
		Build()

	validator := validator.NewValidator(rules)

	if err := validator.Validate(); err != nil {
		return err
	}

	return nil
}
