package providers

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
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

func (a *ArcanaProvider) GetRandomArcana() Arcana {
	return a.arcanaSlice[getRandomIndex(len(a.arcanaSlice))]
}

func getRandomIndex(length int) int {
	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Intn(length)
}
