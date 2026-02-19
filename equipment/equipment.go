package equipment

import (
	_ "embed"
	"encoding/json"
	"log"
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
	EquipmentProvider map[string]Equipment
)

func getEquipmentMap() map[string]Equipment {
	var equipmentSlice []Equipment
	err := json.Unmarshal(equipmentJSON, &equipmentSlice)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	equipmentMap := make(map[string]Equipment, len(equipmentSlice))

	for _, item := range equipmentSlice {
		equipmentMap[item.Name] = item
	}
	return equipmentMap
}

func InitEquipmentProvider() EquipmentProvider {
	return getEquipmentMap()
}
