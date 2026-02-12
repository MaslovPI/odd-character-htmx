package models

type Character struct {
	Stats   Stats
	Starter Starter
}

func GetEmptyChar() Character {
	character := Character{
		Stats:   Stats{},
		Starter: Starter{},
	}
	return character
}
