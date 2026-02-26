package providers

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/maslovpi/odd-character-htmx/functions"
)

//go:embed data/pets.json
var petJSON []byte

type (
	Pet struct {
		Type     string `json:"type"`
		Cost     string `json:"cost"`
		Strength string `json:"str"`
		Attack   string `json:"attack"`
	}
	PetProvider struct {
		petMap map[string]Pet
	}
)

func InitPetProvider() (PetProvider, error) {
	petMap, err := getPetMap()
	return PetProvider{petMap: petMap}, err
}

func getPetMap() (map[string]Pet, error) {
	var equipmentSlice []Pet
	err := json.Unmarshal(petJSON, &equipmentSlice)
	if err != nil {
		return nil, err
	}

	petMap := make(map[string]Pet, len(equipmentSlice))

	for _, item := range equipmentSlice {
		petMap[item.Type] = item
	}
	return petMap, nil
}

func (p *PetProvider) GetPetDescription(petToFind string) (string, error) {
	item, exists := p.petMap[petToFind]
	if !exists {
		return "", nil
	}
	rolledStrength, err := functions.RollDice(item.Strength)
	if err != nil {
		return "", err
	}
	description := fmt.Sprintf(
		"Cost: %s, Strength: %d, Attack: %s",
		item.Cost,
		rolledStrength,
		item.Attack,
	)
	return description, nil
}
