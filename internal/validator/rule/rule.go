package rule

type Rule interface {
	Validate(value any) error
}
