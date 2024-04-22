package rule

const (
	BRL = "BRL"
)

type RuleInput struct {
	Rule  Rule
	Value any
}

type RuleBuilder struct {
	Rules []RuleInput
}

func NewBuilder() *RuleBuilder {
	return &RuleBuilder{}
}

func (rb *RuleBuilder) addRule(rule Rule, value any) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  rule,
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) StringNotEmpty(value string) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewStringNotEmpty(),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) StringMinLength(value string, minLength int) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewStringMinLength(minLength),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) StringMaxLength(value string, maxLength int) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewStringMaxLength(maxLength),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) StringMinMaxLength(value string, minLength int, maxLength int) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewStringMinLength(minLength),
		Value: value,
	})

	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewStringMaxLength(maxLength),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) CurrencyMinAmount(value float64, minAmount float64) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewCurrencyMinAmount(minAmount, BRL),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) CurrencyMaxAmount(value float64, maxAmount float64) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewCurrencyMaxAmount(maxAmount, BRL),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) CurrencyMinMaxAmount(value float64, minAmount float64, maxAmount float64) *RuleBuilder {
	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewCurrencyMinAmount(minAmount, BRL),
		Value: value,
	})

	rb.Rules = append(rb.Rules, RuleInput{
		Rule:  NewCurrencyMaxAmount(maxAmount, BRL),
		Value: value,
	})

	return rb
}

func (rb *RuleBuilder) Build() []RuleInput {
	return rb.Rules
}
