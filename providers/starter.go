package providers

import (
	_ "embed"
	"encoding/json"
	"errors"
	"strings"

	"github.com/maslovpi/odd-character-htmx/models"
)

//go:embed data/starters.json
var starterJSON []byte

type (
	Starter struct {
		Hire    string    `json:"hire"`
		Pet     string    `json:"pet"`
		Content []Content `json:"content"`
		Max     int       `json:"max"`
		Arcana  bool      `json:"arcana"`
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
		starterMap        map[Key]Starter
		equipmentProvider *EquipmentProvider
		hireProvider      *HireProvider
		arcanaProvider    *ArcanaProvider
		petProvider       *PetProvider
	}
)

func InitStarterProvider() (StarterProvider, error) {
	starterMap, err := getStarterMap()
	if err != nil {
		return StarterProvider{}, err
	}

	equipmentProvider, err := InitEquipmentProvider()
	if err != nil {
		return StarterProvider{}, err
	}

	hireProvider, err := InitHireProvider(&equipmentProvider)
	if err != nil {
		return StarterProvider{}, err
	}

	arcanaProvider, err := InitArcanaProvider()
	if err != nil {
		return StarterProvider{}, err
	}

	petProvider, err := InitPetProvider()
	if err != nil {
		return StarterProvider{}, err
	}

	return StarterProvider{
		starterMap:        starterMap,
		equipmentProvider: &equipmentProvider,
		hireProvider:      &hireProvider,
		arcanaProvider:    &arcanaProvider,
		petProvider:       &petProvider,
	}, nil
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

func (sp *StarterProvider) GenerateStarter(hp, maxStat int) (models.Description, error) {
	item, found := sp.starterMap[Key{HitProtection: hp, Max: max(maxStat, 9)}]
	if !found {
		return models.Description{}, errors.New("starter not found")
	}

	contentDescription := sp.getDescriptionFromContentList(item.Content)

	arcana := models.NamedItem{}
	if item.Arcana {
		retreivedArcana := sp.arcanaProvider.GetRandomArcana()
		arcana.Name = retreivedArcana.Name
		arcana.Description = retreivedArcana.Description
		arcana.NotEmpty = true
	}

	pet := models.NamedItem{}
	if item.Pet != "" {
		description, err := sp.petProvider.GetPetDescription(item.Pet)
		if err != nil {
			return models.Description{}, err
		}
		pet.Name = item.Pet
		pet.Description = description
		pet.NotEmpty = true
	}

	hire := models.NamedItem{}
	if item.Hire != "" {
		description, err := sp.hireProvider.GetHireDescription(item.Hire)
		if err != nil {
			return models.Description{}, err
		}
		hire.Name = item.Hire
		hire.Description = description
		hire.NotEmpty = true
	}

	starter := models.Description{
		Content: contentDescription,
		Arcana:  arcana,
		Pet:     pet,
		Hire:    hire,
	}
	return starter, nil
}

func (sp *StarterProvider) getDescriptionFromContentList(contentSlice []Content) string {
	var description strings.Builder
	for _, content := range contentSlice {
		description.WriteString(
			sp.equipmentProvider.GetEquipmentDescription(content.Name, content.ExtraInfo),
		)
	}
	return description.String()
}
