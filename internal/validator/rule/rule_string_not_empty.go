package rule

import "fmt"

var ErrStringNotEmpty = "value cannot be empty"

type StringNotEmpty struct{}

func NewStringNotEmpty() *StringNotEmpty {
	return &StringNotEmpty{}
}

func (r *StringNotEmpty) Validate(value any) error {
	if _, ok := value.(string); !ok {
		panic("value must be a string")
	}

	if value == "" {
		return fmt.Errorf(ErrStringNotEmpty)
	}
	return nil
}
