package models

type Character struct {
	Starter Starter
	Stats   Stats
}

func GetEmptyChar() Character {
	character := Character{
		Stats:   Stats{},
		Starter: Starter{},
	}
	return character
}
