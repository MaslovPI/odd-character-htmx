package models

type ItemType int
type (
	NamedItem struct {
		Name        string
		Description string
		Type        ItemType
	}
	Description struct {
		Content []NamedItem
	}
)

const (
	Equipment ItemType = iota
	Arcana
	Pet
	Hire
)

func (ni *NamedItem) IsEmpty() bool {
	return ni.Name == ""
}
