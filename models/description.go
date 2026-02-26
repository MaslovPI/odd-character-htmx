package models

type (
	NamedItem struct {
		Name        string
		Description string
	}
	Description struct {
		Content []NamedItem
		Arcana  NamedItem
		Pet     NamedItem
		Hire    NamedItem
	}
)

func (ni *NamedItem) IsEmpty() bool {
	return ni.Name == ""
}
