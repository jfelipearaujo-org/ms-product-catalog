package get_product

import (
	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/shared/errors"
)

type GetProductDto struct {
	UUID string `param:"id"`
}

func (r GetProductDto) Validate() error {
	if _, err := uuid.Parse(r.UUID); err != nil {
		return errors.ErrRequestNotValid
	}

	return nil
}
