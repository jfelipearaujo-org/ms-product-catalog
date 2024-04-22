package rule

import "fmt"

var ErrStringMaxLength = "value must be at most %d characters"

type StringMaxLength struct {
	MaxLength int
}

func NewStringMaxLength(maxLength int) *StringMaxLength {
	return &StringMaxLength{
		MaxLength: maxLength,
	}
}

func (r *StringMaxLength) Validate(value any) error {
	if _, ok := value.(string); !ok {
		panic("value must be a string")
	}

	if len(value.(string)) > r.MaxLength {
		return fmt.Errorf(ErrStringMaxLength, r.MaxLength)
	}

	return nil
}
