package rule

import "fmt"

var ErrStringMinLength = "value must be at least %d characters"

type StringMinLength struct {
	MinLength int
}

func NewStringMinLength(minLength int) *StringMinLength {
	return &StringMinLength{
		MinLength: minLength,
	}
}

func (r *StringMinLength) Validate(value any) error {
	if _, ok := value.(string); !ok {
		panic("value must be a string")
	}

	if len(value.(string)) < r.MinLength {
		return fmt.Errorf(ErrStringMinLength, r.MinLength)
	}

	return nil
}
