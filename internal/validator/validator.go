package validator

import (
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/validator/rule"
)

type Validator struct {
	Rules []rule.RuleInput
}

func NewValidator(rules []rule.RuleInput) *Validator {
	return &Validator{
		Rules: rules,
	}
}

func (v *Validator) Validate() error {
	for _, rule := range v.Rules {
		if err := rule.Rule.Validate(rule.Value); err != nil {
			return err
		}
	}
	return nil
}
