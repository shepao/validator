package rule

type FullTag struct {
	fullTag string
}

func (ft *FullTag) SetFullTag(tag string) {
	ft.fullTag = tag
}
func (ft *FullTag) GetFullTag() string {
	return ft.fullTag
}
