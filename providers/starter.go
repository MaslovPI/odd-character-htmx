package providers

import (
	_ "embed"
	"encoding/json"
)

//go:embed data/starters.json
var starterJSON []byte

type (
	Starter struct {
		Max     int       `json:"max"`
		Arcana  byte      `json:"arcana"`
		Content []Content `json:"content"`
		Hire    string    `json:"hire"`
		Pet     string    `json:"pet"`
	}
	Content struct {
		Name      string `json:"name"`
		ExtraInfo string `json:"extra_info"`
	}
	Column struct {
		HitProtection int       `json:"hp"`
		Starters      []Starter `json:"starters"`
	}
	Key struct {
		HitProtection, Max int
	}
	StarterProvider struct {
		starterMap map[Key]Starter
	}
)

func InitStarterProvider() (StarterProvider, error) {
	starterMap, err := getStarterMap()
	return StarterProvider{starterMap: starterMap}, err
}

func getStarterMap() (map[Key]Starter, error) {
	var columnSlice []Column
	err := json.Unmarshal(starterJSON, &columnSlice)
	if err != nil {
		return nil, err
	}

	starterMap := make(map[Key]Starter, len(columnSlice)*len(columnSlice[0].Starters))

	for _, column := range columnSlice {
		for _, starter := range column.Starters {
			starterMap[Key{HitProtection: column.HitProtection, Max: starter.Max}] = starter
		}
	}
	return starterMap, nil
}
