package models

type (
	NamedItem struct {
		Name        string
		Description string
		NotEmpty    bool
	}
	Description struct {
		Content string
		Arcana  NamedItem
		Pet     NamedItem
		Hire    NamedItem
	}
)
