package providers

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/maslovpi/odd-character-htmx/models"
)

//go:embed data/equipment.json
var equipmentJSON []byte

type (
	Equipment struct {
		Name        string   `json:"name"`
		Cost        string   `json:"cost"`
		Description string   `json:"description"`
		Examples    []string `json:"examples"`
	}
	EquipmentProvider struct {
		equipmentMap map[string]Equipment
	}
)

func InitEquipmentProvider() (EquipmentProvider, error) {
	equipmentMap, err := getEquipmentMap()
	return EquipmentProvider{equipmentMap: equipmentMap}, err
}

func getEquipmentMap() (map[string]Equipment, error) {
	var equipmentSlice []Equipment
	err := json.Unmarshal(equipmentJSON, &equipmentSlice)
	if err != nil {
		return nil, err
	}

	equipmentMap := make(map[string]Equipment, len(equipmentSlice))

	for _, item := range equipmentSlice {
		equipmentMap[item.Name] = item
	}
	return equipmentMap, nil
}

func (e *EquipmentProvider) GetEquipmentDescription(
	name, givenDescription string,
) models.NamedItem {
	equipment, exists := e.getByName(name)
	if !exists {
		equipment, exists = e.getByExample(name)
	}

	if exists {
		return models.NamedItem{Name: name, Description: equipment.constructDescription()}
	}

	return models.NamedItem{Name: name, Description: givenDescription}
}

func (e *Equipment) constructDescription() string {
	cost := "0"
	if e.Cost != "" {
		cost = e.Cost
	}

	parts := []string{fmt.Sprintf("Cost: %s", cost)}
	if e.Description != "" {
		parts = append(parts, fmt.Sprintf("Description: %s", e.Description))
	}

	return strings.Join(parts, ", ")
}

func (e *EquipmentProvider) getByName(name string) (Equipment, bool) {
	item, exists := e.equipmentMap[name]
	return item, exists
}

func (e *EquipmentProvider) getByExample(example string) (Equipment, bool) {
	for _, item := range e.equipmentMap {
		if slices.Contains(item.Examples, example) {
			return item, true
		}
	}
	return Equipment{}, false
}
