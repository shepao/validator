package rule

type Rule interface {
	Tag() string
	Generate(value interface{}, tagValue string) error
	Valid() error
	SetFullTag(tag string)
	GetFullTag() string
	Clone() Rule
}
