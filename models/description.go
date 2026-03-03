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

func (t ItemType) CSSClass() string {
	switch t {
	case Arcana:
		return "item-arcana"
	case Pet:
		return "item-pet"
	case Hire:
		return "item-hire"
	default:
		return ""
	}
}
