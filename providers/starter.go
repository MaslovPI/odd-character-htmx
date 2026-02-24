package providers

import (
	_ "embed"
	"encoding/json"
)

//go:embed data/starters.json
var starterJSON []byte

type (
	Starter struct {
		Hire    string    `json:"hire"`
		Pet     string    `json:"pet"`
		Content []Content `json:"content"`
		Max     int       `json:"max"`
		Arcana  byte      `json:"arcana"`
	}
	Content struct {
		Name      string `json:"name"`
		ExtraInfo string `json:"extra_info"`
	}
	Column struct {
		Starters      []Starter `json:"starters"`
		HitProtection int       `json:"hp"`
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
