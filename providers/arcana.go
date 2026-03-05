package providers

import (
	_ "embed"
	"encoding/json"

	"github.com/maslovpi/odd-character-htmx/functions"
	"github.com/maslovpi/odd-character-htmx/models"
)

//go:embed data/arcana.json
var arcanaJSON []byte

type (
	Arcana struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	ArcanaProvider struct {
		arcanaSlice []Arcana
	}
)

func InitArcanaProvider() (ArcanaProvider, error) {
	arcanaSlice, err := getArcanaSlice()
	return ArcanaProvider{arcanaSlice: arcanaSlice}, err
}

func getArcanaSlice() ([]Arcana, error) {
	var arcanaSlice []Arcana
	err := json.Unmarshal(arcanaJSON, &arcanaSlice)
	return arcanaSlice, err
}

func (ap *ArcanaProvider) GetRandomArcana() models.NamedItem {
	return ap.arcanaSlice[getRandomIndex(len(ap.arcanaSlice))].ToNamedItem()
}

func (a *Arcana) ToNamedItem() models.NamedItem {
	return models.NamedItem{Name: a.Name, Description: a.Description, Type: models.Arcana}
}

func getRandomIndex(length int) int {
	return functions.GetRandomInt(length)
}
