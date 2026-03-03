package models

type Character struct {
	Name        string
	Description Description
	Stats       Stats
}

func GetEmptyChar() Character {
	character := Character{
		Stats:       Stats{},
		Description: Description{},
	}
	return character
}
