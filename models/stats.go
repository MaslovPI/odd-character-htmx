package models

import (
	"log"

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

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func rollStat() int {
	result, err := functions.RollMultipleDice(3, 6)
	check(err)
	return result
}

func rollHP() int {
	result, err := functions.Roll(6)
	check(err)
	return result
}

func RollStats() Stats {
	stats := Stats{
		Strength:      rollStat(),
		Dexterity:     rollStat(),
		Willpower:     rollStat(),
		HitProtection: rollHP(),
	}
	stats.populateMax()
	return stats
}
