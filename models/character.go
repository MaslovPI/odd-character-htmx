package models

type Character struct {
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
