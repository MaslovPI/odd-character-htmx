package models

import (
	"github.com/maslovpi/odd-character-htmx/functions"
)

type Stats struct {
	Strength      int
	Dexterity     int
	Willpower     int
	Max           int
	HitProtection int
}

func (s *Stats) populateMax() {
	s.Max = max(s.Strength, s.Dexterity, s.Willpower)
}

func rollStat() (int, error) {
	return functions.RollMultipleDice(3, 6)
}

func rollHP() (int, error) {
	return functions.Roll(6)
}

func RollStats() (Stats, error) {
	str, err := rollStat()
	if err != nil {
		return Stats{}, err
	}
	dex, err := rollStat()
	if err != nil {
		return Stats{}, err
	}
	wil, err := rollStat()
	if err != nil {
		return Stats{}, err
	}
	hp, err := rollHP()
	if err != nil {
		return Stats{}, err
	}

	stats := Stats{
		Strength:      str,
		Dexterity:     dex,
		Willpower:     wil,
		HitProtection: hp,
	}
	stats.populateMax()
	return stats, nil
}
